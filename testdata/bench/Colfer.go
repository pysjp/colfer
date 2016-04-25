package bench

// This file was generated by colf(1); DO NOT EDIT

import (
	"fmt"
	"io"
	"math"
	"time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = math.E
var _ = time.RFC3339

// ColferContinue signals a data continuation as a byte index.
type ColferContinue int

func (i ColferContinue) Error() string {
	return fmt.Sprintf("colfer: data continuation at byte %d", i)
}

// ColferError signals a data mismatch as as a byte index.
type ColferError int

func (i ColferError) Error() string {
	return fmt.Sprintf("colfer: unknown header at byte %d", i)
}

type Colfer struct {
	Key	int64
	Host	string
	Addr	[]byte
	Port	int32
	Size	int64
	Hash	uint64
	Ratio	float64
	Route	bool
}

// MarshalTo encodes o as Colfer into buf and returns the number of bytes written.
// If the buffer is too small, MarshalTo will panic.
func (o *Colfer) MarshalTo(buf []byte) int {
	if o == nil {
		return 0
	}

	var i int

	if v := o.Key; v != 0 {
		x := uint64(v)
		if v >= 0 {
			buf[i] = 0
		} else {
			x = ^x + 1
			buf[i] = 0 | 0x80
		}
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if v := o.Host; len(v) != 0 {
		buf[i] = 1
		i++
		x := uint(len(v))
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		copy(buf[i:], v)
		i += len(v)
	}

	if v := o.Addr; len(v) != 0 {
		buf[i] = 2
		i++
		x := uint(len(v))
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		copy(buf[i:], v)
		i += len(v)
	}

	if v := o.Port; v != 0 {
		x := uint32(v)
		if v >= 0 {
			buf[i] = 3
		} else {
			x = ^x + 1
			buf[i] = 3 | 0x80
		}
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if v := o.Size; v != 0 {
		x := uint64(v)
		if v >= 0 {
			buf[i] = 4
		} else {
			x = ^x + 1
			buf[i] = 4 | 0x80
		}
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if x := o.Hash; x != 0 {
		buf[i] = 5
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if v := o.Ratio; v != 0.0 {
		buf[i] = 6
		x := math.Float64bits(v)
		buf[i+1], buf[i+2], buf[i+3], buf[i+4] = byte(x>>56), byte(x>>48), byte(x>>40), byte(x>>32)
		buf[i+5], buf[i+6], buf[i+7], buf[i+8] = byte(x>>24), byte(x>>16), byte(x>>8), byte(x)
		i += 9
	}

	if o.Route {
		buf[i] = 7
		i++
	}

	buf[i] = 0x7f
	i++
	return i
}

// MarshalLen returns the Colfer serial byte size.
func (o *Colfer) MarshalLen() int {
	if o == nil {
		return 0
	}

	l := 1

	if v := o.Key; v != 0 {
		x := uint64(v)
		if v < 0 {
			x = ^x + 1
		}
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
	}

	if x := len(o.Host); x != 0 {
		l += x
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
	}

	if x := len(o.Addr); x != 0 {
		l += x
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
	}

	if v := o.Port; v != 0 {
		x := uint32(v)
		if v < 0 {
			x = ^x + 1
		}
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
	}

	if v := o.Size; v != 0 {
		x := uint64(v)
		if v < 0 {
			x = ^x + 1
		}
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
	}

	if x := o.Hash; x != 0 {
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
	}

	if o.Ratio != 0.0 {
		l += 9
	}

	if o.Route {
		l++
	}

	return l
}

// MarshalBinary encodes o as Colfer conform encoding.BinaryMarshaler.
// The error return is always nil.
func (o *Colfer) MarshalBinary() (data []byte, err error) {
	data = make([]byte, o.MarshalLen())
	o.MarshalTo(data)
	return data, nil
}

// UnmarshalBinary decodes data as Colfer conform encoding.BinaryUnmarshaler.
// The error return options are io.EOF, testdata/bench.ColferError, and testdata/bench.ColferContinue.
func (o *Colfer) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return io.EOF
	}

	header := data[0]
	i := 1

	if header == 0 || header == 0|0x80 {
		var x uint64
		for shift := uint(0); ; shift += 7 {
			if i == len(data) {
				return io.EOF
			}
			b := data[i]
			i++
			if shift == 63 {
				x |= 1 << 63
				break
			}
			x |= (uint64(b) & 0x7f) << shift
			if b < 0x80 {
				break
			}
		}
		if header&0x80 != 0 {
			x = ^x + 1
		}
		o.Key = int64(x)

		if i == len(data) {
			return io.EOF
		}
		header = data[i]
		i++
	}

	if header == 1 {
		var x uint32
		for shift := uint(0); ; shift += 7 {
			if i == len(data) {
				return io.EOF
			}
			b := data[i]
			i++
			if shift == 28 {
				x |= uint32(b) << 28
				break
			}
			x |= (uint32(b) & 0x7f) << shift
			if b < 0x80 {
				break
			}
		}

		to := i + int(x)
		if to >= len(data) {
			return io.EOF
		}
		o.Host = string(data[i:to])

		header = data[to]
		i = to + 1
	}

	if header == 2 {
		var x uint32
		for shift := uint(0); ; shift += 7 {
			if i == len(data) {
				return io.EOF
			}
			b := data[i]
			i++
			if shift == 28 {
				x |= uint32(b) << 28
				break
			}
			x |= (uint32(b) & 0x7f) << shift
			if b < 0x80 {
				break
			}
		}

		length := int(x)
		to := i + length
		if to >= len(data) {
			return io.EOF
		}
		v := make([]byte, length)
		copy(v, data[i:])
		o.Addr = v

		header = data[to]
		i = to + 1
	}

	if header == 3 || header == 3|0x80 {
		var x uint32
		for shift := uint(0); ; shift += 7 {
			if i == len(data) {
				return io.EOF
			}
			b := data[i]
			i++
			if shift == 28 {
				x |= uint32(b) << 28
				break
			}
			x |= (uint32(b) & 0x7f) << shift
			if b < 0x80 {
				break
			}
		}
		if header&0x80 != 0 {
			x = ^x + 1
		}
		o.Port = int32(x)

		if i == len(data) {
			return io.EOF
		}
		header = data[i]
		i++
	}

	if header == 4 || header == 4|0x80 {
		var x uint64
		for shift := uint(0); ; shift += 7 {
			if i == len(data) {
				return io.EOF
			}
			b := data[i]
			i++
			if shift == 63 {
				x |= 1 << 63
				break
			}
			x |= (uint64(b) & 0x7f) << shift
			if b < 0x80 {
				break
			}
		}
		if header&0x80 != 0 {
			x = ^x + 1
		}
		o.Size = int64(x)

		if i == len(data) {
			return io.EOF
		}
		header = data[i]
		i++
	}

	if header == 5 {
		var x uint64
		for shift := uint(0); ; shift += 7 {
			if i == len(data) {
				return io.EOF
			}
			b := data[i]
			i++
			if shift == 63 {
				x |= 1 << 63
				break
			}
			x |= (uint64(b) & 0x7f) << shift
			if b < 0x80 {
				break
			}
		}
		o.Hash = x

		if i == len(data) {
			return io.EOF
		}
		header = data[i]
		i++
	}

	if header == 6 {
		if i+8 >= len(data) {
			return io.EOF
		}
		x := uint64(data[i])<<56 | uint64(data[i+1])<<48 | uint64(data[i+2])<<40 | uint64(data[i+3])<<32
		x |= uint64(data[i+4])<<24 | uint64(data[i+5])<<16 | uint64(data[i+6])<<8 | uint64(data[i+7])
		o.Ratio = math.Float64frombits(x)

		header = data[i+8]
		i += 9
	}

	if header == 7 {
		o.Route = true

		if i == len(data) {
			return io.EOF
		}
		header = data[i]
		i++
	}

	if header != 0x7f {
		return ColferError(i - 1)
	}
	if i != len(data) {
		return ColferContinue(i)
	}
	return nil
}
