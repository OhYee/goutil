// Code generated by GCG. DO NOT EDIT.
// Go Code Generator 0.0.7 (https://github.com/OhYee/gcg)

package fp

import ()

func AllString(f func(string) bool, input []string) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllInt(f func(int) bool, input []int) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllInt8(f func(int8) bool, input []int8) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllInt16(f func(int16) bool, input []int16) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllInt32(f func(int32) bool, input []int32) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllInt64(f func(int64) bool, input []int64) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllUint8(f func(uint8) bool, input []uint8) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllUint16(f func(uint16) bool, input []uint16) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllUint32(f func(uint32) bool, input []uint32) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllUint64(f func(uint64) bool, input []uint64) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllFloat32(f func(float32) bool, input []float32) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllFloat64(f func(float64) bool, input []float64) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func AllByte(f func(byte) bool, input []byte) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}
func All(f func(any) bool, input []any) (output bool) {
	output = true
	for _, data := range input {
		output = output && f(data)
		if !output {
			break
		}
	}
	return
}