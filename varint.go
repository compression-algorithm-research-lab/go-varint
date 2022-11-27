package varint

import "github.com/golang-infrastructure/go-gtypes"

// VarInt 用于表示一个可变长无符号整数
type VarInt []byte

// From 从无符号类型创建一个variant
func From[T gtypes.Unsigned](value T) VarInt {
	return Encode[T](value)
}

func (x *VarInt) ToUint() uint {
	return Decode[uint](*x)
}

func (x *VarInt) ToUint64() uint64 {
	return Decode[uint64](*x)
}

func (x *VarInt) Add(v VarInt) VarInt {
	newValue := x.ToUint64() + v.ToUint64()
	return Encode[uint64](newValue)
}

func (x *VarInt) Sub(v VarInt) VarInt {
	newValue := x.ToUint64() - v.ToUint64()
	return Encode[uint64](newValue)
}

func (x *VarInt) Multi(v VarInt) VarInt {
	newValue := x.ToUint64() * v.ToUint64()
	return Encode[uint64](newValue)
}

func (x *VarInt) Divide(v VarInt) VarInt {
	newValue := x.ToUint64() / v.ToUint64()
	return Encode[uint64](newValue)
}

func (x *VarInt) GreatThan(v VarInt) bool {
	return x.ToUint64() > v.ToUint64()
}

func (x *VarInt) GreatThanOrEquals(v VarInt) bool {
	return x.ToUint64() >= v.ToUint64()
}

func (x *VarInt) LessThan(v VarInt) bool {
	return x.ToUint64() < v.ToUint64()
}

func (x *VarInt) LessThanOrEqual(v VarInt) bool {
	return x.ToUint64() <= v.ToUint64()
}

func (x *VarInt) Equals(v VarInt) bool {
	return x.ToUint64() == v.ToUint64()
}
