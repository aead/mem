// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

// Common bandwidths for measuring internet / network speed.
const (
	BitPerSecond  Bandwidth = 1
	KBitPerSecond           = 1000 * BitPerSecond
	MBitPerSecond           = 1000 * KBitPerSecond
	GBitPerSecond           = 1000 * MBitPerSecond
	TBitPerSecond           = 1000 * GBitPerSecond
)

// Common bandwidths for drive and memory throughput.
//
// To count the number of units in a Bandwidth, divide:
//
//	mbps := mem.MBytePerSecond
//	fmt.Print(int64(mbps / mem.KBytePerSecond)) // prints 1000
//
// To convert an integer of units to a Bandwidth, multiply:
//
//	mbps := 10
//	fmt.Print(mem.Bandwidth(mbps)*mem.MBytePerSecond) // prints 10MB
const (
	BytePerSecond Bandwidth = 8 * BitPerSecond

	KBytePerSecond Bandwidth = 1000 * BytePerSecond
	MBytePerSecond           = 1000 * KBytePerSecond
	GBytePerSecond           = 1000 * MBytePerSecond
	TBytePerSecond           = 1000 * GBytePerSecond

	KiBytePerSecond Bandwidth = 1024 * BytePerSecond
	MiBytePerSecond           = 1024 * KiBytePerSecond
	GiBytePerSecond           = 1024 * MiBytePerSecond
	TiBytePerSecond           = 1024 * GiBytePerSecond
)

// Bandwidth represents an amount of data per second as
// int64 number of bits/s.
type Bandwidth int64

// Kilobits returns the bandwidth as floating point number
// of kilobits per second (Kbit/s).
func (b Bandwidth) Kilobits() float64 {
	return Size(b).Kilobits()
}

// Megabits returns the bandwidth as floating point number
// of megabits per second (Mbit/s).
func (b Bandwidth) Megabits() float64 {
	return Size(b).Megabits()
}

// Gigabits returns the bandwidth as floating point number
// of gigabits per second (Gbit/s).
func (b Bandwidth) Gigabits() float64 {
	return Size(b).Gigabits()
}

// Terabits returns the bandwidth as floating point number
// of terabits per second (Tbit/s).
func (b Bandwidth) Terabits() float64 {
	return Size(b).Terabits()
}

// Bytes returns the bandwidth as floating point number
// of bytes per second (B/s).
func (b Bandwidth) Bytes() float64 {
	return Size(b).Bytes()
}

// Kilobytes returns the bandwidth as floating point number
// of kilobytes per second (KB/s).
func (b Bandwidth) Kilobytes() float64 {
	return Size(b).Kilobytes()
}

// Megabytes returns the bandwidth as floating point number
// of megabytes per second (MB/s).
func (b Bandwidth) Megabytes() float64 {
	return Size(b).Megabytes()
}

// Gigabytes returns the bandwidth as floating point number
// of gigabytes per second (GB/s).
func (b Bandwidth) Gigabytes() float64 {
	return Size(b).Abs().Gigabytes()
}

// Terabytes returns the bandwidth as floating point number
// of megabytes per second (TB/s).
func (b Bandwidth) Terabytes() float64 {
	return Size(b).Terabytes()
}

// Kibibytes returns the bandwidth as floating point number
// of kibibytes per second (KiB/s).
func (b Bandwidth) Kibibytes() float64 {
	return Size(b).Kibibytes()
}

// Mebibytes returns the bandwidth as floating point number
// of mebibytes per second (MiB/s).
func (b Bandwidth) Mebibytes() float64 {
	return Size(b).Mebibytes()
}

// Gibibytes returns the bandwidth as floating point number
// of Gibibytes per second (GiB/s).
func (b Bandwidth) Gibibytes() float64 {
	return Size(b).Gibibytes()
}

// Tebibytes returns the bandwidth as floating point number
// of Tebibytes per second (TiB/s).
func (b Bandwidth) Tebibytes() float64 {
	return Size(b).Tebibytes()
}

// Truncate returns the result of rounding b towards zero to a
// multiple of m. If m <= 0, Truncate returns b unchanged.
func (b Bandwidth) Truncate(m Bandwidth) Bandwidth {
	return Bandwidth(Size(b).Truncate(Size(m)))
}

// Round returns the result of rounding b to the nearest multiple of m.
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum) value that can be
// stored in a Bandwidth, Round returns the maximum (or minimum) bandwidth.
// If m <= 0, Round returns b unchanged.
func (b Bandwidth) Round(m Bandwidth) Bandwidth {
	return Bandwidth(Size(b).Round(Size(m)))
}

// Abs returns the absolute value of b. As a special case,
// math.MinInt64 is converted to math.MaxInt64.
func (b Bandwidth) Abs() Bandwidth {
	return Bandwidth(Size(b).Abs())
}

// String returns a string representing the bandwidth in the form "1.25MB/s".
// The zero bandwidth formats as 0B/s.
func (b Bandwidth) String() string { return FormatBandwidth(b, 'D', -1) }
