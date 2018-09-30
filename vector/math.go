package vector

import (
	"math"
)

const (
	DATATYPE_PRECISE = 1e-6
)

// DataType represent data type.
type DataType float32

// Min return min value between d1 and d2.
func Min(d1, d2 DataType) DataType {
	if d1 < d2 {
		return d1
	}
	return d2
}

// Max return max value between d1 and d2.
func Max(d1, d2 DataType) DataType {
	if d1 > d2 {
		return d1
	}
	return d2
}

// IsZero return value v is zero.
func IsZero(v DataType) bool {
	return math.Abs(float64(v)) < DATATYPE_PRECISE
}

// MinVec2 return min value between vector d1 and d2.
func MinVec2(v1, v2 *Vector2) *Vector2 {
	return NewVector2(Min(v1.GetX(), v2.GetX()), Min(v1.GetY(), v2.GetY()))
}

// MaxVec2 return max value between vector d1 and d2.
func MaxVec2(v1, v2 *Vector2) *Vector2 {
	return NewVector2(Max(v1.GetX(), v2.GetX()), Max(v1.GetY(), v2.GetY()))
}
