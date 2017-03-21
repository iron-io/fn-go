package fn

import (
	"fmt"
	"os"
)

type manager struct {
	Handler
	IOHandler
}

func (m *manager) Run() error {
	for {
		payload, err := m.Receive()
		resp := newResponse()
		if err == nil {
			c := newContext(payload, resp)
			err := m.Handler.Handle(c)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				// TODO: error response
			}
			m.Send(c.Response)
		} else {
			fmt.Fprintln(os.Stderr, err)
			// TODO: error response
		}
		if !IsHotContainer() {
			break
		}
	}

	return nil
}
