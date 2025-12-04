package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env"`
	HTTPServer HTTPServer `yaml:"http_server"`
	PostgreSQL PostgreSQL `yaml:"postgre_sql"`
	Clinet     Client     `yaml:"client"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type PostgreSQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"db_name"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password" env_required:"PGSQL_PASSWORD"`
}

type Client struct {
	Address string        `yaml:"address"`
	Timeout time.Duration `yaml:"timeout"`
}

func Load() *Config {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(path); os.IsExist(err) {
		log.Fatalf("config file is not exists: %s", path)
	}

	var cfg Config

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
