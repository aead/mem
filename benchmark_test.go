// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import "testing"

func BenchmarkFormatSize(b *testing.B) {
	formatSize := func(s Size, fmt byte, prec int, b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			FormatSize(s, fmt, prec)
		}
	}
	b.Run("0b-d-∞", func(b *testing.B) { formatSize(0, 'd', -1, b) })
	b.Run("10b-d-∞", func(b *testing.B) { formatSize(10, 'd', -1, b) })
	b.Run("1mb-d-∞", func(b *testing.B) { formatSize(MB, 'd', -1, b) })
	b.Run("-1mb-d-∞", func(b *testing.B) { formatSize(MB, 'd', -1, b) })
	b.Run("1mb-b-∞", func(b *testing.B) { formatSize(MB, 'b', -1, b) })
	b.Run("1mb-b-4", func(b *testing.B) { formatSize(MB, 'd', 4, b) })
}

func BenchmarkFormatBitSize(b *testing.B) {
	formatSize := func(s BitSize, fmt byte, prec int, b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			FormatBitSize(s, fmt, prec)
		}
	}
	b.Run("0bit-d-∞", func(b *testing.B) { formatSize(0, 'd', -1, b) })
	b.Run("10bit-d-∞", func(b *testing.B) { formatSize(10, 'd', -1, b) })
	b.Run("1mbit-d-∞", func(b *testing.B) { formatSize(MBit, 'd', -1, b) })
	b.Run("-1mbit-d-∞", func(b *testing.B) { formatSize(MBit, 'd', -1, b) })
}

func BenchmarkParseSize(b *testing.B) {
	parseSize := func(s string, b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if _, err := ParseSize(s); err != nil {
				b.Fatal(err)
			}
		}
	}
	b.Run("0b", func(b *testing.B) { parseSize("0b", b) })
	b.Run("1mb", func(b *testing.B) { parseSize("1mb", b) })
}

func BenchmarkParseBitSize(b *testing.B) {
	parseSize := func(s string, b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if _, err := ParseBitSize(s); err != nil {
				b.Fatal(err)
			}
		}
	}
	b.Run("0bit", func(b *testing.B) { parseSize("0bit", b) })
	b.Run("8.888Kbit", func(b *testing.B) { parseSize("8.888Kbit", b) })
}
