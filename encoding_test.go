package varint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {

	// case: value 0
	{
		v := Encode[uint64](0)
		assert.Equal(t, []byte{0x0}, v)
	}

	// case: value 1
	{
		v := Encode[uint64](1)
		assert.Equal(t, []byte{0x1}, v)
	}

	// case: value 2
	{
		v := Encode[uint64](2)
		assert.Equal(t, []byte{0x2}, v)
	}

	// case: value 127
	{
		v := Encode[uint64](127)
		assert.Equal(t, []byte{0x7f}, v)
	}

	// case: value 128
	{
		v := Encode[uint64](128)
		assert.Equal(t, []byte{0x80, 0x1}, v)
	}

	// case: value 255
	{
		v := Encode[uint64](255)
		assert.Equal(t, []byte{0xff, 0x1}, v)
	}

	// case: case: value 256
	{
		v := Encode[uint64](256)
		assert.Equal(t, []byte{0x80, 0x2}, v)
	}

}

func TestEncodeSlice(t *testing.T) {
	slice := EncodeSlice([]uint64{1, 127, 0, 255, 256})
	assert.Equal(t, []byte{
		0x1,       // 1
		0x7f,      // 127
		0x0,       // 0
		0xff, 0x1, // 255
		0x80, 0x2, // 256
	}, slice)
}

func TestDecode(t *testing.T) {
	var a uint64
	a = 123456
	bytes := Encode(a)
	t.Log(bytes)
	t.Log(Decode[uint64](bytes))
}

func TestDecodeSlice(t *testing.T) {
	ints := []uint64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	bytes := EncodeSlice(ints)
	t.Log(len(bytes))
	t.Log(bytes)
	slice := DecodeSlice[uint64](bytes)
	t.Log(slice)
}
