package vector

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Vector3 represent point or direction in 3D plane.
type Vector3 struct {
	x, y, z DataType
}

// NewZeroVector3 return the zero vector3.
func NewZeroVector3() *Vector3 {
	return &Vector3{}
}

// NewVector3 create Vector struct.
func NewVector3(x, y, z DataType) *Vector3 {
	return &Vector3{
		x: x,
		y: y,
	}
}

// NewVector3ByVector2 copy value x, z and return a new struct.
func NewVector3ByVector2(v *Vector2) *Vector3 {
	return NewVector3(v.x, 0, v.y)
}

// Clone clone the ori Vector2 into the new memory.
func (vec3 Vector3) Clone() *Vector3 {
	return &vec3
}

// MagnitudeSquare return the length's square of the vector.
func (vec3 *Vector3) MagnitudeSquare() DataType {
	return vec3.x*vec3.x + vec3.y*vec3.y + vec3.z*vec3.z
}

// GetX get value x.
func (vec3 *Vector3) GetX() DataType {
	return vec3.x
}

// SetX set value x.
func (vec3 *Vector3) SetX(x DataType) {
	vec3.x = x
}

// GetY get Value y.
func (vec3 *Vector3) GetY() DataType {
	return vec3.y
}

// SetY set value y.
func (vec3 *Vector3) SetY(y DataType) {
	vec3.y = y
}

// GetZ get Value z.
func (vec3 *Vector3) GetZ() DataType {
	return vec3.z
}

// SetZ set value z.
func (vec3 *Vector3) SetZ(z DataType) {
	vec3.z = z
}

// Magnitude return the length of the vector.
func (vec3 *Vector3) Magnitude() DataType {
	return DataType(math.Sqrt(float64(vec3.MagnitudeSquare())))
}

// Add represent the other vector is added to self.
func (vec3 *Vector3) Add(v *Vector3) *Vector3 {
	vec3.x += v.x
	vec3.y += v.y
	vec3.z += v.z
	return vec3
}

// Sub represent self vector sub the other vector.
func (vec3 *Vector3) Sub(v *Vector3) *Vector3 {
	vec3.x -= v.x
	vec3.y -= v.y
	vec3.z -= v.z
	return vec3
}

// MulScalar represent the scalar multiply the vector.
func (vec3 *Vector3) MulScalar(d DataType) *Vector3 {
	vec3.x *= d
	vec3.y *= d
	vec3.z *= d
	return vec3
}

// Scale scale axis value
func (vec3 *Vector3) Scale(d DataType) *Vector3 {
	return vec3.MulScalar(d)
}

// Normalize the vector normalize
func (vec3 *Vector3) Normalize() *Vector3 {
	if size := vec3.Magnitude(); size != 0 {
		vec3.MulScalar(1 / size)
	}
	return vec3
}

// Rotate the vector
func (vec3 *Vector3) Rotate(angle DataType) {
	cos := DataType(math.Cos(float64(angle)))
	sin := DataType(math.Sin(float64(angle)))
	/*
	 *  rotate matrix
	 * --            --
	 * |   cos   sin   |
	 * |  -sin   cos   |
	 * --            --
	 */
	vec3.x = cos*vec3.x + sin*vec3.y
	vec3.y = -sin*vec3.x + cos*vec3.y
}

// 向量分解

// String implement print operator.
func (vec3 *Vector3) String() string {
	return fmt.Sprintf("Vector3(%v, %v, %v)", vec3.x, vec3.y, vec3.z)
}

// UnmarshalJSON implement the json interface.
func (vec3 *Vector3) UnmarshalJSON(bts []byte) error {
	bts = bts[1 : len(bts)-1] // 去除首位的双引号
	dataSlc := strings.Split(string(bts), "#")
	if len(dataSlc) != 3 {
		return fmt.Errorf("unmarshal struct Vector3 error:%v", string(bts))
	}
	x, err := strconv.ParseFloat(dataSlc[0], 64)
	if err != nil {
		return fmt.Errorf("unmarshal struct Vector3 X error:%v, value:%v", err, string(bts))
	}
	vec3.x = DataType(x)

	y, err := strconv.ParseFloat(dataSlc[1], 64)
	if err != nil {
		return fmt.Errorf("unmarshal struct Vector3 Y error:%v, value:%v", err, string(bts))
	}
	vec3.y = DataType(y)

	z, err := strconv.ParseFloat(dataSlc[1], 64)
	if err != nil {
		return fmt.Errorf("unmarshal struct Vector3 Z error:%v, value:%v", err, string(bts))
	}
	vec3.z = DataType(z)
	return nil
}

// MarshalJSON implement the json interface.
func (vec3 Vector3) MarshalJSON() ([]byte, error) {
	jsonStr := fmt.Sprintf("\"%v#%v#%v\"", vec3.x, vec3.y, vec3.z)
	return []byte(jsonStr), nil
}

// MarshalBinary implement gob encoding.
func (vec3 Vector3) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintln(&b, vec3.x, vec3.y, vec3.z)
	return b.Bytes(), nil
}

// UnmarshalBinary implement gob decoding.
func (vec3 *Vector3) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &vec3.x, &vec3.y, &vec3.z)
	return err
}
