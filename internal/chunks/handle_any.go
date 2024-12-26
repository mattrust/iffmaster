// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package chunks

import (
	"golang.org/x/text/encoding/charmap"
)

// handleAnyIso8859 processes any chunk with ISO-8859-1 encoding.
func handleAnyIso8859(data []byte) (StructResult, error) {
	var result StructResult

	decoded, err := charmap.ISO8859_1.NewDecoder().Bytes(data)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"String", string(decoded)})

	return result, nil
}

// handleAnyUtf8 processes any chunk with UTF-8 encoding.
func handleAnyUtf8(data []byte) (StructResult, error) {
	var result StructResult

	result = append(result, [2]string{"String", string(data)})

	return result, nil
}
