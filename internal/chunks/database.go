package chunks

import "log"

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

func NewChunkData() {

	//DOC Document formatting information
	//FOOT Footer information of a document
	//HEAD Header information of a document
	//PAGE Page break indicator
	//PARA Paragraph formatting information
	//PDEF Deluxe Print page definition
	//TABS Tab positions
	//TEXT Text for a paragraph

}

func GetChunkDescription(chType string) string {
	desc := chunkDescriptions[chType]
	log.Printf("Chunk %s: %s\n", chType, desc)
	return desc
}
