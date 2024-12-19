package chunks

import "log"

var chunkDescriptions = map[string]string{
	"ILBM":      "InterLeaved BitMap",
	"ILBM.BMHD": "Bitmap Header",
	"ILBM.CMAP": "Color Map",
	"ILBM.GRAB": "Grab (Hotspot)",
	"ILBM.DEST": "Destination",
	"ILBM.SPRT": "Sprite",
	"ILBM.CAMG": "Amiga Display Mode",
	"ILBM.BODY": "Bitmap Body",
	"ILBM.CRNG": "Color Range",
	"ILBM.CCRT": "Color Cycling",
	"ILBM.TINY": "Thumbnail",
	"ILBM.DPPS": "DPaint Page State",
	"ILBM.DRNG": "DPaint Range",
}

func NewChunkData() {
	//(C) Copyright notice and license
	//ANNO Annotation or comment
	//AUTH Author name
	//DOC Document formatting information
	//FOOT Footer information of a document
	//HEAD Header information of a document
	//PAGE Page break indicator
	//PARA Paragraph formatting information
	//PDEF Deluxe Print page definition
	//TABS Tab positions
	//TEXT Text for a paragraph
	//VERS File version
}

func GetChunkDescription(chType string) string {
	desc := chunkDescriptions[chType]
	log.Printf("Chunk %s: %s\n", chType, desc)
	return desc
}
