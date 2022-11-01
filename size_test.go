// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import (
	"math"
	"testing"
)

func TestSize_String(t *testing.T) {
	for i, test := range sizeStringTests {
		if s := test.Size.String(); s != test.String {
			t.Fatalf("Test %d: got %s - want %s", i, s, test.String)
		}
	}
}

var sizeStringTests = []struct {
	Size   Size
	String string
}{
	{Size: 0, String: "0B"},                                  // 0
	{Size: Byte, String: "1B"},                               // 1
	{Size: MB, String: "1MB"},                                // 2
	{Size: -MB, String: "-1MB"},                              // 3
	{Size: MiB, String: "1.048576MB"},                        // 4
	{Size: 5*TB + 640*GB + 509*MB, String: "5.640509TB"},     // 5
	{Size: -1*MB - 825*KB, String: "-1.825MB"},               // 6
	{Size: 1000*PB + Byte, String: "1000.000000000000001PB"}, // 7
}

func TestSize_Bits(t *testing.T) {
	for i, test := range sizeBitsTests {
		if bits := test.Size.Bits(); bits != test.Bits {
			t.Fatalf("Test %d: got %v - want %v", i, bits, test.Bits)
		}
	}
}

var sizeBitsTests = []struct {
	Size Size
	Bits BitSize
}{
	{Size: 0, Bits: 0},                                 // 0
	{Size: 1, Bits: 8},                                 // 1
	{Size: -1, Bits: -8},                               // 2
	{Size: MB, Bits: 8 * MBit},                         // 3
	{Size: math.MaxInt64, Bits: math.MaxInt64},         // 4
	{Size: math.MinInt64, Bits: math.MinInt64},         // 5
	{Size: math.MaxInt64 / 8, Bits: math.MaxInt64 - 7}, // 6
}

func TestSize_Kilobytes(t *testing.T) {
	for i, test := range sizeConvertTests {
		if bytes := test.Size.Kilobytes(); bytes != test.KB {
			t.Fatalf("Test %d: got %f - want %f", i, bytes, test.KB)
		}
	}
}

func TestSize_Megabytes(t *testing.T) {
	for i, test := range sizeConvertTests {
		if bytes := test.Size.Megabytes(); bytes != test.MB {
			t.Fatalf("Test %d: got %f - want %f", i, bytes, test.MB)
		}
	}
}

func TestSize_Gigabytes(t *testing.T) {
	for i, test := range sizeConvertTests {
		if bytes := test.Size.Gigabytes(); bytes != test.GB {
			t.Fatalf("Test %d: got %f - want %f", i, bytes, test.GB)
		}
	}
}

func TestSize_Terabytes(t *testing.T) {
	for i, test := range sizeConvertTests {
		if bytes := test.Size.Terabytes(); bytes != test.TB {
			t.Fatalf("Test %d: got %f - want %f", i, bytes, test.TB)
		}
	}
}

func TestSize_Petabytes(t *testing.T) {
	for i, test := range sizeConvertTests {
		if bytes := test.Size.Petabytes(); bytes != test.PB {
			t.Fatalf("Test %d: got %f - want %f", i, bytes, test.PB)
		}
	}
}

var sizeConvertTests = []struct {
	Size Size
	KB   float64
	MB   float64
	GB   float64
	TB   float64
	PB   float64
}{
	{ // 0
		Size: 1 * Byte,
		KB:   0.001,
		MB:   0.000001,
		GB:   0.000000001,
		TB:   0.000000000001,
		PB:   0.000000000000001,
	},
	{ // 1
		Size: 1*KB + 172*Byte,
		KB:   1.172,
		MB:   0.001172,
		GB:   0.000001172,
		TB:   0.000000001172,
		PB:   0.000000000001172,
	},
	{ // 2
		Size: 954*KB + 744*Byte,
		KB:   954.744,
		MB:   0.954744,
		GB:   0.000954744,
		TB:   0.000000954744,
		PB:   0.000000000954744,
	},
	{ // 3
		Size: 12*MB + 271*Byte,
		KB:   12000.271,
		MB:   12.000271,
		GB:   0.012000271,
		TB:   0.000012000271,
		PB:   0.000000012000271,
	},
	{ // 4
		Size: 117*TB + 4*KB,
		KB:   117000000004,
		MB:   117000000.004,
		GB:   117000.000004,
		TB:   117.000000004,
		PB:   0.117000000004,
	},
	{ // 5
		Size: 512*PB + 813*TB + 4*GB + 996*MB + 16*KB + 672*Byte,
		KB:   512813004996016.672,
		MB:   512813004996.016672,
		GB:   512813004.996016672,
		TB:   512813.004996016672,
		PB:   512.813004996016672,
	},
}
