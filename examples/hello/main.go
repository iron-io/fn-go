package main

import (
	"fmt"
	"log"

	"github.com/iron-io/fn-go"
)

type payload struct {
	Name string
}

func main() {
	err := fn.HandleRequest(func(c *fn.Context) error {
		var p payload
		c.Unmarshal(&p)
		if p.Name == "" {
			p.Name = "Unknown"
		}
		fmt.Fprintf(c.Response, "Hello, %s\n", p.Name)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
