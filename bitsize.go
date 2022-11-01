// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

// Common sizes when measuring amounts of data in bits.
//
// To count the number of units in a BitSize, divide:
//
//	mbits := mem.MBit
//	fmt.Print(int64(mbits / mem.KBit)) // prints 1000
//
// To convert an integer of units to a BitSize, multiply:
//
//	mbits := 10
//	fmt.Print(mem.BitSize(mbits)*mem.MBit) // prints 10MBit
const (
	Bit  BitSize = 1
	KBit         = 1000 * Bit
	MBit         = 1000 * KBit
	GBit         = 1000 * MBit
	TBit         = 1000 * GBit
)

// BitSize represents an amount of data as int64 number of bits.
// The largest representable size is approximately 9223372 Tbit.
type BitSize int64

// Bytes returns b as number of bytes and any remaining bits,
// if any. It guarantees that:
//
//	bytes, bits := b.Bytes()
//	fmt.Print(b == bytes.Bits() + bits) // true
//	fmt.Print(-7 <= bits && bits <= 7) // true
func (b BitSize) Bytes() (Size, BitSize) {
	return Size(b / 8), b % 8
}

// Kilobits returns the size as floating point number of kilobits (Kbit).
func (b BitSize) Kilobits() float64 {
	m := b / KBit
	r := b % KBit
	return float64(m) + float64(r)/1e3
}

// Megabits returns the size as floating point number of megabits (Mbit).
func (b BitSize) Megabits() float64 {
	m := b / MBit
	r := b % MBit
	return float64(m) + float64(r)/1e6
}

// Gigabits returns the size as floating point number of gigabits (Gbit).
func (b BitSize) Gigabits() float64 {
	m := b / GBit
	r := b % GBit
	return float64(m) + float64(r)/1e9
}

// Terabits returns the size as floating point number of terabits (Tbit).
func (b BitSize) Terabits() float64 {
	m := b / TBit
	r := b % TBit
	return float64(m) + float64(r)/1e12
}

// Abs returns the absolute value of b. As a special case, math.MinInt64 is
// converted to math.MaxInt64.
func (b BitSize) Abs() BitSize {
	return BitSize(abs(int64(b)))
}

// Truncate returns the result of rounding b towards zero to a multiple of m.
// If m <= 0, Truncate returns b unchanged.
func (b BitSize) Truncate(m BitSize) BitSize {
	return BitSize(truncate(int64(b), int64(m)))
}

// Round returns the result of rounding b to the nearest multiple of m.
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum) value that can be
// stored in a Size, Round returns the maximum (or minimum) size.
// If m <= 0, Round returns b unchanged.
func (b BitSize) Round(m BitSize) BitSize {
	return BitSize(round(int64(b), int64(m)))
}

// String returns a string representing the bit size in the form "1.25Mbit".
// The zero size formats as 0Bit.
func (b BitSize) String() string { return FormatBitSize(b, 'D', -1) }
