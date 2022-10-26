// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import (
	"errors"
	"strconv"
)

// ParseSize parses a size string. A size string is a
// possibly signed decimal number with an optional
// fraction and a unit suffix, such as "64KB", "1Mbit"
// or "-1.05GiB".
//
// A string may be a bit, decimal or binary size representation.
// Valid units are:
//   - bit:     "bits", "kbits", "mbits", "gbits" and "tbits"
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

// FormatSize converts the size s to a string, according to the
// format fmt and precision prec.
//
// The format fmt specifies how to format the size s. Valid values
// are:
//   - 'd' formats s as "-ddd.dddddMB" using the decimal byte units.
//   - 'b' formats s as "-ddd.dddddMiB" using the binary byte units.
//   - 'i' formats s as "-ddd.dddddMbit" using the decimal bit units.
//
// The precision prec controls the number of digits after the decimal
// point printed by the 'd', 'b' and 'i' formats. The special precision
// -1 uses the smallest number of digits necessary such that ParseSize
// will return s exactly.
func FormatSize(s Size, fmt byte, prec int) string {
	switch fmt {
	case 'd':
		switch {
		case s >= PB || s <= -PB:
			return string(fmtSize(s, PB, prec, "PB"))
		case s >= TB || s <= -TB:
			return string(fmtSize(s, TB, prec, "TB"))
		case s >= GB || s <= -GB:
			return string(fmtSize(s, GB, prec, "GB"))
		case s >= MB || s <= -MB:
			return string(fmtSize(s, MB, prec, "MB"))
		case s >= KB || s <= -KB:
			return string(fmtSize(s, KB, prec, "KB"))
		case s == 0:
			return "0B"
		default:
			return strconv.FormatFloat(s.Bytes(), 'f', prec, 64) + "B"
		}
	case 'b':
		switch {
		case s >= PiB || s <= -PiB:
			return string(fmtSize(s, PiB, prec, "PiB"))
		case s >= TiB || s <= -TiB:
			return string(fmtSize(s, TiB, prec, "TiB"))
		case s >= GiB || s <= -GiB:
			return string(fmtSize(s, GiB, prec, "GiB"))
		case s >= MiB || s <= -MiB:
			return string(fmtSize(s, MiB, prec, "MiB"))
		case s >= KiB || s <= -KiB:
			return string(fmtSize(s, KiB, prec, "KiB"))
		case s == 0:
			return "0B"
		default:
			return strconv.FormatFloat(s.Bytes(), 'f', prec, 64) + "B"
		}
	case 'i':
		switch {
		case s >= TBit || s <= -TBit:
			return string(fmtSize(s, TBit, prec, "Tbit"))
		case s >= GBit || s <= -GBit:
			return string(fmtSize(s, GBit, prec, "Gbit"))
		case s >= MBit || s <= -MBit:
			return string(fmtSize(s, MBit, prec, "Mbit"))
		case s >= KBit || s <= -KBit:
			return string(fmtSize(s, KBit, prec, "Kbit"))
		case s == 0:
			return "0bit"
		default:
			return strconv.FormatInt(int64(s), 10) + "bit"
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
		case "bit":
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
