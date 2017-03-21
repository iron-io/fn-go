package protocol

import (
	"io"
	"io/ioutil"
)

type Protocol interface {
	IsStreamable() bool
	Receive() ([]byte, error)
	Send([]byte) error
}

type NoProtocol struct {
	In  io.Reader
	Out io.Writer
}

func (p *NoProtocol) IsStreamable() bool {
	return false
}

func (p *NoProtocol) Receive() ([]byte, error) {
	body, err := ioutil.ReadAll(p.In)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p *NoProtocol) Send(b []byte) error {
	_, err := p.Out.Write(b)
	return err
}
