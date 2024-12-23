package chunks

import (
	"fmt"
	"log"
)

func handle8svxVhdr(data []byte) StructResult {
	log.Println("Handling 8SVX.VHDR chunk")

	//typedef struct {
	//	ULONG oneShotHiSamples,
	//		  repeatHiSamples,
	//		  samplesPerHiCycle;
	//	UWORD samplesPerSec;
	//	UBYTE ctOctave,
	//		  sCompression;
	//	Fixed volume;
	//	} Voice8Header;

	var offset uint32
	var result StructResult

	oneShotHiSamples := getULONG(data, &offset)
	result = append(result, [2]string{"One Shot Hi Samples", fmt.Sprintf("%d", oneShotHiSamples)})

	repeatHiSamples := getULONG(data, &offset)
	result = append(result, [2]string{"Repeat Hi Samples", fmt.Sprintf("%d", repeatHiSamples)})

	samplesPerHiCycle := getULONG(data, &offset)
	result = append(result, [2]string{"Samples Per Hi Cycle", fmt.Sprintf("%d", samplesPerHiCycle)})

	samplesPerSec := getUWORD(data, &offset)
	result = append(result, [2]string{"Samples Per Sec", fmt.Sprintf("%d", samplesPerSec)})

	ctOctave := getUBYTE(data, &offset)
	result = append(result, [2]string{"Octave", fmt.Sprintf("%d", ctOctave)})

	sCompression := getUBYTE(data, &offset)
	switch sCompression {
	case 0:
		result = append(result, [2]string{"Compression", "None"})
	case 1:
		result = append(result, [2]string{"Compression", "Fibonacci-Delta-Encoded"})
	}

	volume := getLONG(data, &offset) // TODO: handle Fixed type (16 bit left, 16 bit right)
	result = append(result, [2]string{"Volume", fmt.Sprintf("%d", volume)})
	return result

}

func handle8svxAtakRlse(data []byte) StructResult {
	log.Println("Handling 8SVX.ATAK or 8SVX.RLSE chunk")

	//typedef struct {
	//	UWORD duration;
	//	Fixed dest;
	//	} EGPoint;

	var offset uint32
	var result StructResult

	duration := getUWORD(data, &offset)
	result = append(result, [2]string{"Duration", fmt.Sprintf("%d", duration)})

	dest := getLONG(data, &offset) // TODO: handle Fixed type (16 bit left, 16 bit right)
	result = append(result, [2]string{"Dest", fmt.Sprintf("%d", dest)})

	return result
}
