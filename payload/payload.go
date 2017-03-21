package payload

import (
	"encoding/json"
)

type Payload struct {
	data []byte
}

func (p *Payload) Data() []byte {
	return p.data
}

func (p *Payload) Unmarshal(v interface{}) {
	json.Unmarshal(p.data, v)
}

func NewPayload(d []byte) *Payload {
	// TODO: buffer pool
	return &Payload{d}
}
