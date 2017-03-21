package protocol

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type HTTPProtocol struct {
	In  io.Reader
	Out io.Writer
}

func (p *HTTPProtocol) IsStreamable() bool {
	return true
}

func (p *HTTPProtocol) Receive() ([]byte, error) {
	r := bufio.NewReader(p.In)
	req, err := http.ReadRequest(r)

	if err != nil {
		return nil, errors.New("Input is not using HTTP protocol")
	}

	l, _ := strconv.Atoi(req.Header.Get("Content-Length"))
	body := make([]byte, l)
	r.Read(body)

	return body, nil
}

func (p *HTTPProtocol) Send(b []byte) error {
	res := http.Response{
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		StatusCode: 200,
		Status:     "OK",
	}

	buf := bytes.NewBuffer(b)
	res.Body = ioutil.NopCloser(buf)
	res.ContentLength = int64(buf.Len())
	res.Write(p.Out)

	return nil
}
