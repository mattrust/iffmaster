// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package chunks

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"slices"

	"fyne.io/fyne/v2"
)

type IFFChunk struct {
	// the chunk data from the file
	ID    string
	Size  uint32
	SubID string
	Data  []byte

	// the children of the chunk
	Childs []*IFFChunk

	// the sum of the size of the chunk and all its children
	SumSize int64

	// SubID for group chunks, e.g. ILBM
	// parent's SubID + ID for data chunks, e.g. ILBM.BMHD
	ChType string
}

func ReadIFFFile(reader fyne.URIReadCloser) (*IFFChunk, error) {
	var chunk *IFFChunk

	file, err := os.Open(reader.URI().Path())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileLen := fileInfo.Size()

	chunk, err = readChunk(reader, nil, fileLen, 0)

	return chunk, err
}

func readChunkID(reader io.Reader) (string, error) {
	var id [4]byte
	_, err := reader.Read(id[:])
	if err != nil {
		return "", err
	}

	return string(id[:]), nil
}

func readChunkSize(reader io.Reader) (uint32, error) {
	var size uint32
	err := binary.Read(reader, binary.BigEndian, &size)
	if err != nil {
		return 0, err
	}

	return size, nil
}

func readChunk(reader io.Reader, parentChunk *IFFChunk, maxSize int64, level int) (*IFFChunk, error) {
	var chunk IFFChunk
	var err error

	if maxSize == 0 {
		return nil, nil
	} else if maxSize < 8 {
		// we need at least 8 bytes for ID and Size
		return nil, fmt.Errorf("maxSize is < 8")
	}

	chunk.ID, err = readChunkID(reader)
	if err != nil {
		return nil, err
	}
	chunk.SumSize = 4

	chunk.Size, err = readChunkSize(reader)
	if err != nil {
		return nil, err
	}
	chunk.SumSize += 4

	if parentChunk == nil && chunk.ID != "FORM" && chunk.ID != "CAT " && chunk.ID != "LIST" {
		return nil, fmt.Errorf("file doesn't start with FORM, CAT, or LIST")
	}

	if chunk.ID == "FORM" || chunk.ID == "CAT " || chunk.ID == "LIST" || chunk.ID == "PROP" {
		// we have a group chunk
		if chunk.SumSize+4 > maxSize {
			return nil, fmt.Errorf("SumSize+4 > maxSize")
		}
		chunk.SubID, err = readChunkID(reader)
		if err != nil {
			return nil, err
		}
		chunk.SumSize += 4

		chunk.ChType = chunk.SubID
	} else {
		// we have a data chunk

		// for some generic chunks we prefix with (any)
		if isGeneric(chunk.ID) {
			chunk.ChType = "(any)." + chunk.ID
		} else {
			chunk.ChType = parentChunk.SubID + "." + chunk.ID
		}

		if chunk.SumSize+int64(chunk.Size) > maxSize {
			return nil, fmt.Errorf("SumSize+Size > maxSize")
		}
		chunk.Data = make([]byte, chunk.Size)
		_, err = reader.Read(chunk.Data)
		if err != nil {
			return nil, err
		}
		chunk.SumSize += int64(chunk.Size)
		// If chunk size is odd, read an additional byte for padding
		if chunk.Size%2 != 0 {
			if chunk.SumSize+1 > maxSize {
				return nil, fmt.Errorf("SumSize+1 > maxSize")
			}

			var padding [1]byte
			_, err = reader.Read(padding[:])
			if err != nil {
				return nil, err
			}
			chunk.SumSize++
		}
	}
	//fmt.Printf("ID: %s, Size: %d, SubID: %s\n", chunk.ID, chunk.Size, chunk.SubID)
	if chunk.ID == "FORM" || chunk.ID == "CAT " || chunk.ID == "LIST" || chunk.ID == "PROP" {
		for chunk.SumSize < int64(chunk.Size)+8 {
			child, err := readChunk(reader, &chunk, maxSize-chunk.SumSize, level+1)
			if err != nil {
				return nil, err
			}
			if child != nil {
				chunk.Childs = append(chunk.Childs, child)
				chunk.SumSize += child.SumSize
			}
		}
	}
	return &chunk, nil
}

func PrintIffChunk(chunk *IFFChunk, level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("%s %d %s %d\n", chunk.ID, chunk.Size, chunk.SubID, chunk.SumSize)
	for _, child := range chunk.Childs {
		PrintIffChunk(child, level+1)
	}
}

func isGeneric(id string) bool {

	return slices.Contains([]string{"ANNO", "AUTH", "CHRS",
		"CSET", "FRED", "FVER", "HLID", "INFO", "JUNK", "UTF8",
		"NAME", "TEXT", "(c) "}, id)
}
