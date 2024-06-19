package client

import (
	"fmt"
	"io/ioutil"
	"log"
)

// yml/client/client.go
func ReadYMLConfig(filename string) ([]byte, error) {
	// 读取文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading YAML file: %s\n", err)
		return nil, err
	}
	return data, nil
}
func NewReader(filename string) *ymlReader {
	reader := &ymlReader{
		decoder: defaultDecoder,
	}
	data, err := ReadYMLConfig(filename)
	if err != nil {
		log.Printf("Error reading YAML file: %s\n", err)
		return nil
	}
	err = reader.ReadToConfig(data)
	if err != nil {
		log.Printf("Error decoding YAML file: %s\n", err)
		return nil
	}
	fmt.Println("Read config from file: ", filename)
	//fmt.Println("Config: ", reader.config)
	return reader
}
