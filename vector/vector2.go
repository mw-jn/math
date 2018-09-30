package vector

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Vector2 represent point or direction in 2D plane.
type Vector2 struct {
	x, y DataType

	// temp data, reduce calc.
	magnitude      DataType
	sqrMagnitude   DataType
	dirtyMagnitude bool
	normalized     *Vector2
	dirtyNormal    bool
}

// NewZeroVector2 return zero vector2.
func NewZeroVector2() *Vector2 {
	return &Vector2{}
}

// NewVector2 create Vector struct.
func NewVector2(x, y DataType) *Vector2 {
	v := &Vector2{
		x:              x,
		y:              y,
		dirtyMagnitude: true,
		dirtyNormal:    true,
	}
	return v
}

// NewVector2ByVector3 copy value x, z and return a new struct.
func NewVector2ByVector3(v *Vector3) *Vector2 {
	return NewVector2(v.x, v.z)
}

// Clone clone Vector2
func (vec2 Vector2) Clone() *Vector2 {
	return &vec2
}

// GetX get value x.
func (vec2 *Vector2) GetX() DataType {
	return vec2.x
}

// SetX set value x.
func (vec2 *Vector2) SetX(x DataType) {
	vec2.x = x
	vec2.setDirtyMagitude(true)
}

// GetY get Value y.
func (vec2 *Vector2) GetY() DataType {
	return vec2.y
}

// SetY set value y.
func (vec2 *Vector2) SetY(y DataType) {
	vec2.y = y
	vec2.setDirtyMagitude(true)
}

// Set set value x and y of vector.
func (vec2 *Vector2) Set(x, y DataType) {
	vec2.x = x
	vec2.y = y
	vec2.setDirtyMagitude(true)
}

func (vec2 *Vector2) calcMagnitudeValue() {
	vec2.sqrMagnitude = vec2.x*vec2.x + vec2.y*vec2.y
	vec2.magnitude = DataType(math.Sqrt(float64(vec2.sqrMagnitude)))
	vec2.setDirtyMagitude(false)
}

// MagnitudeSquare return the length's square of the vector.
func (vec2 *Vector2) MagnitudeSquare() DataType {
	if vec2.dirtyMagnitude {
		vec2.calcMagnitudeValue()
	}
	return vec2.sqrMagnitude
}

// Magnitude return the length of the vector.
func (vec2 *Vector2) Magnitude() DataType {
	if vec2.dirtyMagnitude {
		vec2.calcMagnitudeValue()
	}
	return vec2.magnitude
}

// MulScalar represent the scalar multiply the vector.
func (vec2 *Vector2) MulScalar(d DataType) *Vector2 {
	vec2.x *= d
	vec2.y *= d
	vec2.setDirtyMagitude(true)
	return vec2
}

// Scale scale axis value
func (vec2 *Vector2) Scale(d DataType) *Vector2 {
	return vec2.MulScalar(d)
}

// Normalize the vector normalize
func (vec2 *Vector2) Normalize() *Vector2 {
	if vec2.dirtyNormal {
		if size := vec2.Magnitude(); size != 0 {
			vec2.MulScalar(1 / size)
		}
	}
	return vec2
}

// Rotate the vector
func (vec2 *Vector2) Rotate(angle DataType) *Vector2 {
	cos := DataType(math.Cos(float64(angle)))
	sin := DataType(math.Sin(float64(angle)))
	/*
	 * theory: decompose vector, then merge vector.
	 *  rotate matrix
	 * --            --
	 * |   cos   sin   |
	 * |  -sin   cos   |
	 * --            --
	 */
	vec2.x = cos*vec2.x - sin*vec2.y
	vec2.y = sin*vec2.x + cos*vec2.y
	vec2.setDirtyMagitude(true)
	return vec2
}

// Invert inverts the vector.
func (vec2 *Vector2) Invert() *Vector2 {
	vec2.x = -vec2.x
	vec2.y = -vec2.y
	return vec2
}

