package varint

import "github.com/golang-infrastructure/go-gtypes"

// ------------------------------------------------ ---------------------------------------------------------------------

// Encode 对无符号类型进行编码，注意，如果被编码的数字位数较小可能会越编码越大
func Encode[T gtypes.Unsigned](value T) []byte {
	bytes := make([]byte, 0)
	for value > 127 {
		bit := (value & 0x7F) | 0x80
		value >>= 7
		bytes = append(bytes, uint8(bit))
	}
	if value > 0 {
		bytes = append(bytes, uint8(value&0x7F))
	}
	return bytes
}

func EncodeSlice[T gtypes.Unsigned](valueSlice []T) []byte {
	slice := make([]byte, 0)
	for _, value := range valueSlice {
		slice = append(slice, Encode(value)...)
	}
	return slice
}

// ------------------------------------------------ ---------------------------------------------------------------------

// Decode 对varint编码的无符号整数进行解码，一次解码一个
func Decode[T gtypes.Unsigned](bytes []byte) T {
	var r T
	weight := 0
	for _, b := range bytes {
		r = r | (T(b&0x7F) << weight)
		weight += 7
		// 判断是否是最后一个字节，如果是最后一个字节说明当前这个数字读取完毕了
		if b&0x80 == 0 {
			break
		}
	}
	return r
}

// DecodeSlice 解码varint编码的一整个个无符号切片，即一次解码多个
func DecodeSlice[T gtypes.Unsigned](bytes []byte) []T {
	slice := make([]T, 0)
	lastIndex := 0
	for index, b := range bytes {
		if b&0x80 == 0 {
			value := Decode[T](bytes[lastIndex : index+1])
			slice = append(slice, value)
			lastIndex = index + 1
		}
	}
	return slice
}

// ------------------------------------------------ ---------------------------------------------------------------------
