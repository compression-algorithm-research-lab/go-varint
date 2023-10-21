package varint

import "github.com/golang-infrastructure/go-gtypes"

// VarInt 用于表示一个可变长整数，可以根据自己的需要选择是否支持符号存储（最高位可以是符号位，也可以是数据位）
type VarInt []byte

// From 从无符号整数类型创建一个variant
func From[T gtypes.Integer](value T) VarInt {
	return Encode[T](value)
}

// IsZero 此整数是否为0
func (x *VarInt) IsZero() bool {
	return x.ToInt() == 0
}

// 转换为
func (x *VarInt) ToInt() uint {
	return DecodeFromVarInt[uint](*x)
}

func (x *VarInt) ToInt64() int64 {
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
