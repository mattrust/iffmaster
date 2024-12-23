package chunks

import (
	"fmt"
	"log"
)

func handleIlbmBmhd(data []byte) StructResult {
	log.Println("Handling ILBM.BMHD chunk")

	//typedef struct {
	//  UWORD       w, h;
	//	WORD        x, y;
	//	UBYTE       nPlanes;
	//	Masking     masking;
	//	Compression compression;
	//	UBYTE       pad1;
	//	UWORD       transparentColor;
	//	UBYTE       xAspect, yAspect;
	//	WORD        pageWidth, pageHeight;
	//  } BitmapHeader;

	var offset uint32
	var result StructResult

	w := getUWORD(data, &offset)
	h := getUWORD(data, &offset)
	result = append(result, [2]string{"Width : Height", fmt.Sprintf("%d : %d", w, h)})

	x := getWORD(data, &offset)
	y := getWORD(data, &offset)
	result = append(result, [2]string{"Position x : y", fmt.Sprintf("%d : %d", x, y)})

	nPlanes := getUBYTE(data, &offset)
	result = append(result, [2]string{"Number of planes", fmt.Sprintf("%d", nPlanes)})

	masking := getUBYTE(data, &offset)
	switch masking {
	case 0:
		result = append(result, [2]string{"Masking", "None"})
	case 1:
		result = append(result, [2]string{"Masking", "Has Mask"})
	case 2:
		result = append(result, [2]string{"Masking", "Has Transparent Color"})
	case 3:
		result = append(result, [2]string{"Masking", "Lasso"})
	}

	compression := getUBYTE(data, &offset)
	switch compression {
	case 0:
		result = append(result, [2]string{"Compression", "None"})
	case 1:
		result = append(result, [2]string{"Compression", "Byte Run 1"})
	}

	offset++ // ignore pad1

	transparentColor := getUWORD(data, &offset)
	result = append(result, [2]string{"Transparent Color", fmt.Sprintf("%d", transparentColor)})

	xAspect := getUBYTE(data, &offset)
	yAspect := getUBYTE(data, &offset)
	result = append(result, [2]string{"Aspect Ratio x : y", fmt.Sprintf("%d : %d", xAspect, yAspect)})

	pageWidth := getWORD(data, &offset)
	pageHeight := getWORD(data, &offset)
	result = append(result, [2]string{"Page Width : Height", fmt.Sprintf("%d : %d", pageWidth, pageHeight)})

	return result
}

func handleIlbmCmap(data []byte) StructResult {
	log.Println("Handling ILBM.CMAP chunk")

	//typedef struct {
	//	UBYTE red, green, blue;
	//	} ColorRegister;
	//typedef ColorRegister ColorMap[n];

	var offset uint32
	var result StructResult

	n := len(data) / 3

	for i := 0; i < n; i++ {
		red := data[offset]
		green := data[offset+1]
		blue := data[offset+2]
		result = append(result, [2]string{fmt.Sprintf("Color %d", i),
			fmt.Sprintf("%d : %d : %d", red, green, blue)})
		offset += 3
	}
	return result
}

func handleIlbmGrab(data []byte) StructResult {
	log.Println("Handling ILBM.GRAB chunk")

	//typedef struct {
	//	WORD x, y;
	//} Point2D;

	var offset uint32
	var result StructResult

	x := getWORD(data, &offset)
	y := getWORD(data, &offset)
	result = append(result, [2]string{"Position x : y", fmt.Sprintf("%d : %d", x, y)})

	return result
}

func handleIlbmCamg(data []byte) StructResult {
	log.Println("Handling ILBM.CAMG chunk")

	var offset uint32
	var result StructResult

	viewMode := getULONG(data, &offset)

	// TODO: decode viewMode
	result = append(result, [2]string{"View Mode", fmt.Sprintf("0x%x", viewMode)})

	return result
}

func handleIlbmDpi(data []byte) StructResult {
	log.Println("Handling ILBM.DPI chunk")

	// typedef struct {
	//	UWORD dpi_x;
	//	UWORD dpi_y;
	// } DPIHeader ;

	var offset uint32
	var result StructResult

	hDPI := getUWORD(data, &offset)
	vDPI := getUWORD(data, &offset)
	result = append(result, [2]string{"Horizontal DPI", fmt.Sprintf("%d", hDPI)})
	result = append(result, [2]string{"Vertical DPI", fmt.Sprintf("%d", vDPI)})

	return result
}
