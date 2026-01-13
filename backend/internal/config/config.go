package config

import "os"

type Config struct {
	Addr string
}

func ConfigFromEnv() Config {
	config := Config{
		Addr: os.Getenv("LISTEN_ADDR"),
	}
	return config
}
