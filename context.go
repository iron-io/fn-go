package fn

import (
	"github.com/iron-io/fn-go/payload"
)

type Context struct {
	*payload.Payload
	*Response
}

func newContext(p *payload.Payload, r *Response) *Context {
	// TODO: Context pool
	return &Context{p, r}
}
