// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package chunks

import (
	"fmt"
	"log"
)

// handle8svxVhdr processes the 8SVX.VHDR chunk.
func handle8svxVhdr(data []byte) (StructResult, error) {
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

	oneShotHiSamples, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"One Shot Hi Samples", fmt.Sprintf("%d", oneShotHiSamples)})

	repeatHiSamples, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Repeat Hi Samples", fmt.Sprintf("%d", repeatHiSamples)})

	samplesPerHiCycle, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Samples Per Hi Cycle", fmt.Sprintf("%d", samplesPerHiCycle)})

	samplesPerSec, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Samples Per Sec", fmt.Sprintf("%d", samplesPerSec)})

	ctOctave, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Octave", fmt.Sprintf("%d", ctOctave)})

	sCompression, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	switch sCompression {
	case 0:
		result = append(result, [2]string{"Compression", "None"})
	case 1:
		result = append(result, [2]string{"Compression", "Fibonacci-Delta-Encoded"})
	}

	volume, err := getBeLong(data, &offset) // TODO: handle Fixed type (16 bit left, 16 bit right)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Volume", fmt.Sprintf("%d", volume)})

	return result, nil
}

// handle8svxAtakRlse processes the 8SVX.ATAK or 8SVX.RLSE chunk.
func handle8svxAtakRlse(data []byte) (StructResult, error) {
	log.Println("Handling 8SVX.ATAK or 8SVX.RLSE chunk")

	//typedef struct {
	//	UWORD duration;
	//	Fixed dest;
	//	} EGPoint;

	var offset uint32
	var result StructResult

	duration, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Duration", fmt.Sprintf("%d", duration)})

	dest, err := getBeLong(data, &offset) // TODO: handle Fixed type (16 bit left, 16 bit right)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Dest", fmt.Sprintf("%d", dest)})

	return result, nil
}