// Equal return if two vectors equal.
func (vec2 *Vector2) Equal(v *Vector2) bool {
	return IsZero(vec2.GetX()-v.GetX()) && IsZero(vec2.GetY()-v.GetY())
}

// Add represent the other vector is added to self.
func (vec2 *Vector2) Add(v *Vector2) *Vector2 {
	vec2.x += v.x
	vec2.y += v.y
	vec2.setDirtyMagitude(true)
	return vec2
}

// Sub represent self vector sub the other vector.
func (vec2 *Vector2) Sub(v *Vector2) *Vector2 {
	vec2.x -= v.x
	vec2.y -= v.y
	vec2.setDirtyMagitude(true)
	return vec2
}

// Dot return two vector similar value.
// if valye equal 0, then the two vectors are vertical.
func (vec2 *Vector2) Dot(v *Vector2) DataType {
	return vec2.GetX()*v.GetX() + vec2.GetY()*v.GetY()
}

// Cos return two vector cosine value.
func (vec2 *Vector2) Cos(v *Vector2) DataType {
	if vec2.Magnitude() == 0 || v.Magnitude() == 0 {
		panic("zero vector not cos value")
	}
	return vec2.Dot(v) / (vec2.Magnitude() * vec2.Magnitude())
}

// Angle return the angle of two vector.
func (vec2 *Vector2) Angle(v *Vector2) DataType {
	cos := vec2.Cos(v)
	if cos > 1.0 {
		cos = 1.0
	} else if cos < -1.0 {
		cos = -1.0
	}
	return DataType(math.Acos(float64(cos)))
}

// Distance return the distance of two vectors.
func (vec2 *Vector2) Distance(v *Vector2) DataType {
	return vec2.Clone().Sub(v).Magnitude()
}

// Lerp return the vector that point in linear between in [vec2, v].
func (vec2 *Vector2) Lerp(v *Vector2, t float32) *Vector2 {
	if t < 0 || t > 1 {
		fmt.Println("param t not in range [0, 1]")
		return nil
	}
	dir := v.Clone().Sub(vec2).Scale(DataType(t))
	return vec2.Clone().Add(dir)
}

// String implement print operator.
func (vec2 *Vector2) String() string {
	return fmt.Sprintf("vector2(%v, %v)", vec2.x, vec2.y)
}

// UnmarshalJSON implement the json interface.
func (vec2 *Vector2) UnmarshalJSON(bts []byte) error {
	bts = bts[1 : len(bts)-1] // 去除首位的双引号
	dataSlc := strings.Split(string(bts), "#")
	if len(dataSlc) != 2 {
		return fmt.Errorf("unmarshal struct Vector2 error:%v", string(bts))
	}
	x, err := strconv.ParseFloat(dataSlc[0], 64)
	if err != nil {
		return fmt.Errorf("unmarshal struct Vector2 X error:%v, value:%v", err, string(bts))
	}
	vec2.x = DataType(x)

	y, err := strconv.ParseFloat(dataSlc[1], 64)
	if err != nil {
		return fmt.Errorf("unmarshal struct Vector2 Y error:%v, value:%v", err, string(bts))
	}
	vec2.y = DataType(y)
	return nil
}

// MarshalJSON implement the json interface.
func (vec2 Vector2) MarshalJSON() ([]byte, error) {
	jsonStr := fmt.Sprintf("\"%v#%v\"", vec2.x, vec2.y)
	return []byte(jsonStr), nil
}

// MarshalBinary implement gob encoding.
func (vec2 Vector2) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintln(&b, vec2.x, vec2.y)
	return b.Bytes(), nil
}

// UnmarshalBinary implement gob decoding.
func (vec2 *Vector2) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &vec2.x, &vec2.y)
	return err
}

func (vec2 *Vector2) setDirtyMagitude(d bool) {
	vec2.dirtyMagnitude = d
}

func (vec2 *Vector2) setDirtyNormal(n bool) {
	vec2.dirtyNormal = n
}
