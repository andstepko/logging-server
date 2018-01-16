package main

import (
    "fmt"
    "net/http"
	"time"
	"gitlab.com/distributed_lab/logan/v3"
	"net"
)

func main() {
	log := logan.New()

	// TODO Read file name from args
	config, err := ReadConfig("config")
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
