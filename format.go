// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import (
	"errors"
	"strconv"
	"strings"
)

// ParseBandwidth parses a bandwidth string. A bandwidth string
// is a possibly signed decimal number with an optional
// fraction and a unit suffix, such as "64KB/s", "1Mbit/s"
// or "-1.05GiB/s".
//
// A string may be a bit, decimal or binary bandwidth representation.
// Valid units are:
//   - bit:     "bit/s", "kbit/s", "mbit/s", "gbit/s" and "tbit/s"
//   - decimal: "b/s", "kb/s", "mb/s", "gb/s", "tb/s", "pb/s"
//   - binary:  "b/s", "kib/s", "mib/s", "gib/s", "tib/s", "pib/s"
func ParseBandwidth(s string) (Bandwidth, error) {
	orig := s
	if !strings.HasSuffix(s, "/s") {
		return 0, errors.New("mem: invalid bandwidth '" + orig + "'")
	}
	size, err := ParseSize(s[:len(s)-2])
	if err != nil {
		return 0, errors.New("mem: invalid bandwidth '" + orig + "'")
	}
	return size.PerSecond(), nil
}

// ParseSize parses a size string. A size string is a
// possibly signed decimal number with an optional
// fraction and a unit suffix, such as "64KB", "1Mbit"
// or "-1.05GiB".
//
// A string may be a bit, decimal or binary size representation.
// Valid units are:
//   - bit:     "bit", "kbit", "mbit", "gbit" and "tbit"
//   - decimal: "b", "kb", "mb", "gb", "tb", "pb"
//   - binary:  "b", "kib", "mib", "gib", "tib", "pib"
func ParseSize(s string) (Size, error) {
	orig := s
	if s == "" {
		return 0, errors.New("mem: invalid size '" + orig + "'")
	}

	var neg bool
	if c := s[0]; c == '+' || c == '-' {
		neg = c == '-'
		s = s[1:]
	}

	var dot bool
	var m, r uint64
	var l uint64 = 1
	for i, c := range s {
		if dot {
			switch {
			case c >= '0' && c <= '9':
				r = r*10 + uint64(c-'0')
				l *= 10
			default:
				unit, ok := parseUnit(s[i:])
				if !ok {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}
				R := uint64(float64(r) / float64(l) * float64(unit))
				if neg {
					if m > 1<<63/uint64(unit) {
						return 0, errors.New("mem: invalid size '" + orig + "'")
					}
					return -1 * (Size(m)*unit + Size(R)), nil
				}
				if m > uint64(MaxSize)/uint64(unit) {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}

				s := Size(m)*unit + Size(R)
				if s == minSize {
					return MaxSize, nil
				}
				return s, nil
			}
		} else {
			switch {
			case c >= '0' && c <= '9':
				m = m*10 + uint64(c-'0')
			case c == '.':
				dot = true
			default:
				if i == 0 {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}
				unit, ok := parseUnit(s[i:])
				if !ok {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}
				if neg {
					if m > 1<<63/uint64(unit) {
						return 0, errors.New("mem: invalid size '" + orig + "'")
					}
					return -1 * Size(m) * unit, nil
				}
				if m > uint64(MaxSize)/uint64(unit) {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}
				return Size(m) * unit, nil
			}
		}
	}
	return 0, errors.New("mem: invalid size '" + orig + "'")
}

