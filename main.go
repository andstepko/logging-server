package main

import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
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

    w.WriteHeader(200)
}

func main() {
    for _, url := range urls {
        http.HandleFunc(url, handler)
    }

    err := http.ListenAndServeTLS(fmt.Sprintf("%s:%d", ADDRESS, PORT), CERTIFICATE_PATH, KEY_PATH, nil)
    if err != nil && !MUST_SSL {
        // Try to run without SSL
        err = http.ListenAndServe(fmt.Sprintf("%s:%d", ADDRESS, PORT), nil)
    }

    fmt.Printf("Error running server==>%s\n", err)
}
