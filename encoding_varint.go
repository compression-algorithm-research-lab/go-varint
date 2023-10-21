package varint

import "github.com/golang-infrastructure/go-gtypes"


// ------------------------------------------------ ---------------------------------------------------------------------

// DecodeFromVarInt 对varint编码的整数进行解码，认为传入进来的字节数组中只会存放一个数字，但是这个数字是可能会有符号的
func DecodeFromVarInt[T gtypes.Signed](bytes VarInt) T {
	var r T
	weight := 0
	for _, b := range bytes {
		// 查看最高有效位，判断是否是最后一个字节，如果是最后一个字节说明当前这个数字读取完毕了
		if b&0b1000_0000 == 0 {
			// 取出数值部分累加，最后一个字节的有效位数只有6
			r = r | (T(b&0b0011_1111) << (weight - 1))
			// 从倒数第二位中取出符号位
			if uint8(b&0b0100_0000) != 0 {
				// 符号位为负数的时候取反
				r *= -1
			}
			break
		} else {
			// 不是最后一个字节，则无脑累加即可
			r = r | (T(b&0b0111_1111) << weight)
			weight += 7
		}

	}
	return r
}

// DecodeVarIntSliceToIntegerSlice 把VarInt数组解码为整数数组
func DecodeVarIntSliceToIntegerSlice[T gtypes.Integer](varIntSlice []VarInt) []T {
	integerSlice :=
	for {

	}
}

//// DecodeVarIntBytes 解码多个varint在一个[]byte的情况，里面可能会包含一个或多个的varint
//func DecodeVarIntBytes[T gtypes.Signed](varIntBytes []byte) []T {
//	slice := make([]T, 0)
//	lastIndex := 0
//	for index, b := range varIntBytes {
//		// 说明读取结束了
//		if b&ByteTopBit == 0 {
//			// 截取一个varint完整的字节数组解码
//			value := DecodeFromVarInt[T](varIntBytes[lastIndex : index+1])
//			slice = append(slice, value)
//			// 指针偏移
//			lastIndex = index + 1
//		}
//	}
//	// 如果有剩下的没结束的也不管了，认为其不是合法的varint数值
//	return slice
//}

// ------------------------------------------------ ---------------------------------------------------------------------
