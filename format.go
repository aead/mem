// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import (
	"errors"
	"math"
	"strconv"
)

// ParseSize parses a size string. A size string is a
// possibly signed decimal number with an optional
// fraction and a unit suffix, such as "64KB" or "1MiB".
//
// A string may be a decimal or binary size representation.
// Valid units are:
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
				unit, ok := sizeUnits[s[i:]]
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
				if m > math.MaxInt64/uint64(unit) {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}

				s := Size(m)*unit + Size(R)
				if s == math.MinInt64 {
					return math.MaxInt64, nil
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
				unit, ok := sizeUnits[s[i:]]
				if !ok {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}
				if neg {
					if m > 1<<63/uint64(unit) {
						return 0, errors.New("mem: invalid size '" + orig + "'")
					}
					return -1 * Size(m) * unit, nil
				}
				if m > math.MaxInt64/uint64(unit) {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}
				return Size(m) * unit, nil
			}
		}
	}
	return 0, errors.New("mem: invalid size '" + orig + "'")
}

// ParseBitSize parses a bit size string. A bit size string
// is a possibly signed decimal number with an optional
// fraction and a unit suffix, such as "64Kbit" or "1mbit".
//
// A string may be a decimal size representation. Valid units
// are "bit", "kbit", "mbit", "gbit" and "tbit".
func ParseBitSize(s string) (BitSize, error) {
	orig := s
	if s == "" {
		return 0, errors.New("mem: invalid bit size '" + orig + "'")
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
				unit, ok := bitsizeUnits[s[i:]]
				if !ok {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}
				R := uint64(float64(r) / float64(l) * float64(unit))

				if neg {
					if m > 1<<63/uint64(unit) {
						return 0, errors.New("mem: invalid size '" + orig + "'")
					}
					return -1 * (BitSize(m)*unit + BitSize(R)), nil
				}
				if m > math.MaxInt64/uint64(unit) {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}

				s := BitSize(m)*unit + BitSize(R)
				if s == math.MinInt64 {
					return math.MaxInt64, nil
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
				unit, ok := bitsizeUnits[s[i:]]
				if !ok {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}
				if neg {
					if m > 1<<63/uint64(unit) {
						return 0, errors.New("mem: invalid size '" + orig + "'")
					}
					return -1 * BitSize(m) * unit, nil
				}
				if m > math.MaxInt64/uint64(unit) {
					return 0, errors.New("mem: invalid size '" + orig + "'")
				}
				return BitSize(m) * unit, nil
			}
		}
	}
	return 0, errors.New("mem: invalid size '" + orig + "'")
}

// FormatSize converts the size s to a string, according to the
// format fmt and precision prec.
//
// The format fmt specifies how to format the size s. Valid values
// are:
//   - 'd' formats s as "-ddd.dddddmb" using the decimal byte units.
//   - 'b' formats s as "-ddd.dddddmib" using the binary byte units.
//
// In addition, 'D' and 'B' format s similar to 'd' and 'b' but with
// partially uppercase unit strings. In particular:
//   - 'D' formats s as "-ddd.dddddMB" using the decimal byte units.
//   - 'B' formats s as "-ddd.dddddMiB" using the binary byte units.
//
// The precision prec controls the number of digits after the decimal
// point printed by the 'd' and 'b' formats. The special precision
// -1 uses the smallest number of digits necessary such that ParseSize
// will return s exactly.
func FormatSize(s Size, fmt byte, prec int) string {
	if s == 0 { // Optimized path for the zero value
		switch fmt {
		case 'd', 'b':
			return "0b"
		case 'D', 'B':
			return "0B"
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
			return string(fmtNum(int64(s), int64(PB), prec, p))
		case s >= TB || s <= -TB:
			return string(fmtNum(int64(s), int64(TB), prec, t))
		case s >= GB || s <= -GB:
			return string(fmtNum(int64(s), int64(GB), prec, g))
		case s >= MB || s <= -MB:
			return string(fmtNum(int64(s), int64(MB), prec, m))
		case s >= KB || s <= -KB:
			return string(fmtNum(int64(s), int64(KB), prec, k))
		default:
			return string(fmtNum(int64(s), int64(Byte), prec, b))
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
			return string(fmtNum(int64(s), int64(PiB), prec, p))
		case s >= TiB || s <= -TiB:
			return string(fmtNum(int64(s), int64(TiB), prec, t))
		case s >= GiB || s <= -GiB:
			return string(fmtNum(int64(s), int64(GiB), prec, g))
		case s >= MiB || s <= -MiB:
			return string(fmtNum(int64(s), int64(MiB), prec, m))
		case s >= KiB || s <= -KiB:
			return string(fmtNum(int64(s), int64(KiB), prec, k))
		default:
			return string(fmtNum(int64(s), int64(Byte), prec, b))
		}
	default:
		return string([]byte{'%', fmt})
	}
}

