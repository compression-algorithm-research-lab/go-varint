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

	// case: value 0
	{
		v := Decode[uint]([]byte{0x0})
		assert.Equal(t, uint(0), v)
	}

	// case: value 1
	{
		v := Decode[uint]([]byte{0x1})
		assert.Equal(t, uint(1), v)
	}

	// case: value 2
	{
		v := Decode[uint]([]byte{0x2})
		assert.Equal(t, uint(2), v)
	}

	// case: value 127
	{
		v := Decode[uint]([]byte{0x7f})
		assert.Equal(t, uint(127), v)
	}

	// case: value 128
	{
		v := Decode[uint]([]byte{0x80, 0x1})
		assert.Equal(t, uint(128), v)
	}

	// case: value 255
	{
		v := Decode[uint]([]byte{0xff, 0x1})
		assert.Equal(t, uint(255), v)
	}

	// case: case: value 256
	{
		v := Decode[uint]([]byte{0x80, 0x2})
		assert.Equal(t, uint(256), v)
	}

	// case: 畸形编码的byte slice，没有设置结束标志位
	{
		// 0x82中的0x80表示期望后面还有数组，但是实际上是没有的
		// 不能报错，并且能够正确处理到期望的值
		v := Decode[uint]([]byte{0x80, 0x82})
		assert.Equal(t, uint(256), v)
	}

	// case: byte slice中只有一部分是varint编码，后面还有其它的内容，但是正确设置了varint的结束标志位
	{
		// 期望：只读取到第一个无符号数字，后边的bytes不管是无符号数字的varint编码还是其它内容都不会读取
		v := Decode[uint]([]byte{
			0x80, 0x2,
			0xFF, 0xFF,
			0xFF, 0xFF,
			0xFF, 0xFF,
			0xFF, 0xFF,
			0xFF, 0xFF,
			0xFF, 0xFF,
		})
		assert.Equal(t, uint(256), v)
	}

	// case: 解码的时候会发生溢出
	{
		v := Decode[uint8]([]byte{
			0xFF, 0xFF,
			0xFF, 0xFF,
			0xFF, 0xFF,
			0xFF, 0xFF,
			0xFF, 0xFF,
			0xFF, 0xFF,
			0xFF, 0xFF,
		})
		assert.Equal(t, uint8(0xff), v)
	}

}

func TestDecodeSlice(t *testing.T) {

	// case: 多个无符号整数解码，varint byte数组是合法的
	{
		slice := DecodeSlice[uint]([]byte{
			0x1,       // 1
			0x7f,      // 127
			0x0,       // 0
			0xff, 0x1, // 255
			0x80, 0x2, // 256
		})
		assert.Equal(t, []uint{1, 127, 0, 255, 256}, slice)
	}

	// case: 1个无符号整数解码，varint byte数组畸形，但是不会发生溢出
	{
		// 期望能够正常解码，不会发生报错，并且返回值是期望值
		slice := DecodeSlice[uint]([]byte{
			0x81, // 1
		})
		assert.Equal(t, []uint{1}, slice)
	}

	// case: 多个无符号整数解码，varint byte数组畸形，但是不会发生溢出
	{
		// 期望能够正常解码，不会发生报错，并且返回值是期望值
		slice := DecodeSlice[uint]([]byte{
			0x1,       // 1
			0x7f,      // 127
			0x0,       // 0
			0xff, 0x1, // 255
			0x80, 0x82, // 256
		})
		assert.Equal(t, []uint{1, 127, 0, 255, 256}, slice)
	}

	// case: varint byte数组是畸形的，会发生溢出
	{
		slice := DecodeSlice[uint8]([]byte{
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
			0xFF,
		})
		assert.Equal(t, []uint8{0xff}, slice)
	}

}

func TestEncodeChannel(t *testing.T) {

	// 测试测试数据
	inputChannel := make(chan uint64, 10)
	inputChannel <- 1
	inputChannel <- 127
	inputChannel <- 0
	inputChannel <- 255
	inputChannel <- 256
	close(inputChannel)

	// 开始对测试数据编码
	outputChannel := make(chan byte, 10000)
	EncodeChannel(inputChannel, outputChannel)

	// 读取编码结果，判断是否正确
	outputChannelBytes := make([]byte, 0)
	for b := range outputChannel {
		outputChannelBytes = append(outputChannelBytes, b)
	}
	assert.Equal(t, []byte{
		0x1,       // 1
		0x7f,      // 127
		0x0,       // 0
		0xff, 0x1, // 255
		0x80, 0x2, // 256
	}, outputChannelBytes)
}

func TestDecodeChannel(t *testing.T) {

	// 测试测试数据
	inputBytes := []byte{
		0x1,       // 1
		0x7f,      // 127
		0x0,       // 0
		0xff, 0x1, // 255
		0x80, 0x2, // 256
	}
	inputChannel := make(chan byte, len(inputBytes))
	for _, b := range inputBytes {
		inputChannel <- b
	}
	close(inputChannel)

	// 开始对测试数据解码
	outputChannel := make(chan uint64, 10000)
	DecodeChannel[uint64](inputChannel, outputChannel)

	// 读取编码结果，判断是否正确
	outputChannelValues := make([]uint64, 0)
	for value := range outputChannel {
		outputChannelValues = append(outputChannelValues, value)
	}
	assert.Equal(t, []uint64{
		1, 127, 0, 255, 256,
	}, outputChannelValues)

}
