package json

import (
	"encoding/json"
	"io"

	"github.com/mirango/framework"
)

type JSON struct{}

func New() *JSON {
	return &JSON{}
}

func (JSON) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (JSON) EncodeTo(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v)
}

func (JSON) EncodingType() string {
	return framework.MIME_JSON
}

func (JSON) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (JSON) DecodeFrom(r io.Reader, v interface{}) error {
	dec := json.NewDecoder(r)
	return dec.Decode(v)
}

func (JSON) DecodingType() string {
	return framework.MIME_JSON
}

func (JSON) MIMEType() string {
	return framework.MIME_JSON
}

type JSONIndented struct{}

func NewIndented() *JSONIndented {
	return &JSONIndented{}
}

func (JSONIndented) Encode(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", " ")
}

func (JSONIndented) EncodeTo(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	return enc.Encode(v)
}

func (JSONIndented) EncodingType() string {
	return framework.MIME_JSON
}

func (JSONIndented) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (JSONIndented) DecodeFrom(r io.Reader, v interface{}) error {
	dec := json.NewDecoder(r)
	return dec.Decode(v)
}

func (JSONIndented) DecodingType() string {
	return framework.MIME_JSON
}

func (JSONIndented) MIMEType() string {
	return framework.MIME_JSON
}
