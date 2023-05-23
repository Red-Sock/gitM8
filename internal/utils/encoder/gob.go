package encoder

import (
	"bytes"
	"encoding/gob"

	"github.com/pkg/errors"
)

func MarshalToGob(req any) ([]byte, error) {
	var payload []byte
	enc := gob.NewEncoder(bytes.NewBuffer(payload))

	err := enc.Encode(req)
	if err != nil {
		return payload, errors.Wrap(err, "error marshalling ticket rule to bytes")
	}

	return payload, nil
}

func MarshalFromGob(b []byte, src any) error {
	enc := gob.NewDecoder(bytes.NewBuffer(b))

	err := enc.Decode(src)
	if err != nil {
		return errors.Wrap(err, "error marshalling ticket rule to bytes")
	}

	return nil
}
