package fn

import (
	"bytes"
)

func newResponse() *Response {
	// TODO: response pool
	return &Response{}
}

type Response struct {
	// TODO: improve this
	bytes.Buffer
}

func (r *Response) Write(b []byte) (int, error) {
	return r.Buffer.Write(b)
}

func (r *Response) Flush() []byte {
	return r.Bytes()
}
