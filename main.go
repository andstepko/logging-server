package main

import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
	"time"
	"gitlab.com/distributed_lab/logan/v3"
	"net"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("HEADERS:")
    for name, headers := range r.Header {
        name = strings.ToLower(name)
       for _, h := range headers {
            fmt.Printf("%v: %v\n", name, h)
       }
     }

    body, _ := ioutil.ReadAll(r.Body)
    fmt.Println("\nBODY:")
    fmt.Println(string(body) + "\n\n")

    w.WriteHeader(http.StatusOK)
}

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
