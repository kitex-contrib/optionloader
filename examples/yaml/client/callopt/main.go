package main

import (
	"fmt"
	"github.com/kitex-contrib/optionloader/configloader/yaml"
	"github.com/kitex-contrib/optionloader/optionloader/client/callopt"
)

func main() {
	filePath := "./examples/yaml/client/callopt/config.yaml"
	cfg, err := yaml.LoadCalloptConfig(filePath)
	if err != nil {
		panic(err)
	}
	loader := callopt.NewOptionLoader()
	options, err := loader.Load(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(options))
}
