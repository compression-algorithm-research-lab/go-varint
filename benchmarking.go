package varint

import (
	"github.com/golang-infrastructure/go-gtypes"
	"unsafe"
)

// 这个文件中存放的是一些计算压缩率相关的方法

// BenchmarkingValueRatio 计算单个值的压缩率
func BenchmarkingValueRatio[T gtypes.Unsigned](value T) float64 {
	rawByteSize := unsafe.Sizeof(value)
	varintBytes := Encode(value)
	return float64(len(varintBytes)) * 100 / float64(rawByteSize)
}

// BenchmarkingValueSliceRatio 计算整个切片的压缩率
func BenchmarkingValueSliceRatio[T gtypes.Unsigned](valueSlice []T) float64 {
	rawByteSizeSum := uint64(0)
	varintBytesSum := uint64(0)
	for _, v := range valueSlice {
		bytes := Encode(v)
		varintBytesSum += uint64(len(bytes))
		rawByteSizeSum += uint64(unsafe.Sizeof(v))
	}
	return float64(varintBytesSum) * 100 / float64(rawByteSizeSum)
}
