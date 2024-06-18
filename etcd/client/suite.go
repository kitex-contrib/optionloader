package etcdclient

import (
	"github.com/cloudwego/kitex/client"
)

type EtcdClientSuite struct {
	opts []client.Option
}

// Options return a list client.Option
func (s *EtcdClientSuite) Options() []client.Option {
	return s.opts
}
