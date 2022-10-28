// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem_test

import (
	"fmt"
	"log"

	"aead.dev/mem"
)

func ExampleParseBandwidth() {
	a, err := mem.ParseBandwidth("1.123MB/s")
	if err != nil {
		log.Fatalln(err)
	}
	b, err := mem.ParseBandwidth("3.877MB/s")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(a + b)
	// Output:
	// 5MB/s
}

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

func ExampleFormatBandwidth() {
	fmt.Println(mem.FormatBandwidth(1*mem.MBytePerSecond, 'D', -1))
	fmt.Println(mem.FormatBandwidth(1*mem.MBytePerSecond+111*mem.KBytePerSecond, 'd', 2))
	fmt.Println(mem.FormatBandwidth(2*mem.TiBytePerSecond+512*mem.MiBytePerSecond, 'B', 4))

	fmt.Println(mem.FormatBandwidth(5*mem.MBitPerSecond, 'I', -1))
	fmt.Println(mem.FormatBandwidth(200*mem.MBitPerSecond, 'D', -1))
	// Output:
	// 1MB/s
	// 1.11mb/s
	// 2.0005TiB/s
	// 5Mbit/s
	// 25MB/s
}

func ExampleFormatSize() {
	fmt.Println(mem.FormatSize(1*mem.MB, 'D', -1))
	fmt.Println(mem.FormatSize(1*mem.MB+111*mem.KB, 'd', 2))
	fmt.Println(mem.FormatSize(2*mem.TiB+512*mem.MiB, 'B', 4))

	fmt.Println(mem.FormatSize(5*mem.MBit, 'I', -1))
	fmt.Println(mem.FormatSize(200*mem.MBit, 'D', -1))
	// Output:
	// 1MB
	// 1.11mb
	// 2.0005TiB
	// 5Mbit
	// 25MB
}

func ExampleSize_PerSecond() {
	size := 1 * mem.MB
	fmt.Println(size.PerSecond())
	// Output:
	// 1MB/s
}

func ExampleSize_String() {
	fmt.Println(1 * mem.MB)
	fmt.Println(1*mem.GB + 500*mem.MB)
	fmt.Println(5*mem.KiB + 880*mem.Byte)
	fmt.Println(40 * mem.GBit)
	// Output:
	// 1MB
	// 1.5GB
	// 6KB
	// 5GB
}

func ExampleBandwidth_String() {
	fmt.Println(1 * mem.MBytePerSecond)
	fmt.Println(1*mem.GBytePerSecond + 500*mem.MBytePerSecond)
	fmt.Println(5*mem.KiBytePerSecond + 880*mem.BytePerSecond)
	fmt.Println(40 * mem.GBitPerSecond)
	// Output:
	// 1MB/s
	// 1.5GB/s
	// 6KB/s
	// 5GB/s
}
