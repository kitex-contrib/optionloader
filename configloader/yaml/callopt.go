package yaml

import (
	"fmt"
	"github.com/kitex-contrib/optionloader/config"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadCalloptConfig(filePath string) (*config.CalloptConfig, error) {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", filePath)
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg config.CalloptConfig

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
