package main

import (
	"github.com/MateuSoares1/gopportunities/config"
	"github.com/MateuSoares1/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	//Initialize router
	router.Initialize()
}
