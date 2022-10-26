// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import (
	"fmt"
	"testing"
)

var formatSizeTests = []struct {
	Size    Size
	Prec    int
	D, B, I string
}{
	{Size: 0, Prec: -1, D: "0b", B: "0b", I: "0bit"},                                                                      // 0
	{Size: Bit, Prec: -1, D: "0.125b", B: "0.125b", I: "1bit"},                                                            // 1
	{Size: Byte, Prec: -1, D: "1b", B: "1b", I: "8bit"},                                                                   // 2
	{Size: -1 * Byte, Prec: -1, D: "-1b", B: "-1b", I: "-8bit"},                                                           // 3
	{Size: 1*MB + 111*KB, Prec: -1, D: "1.111mb", B: "1.05953216552734375mib", I: "8.888mbit"},                            // 4
	{Size: 1*MB + 111*KB, Prec: 2, D: "1.11mb", B: "1.06mib", I: "8.89mbit"},                                              // 5
	{Size: -1*MB - 111*KB, Prec: -1, D: "-1.111mb", B: "-1.05953216552734375mib", I: "-8.888mbit"},                        // 6
	{Size: 1*GiB + 512*MiB, Prec: -1, D: "1.610612736gb", B: "1.5gib", I: "12.884901888gbit"},                             // 7
	{Size: 5 * TBit, Prec: -1, D: "625gb", B: "582.07660913467407227gib", I: "5tbit"},                                     // 8
	{Size: MaxSize, Prec: -1, D: "1152.9215046068469759pb", B: "1023.9999999999999999pib", I: "9223372.036854775807tbit"}, // 9
	{Size: minSize, Prec: -1, D: "-1152.921504606846976pb", B: "-1024pib", I: "-9223372.036854775808tbit"},                // 10
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

var formatParseSizeTests = []Size{
	0, Bit, Byte, 512 * Byte, -Bit, -Byte, -512 * Byte,
	KBit, KB, KiB, 384 * KB, 384 * KiB, -KBit, -KB, -KiB, -732 * KB, -732 * KiB,
	MBit, MB, MiB, 18 * MB, 81 * MiB, -MBit, -MB, -MiB, -963 * MB, -963 * KiB,
	GBit, GB, GiB, 740 * GB, 59 * GiB, -GBit, -GB, -GiB, -64*GB - 837*MB - 848*Byte,
	TBit, TB, TiB, 182 * TB, 485 * TiB, -TBit, -TB, -TiB, 301*TB + 643*MB - 553*Byte,
	PB, PiB, 871 * PB, 131 * PiB, -PB, -PiB, MaxSize, minSize,
}

func TestFormatParseSize(t *testing.T) {
	fmts := []byte{'d', 'b', 'i', 'D', 'B', 'I'}
	precs := []int{-1, 16}
	for _, f := range fmts {
		for _, prec := range precs {
			for _, s := range formatParseSizeTests {
				v := FormatSize(s, f, prec)
				w, err := ParseSize(v)
				if err != nil {
					details := fmt.Sprintf("formatted '%d' with fmt='%c' and prec='%d'", s, f, prec)
					t.Fatalf("Failed to parse size string '%s' - %s", v, details)
				}
				if w != s {
					details := fmt.Sprintf("formatted '%d' with fmt='%c' and prec='%d'", s, f, prec)
					t.Fatalf("Parsed size does not match original size: got '%v' ('%d') - want '%v' ('%d') - %s", w, w, s, s, details)
				}
			}
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
	{String: "+0Bit", Size: 0},                              // 2
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
