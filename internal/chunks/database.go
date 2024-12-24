// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package chunks

import (
	"log"
)

type StructResult [][2]string
type ChunkHandler func(data []byte) StructResult
type ChunkData struct {
	Handler     ChunkHandler
	Description string
}

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

func GetStructData(chType string, data []byte) (string, StructResult) {
	var result StructResult
	var description string

	if chunkData, exists := structData[chType]; exists {
		description = chunkData.Description
		handler := chunkData.Handler
		if handler != nil {
			result = handler(data)
		} else {
			result = append(result, [2]string{"", "(not available)"})
		}
	} else {
		description = "(unknown)"
		result = append(result, [2]string{"", "(unknown)"})
	}

	return description, result
}

func getUWORD(data []byte, offset *uint32) uint16 {
	var result uint16

	if len(data) < int(*offset)+2 {
		log.Printf("Data too short for UWORD")
		return 0
	}
	result = uint16(data[*offset])<<8 | uint16(data[*offset+1])
	*offset += 2
	return result
}

func getWORD(data []byte, offset *uint32) int16 {
	var result int16

	if len(data) < int(*offset)+2 {
		log.Printf("Data too short for WORD")
		return 0
	}
	result = int16(data[*offset])<<8 | int16(data[*offset+1])
	*offset += 2
	return result
}

func getULONG(data []byte, offset *uint32) uint32 {
	var result uint32

	if len(data) < int(*offset)+4 {
		log.Printf("Data too short for ULONG")
		return 0
	}
	result = uint32(data[*offset])<<24 | uint32(data[*offset+1])<<16 |
		uint32(data[*offset+2])<<8 | uint32(data[*offset+3])
	*offset += 4
	return result
}

func getLONG(data []byte, offset *uint32) int32 {
	var result int32

	if len(data) < int(*offset)+4 {
		log.Printf("Data too short for LONG")
		return 0
	}
	result = int32(data[*offset])<<24 | int32(data[*offset+1])<<16 |
		int32(data[*offset+2])<<8 | int32(data[*offset+3])
	*offset += 4
	return result
}

func getUBYTE(data []byte, offset *uint32) uint8 {
	var result uint8

	if len(data) < int(*offset)+1 {
		log.Printf("Data too short for UBYTE")
		return 0
	}
	result = data[*offset]
	*offset++
	return result
}

func getBYTE(data []byte, offset *uint32) int8 {
	var result int8

	if len(data) < int(*offset)+1 {
		log.Printf("Data too short for BYTE")
		return 0
	}

	result = int8(data[*offset])
	*offset++
	return result
}
