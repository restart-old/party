package main

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/dragonfly-on-steroids/party"
)

func main() {
	c := server.DefaultConfig()
	s := server.New(&c, nil)
	s.Start()

	for {
		p, err := s.Accept()
		if err != nil {
			return
		}
		_ = party.New(p, 10)
	}
}
