package main

import (
	"fmt"
	"github.com/compression-algorithm-research-lab/go-varint"
	"unsafe"
)

func main() {

	// 现在有一个无符号数字
	v := uint64(10)
	// 对齐进行varint编码，可以得到一个字节切片
	valueVarintBytes := varint.Encode[uint64](v)
	// 原始的变量的值占用8个字节的内存，在存储的时候一般也是会视数据类型占用不同的字节数，uint64存储的时候也是要占用8个字节
	fmt.Println(unsafe.Sizeof(v)) // Output: 8
	// 编码之后存储的时候只需要占用一个字节即可，节省了7个字节的存储消耗
	fmt.Println(len(valueVarintBytes)) // Output: 1
	// 注意，varint仅仅是在存储的时候节省空间，并不适合在内存中以varint的方式表示，因为Golang表示数组会有额外的元数据开销（数组的头指针之类的）
	fmt.Println(unsafe.Sizeof(valueVarintBytes)) // Output: 24

	// 需要还原值的时候就解码
	decodeValue := varint.Decode[uint64](valueVarintBytes)
	fmt.Println(decodeValue) // Output: 10
}
