// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package chunks

import (
	"fmt"
	"log"
)

func handlePrefPrhd(data []byte) (StructResult, error) {
	log.Println("Handling PREF.PRHD chunk")

	//struct FilePrefHeader
	//{
	//	UBYTE ph_Version;
	//	UBYTE ph_Type;
	//	UBYTE ph_Flags[4];
	//};

	var offset uint32
	var result StructResult

	phVersion, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Version", fmt.Sprintf("%d", phVersion)})

	phType, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Type", fmt.Sprintf("%d", phType)})

	for i := 0; i < 4; i++ {
		phFlags, err := getUbyte(data, &offset)
		if err != nil {
			return result, err
		}
		result = append(result, [2]string{fmt.Sprintf("Flags %d", i), fmt.Sprintf("%0b", phFlags)})
	}
	return result, nil
}
