package client

import (
	"github.com/cloudwego/kitex/client"
)

type ConsulClientSuite struct {
	opts []client.Option
}

// Options return a list client.Option
func (s *ConsulClientSuite) Options() []client.Option {
	return s.opts
}
