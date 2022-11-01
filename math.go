// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import "math"

func abs(v int64) int64 {
	switch {
	case v >= 0:
		return v
	case v == math.MinInt64:
		return math.MaxInt64
	default:
		return -v
	}
}

func truncate(v, m int64) int64 {
	if m <= 0 {
		return v
	}
	return v - v%m
}

func round(v, m int64) int64 {
	if m <= 0 {
		return v
	}
	r := v % m
	if v < 0 {
		r = -r
		if lessThanHalf(v, m) {
			return v + r
		}
		if v1 := v - m + r; v1 < v {
			return v1
		}
		return math.MinInt64 // overflow
	}
	if lessThanHalf(r, m) {
		return v - r
	}
	if v1 := v + m - r; v1 > v {
		return v1
	}
	return math.MaxInt64 // overflow
}

func lessThanHalf(x, y int64) bool { return uint64(x)+uint64(x) < uint64(y) }
