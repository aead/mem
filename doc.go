// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Package mem provides functionality for measuring and displaying
// memory throughput and capacity.
//
// # Units
//
// The fundamental unit of data is the bit. For example, networking
// speed is usually measured in bits per second. Large quantities of
// bits are commonly displayed using the SI / decimal prefixes for
// powers of 10:
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
// Sizes and bandwidths can be formatted and displayed in various
// units and with various precisions. The formats 'd/D', 'b/B' and
// 'i/I' are used for lower and uppercase decimal, binary and deciaml
// bit prefixes. For example:
//
//	d := mem.FormatSize(1*mem.MB, 'd', -1) // "1mb"
//	D := mem.FormatSize(1*mem.MB, 'D', -1) // "1MB"
//	b := mem.FormatSize(1*mem.MB, 'b', -1) // "976.5625kib"
//	B := mem.FormatSize(1*mem.MB, 'B', -1) // "976.5625KiB"
//	i := mem.FormatSize(1*mem.MB, 'i', -1) // "1mbit"
//	I := mem.FormatSize(1*mem.MB, 'I', -1) // "1Mbit"
package mem
