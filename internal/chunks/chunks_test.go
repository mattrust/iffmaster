// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package chunks

import "testing"

var data = []byte{0x01, 0xff, 0x07, 0x5B, 0xCD, 0x15, 0x02}

func TestGetBeUlong(t *testing.T) {
	var tests = []struct {
		name        string
		inputOffset uint32
		wantData    uint32
		wantOffset  uint32
		wantIsErr   bool
	}{
		{"Normal", 2, 123456789, 6, false},
		{"Large", 1, 4278672333, 5, false},
		{"Offside", 4, 0, 4, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ansData, err := getBeUlong(data, &tt.inputOffset)
			ansErr := err != nil

			if ansData != tt.wantData {
				t.Errorf("Data: got %d, want %d", ansData, tt.wantData)
			}
			if tt.inputOffset != tt.wantOffset {
				t.Errorf("Offset: got %d, want %d", tt.inputOffset, tt.wantOffset)
			}
			if ansErr != tt.wantIsErr {
				t.Errorf("Error: got %t, want %t", ansErr, tt.wantIsErr)
			}
		})
	}
}

func TestGetBeLong(t *testing.T) {
	var tests = []struct {
		name        string
		inputOffset uint32
		wantData    int32
		wantOffset  uint32
		wantIsErr   bool
	}{
		{"Normal", 2, 123456789, 6, false},
		{"Large", 1, -16294963, 5, false},
		{"Offside", 4, 0, 4, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ansData, err := getBeLong(data, &tt.inputOffset)
			ansErr := err != nil

			if ansData != tt.wantData {
				t.Errorf("Data: got %d, want %d", ansData, tt.wantData)
			}
			if tt.inputOffset != tt.wantOffset {
				t.Errorf("Offset: got %d, want %d", tt.inputOffset, tt.wantOffset)
			}
			if ansErr != tt.wantIsErr {
				t.Errorf("Error: got %t, want %t", ansErr, tt.wantIsErr)
			}
		})
	}
}

func TestGetBeUword(t *testing.T) {
	var tests = []struct {
		name        string
		inputOffset uint32
		wantData    uint16
		wantOffset  uint32
		wantIsErr   bool
	}{
		{"Normal", 2, 1883, 4, false},
		{"Large", 1, 65287, 3, false},
		{"Offside", 7, 0, 7, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ansData, err := getBeUword(data, &tt.inputOffset)
			ansErr := err != nil

			if ansData != tt.wantData {
				t.Errorf("Data: got %d, want %d", ansData, tt.wantData)
			}
			if tt.inputOffset != tt.wantOffset {
				t.Errorf("Offset: got %d, want %d", tt.inputOffset, tt.wantOffset)
			}
			if ansErr != tt.wantIsErr {
				t.Errorf("Error: got %t, want %t", ansErr, tt.wantIsErr)
			}
		})
	}
}

func TestGetBeWord(t *testing.T) {
	var tests = []struct {
		name        string
		inputOffset uint32
		wantData    int16
		wantOffset  uint32
		wantIsErr   bool
	}{
		{"Normal", 2, 1883, 4, false},
		{"Large", 1, -249, 3, false},
		{"Offside", 7, 0, 7, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ansData, err := getBeWord(data, &tt.inputOffset)
			ansErr := err != nil

			if ansData != tt.wantData {
				t.Errorf("Data: got %d, want %d", ansData, tt.wantData)
			}
			if tt.inputOffset != tt.wantOffset {
				t.Errorf("Offset: got %d, want %d", tt.inputOffset, tt.wantOffset)
			}
			if ansErr != tt.wantIsErr {
				t.Errorf("Error: got %t, want %t", ansErr, tt.wantIsErr)
			}
		})
	}
}

func TestGetUbyte(t *testing.T) {
	var tests = []struct {
		name        string
		inputOffset uint32
		wantData    uint8
		wantOffset  uint32
		wantIsErr   bool
	}{
		{"Normal", 2, 7, 3, false},
		{"Large", 1, 255, 2, false},
		{"Offside", 8, 0, 8, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ansData, err := getUbyte(data, &tt.inputOffset)
			ansErr := err != nil

			if ansData != tt.wantData {
				t.Errorf("Data: got %d, want %d", ansData, tt.wantData)
			}
			if tt.inputOffset != tt.wantOffset {
				t.Errorf("Offset: got %d, want %d", tt.inputOffset, tt.wantOffset)
			}
			if ansErr != tt.wantIsErr {
				t.Errorf("Error: got %t, want %t", ansErr, tt.wantIsErr)
			}
		})
	}
}

func TestGetByte(t *testing.T) {
	var tests = []struct {
		name        string
		inputOffset uint32
		wantData    int8
		wantOffset  uint32
		wantIsErr   bool
	}{
		{"Normal", 2, 7, 3, false},
		{"Large", 1, -1, 2, false},
		{"Offside", 8, 0, 8, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ansData, err := getByte(data, &tt.inputOffset)
			ansErr := err != nil

			if ansData != tt.wantData {
				t.Errorf("Data: got %d, want %d", ansData, tt.wantData)
			}
			if tt.inputOffset != tt.wantOffset {
				t.Errorf("Offset: got %d, want %d", tt.inputOffset, tt.wantOffset)
			}
			if ansErr != tt.wantIsErr {
				t.Errorf("Error: got %t, want %t", ansErr, tt.wantIsErr)
			}
		})
	}
}
