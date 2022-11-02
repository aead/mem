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
// Some computer metrics, like networking speed, are usually measured
// in bits while others, like storage capacity, are measured in bytes.
// Large quantities of bits / bytes are commonly displayed using the
// decimal prefixes for powers of 10. However, some systems use binary
// prefixes (2^10 = 1024).
//
//	BitSize (decimal)     Size (decimal)        Size (binary)
//	┌──────┬───────────┐  ┌──────┬───────────┐  ┌──────┬───────────┐
//	│  Bit │    1      │  │ Byte │    1      │  │ Byte │    1      │
//	│ Kbit │ 1000  Bit │  │ KB   │ 1000 Byte │  │ KiB  │ 1024 Byte │
//	│ Mbit │ 1000 Kbit │  │ MB   │ 1000 KB   │  │ MiB  │ 1024 KiB  │
//	│ Gbit │ 1000 Mbit │  │ GB   │ 1000 MB   │  │ GiB  │ 1024 MiB  │
//	│ Tbit │ 1000 Gbit │  │ TB   │ 1000 GB   │  │ TiB  │ 1024 GiB  │
//	│      │           │  │ PB   │ 1000 TB   │  │ PiB  │ 1024 TiB  │
//	└──────┴───────────┘  └──────┴───────────┘  └──────┴───────────┘
//
// # Formatting
//
// Sizes can be formatted and displayed in various units and with
// various precisions. For example:
//
//	a := mem.FormatSize(1*mem.MB, 'd', -1)      // "1mb"
//	b := mem.FormatSize(1*mem.MB, 'b', -1)      // "976.5625kib"
//	c := mem.FormatBitSize(1*mem.MBit, 'd', -1) // "1mbit"
//
// # Parsing
//
// String representation of sizes can be parsed using the corresponding
// Parse functions. For example:
//
//	a, err := mem.ParseSize("1mb")
//	b, err := mem.ParseSize("976.5625KiB")
//	c, err := mem.ParseBitSize("1mbit")
//
// # Buffers and Limits
//
// Often, code has to pre-allocate a buffer of a certain size or limit
// the amount of data from an io.Reader. With the mem package this can
// be done in a human readable way. For example:
//
//	buffer := make([]byte, 1 * mem.MB)        // Allocate a 1MB buffer
//	reader := io.LimitReader(r, 512 * mem.MB) // Limit the reader to 512MB
//
//	// Limit an HTTP request body to 5 MB.
//	req.Body = http.MaxBytesReader(w, req.Body, 5 * mem.MB)
package mem
