package main

import (
	"github.com/Erick-VP/godin/config"
	"github.com/Erick-VP/godin/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("failed to initialize config: %v", err)
		return
	}

	router.Initialize()
}
