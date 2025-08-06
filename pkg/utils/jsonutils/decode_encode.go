package jsonutils

import (
	"encoding/json"
	"io"
)

func DecodeJSON(body io.ReadCloser, v any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}

func EncodeJSON(v any) ([]byte, error) {
	return json.Marshal(v)
}
