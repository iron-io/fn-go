package fn

import (
	"errors"
	"os"

	"github.com/iron-io/fn-go/payload"
	"github.com/iron-io/fn-go/protocol"
)

var (
	protos = map[string]protocol.Protocol{}

	DefaultProto = &protocol.NoProtocol{os.Stdin, os.Stdout}
)

func init() {
	RegisterProtocol("http", &protocol.HTTPProtocol{os.Stdin, os.Stdout})
}

func RegisterProtocol(f string, p protocol.Protocol) {
	protos[f] = p
}

type IOHandler interface {
	Receive() (*payload.Payload, error)
	Send(*Response) error
}

func NewIOHandler() (IOHandler, error) {
	format := InputFormat()
	var p protocol.Protocol
	if IsHotContainer() {
		var exist bool
		p, exist = protos[format]
		if !exist {
			return nil, errors.New("Invalid function format")
		}
	} else {
		p = DefaultProto
	}

	return &ioHandler{p}, nil
}

type ioHandler struct {
	protocol protocol.Protocol
}

func (d *ioHandler) Receive() (*payload.Payload, error) {
	data, err := d.protocol.Receive()
	if err != nil {
		return nil, err
	}
	return payload.NewPayload(data), nil
}

func (d *ioHandler) Send(r *Response) error {
	err := d.protocol.Send(r.Flush())
	if err != nil {
		return err
	}
	return nil
}
