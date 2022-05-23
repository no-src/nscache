package encoding

// Serializer the serializer for data
type Serializer interface {
	// Serialize serialize the data to byte array
	Serialize(v any) ([]byte, error)
	// Deserialize deserialize the byte array to destination value
	Deserialize(data []byte, v any) error
}

// DefaultSerializer the default Serializer implementation
var DefaultSerializer Serializer = &jsonSerializer{}
