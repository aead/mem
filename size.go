// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import "math"

// Common sizes for measuring memory and disk capacity.
//
// To count the number of units in a Size, divide:
//
//	megabyte := mem.MB
//	fmt.Print(int64(megabyte / mem.KB)) // prints 1000
//
// To convert an integer of units to a Size, multiply:
//
//	megabytes := 10
//	fmt.Print(mem.Size(megabytes)*mem.MB) // prints 10MB
const (
	Byte Size = 1

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

// Size represents an amount of data as int64 number of bytes.
// The largest representable size is approximately 8192 PiB.
type Size int64

// Bits returns s as number of bits. As special cases, if s would
// be greater resp. smaller than the max. resp. min. representable
// BitSize it returns math.MaxInt64 resp. math.MinInt64.
func (s Size) Bits() BitSize {
	switch {
	case s > math.MaxInt64/8:
		return math.MaxInt64
	case s < math.MinInt64/8:
		return math.MinInt64
	default:
		return BitSize(8 * s)
	}
}

// Kilobytes returns the size as floating point number of kilobytes (KB).
func (s Size) Kilobytes() float64 {
	k := s / KB
	r := s % KB
	return float64(k) + float64(r)/1e3
}

// Megabytes returns the size as floating point number of megabytes (MB).
func (s Size) Megabytes() float64 {
	m := s / MB
	r := s % MB
	return float64(m) + float64(r)/1e6
}

// Gigabytes returns the size as floating point number of gigabytes (GB).
func (s Size) Gigabytes() float64 {
	g := s / GB
	r := s % GB
	return float64(g) + float64(r)/1e9
}

// Terabytes returns the size as floating point number of terabytes (TB).
func (s Size) Terabytes() float64 {
	t := s / TB
	r := s % TB
	return float64(t) + float64(r)/1e12
}

// Petabytes returns the size as floating point number of petabytes (PB).
func (s Size) Petabytes() float64 {
	p := s / PB
	r := s % PB
	return float64(p) + float64(r)/1e15
}

// Kibibytes returns the size as floating point number of kibibytes (KiB).
func (s Size) Kibibytes() float64 {
	k := s / KiB
	r := s % KiB
	return float64(k) + float64(r)/(1<<10)
}

// Mebibytes returns the size as floating point number of mebibytes (MiB).
func (s Size) Mebibytes() float64 {
	m := s / MiB
	r := s % MiB
	return float64(m) + float64(r)/(1<<20)
}

// Gibibytes returns the size as floating point number of gibibytes (GiB).
func (s Size) Gibibytes() float64 {
	g := s / GiB
	r := s % GiB
	return float64(g) + float64(r)/(1<<30)
}

// Tebibytes returns the size as floating point number of tebibytes (TiB).
func (s Size) Tebibytes() float64 {
	t := s / TiB
	r := s % TiB
	return float64(t) + float64(r)/(1<<40)
}

// Pebibytes returns the size as floating point number of pebibytes (PiB).
func (s Size) Pebibytes() float64 {
	p := s / PiB
	r := s % PiB
	return float64(p) + float64(r)/(1<<50)
}

// Abs returns the absolute value of s. As a special case, math.MinInt64 is
// converted to math.MaxInt64.
func (s Size) Abs() Size {
	return Size(abs(int64(s)))
}

// Truncate returns the result of rounding s towards zero to a multiple of m.
// If m <= 0, Truncate returns s unchanged.
func (s Size) Truncate(m Size) Size {
	return Size(truncate(int64(s), int64(m)))
}

// Round returns the result of rounding s to the nearest multiple of m.
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum) value that can be
// stored in a Size, Round returns the maximum (or minimum) size.
// If m <= 0, Round returns s unchanged.
func (s Size) Round(m Size) Size {
	return Size(round(int64(s), int64(m)))
}

// String returns a string representing the size in the form "1.25MB".
// The zero size formats as 0B.
func (s Size) String() string { return FormatSize(s, 'D', -1) }
