// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	for i, test := range absTests {
		if abs := abs(test.Size); abs != test.Abs {
			t.Fatalf("Test %d: got %d - want %d", i, abs, test.Abs)
		}
	}
}

var absTests = []struct {
	Size int64
	Abs  int64
}{
	{Size: 0, Abs: 0},                             // 0
	{Size: 1, Abs: 1},                             // 1
	{Size: -1, Abs: 1},                            // 2
	{Size: math.MaxInt64, Abs: math.MaxInt64},     // 3
	{Size: math.MinInt64, Abs: math.MaxInt64},     // 4
	{Size: math.MinInt64 + 1, Abs: math.MaxInt64}, // 5
}

func TestTruncate(t *testing.T) {
	for i, test := range truncateTests {
		if trunc := truncate(test.Size, test.Mod); trunc != test.Trunc {
			t.Fatalf("Test %d: got %d - want %d", i, trunc, test.Trunc)
		}
	}
}

var truncateTests = []struct {
	Size  int64
	Mod   int64
	Trunc int64
}{
	{Size: 0, Mod: 0, Trunc: 0},                                  // 0
	{Size: 0, Mod: 1, Trunc: 0},                                  // 1
	{Size: 1, Mod: 1, Trunc: 1},                                  // 2
	{Size: 26, Mod: 8, Trunc: 24},                                // 3
	{Size: 1111, Mod: 10, Trunc: 1110},                           // 4
	{Size: 1111, Mod: -10, Trunc: 1111},                          // 5
	{Size: -1111, Mod: 10, Trunc: -1110},                         // 6
	{Size: math.MaxInt64, Mod: 1111, Trunc: 9223372036854775415}, // 7
}

func TestRound(t *testing.T) {
	for i, test := range roundTests {
		if round := round(test.Size, test.Mod); round != test.Round {
			t.Fatalf("Test %d: got %d - want %d", i, round, test.Round)
		}
	}
}

var roundTests = []struct {
	Size  int64
	Mod   int64
	Round int64
}{
	{Size: 0, Mod: 0, Round: 0},
	{Size: 1, Mod: 0, Round: 1},
	{Size: 1, Mod: 1, Round: 1},
	{Size: 8, Mod: 4, Round: 8},
	{Size: 26, Mod: 8, Round: 24},
	{Size: -999, Mod: 1000, Round: -1000},
	{Size: -1500, Mod: 1000, Round: -2000},
	{Size: 1001, Mod: 1000, Round: 1000},
	{Size: 1499, Mod: 1000, Round: 1000},
	{Size: math.MaxInt64, Mod: 2, Round: math.MaxInt64},
	{Size: math.MaxInt64, Mod: 3, Round: math.MaxInt64 - 1},
}
