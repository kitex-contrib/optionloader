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
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
)

type ConfigParser interface {
	Decode(configType ConfigType, data []byte, config *ConsulConfig) error
}

type defaultParser struct {
}

func (p *defaultParser) Decode(configType ConfigType, data []byte, config *ConsulConfig) error {
	switch configType {
	case JSON:
		return json.Unmarshal(data, config)
	case YAML:
		return yaml.Unmarshal(data, config)
	default:
		return fmt.Errorf("unsupported config data type %s", configType)
	}
}
