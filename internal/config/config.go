package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	ConfigPath string `yaml:"config_path"`
	App        struct {
		Name string `yaml:"name"`
		Env  string `yaml:"env"`
		GRPC struct {
			Port int `yaml:"port"`
		} `yaml:"grpc"`
	} `yaml:"app"`

	Mongo struct {
		ConnectString string `yaml:"connect_string"`
		Database      string `yaml:"database"`
		Collections   struct {
			Movies string `yaml:"movies"`
		} `yaml:"collections"`
	} `yaml:"mongo"`

	Log struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"log"`
}

func LoadConfig() (*Config, error) {
	f, err := os.Open("local.yaml")
	if err != nil {
		return nil, fmt.Errorf("Ошибка при  открытия файла конфигурации: %w", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("Ошибка декодирования base config: %w", err)
	}

	if cfg.ConfigPath != "" {
		f, err = os.Open(cfg.ConfigPath)
		if err != nil {
			return nil, fmt.Errorf("Ошибка открытия файла конфигурации: %w", err)
		}
		defer f.Close()

		decoder = yaml.NewDecoder(f)
		if err = decoder.Decode(&cfg); err != nil {
			return nil, fmt.Errorf("Ошибка декодирования основной конфигурации: %w", err)
		}
	}

	return &cfg, nil
}

func MustLoad() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	return cfg
}
