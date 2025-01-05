// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package chunks

import (
	"fmt"
)

// StructResult is a list of key-value pairs.
type StructResult [][2]string

// ChunkHandler is a function that processes a chunk and returns the structured data.
type ChunkHandler func(data []byte) (StructResult, error)

// ChunkData contains the handler and the description for each chunk type.
type ChunkData struct {
	Handler     ChunkHandler
	Description string
}

// structData contains the handler and the description for each chunk type.
var structData = map[string]ChunkData{
	// generic chunks
	"(any).ANNO": {handleAnyIso8859, "Annotation"},
	"(any).AUTH": {handleAnyIso8859, "Author"},
	"(any).CHRS": {nil, "Character String"}, // can contain ANSI codes
	"(any).CSET": {nil, "Character Set"},    // binary
	"(any).FRED": {nil, "ASDG Private"},
	"(any).FVER": {handleAnyIso8859, "Version"},
	"(any).HLID": {nil, "Hotlink"},
	"(any).INFO": {nil, "Icon Data"},
	"(any).JUNK": {nil, "To Be Ignored"},
	"(any).UTF8": {handleAnyUtf8, "UTF-8 Character Text"},
	"(any).NAME": {handleAnyIso8859, "Name"},
	"(any).TEXT": {handleAnyIso8859, "ASCII Text"},
	"(any).(c) ": {handleAnyIso8859, "Copyright"},

	"8SVX":      {nil, "8-Bit Sampled Voice"},
	"8SVX.VHDR": {handle8svxVhdr, "Voice Header"},
	"8SVX.ATAK": {handle8svxAtakRlse, "Attack"},
	"8SVX.RLSE": {handle8svxAtakRlse, "Release"},

	"ACBM":      {nil, "Amiga Continuous Bitmap"},
	"ACBM.ABIT": {nil, "Bitmap Body"},
	"ACBM.BMHD": {handleIlbmBmhd, "Bitmap Header"},      // reusing ILBM
	"ACBM.CMAP": {handleIlbmCmap, "Color Map"},          // reusing ILBM
	"ACBM.GRAB": {handleIlbmGrab, "Grab (Hotspot)"},     // reusing ILBM
	"ACBM.DEST": {handleIlbmDest, "Destination"},        // reusing ILBM
	"ACBM.SPRT": {handleIlbmSprt, "Sprite"},             // reusing ILBM
	"ACBM.CAMG": {handleIlbmCamg, "Amiga Display Mode"}, // reusing ILBM

	"AIFF": {nil, "Audio Samples"},
	"ANBM": {nil, "Animated Bitmap"},

	"ANIM":      {nil, "CEL Animations"},
	"ILBM.ANHD": {handleAnimAnhd, "Animation Header"},   // parent is ILBM!
	"ILBM.DLTA": {nil, "Delta Compression"},             // parent is ILBM!
	"ILBM.DPAN": {handleAnimDpan, "Display Parameters"}, // parent is ILBM!

	"CMUS": {nil, "Musical Score"},
	"CSET": {nil, "Text Character Set"},

	"CTLG":      {nil, "Catalog"},
	"CTLG.LANG": {handleAnyIso8859, "Language"},
	"CTLG.STRS": {nil, "Strings"},

	"DEEP": {nil, "Chunky Pixel Image"},
	"DTYP": {nil, "DataType Identification"},
	"DR2D": {nil, "2-D Objects"},
	"EXEC": {nil, "Executable Code"},
	"FANT": {nil, "Movie Format"},
	"FAXX": {nil, "Facsimile Image"},
	"FTXT": {nil, "Formatted Text"},
	"FVER": {nil, "Version String"},
	"HEAD": {nil, "Flow Idea Processor Format"},
	"HLID": {nil, "Hotlink Identification"},

	"ILBM":      {nil, "InterLeaved BitMap"},
	"ILBM.BMHD": {handleIlbmBmhd, "Bitmap Header"},
	"ILBM.BODY": {nil, "Bitmap Body"},
	"ILBM.CAMG": {handleIlbmCamg, "Amiga Display Mode"},
	"ILBM.CCRT": {nil, "Color Cycling"},
	"ILBM.CMAP": {handleIlbmCmap, "Color Map"},
	"ILBM.CLUT": {nil, "Color Look Up Table"},
	"ILBM.CMYK": {nil, "Cyan Magenta Yellow Black"},
	"ILBM.CNAM": {nil, "Color Naming"},
	"ILBM.CTBL": {nil, "Dynamic Color Palette"},
	"ILBM.CRNG": {handleIlbmCrng, "Color Range"},
	"ILBM.DPPS": {nil, "DPaint Page State"},
	"ILBM.DRNG": {nil, "DPaint Range"},
	"ILBM.DYCP": {nil, "Dynamic Color Palette"},
	"ILBM.DPI ": {handleIlbmDpi, "Dots Per Inch"},
	"ILBM.DPPV": {nil, "DPaint Perspective"},
	"ILBM.DEST": {handleIlbmDest, "Destination"},
	"ILBM.EPSF": {nil, "Encapsulated Postscript"},
	"ILBM.GRAB": {handleIlbmGrab, "Grab (Hotspot)"},
	"ILBM.PCHG": {nil, "Line By line Palette"},
	"ILBM.PRVW": {nil, "Preview"},
	"ILBM.SPRT": {handleIlbmSprt, "Sprite"},
	"ILBM.TINY": {nil, "Thumbnail"},
	"ILBM.XBMI": {nil, "Extended BitMap Information"},
	"ILBM.XSSL": {nil, "3D X-Specs Image"},

	"INFO": {nil, "Icon Information"},
	"JUNK": {nil, "Junk Data"},
	"MTRX": {nil, "Matrix Data Storage"},
	"OB3D": {nil, "3-D Object Format"},
	"PGTB": {nil, "Program Traceback"},
	"PMBC": {nil, "High-color Image Format"},

	"PREF":      {nil, "Preferences"},
	"PREF.PRHD": {handlePrefPrhd, "Preferences Header"},
	"PREF.ASL ": {handlePrefAsl, "ASL Preferences"},
	"PREF.FONT": {handlePrefFont, "Font Preferences"},
	"PREF.ICTL": {handlePrefIctl, "IControl Preferences"},
	"PREF.INPT": {handlePrefInpt, "Input Preferences"},
	"PREF.KMSW": {handlePrefKmsw, "Keyboard/Mouse Preferences"},
	"PREF.LCLE": {nil, "Locale Preferences"},
	"PREF.PALT": {nil, "Palette Preferences"},
	"PREF.CMAP": {handleIlbmCmap, "Color Map"},
	"PREF.NPTR": {nil, "Pointer Preferences"},
	"PREF.PTXT": {nil, "Printer Preferences"},
	"PREF.PUNT": {nil, "Printer Unit Preferences"},
	"PREF.PDEV": {nil, "Printer Device Preferences"},
	"PREF.PGFX": {nil, "Printer Graphics Preferences"},
	"PREF.SCRM": {nil, "Screen Mode Preferences"},
	"PREF.SERL": {nil, "Serial Preferences"},
	"PREF.WANR": {nil, "Wanderer Preferences"}, // AROS Wanderer

	"PRSP": {nil, "Perspective Move"},
	"RGBN": {nil, "Image Data"},
	"RGB8": {nil, "Image Data"},
	"SAMP": {nil, "Sampled Sound"},
	"SMUS": {nil, "Simple Musical Score"},
	"SPLT": {nil, "File Splitting"},
	"TDDD": {nil, "3-D Rendering Data"},
	"TMUI": {nil, "Project File Format"},
	"TREE": {nil, "Tree Data Structure"},
	"TRKR": {nil, "Tracker Music Module"},
	"UTF8": {nil, "UTF-8 Unicode Text"},
	"WORD": {nil, "Document Storage"},
	"YUVN": {nil, "YUV Image Data"},
}