// FormatBandwidth converts the bandwidth b to a string, according to
// the format fmt and precision prec.
//
// The format fmt specifies how to format the bandwidth b. Valid values
// are:
//   - 'd' formats s as "-ddd.dddddmb/s" using the decimal byte units.
//   - 'b' formats s as "-ddd.dddddmib/s" using the binary byte units.
//   - 'i' formats s as "-ddd.dddddmbit/s" using the decimal bit units.
//
// In addition, 'D', 'B', 'I' format b similar to 'd', 'b' and 'i'
// but with partially uppercase unit strings. In particular:
//   - 'D' formats s as "-ddd.dddddMB/s" using the decimal byte units.
//   - 'B' formats s as "-ddd.dddddMiB/s" using the binary byte units.
//   - 'I' formats s as "-ddd.dddddMbit/s" using the decimal bit units.
//
// The precision prec controls the number of digits after the decimal
// point printed by the 'd', 'b' and 'i' formats. The special precision
// -1 uses the smallest number of digits necessary such that
// ParseBandwidth will return b exactly.
func FormatBandwidth(b Bandwidth, fmt byte, prec int) string {
	return FormatSize(Size(b), fmt, prec) + "/s"
}

// FormatSize converts the size s to a string, according to the
// format fmt and precision prec.
//
// The format fmt specifies how to format the size s. Valid values
// are:
//   - 'd' formats s as "-ddd.dddddmb" using the decimal byte units.
//   - 'b' formats s as "-ddd.dddddmib" using the binary byte units.
//   - 'i' formats s as "-ddd.dddddmbit" using the decimal bit units.
//
// In addition, 'D', 'B', 'I' format s similar to 'd', 'b' and 'i'
// but with partially uppercase unit strings. In particular:
//   - 'D' formats s as "-ddd.dddddMB" using the decimal byte units.
//   - 'B' formats s as "-ddd.dddddMiB" using the binary byte units.
//   - 'I' formats s as "-ddd.dddddMbit" using the decimal bit units.
//
// The precision prec controls the number of digits after the decimal
// point printed by the 'd', 'b' and 'i' formats. The special precision
// -1 uses the smallest number of digits necessary such that ParseSize
// will return s exactly.
func FormatSize(s Size, fmt byte, prec int) string {
	if s == 0 { // Optimized path for the zero value
		switch fmt {
		case 'd', 'b':
			return "0b"
		case 'D', 'B':
			return "0B"
		case 'i':
			return "0bit"
		case 'I':
			return "0Bit"
		default:
			return string([]byte{'%', fmt})
		}
	}

	switch fmt {
	case 'd', 'D':
		var p, t, g, m, k, b string
		if fmt == 'D' {
			p, t, g, m, k, b = "PB", "TB", "GB", "MB", "KB", "B"
		} else {
			p, t, g, m, k, b = "pb", "tb", "gb", "mb", "kb", "b"
		}
		switch {
		case s >= PB || s <= -PB:
			return string(fmtSize(s, PB, prec, p))
		case s >= TB || s <= -TB:
			return string(fmtSize(s, TB, prec, t))
		case s >= GB || s <= -GB:
			return string(fmtSize(s, GB, prec, g))
		case s >= MB || s <= -MB:
			return string(fmtSize(s, MB, prec, m))
		case s >= KB || s <= -KB:
			return string(fmtSize(s, KB, prec, k))
		default:
			return string(fmtSize(s, Byte, prec, b))
		}
	case 'b', 'B':
		var p, t, g, m, k, b string
		if fmt == 'B' {
			p, t, g, m, k, b = "PiB", "TiB", "GiB", "MiB", "KiB", "B"
		} else {
			p, t, g, m, k, b = "pib", "tib", "gib", "mib", "kib", "b"
		}
		switch {
		case s >= PiB || s <= -PiB:
			return string(fmtSize(s, PiB, prec, p))
		case s >= TiB || s <= -TiB:
			return string(fmtSize(s, TiB, prec, t))
		case s >= GiB || s <= -GiB:
			return string(fmtSize(s, GiB, prec, g))
		case s >= MiB || s <= -MiB:
			return string(fmtSize(s, MiB, prec, m))
		case s >= KiB || s <= -KiB:
			return string(fmtSize(s, KiB, prec, k))
		default:
			return string(fmtSize(s, Byte, prec, b))
		}
	case 'i', 'I':
		var t, g, m, k, b string
		if fmt == 'I' {
			t, g, m, k, b = "Tbit", "Gbit", "Mbit", "Kbit", "Bit"
		} else {
			t, g, m, k, b = "tbit", "gbit", "mbit", "kbit", "bit"
		}
		switch {
		case s >= TBit || s <= -TBit:
			return string(fmtSize(s, TBit, prec, t))
		case s >= GBit || s <= -GBit:
			return string(fmtSize(s, GBit, prec, g))
		case s >= MBit || s <= -MBit:
			return string(fmtSize(s, MBit, prec, m))
		case s >= KBit || s <= -KBit:
			return string(fmtSize(s, KBit, prec, k))
		default:
			return strconv.FormatInt(int64(s), 10) + b
		}
	default:
		return string([]byte{'%', fmt})
	}
}

