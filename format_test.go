// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import (
	"fmt"
	"testing"
)

var formatSizeTests = []struct {
	Size Size
	Prec int
	D, B string
}{
	{Size: 0, Prec: -1, D: "0b", B: "0b"},                                                 // 0
	{Size: Byte, Prec: -1, D: "1b", B: "1b"},                                              // 2
	{Size: -1 * Byte, Prec: -1, D: "-1b", B: "-1b"},                                       // 3
	{Size: 1*MB + 111*KB, Prec: -1, D: "1.111mb", B: "1.05953216552734375mib"},            // 4
	{Size: 1*MB + 111*KB, Prec: 2, D: "1.11mb", B: "1.06mib"},                             // 5
	{Size: -1*MB - 111*KB, Prec: -1, D: "-1.111mb", B: "-1.05953216552734375mib"},         // 6
	{Size: 1*GiB + 512*MiB, Prec: -1, D: "1.610612736gb", B: "1.5gib"},                    // 7
	{Size: 5 * TBit, Prec: -1, D: "625gb", B: "582.07660913467407227gib"},                 // 8
	{Size: MaxSize, Prec: -1, D: "9223.372036854775807pb", B: "8191.9999999999999991pib"}, // 8
}

func TestFormatSize(t *testing.T) {
	for i, test := range formatSizeTests {
		if d := FormatSize(test.Size, 'd', test.Prec); d != test.D {
			t.Fatalf("Test %d: format 'd': got %s - want %s", i, d, test.D)
		}
		if b := FormatSize(test.Size, 'b', test.Prec); b != test.B {
			t.Fatalf("Test %d: format 'b': got %s - want %s", i, b, test.B)
		}
	}
}

var formatParseSizeTests = []Size{
	0, Byte, 512 * Byte, -Byte, -512 * Byte,
	KBit, KB, KiB, 384 * KB, 384 * KiB, -KBit, -KB, -KiB, -732 * KB, -732 * KiB,
	MBit, MB, MiB, 18 * MB, 81 * MiB, -MBit, -MB, -MiB, -963 * MB, -963 * KiB,
	GBit, GB, GiB, 740 * GB, 59 * GiB, -GBit, -GB, -GiB, -64*GB - 837*MB - 848*Byte,
	TBit, TB, TiB, 182 * TB, 485 * TiB, -TBit, -TB, -TiB, 301*TB + 643*MB - 553*Byte,
	PB, PiB, 871 * PB, 131 * PiB, -PB, -PiB, MaxSize, minSize,
}

func TestFormatParseSize(t *testing.T) {
	fmts := []byte{'d', 'b', 'D', 'B'}
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
	{String: "+1b", Size: Byte},                             // 1
	{String: "-1b", Size: -Byte},                            // 2
	{String: "1B", Size: Byte},                              // 3
	{String: "-8B", Size: -8 * Byte},                        // 4
	{String: "1MB", Size: MB},                               // 5
	{String: "1.111MB", Size: 1*MB + 111*KB},                // 6
	{String: "1.05953216552734375MiB", Size: 1*MB + 111*KB}, // 7
	{String: "1.610612736gb", Size: 1*GiB + 512*MiB},        // 8
	{String: "1.5gib", Size: 1*GiB + 512*MiB},               // 9
	{String: "582.07660913467407227GiB", Size: 5 * TBit},    // 10
	{String: "8191.99999999999999991PiB", Size: MaxSize},    // 11

	{String: "0", ShouldFail: true},          // 12
	{String: "--0b", ShouldFail: true},       // 13
	{String: "+-0b", ShouldFail: true},       // 14
	{String: " 0B", ShouldFail: true},        // 15
	{String: "0B ", ShouldFail: true},        // 16
	{String: "1.125.0KB ", ShouldFail: true}, // 17
	{String: "1.25.0KB ", ShouldFail: true},  // 18
	{String: "8bit ", ShouldFail: true},      // 19
	{String: "8Kbit ", ShouldFail: true},     // 20
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
