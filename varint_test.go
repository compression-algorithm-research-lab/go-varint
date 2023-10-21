package varint

//import (
//	"testing"
//)
//
//func TestEncode(t *testing.T) {
//	var a uint64
//	a = 123456
//	bytes := Encode(a)
//	t.Log(bytes)
//	t.Log(Decode[uint64](bytes))
//}
//
//func TestDecodeSlice(t *testing.T) {
//	ints := []uint64{
//		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
//	}
//	bytes := EncodeSlice(ints)
//	t.Log(len(bytes))
//	t.Log(bytes)
//	slice := DecodeSlice[uint64](bytes)
//	t.Log(slice)
//}
