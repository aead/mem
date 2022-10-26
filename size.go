// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

// Common sizes used to measure internet and network speed.
const (
	Bit  Size = 1
	KBit      = 1000 * Bit
	MBit      = 1000 * KBit
	GBit      = 1000 * MBit
	TBit      = 1000 * GBit
)

// Common sizes used to measure memory and disk capacity.
//
// There are two commonly used unit systems for measuring capacities.
// The decimal unit of data uses the kilo, mega giga prefixes as
// multipliers of 1000. For example, 1000 byte = 1 kilobyte (KB) and
// 1000 KB = 1 megabyte (MB). This definition has been incorporated
// into the International System of Quantities. It is consistent
// with other unit systems for computers, like CPU clock or networking
// speeds.
//
// In contrast, the binary unit of data uses the kibi, mebi, gibi
// prefixes as multipliers of 1024 or 2^10. For example,
// 1024 byte = 1 kibibyte (KiB) and 1024 kibibyte = 1 mebibyte (MiB).
// The binary unit of data is used by e.g. MS Windows for computer
// memory like RAM or caches.
//
// Both, the binary and decimal unit of data, use the one byte (8 bit)
// as base unit.
//
// To count the number of units in a Size, divide:
//
//	megabyte := mem.MB
//	fmt.Print(int64(megabyte / mem.KB)) // prints 1000
//
// To convert an integer of units to a Size, multiply:
//
//	megabytes := 10
//	fmt.Print(mem.Size(megabytes)*mem.MB) // prints 10mb
const (
	Byte = 8 * Bit

	KB Size = 1000 * Byte
	MB      = 1000 * KB
	GB      = 1000 * MB
	TB      = 1000 * GB
	PB      = 1000 * TB

	KiB Size = 1024 * Byte
	MiB      = 1024 * KiB
	GiB      = 1024 * MiB
	TiB      = 1024 * GiB
	PiB      = 1024 * TiB
)

const (
	// MaxSize is the largest representable size: 1024PiB - 1bit.
	MaxSize Size = 1<<63 - 1

	minSize Size = -1 << 63
)

// Size represents an amount of memory as int64 number of bits.
// The largest representable size is approximately one exabyte.
type Size int64

// Kilobits returns the size as floating point number of kilobits (KBit).
func (s Size) Kilobits() float64 {
	k := s / KBit
	r := s % KBit
	return float64(k) + float64(r)/1e3
}

// Megabits returns the size as floating point number of megabits (MBit).
func (s Size) Megabits() float64 {
	m := s / MBit
	r := s % MBit
	return float64(m) + float64(r)/1e6
}

// Gigabits returns the size as floating point number of gigabits (GBit).
func (s Size) Gigabits() float64 {
	g := s / GBit
	r := s % GBit
	return float64(g) + float64(r)/1e9
}

// Terabits returns the size as floating point number of terabits (TBit).
func (s Size) Terabits() float64 {
	t := s / TBit
	r := s % TBit
	return float64(t) + float64(r)/1e12
}

// Bytes returns the size as floating point number of bytes.
func (s Size) Bytes() float64 {
	b := s / Byte
	r := s % Byte
	return float64(b) + float64(r)/8
}

// Kilobytes returns the size as floating point number of kilobytes (KB).
func (s Size) Kilobytes() float64 {
	k := s / KB
	r := s % KB
	return float64(k) + float64(r)/(8*1e3)
}

// Megabytes returns the size as floating point number of megabytes (MB).
func (s Size) Megabytes() float64 {
	m := s / MB
	r := s % MB
	return float64(m) + float64(r)/(8*1e6)
}

// Gigabytes returns the size as floating point number of gigabytes (GB).
func (s Size) Gigabytes() float64 {
	g := s / GB
	r := s % GB
	return float64(g) + float64(r)/(8*1e9)
}

// Terabytes returns the size as floating point number of terabytes (TB).
func (s Size) Terabytes() float64 {
	t := s / TB
	r := s % TB
	return float64(t) + float64(r)/(8*1e12)
}

// Petabytes returns the size as floating point number of petabytes (PB).
func (s Size) Petabytes() float64 {
	p := s / PB
	r := s % PB
	return float64(p) + float64(r)/(8*1e15)
}

// Kibibytes returns the size as floating point number of kibibytes (KiB).
func (s Size) Kibibytes() float64 {
	k := s / KiB
	r := s % KiB
	return float64(k) + float64(r)/(8*(1<<10))
}

// Mebibytes returns the size as floating point number of mebibytes (MiB).
func (s Size) Mebibytes() float64 {
	m := s / MiB
	r := s % MiB
	return float64(m) + float64(r)/(8*(1<<20))
}

// Gibibytes returns the size as floating point number of gibibytes (GiB).
func (s Size) Gibibytes() float64 {
	g := s / GiB
	r := s % GiB
	return float64(g) + float64(r)/(8*(1<<30))
}

// Tebibytes returns the size as floating point number of tebibytes (TiB).
func (s Size) Tebibytes() float64 {
	t := s / TiB
	r := s % TiB
	return float64(t) + float64(r)/(8*(1<<40))
}

// Pebibytes returns the size as floating point number of pebibytes (PiB).
func (s Size) Pebibytes() float64 {
	p := s / PiB
	r := s % PiB
	return float64(p) + float64(r)/(8*(1<<50))
}

// Abs returns the absolute value of s. As a special case, math.MinInt64 is
// converted to math.MaxInt64.
func (s Size) Abs() Size {
	switch {
	case s >= 0:
		return s
	case s == minSize:
		return MaxSize
	default:
		return -s
	}
}

// Truncate returns the result of rounding s towards zero to a multiple of m.
// If m <= 0, Truncate returns s unchanged.
func (s Size) Truncate(m Size) Size {
	if m <= 0 {
		return s
	}
	return s - s%m
}

// Round returns the result of rounding s to the nearest multiple of m.
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum) value that can be
// stored in a Size, Round returns the maximum (or minimum) size.
// If m <= 0, Round returns s unchanged.
func (s Size) Round(m Size) Size {
	if m <= 0 {
		return s
	}
	r := s % m
	if s < 0 {
		r = -r
		if lessThanHalf(s, m) {
			return s + r
		}
		if s1 := s - m + r; s1 < s {
			return s1
		}
		return minSize // overflow
	}
	if lessThanHalf(r, m) {
		return s - r
	}
	if s1 := s + m - r; s1 > s {
		return s1
	}
	return MaxSize // overflow
}

// String returns a string representing the size in the form "1.25MB".
// The zero size formats as 0B.
func (s Size) String() string { return FormatSize(s, 'd', -1) }

func lessThanHalf(x, y Size) bool {
	return uint64(x)+uint64(x) < uint64(y)
}