// GetStructData returns the description and the structured data of a chunk.
// - chType is the chunk type, e.g. "ILBM", "ILBM.BMHD"
// - data is the chunk data
// It returns the description and the structured data as a list of key-value pairs.
// In case of an error the incomplete result is returned.
func GetStructData(chType string, data []byte) (string, StructResult, error) {
	var result StructResult
	var description string
	var err error

	if chunkData, exists := structData[chType]; exists {
		description = chunkData.Description
		handler := chunkData.Handler
		if handler != nil {
			result, err = handler(data)
			if err != nil {
				result = append(result, [2]string{"", fmt.Sprintf("(error: %s)", err)})
			}
		} else {
			result = append(result, [2]string{"", "(not available)"})
		}
	} else {
		description = "(unknown)"
		result = append(result, [2]string{"", "(unknown)"})
	}

	return description, result, err
}

// getBeUword reads a big-endian unsigned WORD from the data at the given offset.
// The offset is incremented by 2.
// In case of an error, it returns 0 and the error. The offset is unchanged.
func getBeUword(data []byte, offset *uint32) (uint16, error) {
	var result uint16

	if len(data) < int(*offset)+2 {
		return 0, fmt.Errorf("data too short for UWORD")
	}
	result = uint16(data[*offset])<<8 | uint16(data[*offset+1])
	*offset += 2
	return result, nil
}

