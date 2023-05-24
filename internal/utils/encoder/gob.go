package encoder

import (
	"encoding/json"
)

func MarshalTo(req any) ([]byte, error) {
	return json.Marshal(req)
}

func MarshalFrom(b []byte, src any) error {
	return json.Unmarshal(b, src)
}
