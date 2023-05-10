package store

import (
	"bytes"
	"testing"
	"time"
)

func TestData(t *testing.T) {
	testCases := []struct {
		name       string
		data       []byte
		expiration time.Duration
	}{
		{"normal", []byte("hello world"), time.Second},
		{"empty data", nil, time.Second},
		{"never expiration", []byte("hello world"), 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sd := NewData(tc.data, tc.expiration)
			if !bytes.Equal(tc.data, sd.Data) {
				t.Errorf("expect to get data %s, but actual %s", string(tc.data), string(sd.Data))
				return
			}
			time.Sleep(tc.expiration)
			expectExpired := tc.expiration > 0
			if sd.IsExpired() != expectExpired {
				t.Errorf("expect to get data expired =>%v, but actual %v", expectExpired, sd.IsExpired())
			}
		})
	}
}
