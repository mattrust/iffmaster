package chunks

import (
	"fmt"
	"log"
)

func handleAnimAnhd(data []byte) StructResult {
	log.Println("Handling ANIM.ANHD chunk")

	//typedef struct {
	//	UBYTE operation;
	//	UBYTE mask;
	//	UWORD w,h;
	//	WORD x,y;
	//	ULONG abstime;
	//	ULONG reltime;
	//	UBYTE interleave;
	//	UBYTE pad0;
	//	ULONG bits;
	//	ULONG pad[16];

	//	} AnimHeader;

	var offset uint32
	var result StructResult

	operation := getUBYTE(data, &offset)
	switch operation {
	case 0:
		result = append(result, [2]string{"Operation", "Direct"})
	case 1:
		result = append(result, [2]string{"Operation", "XOR"})
	case 2:
		result = append(result, [2]string{"Operation", "Long Delta"})
	case 3:
		result = append(result, [2]string{"Operation", "Short Delta"})
	case 4:
		result = append(result, [2]string{"Operation", "Short/Long Delta"})
	case 5:
		result = append(result, [2]string{"Operation", "Byte Vertical Delta"})
	case 6:
		result = append(result, [2]string{"Operation", "Stereo Op 5"})
	case 7:
		result = append(result, [2]string{"Operation", "Short/Long Vertical Delta"})
	case 74:
		result = append(result, [2]string{"Operation", "Graham"})
	}
	mask := getUBYTE(data, &offset)
	result = append(result, [2]string{"Mask", fmt.Sprintf("%d", mask)})

	w := getUWORD(data, &offset)
	h := getUWORD(data, &offset)
	result = append(result, [2]string{"Width : Height", fmt.Sprintf("%d : %d", w, h)})

	x := getWORD(data, &offset)
	y := getWORD(data, &offset)
	result = append(result, [2]string{"Position x : y", fmt.Sprintf("%d : %d", x, y)})

	abstime := getULONG(data, &offset)
	result = append(result, [2]string{"Absolute Time", fmt.Sprintf("%d", abstime)})

	reltime := getULONG(data, &offset)
	result = append(result, [2]string{"Relative Time", fmt.Sprintf("%d", reltime)})

	interleave := getUBYTE(data, &offset)
	result = append(result, [2]string{"Interleave", fmt.Sprintf("%d", interleave)})

	offset++ // ignore pad0

	bits := getULONG(data, &offset)
	/*
	   bit#    =0                   =1
	   0      short data          long data
	   1      store               XOR
	   2      separate info       one info for
	          for each plane      for all planes
	   3      not RLC             RLC (run length encoded)
	   4      horizontal          vertical
	   5      short info offsets  long info offsets
	*/
	if bits&(1<<0) == 0 { // bit 0 not set
		result = append(result, [2]string{"Bit 0", "Short Data"})
	} else {
		result = append(result, [2]string{"Bit 0", "Long Data"})
	}

	if bits&(1<<1) == 0 {
		result = append(result, [2]string{"Bit 1", "Store"})
	} else {
		result = append(result, [2]string{"Bit 1", "XOR"})
	}

	if bits&(1<<2) == 0 {
		result = append(result, [2]string{"Bit 2", "Separate Info"})
	} else {
		result = append(result, [2]string{"Bit 2", "One Info for All"})
	}

	if bits&(1<<3) == 0 {
		result = append(result, [2]string{"Bit 3", "Not RLC"})
	} else {
		result = append(result, [2]string{"Bit 3", "RLC"})
	}

	if bits&(1<<4) == 0 {
		result = append(result, [2]string{"Bit 4", "Horizontal"})
	} else {
		result = append(result, [2]string{"Bit 4", "Vertical"})
	}

	if bits&(1<<5) == 0 {
		result = append(result, [2]string{"Bit 5", "Short Info Offsets"})
	} else {
		result = append(result, [2]string{"Bit 5", "Long Info Offsets"})
	}

	return result
}

func handleAnimDpan(data []byte) StructResult {
	log.Println("Handling ANIM.DPAN chunk")

	//typedef struct {
	//	UWORD version;
	//	UWORD nframes;
	//	ULONG flags;
	//} DPAnimChunk;

	var offset uint32
	var result StructResult

	version := getUWORD(data, &offset)
	result = append(result, [2]string{"Version", fmt.Sprintf("%d", version)})

	nframes := getUWORD(data, &offset)
	result = append(result, [2]string{"Number of Frames", fmt.Sprintf("%d", nframes)})

	flags := getULONG(data, &offset)
	result = append(result, [2]string{"Flags", fmt.Sprintf("%032b", flags)})

	return result
}
