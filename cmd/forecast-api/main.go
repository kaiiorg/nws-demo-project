package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/kaiiorg/nws-demo-project/internal/api"
	"github.com/kaiiorg/nws-demo-project/internal/config"

	"github.com/rs/zerolog/log"
)

var (
	configPath = flag.String("config-path", "", "Path to config file")
)

func main() {
	flag.Parse()
	configureLogging()
	c := config.LoadConfig(*configPath)

	go func() {
		api.Run(c)
	}()

	waitForInterrupt()
}

func waitForInterrupt() {
	log.Warn().Msg("ctrl+c to exit")
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan
	log.Warn().Str("signal", sig.String()).Msg("Received signal, exiting")
}
