package json

import (
	"encoding/json"
	"io"

	"github.com/wlMalk/mirango/framework"
)

type Renderer struct{}

func New() *Renderer {
	return &Renderer{}
}

func (r *Renderer) MimeType() string {
	return framework.MIME_JSON
}

func (r *Renderer) RenderToWriter(wr io.Writer, c framework.Context, data interface{}) error {
	b, err := r.Render(c, data)
	if err != nil {
		return err
	}
	_, err = wr.Write(b)
	return err
}

func (r *Renderer) Render(c framework.Context, data interface{}) ([]byte, error) {
	if c.IsIndented() {
		return json.MarshalIndent(data, "", " ")
	}
	return json.Marshal(data)
}