// getBeWord reads a big-endian signed WORD from the data at the given offset.
// The offset is incremented by 2.
// In case of an error, it returns 0 and the error. The offset is unchanged.
func getBeWord(data []byte, offset *uint32) (int16, error) {
	var result int16

	if len(data) < int(*offset)+2 {
		return 0, fmt.Errorf("data too short for WORD")
	}
	result = int16(data[*offset])<<8 | int16(data[*offset+1])
	*offset += 2
	return result, nil
}

// getBeUlong reads a big-endian unsigned LONG from the data at the given offset.
// The offset is incremented by 4.
// In case of an error, it returns 0 and the error. The offset is unchanged.
func getBeUlong(data []byte, offset *uint32) (uint32, error) {
	var result uint32

	if len(data) < int(*offset)+4 {
		return 0, fmt.Errorf("data too short for ULONG")
	}
	result = uint32(data[*offset])<<24 | uint32(data[*offset+1])<<16 |
		uint32(data[*offset+2])<<8 | uint32(data[*offset+3])
	*offset += 4
	return result, nil
}

// getBeLong reads a big-endian signed LONG from the data at the given offset.
// The offset is incremented by 4.
// In case of an error, it returns 0 and the error. The offset is unchanged.
func getBeLong(data []byte, offset *uint32) (int32, error) {
	var result int32

	if len(data) < int(*offset)+4 {
		return 0, fmt.Errorf("data too short for LONG")
	}
	result = int32(data[*offset])<<24 | int32(data[*offset+1])<<16 |
		int32(data[*offset+2])<<8 | int32(data[*offset+3])
	*offset += 4
	return result, nil
}

// getUbyte reads an unsigned BYTE from the data at the given offset.
// The offset is incremented by 1.
// In case of an error, it returns 0 and the error. The offset is unchanged.
func getUbyte(data []byte, offset *uint32) (uint8, error) {
	var result uint8

	if len(data) < int(*offset)+1 {
		return 0, fmt.Errorf("data too short for UBYTE")
	}
	result = data[*offset]
	*offset++
	return result, nil
}

// getByte reads a signed BYTE from the data at the given offset.
// The offset is incremented by 1.
// In case of an error, it returns 0 and the error. The offset is unchanged.
func getByte(data []byte, offset *uint32) (int8, error) {
	var result int8

	if len(data) < int(*offset)+1 {
		return 0, fmt.Errorf("data too short for BYTE")
	}

	result = int8(data[*offset])
	*offset++
	return result, nil
}

// getStringBuffer reads a string from the data at the given offset.
// The numer of bytes to read is given by bufLen. The offset is incremented
// by the bufLen.
// In case of an error, it returns "" and the error. The offset is unchanged.
func getStringBuffer(data []byte, offset *uint32, bufLen uint32) (string, error) {
	var result string

	if len(data) < int(*offset+bufLen) {
		return "", fmt.Errorf("data too short for String")
	}
	high := int(*offset + bufLen)
	result = string(data[*offset:high])

	*offset += bufLen
	return result, nil
}
