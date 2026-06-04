package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type DB struct {
	Port   int    `yaml:"port"`
	Host   string `yaml:"host"`
	Name   string `yaml:"name"`
	User   string `yaml:"user"`
	SSLMod string `yaml:"ssl_mode"`
}

type Config struct {
	Password string `env:"DB_PASS"`
	DB       DB     `yaml:"db"`
}

func MustLoad() *Config {
	//Read env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Load config
	var c Config
	err = cleanenv.ReadConfig("config.yml", &c)
	if err != nil {
		log.Fatal("Error loading config")
	}
	return &c
}
