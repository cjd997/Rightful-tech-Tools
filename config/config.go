package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type (
	ServerConfig struct {
		Port        string `toml:"port"`
		FilesFolder string `toml:"files_folder"`
	}

	UploadConfig struct {
		Bucket  string `toml:"bucket"`
		Key     string `toml:"key"`
		Timeout int  `toml:"timeout"`
	}
)

type Config struct {
	Server  ServerConfig
	Upload  UploadConfig
}

func Load(configFile string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(configFile, &cfg); err != nil {
		return nil, fmt.Errorf("error loading config: %s", err.Error())
	}
	return &cfg, nil
}
