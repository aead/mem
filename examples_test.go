// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem_test

import (
	"fmt"
	"log"

	"aead.dev/mem"
)

func ExampleParseSize() {
	a, err := mem.ParseSize("1.123MB")
	if err != nil {
		log.Fatalln(err)
	}
	b, err := mem.ParseSize("3.877MB")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(a + b)
	// Output:
	// 5MB
}

func ExampleFormatSize() {
	fmt.Println(mem.FormatSize(1*mem.MB, 'D', -1))
	fmt.Println(mem.FormatSize(1*mem.MB+111*mem.KB, 'd', 2))
	fmt.Println(mem.FormatSize(2*mem.TiB+512*mem.MiB, 'B', 4))
	// Output:
	// 1MB
	// 1.11mb
	// 2.0005TiB
}

func ExampleSize_String() {
	fmt.Println(1 * mem.MB)
	fmt.Println(1*mem.GB + 500*mem.MB)
	fmt.Println(5*mem.KiB + 880*mem.Byte)
	// Output:
	// 1MB
	// 1.5GB
	// 6KB
}

func ExampleFormatBitSize() {
	fmt.Println(mem.FormatBitSize(1*mem.MBit, 'D', -1))
	fmt.Println(mem.FormatBitSize(1*mem.MBit+111*mem.KBit, 'd', 2))
	fmt.Println(mem.FormatBitSize(2*mem.TBit+512*mem.MBit, 'D', 4))
	// Output:
	// 1Mbit
	// 1.11mbit
	// 2.0005Tbit
}

func ExampleParseBitSize() {
	a, err := mem.ParseBitSize("1.123Mbit")
	if err != nil {
		log.Fatalln(err)
	}
	b, err := mem.ParseBitSize("3.877Mbit")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(a + b)
	// Output:
	// 5Mbit
}

func ExampleBitSize_Bytes() {
	b := mem.MBit
	fmt.Println(b.Bytes())

	b = 1*mem.MBit + 4*mem.Bit
	bytes, bits := b.Bytes()

	fmt.Println(b == bytes.Bits()+bits && -7 <= bits && bits <= 7)
	// Output:
	// 125KB 0Bit
	// true
}

func ExampleBitSize_String() {
	fmt.Println(1 * mem.MBit)
	fmt.Println(1*mem.GBit + 500*mem.MBit)
	fmt.Println(5*mem.KBit + 880*mem.Bit)
	// Output:
	// 1Mbit
	// 1.5Gbit
	// 5.88Kbit
}
