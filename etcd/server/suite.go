package server

import (
	"github.com/cloudwego/kitex/server"
)

type EtcdServerSuite struct {
	opts []server.Option
}

// Options return a list client.Option
func (s *EtcdServerSuite) Options() []server.Option {
	return s.opts
}
