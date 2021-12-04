package main

import (
	"github.com/RestartFU/party"
	"github.com/df-mc/dragonfly/server"
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
