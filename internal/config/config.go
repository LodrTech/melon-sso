package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string 			`yaml:"env" env-default:"local"`
	StoragePath string 			`yaml:"storage_path" env-required:"./data"`
	TokenTTL	time.Duration	`yaml:"token_ttl" env-required:"true"`
	GRPC 		GRPCConfig		`yaml:"grpc"`
}

type GRPCConfig struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	err = cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}