// FormatBitSize converts the bit size s to a string, according to the
// format fmt and precision prec.
//
// The format fmt specifies how to format the size s. Valid values
// are:
//   - 'd' formats s as "-ddd.dddddmbit" using the decimal byte units.
//   - 'D' formats s as "-ddd.dddddMbit" using the decimal byte units.
//
// The precision prec controls the number of digits after the decimal
// point printed by the 'd' and 'D' formats. The special precision
// -1 uses the smallest number of digits necessary such that ParseBitSize
// will return s exactly.
func FormatBitSize(s BitSize, fmt byte, prec int) string {
	if s == 0 {
		switch fmt {
		case 'd':
			return "0bit"
		case 'D':
			return "0Bit"
		default:
			return string([]byte{'%', fmt})
		}
	}

	var t, g, m, k, b string
	switch fmt {
	case 'd':
		t, g, m, k, b = "tbit", "gbit", "mbit", "kbit", "bit"
	case 'D':
		t, g, m, k, b = "Tbit", "Gbit", "Mbit", "Kbit", "Bit"
	default:
		return string([]byte{'%', fmt})
	}
	switch {
	case s >= TBit || s <= -TBit:
		return string(fmtNum(int64(s), int64(TBit), prec, t))
	case s >= GBit || s <= -GBit:
		return string(fmtNum(int64(s), int64(GBit), prec, g))
	case s >= MBit || s <= -MBit:
		return string(fmtNum(int64(s), int64(MBit), prec, m))
	case s >= KBit || s <= -KBit:
		return string(fmtNum(int64(s), int64(KBit), prec, k))
	default:
		return string(fmtNum(int64(s), int64(Bit), prec, b))
	}
}

func fmtNum(v, base int64, prec int, unit string) []byte {
	m := v / base
	r := v % base

	// Usually a formatted string consists of
	// a potential minus sign, at most three
	// digits and any precision digits followed
	// by the unit.
	// For example: -999Tbit or 512.125GiB
	//
	// We don't optimize for edge cases
	// like: 4096000Tbit
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
		buf = strconv.AppendInt(buf, m, 10)
	case r == 0:
		buf = make([]byte, 0, 4+prec+1+len(unit))
		buf = strconv.AppendInt(buf, m, 10)
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
		buf = strconv.AppendInt(buf, m, 10)
		buf = append(buf, rbuf[1:]...)
	}
	buf = append(buf, unit...)
	buf = buf[:len(buf):len(buf)]
	return buf
}

var sizeUnits = map[string]Size{
	"b": Byte, "B": Byte,

	"kb": KB, "KB": KB,
	"mb": MB, "MB": MB,
	"gb": GB, "GB": GB,
	"tb": TB, "TB": TB,
	"pb": PB, "PB": PB,

	"kib": KiB, "KiB": KiB,
	"mib": MiB, "MiB": MiB,
	"gib": GiB, "GiB": GiB,
	"tib": TiB, "TiB": TiB,
	"pib": PiB, "PiB": PiB,
}

var bitsizeUnits = map[string]BitSize{
	"bit": Bit, "Bit": Bit,
	"kbit": KBit, "Kbit": KBit,
	"mbit": MBit, "Mbit": MBit,
	"gbit": GBit, "Gbit": GBit,
	"tbit": TBit, "Tbit": TBit,
}
