package main

import (
	"github.com/jceatwell/bankHexArch/app"
	"github.com/jceatwell/bankHexArch/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
