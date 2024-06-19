package server

import (
	"github.com/cloudwego/kitex/server"
)

type ConsulServerSuite struct {
	opts []server.Option
}

// Options return a list client.Option
func (s *ConsulServerSuite) Options() []server.Option {
	return s.opts
}
