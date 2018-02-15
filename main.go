package main

import (
    "fmt"
    "net/http"
	"time"
	"gitlab.com/distributed_lab/logan/v3"
	"net"
	"os"
)

func main() {
	log := logan.New()

	args := os.Args[1:]

	var configFileName string
	if len(args) < 1 {
		configFileName = "config"
	} else {
		configFileName = args[0]
		if len(configFileName) > 5 && configFileName[len(configFileName) - 5:] == ".yaml" {
			configFileName = configFileName[:len(configFileName) - 5]
		}
	}

	config, err := ReadConfig(configFileName)
	if err != nil {
		panic(err)
	}

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Address, config.Port))
	if err != nil {
		panic(err)
	}

	err = http.ServeTLS(ln, http.HandlerFunc(handler), config.CertificatePath, config.KeyPath)
    if err != nil && !config.SSLOnly {
    	log.WithError(err).Warn("Failed to run server with SSL.")

        time.Sleep(1 * time.Second)

        log.Info("Trying to run server without SSL.")
        err = http.Serve(ln, http.HandlerFunc(handler))
    }

    log.WithError(err).Error("Error running server.")
}
