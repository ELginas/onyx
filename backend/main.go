package main

import (
	"github.com/elginas/onyx-backend/internal/config"
	"github.com/elginas/onyx-backend/internal/http"
)

func main() {
	config := config.ConfigFromEnv()
	http.Run(config)
}
