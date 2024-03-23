package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

//const CONFIG_PATH = "./config/local.yaml"

type Config struct {
	Env         string `yaml:"env"  env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	os.Setenv("CONFIG_PATH", "./config/local.yaml")

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}

func LoadEnv() (string, string, int, string, string) {
	err := godotenv.Load("C:/Users/DinarKhayrutdinov/GolandProjects/GolangQuizlet/.env")
	if err != nil {
		HandleError(err, true)
	}

	port, _ := strconv.Atoi(os.Getenv("port"))
	host := os.Getenv("host")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	return user, host, port, password, dbname
}
