// Code generated by GCG. DO NOT EDIT.
// Go Code Generator 0.0.8 (https://github.com/OhYee/gcg)

package fp

import (
	"reflect"
	"testing"
)

func TestAnyString(t *testing.T) {
	type args struct {
		f     func(string, int) bool
		input []string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(s string, idx int) bool { return len(s) > 0 }, input: []string{"a", "bb", "ccc", "dddd"}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(s string, idx int) bool { return len(s) > 2 }, input: []string{"a", "bb", "ccc", "dddd"}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(s string, idx int) bool { return len(s) > 5 }, input: []string{"a", "bb", "ccc", "dddd"}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyString(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyInt(t *testing.T) {
	type args struct {
		f     func(int, int) bool
		input []int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n int, idx int) bool { return n > 0 }, input: []int{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n int, idx int) bool { return n > 3 }, input: []int{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n int, idx int) bool { return n > 5 }, input: []int{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyInt(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyInt() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyInt8(t *testing.T) {
	type args struct {
		f     func(int8, int) bool
		input []int8
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n int8, idx int) bool { return n > 0 }, input: []int8{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n int8, idx int) bool { return n > 3 }, input: []int8{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n int8, idx int) bool { return n > 5 }, input: []int8{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyInt8(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyInt8() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyInt16(t *testing.T) {
	type args struct {
		f     func(int16, int) bool
		input []int16
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n int16, idx int) bool { return n > 0 }, input: []int16{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n int16, idx int) bool { return n > 3 }, input: []int16{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n int16, idx int) bool { return n > 5 }, input: []int16{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyInt16(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyInt16() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyInt32(t *testing.T) {
	type args struct {
		f     func(int32, int) bool
		input []int32
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n int32, idx int) bool { return n > 0 }, input: []int32{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n int32, idx int) bool { return n > 3 }, input: []int32{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n int32, idx int) bool { return n > 5 }, input: []int32{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyInt32(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyInt32() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyInt64(t *testing.T) {
	type args struct {
		f     func(int64, int) bool
		input []int64
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n int64, idx int) bool { return n > 0 }, input: []int64{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n int64, idx int) bool { return n > 3 }, input: []int64{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n int64, idx int) bool { return n > 5 }, input: []int64{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyInt64(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyInt64() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyUint8(t *testing.T) {
	type args struct {
		f     func(uint8, int) bool
		input []uint8
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n uint8, idx int) bool { return n > 0 }, input: []uint8{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n uint8, idx int) bool { return n > 3 }, input: []uint8{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n uint8, idx int) bool { return n > 5 }, input: []uint8{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyUint8(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyUint8() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyUint16(t *testing.T) {
	type args struct {
		f     func(uint16, int) bool
		input []uint16
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n uint16, idx int) bool { return n > 0 }, input: []uint16{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n uint16, idx int) bool { return n > 3 }, input: []uint16{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n uint16, idx int) bool { return n > 5 }, input: []uint16{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyUint16(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyUint16() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyUint32(t *testing.T) {
	type args struct {
		f     func(uint32, int) bool
		input []uint32
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n uint32, idx int) bool { return n > 0 }, input: []uint32{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n uint32, idx int) bool { return n > 3 }, input: []uint32{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n uint32, idx int) bool { return n > 5 }, input: []uint32{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyUint32(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyUint32() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyUint64(t *testing.T) {
	type args struct {
		f     func(uint64, int) bool
		input []uint64
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n uint64, idx int) bool { return n > 0 }, input: []uint64{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n uint64, idx int) bool { return n > 3 }, input: []uint64{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n uint64, idx int) bool { return n > 5 }, input: []uint64{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyUint64(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyUint64() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyFloat32(t *testing.T) {
	type args struct {
		f     func(float32, int) bool
		input []float32
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n float32, idx int) bool { return n > 0 }, input: []float32{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n float32, idx int) bool { return n > 3 }, input: []float32{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n float32, idx int) bool { return n > 5 }, input: []float32{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyFloat32(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyFloat32() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyFloat64(t *testing.T) {
	type args struct {
		f     func(float64, int) bool
		input []float64
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n float64, idx int) bool { return n > 0 }, input: []float64{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n float64, idx int) bool { return n > 3 }, input: []float64{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n float64, idx int) bool { return n > 5 }, input: []float64{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyFloat64(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyFloat64() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAnyByte(t *testing.T) {
	type args struct {
		f     func(byte, int) bool
		input []byte
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(n byte, idx int) bool { return n > 0 }, input: []byte{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(n byte, idx int) bool { return n > 3 }, input: []byte{1, 2, 3, 4, 5}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(n byte, idx int) bool { return n > 5 }, input: []byte{1, 2, 3, 4, 5}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AnyByte(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("AnyByte() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAny(t *testing.T) {
	type args struct {
		f     func(any, int) bool
		input []any
	}
	tests := []struct {
		name       string
		args       args
		wantOutput bool
	}{
		{
			name:       "all item is true",
			args:       args{f: func(s any, idx int) bool { return reflect.TypeOf(s).Kind() == reflect.Int }, input: []any{1, 2}},
			wantOutput: true,
		},
		{
			name:       "some items are true and others are false",
			args:       args{f: func(s any, idx int) bool { return reflect.TypeOf(s).Kind() == reflect.Int }, input: []any{1, 2, "A"}},
			wantOutput: true,
		},
		{
			name:       "all item is false",
			args:       args{f: func(s any, idx int) bool { return reflect.TypeOf(s).Kind() == reflect.Int }, input: []any{"A", "B"}},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Any(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Any() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
