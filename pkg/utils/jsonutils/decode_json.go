package jsonutils

import (
	"encoding/json"
	"io"
)

func DecodeJSON(body io.ReadCloser, v any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(v)

	if err != nil {
		return err
	}

	return nil
}
