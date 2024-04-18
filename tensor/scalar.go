package tensor

import "reflect"

// NumericScalar is a type that is a numeric scalar. For example, int, float64, etc.
type NumericScalar interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		complex64 | complex128
}

/*
Scalar is a type that is only a single value, not a collection of values. For example, bool, int, float64, etc.
It's a union type of bool and NumericScalar.

The reason I've to separate boolean and numbers is that T(0), where T is constrained to Scalar would fail because
bool(0) doesn't work in Go. Also for mathematical operations, we need to differentiate between boolean and numbers
as math operators like +, -, *, /, etc. don't work on boolean values.

Hence the two separate interfaces: Scalar and NumericScalar.
*/
type Scalar interface {
	bool | NumericScalar
}

// Checks if the provided value is a Scalar or not.
func IsScalar(value interface{}) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		return true
	default:
		return false
	}
}