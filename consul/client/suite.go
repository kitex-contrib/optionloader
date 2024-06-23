// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/streamclient"
)

type ConsulClientSuite struct {
	opts []client.Option
}

// Options return a list client.Option
func (s *ConsulClientSuite) Options() []client.Option {
	newOpts := make([]client.Option, len(s.opts))
	copy(newOpts, s.opts)
	return newOpts
}

type ConsulClientStreamSuite struct {
	opts []streamclient.Option
}

// Options return a list streamclient.Option
func (s *ConsulClientStreamSuite) Options() []streamclient.Option {
	newOpts := make([]streamclient.Option, len(s.opts))
	copy(newOpts, s.opts)
	return newOpts
}
