package testutil

import (
	"errors"
)

var (
	ErrMockSerialize   = errors.New("mock serialize error")
	ErrMockDeserialize = errors.New("mock deserialize error")
)

type MockErrSerializer struct {
}

func (s *MockErrSerializer) Serialize(v any) ([]byte, error) {
	return nil, ErrMockSerialize
}

func (s *MockErrSerializer) Deserialize(data []byte, v any) error {
	return ErrMockDeserialize
}
