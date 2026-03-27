package main

import (
	"flag"

	"github.com/kaiiorg/nws-demo-project/internal/api"
	"github.com/kaiiorg/nws-demo-project/internal/config"
)

var (
	configPath = flag.String("config-path", "./config.json", "Path to config file")
)

func main() {
	flag.Parse()
	configureLogging()
	c := config.LoadConfig(*configPath)
	api.Run(c)
}
