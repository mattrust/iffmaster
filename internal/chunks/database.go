// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package chunks

import (
	"log"
)

type ChunkHandler func(data []byte) StructResult
type StructResult [][2]string

var chunkHandlers = map[string]ChunkHandler{
	"ILBM.BMHD": handleIlbmBmhd,
	"ILBM.CMAP": handleIlbmCmap,
	"ILBM.GRAB": handleIlbmGrab,
	"ILBM.CAMG": handleIlbmCamg,
	"ILBM.DPI ": handleIlbmDpi,
}

var chunkDescriptions = map[string]string{
	"(ANY).(C) ": "Copyright notice and license",
	"(ANY).ANNO": "Annotation or comment",
	"(ANY).AUTH": "Author name",
	"(ANY).VERS": "File version",

	"8SVX": "8-Bit Sampled Voice",
	"ACBM": "Amiga Continuous Bitmap",
	"AIFF": "Audio Samples",
	"ANBM": "Animated Bitmap",
	"ANIM": "CEL Animations",
	"CMUS": "Musical Score",
	"CSET": "Text Character Set",
	"DEEP": "Chunky Pixel Image",
	"DTYP": "DataType Identification",
	"DR2D": "2-D Objects",
	"EXEC": "Executable Code",
	"FANT": "Movie Format",
	"FAXX": "Facsimile Image",
	"FTXT": "Formatted Text",
	"FVER": "Version String",
	"HEAD": "Flow Idea Processor Format",
	"HLID": "Hotlink Identification",

	"ILBM":      "InterLeaved BitMap",
	"ILBM.BMHD": "Bitmap Header",
	"ILBM.BODY": "Bitmap Body",
	"ILBM.CAMG": "Amiga Display Mode",
	"ILBM.CCRT": "Color Cycling",
	"ILBM.CMAP": "Color Map",
	"ILBM.CLUT": "Color Look Up Table",
	"ILBM.CMYK": "Cyan Magenta Yellow Black",
	"ILBM.CNAM": "Color Naming",
	"ILBM.CTBL": "Dynamic Color Palette",
	"ILBM.CRNG": "Color Range",
	"ILBM.DPPS": "DPaint Page State",
	"ILBM.DRNG": "DPaint Range",
	"ILBM.DYCP": "Dynamic Color Palette",
	"ILBM.DPI ": "Dots Per Inch",
	"ILBM.DPPV": "DPaint Perspective",
	"ILBM.DEST": "Destination",
	"ILBM.EPSF": "Encapsulated Postscript",
	"ILBM.GRAB": "Grab (Hotspot)",
	"ILBM.PCHG": "Line By line Palette",
	"ILBM.PRVW": "Preview",
	"ILBM.SPRT": "Sprite",
	"ILBM.TINY": "Thumbnail",
	"ILBM.XBMI": "Extended BitMap Information",
	"ILBM.XSSL": "3D X-Specs Image",

	"INFO": "Icon Information",
	"JUNK": "Junk Data",
	"MTRX": "Matrix Data Storage",
	"OB3D": "3-D Object Format",
	"PGTB": "Program Traceback",
	"PMBC": "High-color Image Format",
	"PRSP": "Perspective Move",
	"RGBN": "Image Data",
	"RGB8": "Image Data",
	"SAMP": "Sampled Sound",
	"SMUS": "Simple Musical Score",
	"SPLT": "File Splitting",
	"TDDD": "3-D Rendering Data",
	"TMUI": "Project File Format",
	"TREE": "Tree Data Structure",
	"TRKR": "Tracker Music Module",
	"UTF8": "UTF-8 Unicode Text",
	"WORD": "Document Storage",
	"YUVN": "YUV Image Data",
}

func GetChunkDescription(chType string) string {
	desc := chunkDescriptions[chType]
	log.Printf("Chunk %s: %s\n", chType, desc)
	return desc
}

func GetChunkStructure(chType string, data []byte) StructResult {
	var result StructResult

	if handler, exists := chunkHandlers[chType]; exists {
		result = handler(data)
	}
	return result
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
