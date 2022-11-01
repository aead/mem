// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Package mem provides functionality for measuring and displaying
// memory throughput and capacity.
//
// # Units
//
// Package mem defines two different types that represent quantities
// of data - Size and BitSize. The former represents an amount of
// data as bytes while the later represents an amount of data as
// bits.
//
// Not each BitSize can be represented as Size and vice versa. For
// example 65 bit is neither 8 bytes (64 bit) nor 9 bytes (72 bit).
// However, Size and BitSize provide APIs for converting one to the
// other.
//
// Some computer metrics - like networking speed - are usually
// measured in bits. Large quantities of bits are commonly displayed
// using the SI / decimal prefixes for powers of 10:
//
//	Unit |    Amount
//	----------------
//	 Bit |    1
//	Kbit | 1000 Bit
//	Mbit | 1000 Kbit
//	Gbit | 1000 Mbit
//	Tbit | 1000 Gbit
//
// In contrast, storage capacity is usually measured in bytes.
// For large quantities of bytes there are two commonly used
// prefix scales - the SI / decimal prefixes for powers of 10
// and the binary prefixes for powers of 2:
//
//	Unit (decimal) |    Amount      Unit (binary) |    Amount
//	--------------------------      -------------------------
//	          Byte |    8 Bit                Byte |    8 Bit
//	            KB | 1000 Byte                KiB | 1024 Byte
//	            MB | 1000 KB                  MiB | 1024 KiB
//	            GB | 1000 MB                  GiB | 1024 MiB
//	            TB | 1000 GB                  TiB | 1024 GiB
//	            PB | 1000 TB                  PiB | 1024 TiB
//
// Most software and operating systems, like macOS or linux report
// file sizes in decimal units. This is consistent with other units
// for computers, like CPU clock or networking speeds.
// One prominent example that uses the binary instead of decimal
// units is MS Windows.
//
// # Formatting
//
// Sizes can be formatted and displayed in various units and with
// various precisions. The formats 'd/D' and 'b/B' are used for
// lower and uppercase decimal and binary prefixes.
// For example:
//
// a := mem.FormatSize(1*mem.MB, 'd', -1)      // "1mb"
// b := mem.FormatSize(1*mem.MB, 'D', -1)      // "1MB"
// c := mem.FormatSize(1*mem.MB, 'b', -1)      // "976.5625kib"
// d := mem.FormatSize(1*mem.MB, 'B', -1)      // "976.5625KiB"
// e := mem.FormatBitSize(1*mem.MBit, 'd', -1) // 1mbit
// f := mem.FormatBitSize(1*mem.MBit, 'D', -1) // 1Mbit
//
// # Parsing
//
// A string can be parsed as Size or BitSize using the
// ParseSize or ParseBitSize functions. For example:
//
//	a, err := mem.ParseSize("1mb")
//	b, err := mem.ParseSize("976.5625KiB")
//	c, err := mem.ParseBitSize("1mbit")
//	d, err := mem.ParseBitSize("25.473GBit")
package mem
