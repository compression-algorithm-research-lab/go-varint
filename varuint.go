package varint

//import "github.com/golang-infrastructure/go-gtypes"
//
//// VarUInt 用于表示一个可变长无符号整数
//type VarUInt []byte
//
//// VarUIntFrom 从无符号类型创建一个variant
//func VarUIntFrom[T gtypes.Unsigned](value T) VarUInt {
//	return Encode[T](value)
//}
//
//func (x *VarUInt) ToUint() uint {
//	return Decode[uint](*x)
//}
//
//func (x *VarUInt) ToUint64() uint64 {
//	return Decode[uint64](*x)
//}
//
//func (x *VarUInt) Add(v VarUInt) VarUInt {
//	newValue := x.ToUint64() + v.ToUint64()
//	return Encode[uint64](newValue)
//}
//
//func (x *VarUInt) Sub(v VarUInt) VarUInt {
//	newValue := x.ToUint64() - v.ToUint64()
//	return Encode[uint64](newValue)
//}
//
//func (x *VarUInt) Multi(v VarUInt) VarUInt {
//	newValue := x.ToUint64() * v.ToUint64()
//	return Encode[uint64](newValue)
//}
//
//func (x *VarUInt) Divide(v VarUInt) VarUInt {
//	newValue := x.ToUint64() / v.ToUint64()
//	return Encode[uint64](newValue)
//}
//
//func (x *VarUInt) GreatThan(v VarUInt) bool {
//	return x.ToUint64() > v.ToUint64()
//}
//
//func (x *VarUInt) GreatThanOrEquals(v VarUInt) bool {
//	return x.ToUint64() >= v.ToUint64()
//}
//
//func (x *VarUInt) LessThan(v VarUInt) bool {
//	return x.ToUint64() < v.ToUint64()
//}
//
//func (x *VarUInt) LessThanOrEqual(v VarUInt) bool {
//	return x.ToUint64() <= v.ToUint64()
//}
//
//func (x *VarUInt) Equals(v VarUInt) bool {
//	return x.ToUint64() == v.ToUint64()
//}
