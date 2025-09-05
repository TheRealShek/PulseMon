package configs

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

// Intervals --> type
// IntervalConfig --> data type
// `yaml:"intervals"` --> mapping
type config struct {
	Intervals IntervalConfig `yaml:"intervals"`
	History   HistoryConfig  `yaml:"history"`
	Alerts    AlertConfig    `yaml:"alerts"`
}

type IntervalConfig struct {
	Cpu     time.Duration `yaml:"cpu"`
	Ram     time.Duration `yaml:"ram"`
	Disk    time.Duration `yaml:"disk"`
	Network time.Duration `yaml:"network"`
}

type HistoryConfig struct {
	Size int `yaml:"size"`
}

type AlertConfig struct {
	Cpu  int `yaml:"cpu"`
	Ram  int `yaml:"ram"`
	Disk int `yaml:"disk"`
}

func LoadConfig(path string) (*config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg config
	//It takes raw YAML bytes (data) and fills a Go struct (cfg).
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
