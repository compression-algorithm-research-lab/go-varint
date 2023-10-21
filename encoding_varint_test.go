package varint

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDecodeVarInt(t *testing.T) {
	// TODO 编码出来的值不对
	for i := 0; i < 1000000; i++ {
		n := rand.Int()
		fmt.Println(n)
		varInt := Encode[int](n)
		decodeVarInt := DecodeFromVarInt[int](varInt)
		if n != decodeVarInt {
			fmt.Println("------")
			fmt.Println(n)
			fmt.Println(decodeVarInt)
			panic(n)
		}
	}
}

func TestEncodeToVarInt(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		n := rand.Int()
		fmt.Println(n)
		varInt := Encode[int](n)
		decodeVarInt := DecodeFromVarInt[int](varInt)
		if n != decodeVarInt {
			fmt.Println("------")
			fmt.Println(n)
			fmt.Println(decodeVarInt)
			panic(n)
		}
	}
}