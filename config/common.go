package config

import "time"

type EndpointBasicInfo struct {
	ServiceName string            `yaml:"service_name"`
	Method      string            `yaml:"method"`
	Tags        map[string]string `yaml:"tags"`
}

// TimeInterval common time interval struct.
// Unit ns, us, ms, s, m, h.
// Value is the timeout value.
type TimeInterval struct {
	Unit  string `yaml:"unit"`
	Value int    `yaml:"value"`
}

func (t *TimeInterval) Transform() time.Duration {
	var unit time.Duration
	switch t.Unit {
	case "ns":
		unit = time.Nanosecond
	case "us":
		unit = time.Microsecond
	case "ms":
		unit = time.Millisecond
	case "s":
		unit = time.Second
	case "m":
		unit = time.Minute
	case "h":
		unit = time.Hour
	}
	return time.Duration(t.Value) * unit
}

type Tag struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