func fmtSize(v, base Size, prec int, unit string) []byte {
	m := v / base
	r := v % base

	// Usually a formatted string consists of
	// a potential minus sign, at most three
	// digits and any precision digits followed
	// by the unit.
	// For example: -999Tbit or 512.125Gib
	//
	// We don't optimize for edge cases
	// like 512PB as Tbit: 4096000Tbit
	//
	// Hence, we allocate an 4+prec+1+unit
	// (+1 for the '.') buffer when r == 0.
	//
	// Otherwise, we format r first with the
	// requested precision: 0.xyz...
	// Then we concat and return the m and r
	// buffers as: abc.xyz

	var buf []byte
	switch {
	case r == 0 && prec <= 0:
		buf = make([]byte, 0, 4+len(unit))
		buf = strconv.AppendInt(buf, int64(m), 10)
	case r == 0:
		buf = make([]byte, 0, 4+prec+1+len(unit))
		buf = strconv.AppendInt(buf, int64(m), 10)
		buf = append(buf, '.')
		for prec > 0 {
			buf = append(buf, '0')
			prec--
		}
	default:
		if r < 0 {
			r *= -1
		}
		rbuf := strconv.AppendFloat(nil, float64(r)/float64(base), 'f', prec, 64)
		buf = make([]byte, 0, 4+prec+len(rbuf)+len(unit)) // The '.' is already included in rbuf
		if v < 0 && m == 0 {
			// When formatting a negative size like -4bit as -0.5b
			// where the abs. value is small than the base,
			// m = v / base will be zero.
			// In this case, we have to add the minus sign manually
			// since strconv.AppendInt(buf, 0, 10) will not add it.
			buf = append(buf, '-')
		}
		buf = strconv.AppendInt(buf, int64(m), 10)
		buf = append(buf, rbuf[1:]...)
	}
	buf = append(buf, unit...)
	buf = buf[:len(buf):len(buf)]
	return buf
}

func parseUnit(s string) (Size, bool) {
	if len(s) == 0 {
		return 0, false
	}
	switch s[0] {
	case 'b', 'B':
		switch s {
		case "b", "B":
			return Byte, true
		case "bit", "Bit":
			return Bit, true
		}
	case 'k', 'K':
		switch s {
		case "kb", "KB":
			return KB, true
		case "kib", "KiB":
			return KiB, true
		case "kbit", "Kbit":
			return KBit, true
		}
	case 'm', 'M':
		switch s {
		case "mb", "MB":
			return MB, true
		case "mib", "MiB":
			return MiB, true
		case "mbit", "Mbit":
			return MBit, true
		}
	case 'g', 'G':
		switch s {
		case "gb", "GB":
			return GB, true
		case "gib", "GiB":
			return GiB, true
		case "gbit", "Gbit":
			return GBit, true
		}
	case 't', 'T':
		switch s {
		case "tb", "TB":
			return TB, true
		case "tib", "TiB":
			return TiB, true
		case "tbit", "Tbit":
			return TBit, true
		}
	case 'p', 'P':
		switch s {
		case "pb", "PB":
			return PB, true
		case "pib", "PiB":
			return PiB, true
		}
	}
	return 0, false
}
