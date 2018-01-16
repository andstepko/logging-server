package main

import (
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const Delimiter = "=============================================================================================="

func handler(w http.ResponseWriter, r *http.Request) {
	reqPrint := buildRequestPrint(r)

	fmt.Println(Delimiter)
	fmt.Println(reqPrint)
	fmt.Println(Delimiter)

	err := saveRequest(r.Method, reqPrint)
	if err != nil {
		logan.New().WithError(err).Error("Failed to save request to file.")
	}

	w.WriteHeader(http.StatusOK)
}

func buildRequestPrint(r *http.Request) string {
	result := r.Method + "\n\n"

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			result += fmt.Sprintf("%v: %v\n", name, h)
		}
	}
	result += "\n"

	body, _ := ioutil.ReadAll(r.Body)
	result += string(body)

	return result
}

func saveRequest(method, text string) error {
	fileName := fmt.Sprintf("./%s %s", method, time.Now().Format("Mon, 02 Jan 2006 15:04:05.000 -0700"))

	err := ioutil.WriteFile(fileName, []byte(text), 0644)
	if err != nil {
		return errors.Wrap(err, "Failed to write file", logan.F{
			//"folder_name": folderName,
			"file_name":   fileName,
		})
	}

	return nil
}
