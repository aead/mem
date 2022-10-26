// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import "testing"

var sizeStringTests = []struct {
	Size   Size
	String string
}{
	{Size: 0, String: "0B"},                                    // 0
	{Size: Bit, String: "0.125B"},                              // 1
	{Size: Byte, String: "1B"},                                 // 2
	{Size: MB, String: "1MB"},                                  // 3
	{Size: -MB, String: "-1MB"},                                // 4
	{Size: MiB, String: "1.048576MB"},                          // 5
	{Size: MBit, String: "125KB"},                              // 6
	{Size: 5*TB + 640*GB + 509*MB, String: "5.640509TB"},       // 7
	{Size: -1*MB - 825*KB, String: "-1.825MB"},                 // 8
	{Size: 1000*PB + Bit, String: "1000.000000000000000125PB"}, // 9
	{Size: 1000*PB + Byte, String: "1000.000000000000001PB"},   // 10
}

func TestSize_String(t *testing.T) {
	for i, test := range sizeStringTests {
		if s := test.Size.String(); s != test.String {
			t.Fatalf("Test %d: got %s - want %s", i, s, test.String)
		}
	}
}

func TestSize_Abs(t *testing.T) {
	for i, test := range sizeAbsTests {
		if abs := test.Size.Abs(); abs != test.Abs {
			t.Fatalf("Test %d: got %d - want %d", i, abs, test.Abs)
		}
	}
}

var sizeAbsTests = []struct {
	Size Size
	Abs  Size
}{
	{Size: 0, Abs: 0},                 // 0
	{Size: 1 * Bit, Abs: 1 * Bit},     // 1
	{Size: -1 * Bit, Abs: 1 * Bit},    // 2
	{Size: -1 * Byte, Abs: 1 * Byte},  // 3
	{Size: minSize, Abs: MaxSize},     // 4
	{Size: minSize + 1, Abs: MaxSize}, // 5
}

func TestSize_Truncate(t *testing.T) {
	for i, test := range sizeTruncateTests {
		if trunc := test.Size.Truncate(test.Mod); trunc != test.Trunc {
			t.Fatalf("Test %d: got %d - want %d", i, trunc, test.Trunc)
		}
	}
}

var sizeTruncateTests = []struct {
	Size, Mod, Trunc Size
}{
	{ // 0
		Size:  1*MB + 500*KB,
		Mod:   0,
		Trunc: 1*MB + 500*KB,
	},
	{ // 1
		Size:  1*MB + 500*KB,
		Mod:   -1 * KB,
		Trunc: 1*MB + 500*KB,
	},
	{ // 2
		Size:  1*MB + 500*KB,
		Mod:   1*MB + 500*KB,
		Trunc: 1*MB + 500*KB,
	},
	{ // 3
		Size:  1*MB + 500*KB,
		Mod:   1 * MB,
		Trunc: 1 * MB,
	},
	{ // 4
		Size:  11*KB + 7*Bit,
		Mod:   11 * KB,
		Trunc: 11 * KB,
	},
	{ // 5
		Size:  -(12*GiB + 5*MiB + 7*Byte),
		Mod:   1 * KiB,
		Trunc: -(12*GiB + 5*MiB),
	},
}

func TestSize_Round(t *testing.T) {
	for i, test := range sizeRoundTests {
		if round := test.Size.Round(test.Mod); round != test.Round {
			t.Fatalf("Test %d: got %d - want %d", i, round, test.Round)
		}
	}
}

var sizeRoundTests = []struct {
	Size, Mod, Round Size
}{
	{ // 0
		Size:  1*MB + 500*KB,
		Mod:   0,
		Round: 1*MB + 500*KB,
	},
	{ // 1
		Size:  1*MB + 500*KB,
		Mod:   -1 * KB,
		Round: 1*MB + 500*KB,
	},
	{ // 2
		Size:  1*MB + 500*KB,
		Mod:   500 * KB,
		Round: 1*MB + 500*KB,
	},
	{ // 3
		Size:  1*MB + 500*KB,
		Mod:   1 * MB,
		Round: 2 * MB,
	},
	{ // 4
		Size:  -(1*MB + 500*KB),
		Mod:   1 * MB,
		Round: -2 * MB,
	},
	{ // 5
		Size:  1*MB + 500*KB,
		Mod:   13 * KB,
		Round: 1*MB + 495*KB,
	},
	{ // 6
		Size:  1*MiB + 512*KiB,
		Mod:   13 * KiB,
		Round: 1*MiB + 510*KiB,
	},
	{ // 7
		Size:  1 * MaxSize,
		Mod:   1 * Byte,
		Round: 1 * MaxSize,
	},
}

func TestSize_Bytes(t *testing.T) {
	for i, test := range sizeByteTests {
		if bytes := test.Size.Bytes(); bytes != test.Bytes {
			t.Fatalf("Test %d: got %f - want %f", i, bytes, test.Bytes)
		}
	}
}

func TestSize_Kilobytes(t *testing.T) {
	for i, test := range sizeByteTests {
		if bytes := test.Size.Kilobytes(); bytes != test.KB {
			t.Fatalf("Test %d: got %f - want %f", i, bytes, test.KB)
		}
	}
}

func TestSize_Megabytes(t *testing.T) {
	for i, test := range sizeByteTests {
		if bytes := test.Size.Megabytes(); bytes != test.MB {
			t.Fatalf("Test %d: got %f - want %f", i, bytes, test.MB)
		}
	}
}

func TestSize_Gigabytes(t *testing.T) {
	for i, test := range sizeByteTests {
		if bytes := test.Size.Gigabytes(); bytes != test.GB {
			t.Fatalf("Test %d: got %f - want %f", i, bytes, test.GB)
		}
	}
}

