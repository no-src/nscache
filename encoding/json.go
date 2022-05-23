package encoding

import "encoding/json"

type jsonSerializer struct {
}

func (s *jsonSerializer) Serialize(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (s *jsonSerializer) Deserialize(data []byte, v any) error {
	return json.Unmarshal(data, &v)
}
