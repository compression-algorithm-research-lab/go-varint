package main

import (
	"encoding/hex"
	"fmt"
	"github.com/compression-algorithm-research-lab/go-varint"
)

func main() {

	// 要编码的无符号整数切片
	uintSlice := []uint{1, 127, 0, 255, 256}
	// 对切片编码之后会返回一个字节切片，切片中挨个存放的就是每个无符号整数的varint编码后的值
	uintVarintByteSlice := varint.EncodeSlice(uintSlice)

	// 打印看一下，字节切片的值是这样子的
	// []byte{
	//		0x1,       // 1
	//		0x7f,      // 127
	//		0x0,       // 0
	//		0xff, 0x1, // 255
	//		0x80, 0x2, // 256
	//	}
	fmt.Println(hex.EncodeToString(uintVarintByteSlice)) // Output: 017f00ff018002

	// 然后解码
	slice := varint.DecodeSlice[uint](uintVarintByteSlice)
	fmt.Println(slice) // Output: [1 127 0 255 256]

}
