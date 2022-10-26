// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import "testing"

var formatSizeTests = []struct {
	Size    Size
	Prec    int
	D, B, I string
}{
	{Size: 0, Prec: -1, D: "0B", B: "0B", I: "0bit"},                                                                      // 0
	{Size: Bit, Prec: -1, D: "0.125B", B: "0.125B", I: "1bit"},                                                            // 1
	{Size: Byte, Prec: -1, D: "1B", B: "1B", I: "8bit"},                                                                   // 2
	{Size: -1 * Byte, Prec: -1, D: "-1B", B: "-1B", I: "-8bit"},                                                           // 3
	{Size: 1*MB + 111*KB, Prec: -1, D: "1.111MB", B: "1.05953216552734375MiB", I: "8.888Mbit"},                            // 4
	{Size: 1*MB + 111*KB, Prec: 2, D: "1.11MB", B: "1.06MiB", I: "8.89Mbit"},                                              // 5
	{Size: -1*MB - 111*KB, Prec: -1, D: "-1.111MB", B: "-1.05953216552734375MiB", I: "-8.888Mbit"},                        // 6
	{Size: 1*GiB + 512*MiB, Prec: -1, D: "1.610612736GB", B: "1.5GiB", I: "12.884901888Gbit"},                             // 7
	{Size: 5 * TBit, Prec: -1, D: "625GB", B: "582.07660913467407227GiB", I: "5Tbit"},                                     // 8
	{Size: MaxSize, Prec: -1, D: "1152.9215046068469759PB", B: "1023.9999999999999999PiB", I: "9223372.036854775807Tbit"}, // 9
	{Size: minSize, Prec: -1, D: "-1152.921504606846976PB", B: "-1024PiB", I: "-9223372.036854775808Tbit"},                // 10
}

func TestFormatSize(t *testing.T) {
	for i, test := range formatSizeTests {
		if d := FormatSize(test.Size, 'd', test.Prec); d != test.D {
			t.Fatalf("Test %d: format 'd': got %s - want %s", i, d, test.D)
		}
		if b := FormatSize(test.Size, 'b', test.Prec); b != test.B {
			t.Fatalf("Test %d: format 'b': got %s - want %s", i, b, test.B)
		}
		if s := FormatSize(test.Size, 'i', test.Prec); s != test.I {
			t.Fatalf("Test %d: format 'i': got %s - want %s", i, s, test.I)
		}
	}
}

var parseSizeTests = []struct {
	String     string
	Size       Size
	ShouldFail bool
}{
	{String: "0B", Size: 0},                                 // 0
	{String: "0bit", Size: 0},                               // 1
	{String: "+0bit", Size: 0},                              // 2
	{String: "-0bit", Size: 0},                              // 3
	{String: "1bit", Size: Bit},                             // 4
	{String: "+1bit", Size: Bit},                            // 5
	{String: "0.125B", Size: Bit},                           // 6
	{String: "8bit", Size: Byte},                            // 7
	{String: "1B", Size: Byte},                              // 8
	{String: "-1B", Size: -Byte},                            // 9
	{String: "1.5B", Size: Byte + 4*Bit},                    // 10
	{String: "-1.75B", Size: -Byte - 6*Bit},                 // 11
	{String: "1MB", Size: MB},                               // 12
	{String: "1.111MB", Size: 1*MB + 111*KB},                // 13
	{String: "1.05953216552734375MiB", Size: 1*MB + 111*KB}, // 14
	{String: "8.888Mbit", Size: 1*MB + 111*KB},              // 15
	{String: "1.610612736gb", Size: 1*GiB + 512*MiB},        // 16
	{String: "1.5gib", Size: 1*GiB + 512*MiB},               // 17
	{String: "12.884901888gbit", Size: 1*GiB + 512*MiB},     // 18
	{String: "582.07660913467407227GiB", Size: 5 * TBit},    // 19
	{String: "1023.9999999999999999PiB", Size: MaxSize},     // 20
	{String: "9223372036854775807bit", Size: MaxSize},       // 21
	{String: "-1024PiB", Size: minSize},                     // 22
	{String: "-1152.921504606846976PB", Size: minSize},      // 23

	{String: "0", ShouldFail: true},
	{String: "0Bit", ShouldFail: true},
	{String: "--0bit", ShouldFail: true},
	{String: "+-0bit", ShouldFail: true},
	{String: " 0B", ShouldFail: true},
	{String: "0B ", ShouldFail: true},
	{String: "0.125.0B ", ShouldFail: true},
	{String: "1.25.0B ", ShouldFail: true},
}

func TestParseSize(t *testing.T) {
	for i, test := range parseSizeTests {
		size, err := ParseSize(test.String)
		if err == nil && test.ShouldFail {
			t.Fatalf("Test %d should have failed", i)
		}
		if err != nil && !test.ShouldFail {
			t.Fatalf("Test %d: failed to parse Size: %v", i, err)
		}
		if err != nil {
			continue
		}
		if size != test.Size {
			t.Fatalf("Test %d: got '%d (%s)' - want %d (%s)", i, size, size, test.Size, test.Size)
		}
	}
}
