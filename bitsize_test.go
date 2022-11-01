// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import (
	"math"
	"testing"
)

func TestBitSize_String(t *testing.T) {
	for i, test := range bitsizeStringTests {
		if s := test.Size.String(); s != test.String {
			t.Fatalf("Test %d: got %s - want %s", i, s, test.String)
		}
	}
}

var bitsizeStringTests = []struct {
	Size   BitSize
	String string
}{
	{Size: 0, String: "0Bit"},                                 // 0
	{Size: Bit, String: "1Bit"},                               // 1
	{Size: 8*KBit + 172*Bit, String: "8.172Kbit"},             // 2
	{Size: MBit, String: "1Mbit"},                             // 3
	{Size: -MBit, String: "-1Mbit"},                           // 4
	{Size: math.MaxInt64, String: "9223372.036854775807Tbit"}, // 5
}

func TestBitSize_Bytes(t *testing.T) {
	for i, test := range bitsizeBytesTests {
		bytes, bits := test.Size.Bytes()
		if bytes != test.Bytes {
			t.Fatalf("Test %d: got %v - want %v", i, bytes, test.Bytes)
		}
		if bits != test.Bits {
			t.Fatalf("Test %d: got %v - want %v", i, bits, test.Bits)
		}
	}
}

var bitsizeBytesTests = []struct {
	Size  BitSize
	Bytes Size
	Bits  BitSize
}{
	{Size: 0, Bytes: 0, Bits: 0},                             // 0
	{Size: 1, Bytes: 0, Bits: 1},                             // 1
	{Size: -1, Bytes: 0, Bits: -1},                           // 2
	{Size: 8, Bytes: 1, Bits: 0},                             // 3
	{Size: -8, Bytes: -1, Bits: 0},                           // 4
	{Size: KBit, Bytes: 125, Bits: 0},                        // 5
	{Size: MBit + 4, Bytes: 125 * KB, Bits: 4},               // 6
	{Size: math.MaxInt64, Bytes: math.MaxInt64 / 8, Bits: 7}, // 7
	{Size: math.MinInt64, Bytes: math.MinInt64 / 8, Bits: 0}, // 8
}

func TestBitSize_Kilobits(t *testing.T) {
	for i, test := range bitsizeConvertTests {
		if bits := test.Size.Kilobits(); bits != test.KBit {
			t.Fatalf("Test %d: got %f - want %f", i, bits, test.KBit)
		}
	}
}

func TestBitSize_Megabits(t *testing.T) {
	for i, test := range bitsizeConvertTests {
		if bits := test.Size.Megabits(); bits != test.MBit {
			t.Fatalf("Test %d: got %f - want %f", i, bits, test.MBit)
		}
	}
}

func TestBitSize_Gigabits(t *testing.T) {
	for i, test := range bitsizeConvertTests {
		if bits := test.Size.Gigabits(); bits != test.GBit {
			t.Fatalf("Test %d: got %f - want %f", i, bits, test.GBit)
		}
	}
}

func TestBitSize_Terabits(t *testing.T) {
	for i, test := range bitsizeConvertTests {
		if bits := test.Size.Terabits(); bits != test.TBit {
			t.Fatalf("Test %d: got %f - want %f", i, bits, test.TBit)
		}
	}
}

var bitsizeConvertTests = []struct {
	Size BitSize
	KBit float64
	MBit float64
	GBit float64
	TBit float64
}{
	{ // 0
		Size: Bit,
		KBit: 0.001,
		MBit: 0.000001,
		GBit: 0.000000001,
		TBit: 0.000000000001,
	},
	{ // 1
		Size: 1*KBit + 172*Bit,
		KBit: 1.172,
		MBit: 0.001172,
		GBit: 0.000001172,
		TBit: 0.000000001172,
	},
	{ // 2
		Size: 954*KBit + 744*Bit,
		KBit: 954.744,
		MBit: 0.954744,
		GBit: 0.000954744,
		TBit: 0.000000954744,
	},
	{ // 3
		Size: 12*MBit + 271*Bit,
		KBit: 12000.271,
		MBit: 12.000271,
		GBit: 0.012000271,
		TBit: 0.000012000271,
	},
	{ // 4
		Size: 117*TBit + 4*KBit,
		KBit: 117000000004,
		MBit: 117000000.004,
		GBit: 117000.000004,
		TBit: 117.000000004,
	},
}
