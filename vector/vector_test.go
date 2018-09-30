package vector

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

type DataStruct struct {
	A int
	B string
	C *Vector2
}

func TestVectorJson(t *testing.T) {
	data := &DataStruct{
		A: 1,
		B: "ma",
		C: NewVector2(4, 5),
	}

	bts, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(bts))

	d := &DataStruct{}

	json.Unmarshal(bts, d)

	fmt.Println("code:", *d)
}

func TestVector2Rotate(t *testing.T) {
	vecs := NewVector2(1, 1)
	vecs.Rotate(math.Pi / 6)
	fmt.Println(vecs)
}