func TestSize_Terabytes(t *testing.T) {
	for i, test := range sizeByteTests {
		if bytes := test.Size.Terabytes(); bytes != test.TB {
			t.Fatalf("Test %d: got %f - want %f", i, bytes, test.TB)
		}
	}
}

func TestSize_Petabytes(t *testing.T) {
	for i, test := range sizeByteTests {
		if bytes := test.Size.Petabytes(); bytes != test.PB {
			t.Fatalf("Test %d: got %f - want %f", i, bytes, test.PB)
		}
	}
}

var sizeByteTests = []struct {
	Size  Size
	Bytes float64
	KB    float64
	MB    float64
	GB    float64
	TB    float64
	PB    float64
}{
	{ // 0
		Size:  1 * Bit,
		Bytes: 0.125,
		KB:    0.000125,
		MB:    0.000000125,
		GB:    0.000000000125,
		TB:    0.000000000000125,
		PB:    0.000000000000000125,
	},
	{ // 1
		Size:  1 * Byte,
		Bytes: 1,
		KB:    0.001,
		MB:    0.000001,
		GB:    0.000000001,
		TB:    0.000000000001,
		PB:    0.000000000000001,
	},
	{ // 2
		Size:  1*Byte + 3*Bit,
		Bytes: 1.375,
		KB:    0.001375,
		MB:    0.000001375,
		GB:    0.000000001375,
		TB:    0.000000000001375,
		PB:    0.000000000000001375,
	},
	{ // 3
		Size:  1*KB + 172*Byte,
		Bytes: 1172,
		KB:    1.172,
		MB:    0.001172,
		GB:    0.000001172,
		TB:    0.000000001172,
		PB:    0.000000000001172,
	},
	{ // 4
		Size:  954*KB + 744*Byte,
		Bytes: 954744,
		KB:    954.744,
		MB:    0.954744,
		GB:    0.000954744,
		TB:    0.000000954744,
		PB:    0.000000000954744,
	},
	{ // 5
		Size:  12*MB + 271*Byte,
		Bytes: 12000271,
		KB:    12000.271,
		MB:    12.000271,
		GB:    0.012000271,
		TB:    0.000012000271,
		PB:    0.000000012000271,
	},
	{ // 6
		Size:  542*GB + 1*MB + 17*KB + 859*Byte + 4*Bit,
		Bytes: 542001017859.5,
		KB:    542001017.8595,
		MB:    542001.0178595,
		GB:    542.0010178595,
		TB:    0.5420010178595,
		PB:    0.0005420010178595,
	},
	{ // 7
		Size:  117*TB + 4*KB,
		Bytes: 117000000004000,
		KB:    117000000004,
		MB:    117000000.004,
		GB:    117000.000004,
		TB:    117.000000004,
		PB:    0.117000000004,
	},
	{ // 8
		Size:  512*PB + 813*TB + 4*GB + 996*MB + 16*KB + 672*Byte,
		Bytes: 512813004996016672,
		KB:    512813004996016.672,
		MB:    512813004996.016672,
		GB:    512813004.996016672,
		TB:    512813.004996016672,
		PB:    512.813004996016672,
	},
}

func TestSize_Kilobits(t *testing.T) {
	for i, test := range sizeBitTests {
		if bits := test.Size.Kilobits(); bits != test.KBits {
			t.Fatalf("Test %d: got %f - want %f", i, bits, test.KBits)
		}
	}
}

func TestSize_Megabits(t *testing.T) {
	for i, test := range sizeBitTests {
		if bits := test.Size.Megabits(); bits != test.MBits {
			t.Fatalf("Test %d: got %f - want %f", i, bits, test.MBits)
		}
	}
}

func TestSize_Gigabits(t *testing.T) {
	for i, test := range sizeBitTests {
		if bits := test.Size.Gigabits(); bits != test.GBits {
			t.Fatalf("Test %d: got %f - want %f", i, bits, test.GBits)
		}
	}
}

func TestSize_Terabits(t *testing.T) {
	for i, test := range sizeBitTests {
		if bits := test.Size.Terabits(); bits != test.TBits {
			t.Fatalf("Test %d: got %f - want %f", i, bits, test.TBits)
		}
	}
}

var sizeBitTests = []struct {
	Size  Size
	KBits float64
	MBits float64
	GBits float64
	TBits float64
}{
	{ // 0
		Size:  1 * Bit,
		KBits: 0.001,
		MBits: 0.000001,
		GBits: 0.000000001,
		TBits: 0.000000000001,
	},
	{ // 1
		Size:  1 * KBit,
		KBits: 1,
		MBits: 0.001,
		GBits: 0.000001,
		TBits: 0.000000001,
	},
	{ // 2
		Size:  1*KBit + 512*Bit,
		KBits: 1.512,
		MBits: 0.001512,
		GBits: 0.000001512,
		TBits: 0.000000001512,
	},
	{ // 3
		Size:  64*KBit + 375*Bit,
		KBits: 64.375,
		MBits: 0.064375,
		GBits: 0.000064375,
		TBits: 0.000000064375,
	},
	{ // 4
		Size:  1*MBit + 784*KBit + 63*Bit,
		KBits: 1784.063,
		MBits: 1.784063,
		GBits: 0.001784063,
		TBits: 0.000001784063,
	},
	{ // 5
		Size:  12*GBit + 632*MBit + 76*KBit + 901*Bit,
		KBits: 12632076.901,
		MBits: 12632.076901,
		GBits: 12.632076901,
		TBits: 0.012632076901,
	},
	{ // 6
		Size:  402*TBit + 551*GBit + 209*MBit + 811*KBit + 4*Bit,
		KBits: 402551209811.004,
		MBits: 402551209.811004,
		GBits: 402551.209811004,
		TBits: 402.551209811004,
	},
}
