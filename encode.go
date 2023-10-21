package varint

import "github.com/golang-infrastructure/go-gtypes"

// ------------------------------------------------ ---------------------------------------------------------------------

// Encode 对有符号整数类型进行编码，注意，是可能会越编码越大的情况的，所以应该根据自己的情况选择
func Encode[T gtypes.Integer](integer T) VarInt {

}

// EncodeForSigned 将有符号数字编码为varint
func EncodeForSigned[T gtypes.Signed](integer T) VarInt {
	sign := 0
	if integer < 0 {
		// 设置符号位，0表示是正数，1表示是负数
		sign = 1
		integer *= -1
	}

	bytes := make([]byte, 0)
	for integer > 0 {
		// 获取value的最低7位的值作为一个字节，并且将这个字节的最高位置为1
		bytePart := uint8(integer&0b0111_1111) | uint8(0b1000_0000)
		bytes = append(bytes, uint8(bytePart))
		// 已经处理过的部分移除掉
		integer >>= 7
	}

	// 0作为一个特殊值，返回一个长度为1的字节数组，而不是空数组
	if len(bytes) == 0 {
		bytes = append(bytes, uint8(0))
		return bytes
	}

	// 处理最后一个字节，同时把符号位设置上
	lastByte := bytes[len(bytes)-1]
	if lastByte&0b0100_0000 == 0 {
		// 说明倒数第二位是可以被使用的，则直接使用这一位皆可以了，把它置位符号位
		bytes[len(bytes)-1] = uint8(lastByte | uint8(sign<<6) | 0b1000_0000)
	} else {
		// 说明倒数第二位已经被占用了，则再追加一个字节来描述
		// 先将现在的最后一个字节的最高位修改为1表示未结束
		bytes[len(bytes)-1] = uint8(bytes[len(bytes)-1] | 0b1000_0000)
		// 然后再追加一个字节专门用来表示符号位
		bytes = append(bytes, uint8(sign<<6))
	}

	return bytes
}

// EncodeForUnsigned 将无符号类型编码为variant
func EncodeForUnsigned[T gtypes.Unsigned](integer T) VarInt {
	
}

// EncodeIntSliceToVarIntSlice 把int切片编码为varint类型的切片
func EncodeIntSliceToVarIntSlice[T gtypes.Integer](integerSlice []T) []VarInt {
	varIntSlice := make([]VarInt, 0)
	for _, integer := range integerSlice {
		varIntSlice = append(varIntSlice, Encode(integer))
	}
	return varIntSlice
}

// EncodeIntSliceToVarIntBytes 把int切片编码为varint的字节数组，同一个字节数组中会紧凑顺序存放着多个VarInt的字节
func EncodeIntSliceToVarIntBytes[T gtypes.Integer](integerSlice []T) []byte {
	bytes := make([]byte, 0)
	for _, value := range integerSlice {
		bytes = append(bytes, Encode(value)...)
	}
	return bytes
}
