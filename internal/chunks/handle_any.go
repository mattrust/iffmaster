package chunks

import (
	"golang.org/x/text/encoding/charmap"
)

func handleAnyIso8859(data []byte) (StructResult, error) {
	var result StructResult

	decoded, err := charmap.ISO8859_1.NewDecoder().Bytes(data)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"String", string(decoded)})

	return result, nil
}

func handleAnyUtf8(data []byte) (StructResult, error) {
	var result StructResult

	result = append(result, [2]string{"String", string(data)})

	return result, nil
}
