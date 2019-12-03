// Code generated by GCG. DO NOT EDIT.
// Go Code Generator 0.0.7 (https://github.com/OhYee/gcg)

package bytes

import (
	"encoding/binary"
	"math"
)

// FromFloat32 transfer from float32 to []byte
func FromFloat32(value float32) (b []byte) {
	b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(value))
	return
}

// ToFloat32 transfer from []byte to float32
func ToFloat32(b []byte) (value float32) {
	value = math.Float32frombits(ToUint32(b))
	return
}

// FromFloat64 transfer from float64 to []byte
func FromFloat64(value float64) (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(value))
	return
}

// ToFloat64 transfer from []byte to float64
func ToFloat64(b []byte) (value float64) {
	value = math.Float64frombits(ToUint64(b))
	return
}