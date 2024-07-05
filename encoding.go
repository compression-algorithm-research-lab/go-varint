package varint

import "github.com/golang-infrastructure/go-gtypes"

// 这个文件中存放的是比较靠底层的方法实现，上层的API都是对这些底层方法的进一步封装
// 约定：
// 1. varint编码的字节数组约定以字节的最高位的值来表示某个无符号整数对应的多个字节是否结束，
//    如果传入的要解码的字节数组没有正确设置结束标志位的话，会一直读取到字节数组结束位置都认为是当前数字的值的一部分，这可能会发生溢出

// ---------------------------------------------------------------------------------------------------------------------

const (

	// VarIntEndByteHighestBitValue 对varint字节数组而言，如果某个字节的最高位的值是0表示这是当前数字variant编码字节数组的最后一个字节
	VarIntEndByteHighestBitValue = 0x0

	// VarIntNotEndByteHighestBitValue 对varint字节数组而言，如果某个字节的最高位的值是1表示这不是当前数字varint编码字节数组的最后一个字节，后面至少还有一个字节
	VarIntNotEndByteHighestBitValue = 0x1
)

// ByteHighestBit 获取byte的最高位
const ByteHighestBit = 0x80

// 0对应的varint编码
var ZeroVarIntBytes = []byte{VarIntNotEndByteHighestBitValue << 7}

// ---------------------------------------------------------------------------------------------------------------------

// Encode 对无符号类型进行编码，注意，如果被编码的数字位数较小可能会越编码越大
func Encode[T gtypes.Unsigned](value T) []byte {
	bytes := make([]byte, 0)

	// 只要不是读取到了最后一个字节就一直读取
	for value > 127 {
		bit := (value & 0x7F) | (VarIntNotEndByteHighestBitValue << 7)
		value >>= 7
		bytes = append(bytes, uint8(bit))
	}

	// 设置结束标志位，结束标志位是强制必须设置的
	v := uint8(value&0x7F) | (VarIntEndByteHighestBitValue << 7)
	bytes = append(bytes, v)

	return bytes
}

// EncodeSlice 对无符号切片编码，切片中的无符号数字会被挨个存储在返回的字节切片中
func EncodeSlice[T gtypes.Unsigned](valueSlice []T) []byte {
	slice := make([]byte, 0)
	for _, value := range valueSlice {
		slice = append(slice, Encode(value)...)
	}
	return slice
}

// ---------------------------------------------------------------------------------------------------------------------

// Decode 对varint编码的无符号整数进行解码，一次解码一个
func Decode[T gtypes.Unsigned](bytes []byte) T {
	var r T
	weight := 0
	for _, b := range bytes {
		r = r | (T(b&0x7F) << weight)
		weight += 7
		// 判断是否是最后一个字节，如果是最后一个字节说明当前这个数字读取完毕了
		if b&ByteHighestBit == VarIntEndByteHighestBitValue {
			break
		}
	}
	return r
}

// DecodeSlice 当多个无符号整数放在同一个字节切片中时，可以调用这个方法解码
func DecodeSlice[T gtypes.Unsigned](bytes []byte) []T {
	slice := make([]T, 0)
	lastIndex := 0
	for index, b := range bytes {
		// 当读取到当前无符号整数的最后一个字节的时候，截取当前无符号整数的全部字节传进去解码
		if b&ByteHighestBit == VarIntEndByteHighestBitValue {
			value := Decode[T](bytes[lastIndex : index+1])
			slice = append(slice, value)
			lastIndex = index + 1
		}
	}

	// 说明有未读取处理完的内容，将剩下的内容一股脑儿的当做一个数字解码
	if lastIndex < len(bytes) {
		value := Decode[T](bytes[lastIndex:len(bytes)])
		slice = append(slice, value)
	}

	return slice
}

// ---------------------------------------------------------------------------------------------------------------------
