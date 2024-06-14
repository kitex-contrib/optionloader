package client

import (
	"go.uber.org/zap"
	"log"
	"time"
)

// Options etcd config options. All the fields have default value.
type Options struct {
	Node             []string
	Prefix           string
	ClientPathFormat string
	Timeout          time.Duration
	LoggerConfig     *zap.Config
	ConfigParser     ConfigParser
}

func NewReader(filename string) (*etcdReader, error) {
	//todo
	reader := &etcdReader{
		decoder: parser{},
	}
	key := Key{}
	err := reader.ReadEtcdConfig(key)
	if err != nil {
		log.Printf("Error reading YAML file: %s\n", err)
		return nil, err
	}
	//err = reader.ReadToConfig(data)
	//if err != nil {
	//	log.Printf("Error decoding YAML file: %s\n", err)
	//	return nil, err
	//}
	//fmt.Println("Read config from file: ", filename)
	//fmt.Println("Config: ", reader.config)
	return reader, err
}
