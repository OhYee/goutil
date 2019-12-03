package bytes

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/OhYee/goutils"
)

func TestReadNBytes(t *testing.T) {
	testCases := []struct {
		name    string
		input   []byte
		length  int
		want    []byte
		wantErr error
	}{
		{
			name:    "simple",
			input:   []byte{0x01, 0x02, 0x03},
			length:  3,
			want:    []byte{0x01, 0x02, 0x03},
			wantErr: nil,
		},
		{
			name:    "read part of data",
			input:   []byte{0x01, 0x02, 0x03},
			length:  2,
			want:    []byte{0x01, 0x02},
			wantErr: nil,
		},
		{
			name:    "read more than data",
			input:   []byte{0x01, 0x02, 0x03},
			length:  4,
			want:    []byte{},
			wantErr: fmt.Errorf("EOF"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			b, err := ReadNBytes(bytes.NewBuffer(testCase.input), testCase.length)
			if !bytes.Equal(b, testCase.want) || 0 != strings.Compare(fmt.Sprintf("%v", err), fmt.Sprintf("%v", testCase.wantErr)) {
				t.Errorf("Except %v %v, but got %v %v", testCase.want, testCase.wantErr, b, err)
			}
		})
	}
	t.Run("dynamic add data", func(t *testing.T) {
		pipe := goutils.NewPipe(16, []byte{0x01, 0x02})
		answer := []byte{0x01, 0x02, 0x03}
		var b []byte
		var err error
		c := make(chan bool, 8)
		go func() {
			b, err = ReadNBytes(pipe, 3)
			c <- true
		}()
		go func() {
			<-time.After(time.Second * 1)
			pipe.Write([]byte{0x03})
			pipe.Close()
		}()

		<-c
		if !goutils.Equal(b, answer) || !goutils.Equal(err, nil) {
			t.Errorf("Except %v %v, but got %v %v", answer, nil, b, err)
		}
	})
}

func TestWriteWithLength(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want []byte
	}{
		{
			name: "0x1234",
			b:    []byte{0x12, 0x34},
			want: []byte{0, 0, 0, 2, 0x12, 0x34},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			WriteWithLength(buf, tt.b)
			if !goutils.Equal(buf.Bytes(), tt.want) {
				t.Errorf("Want %v, got %v", tt.want, buf.Bytes())
			}
		})
	}
}