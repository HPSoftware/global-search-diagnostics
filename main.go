package main

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func main() {
	const PORT string = "8000"
	http.Handle("/", http.FileServer(http.Dir("./app/web/")))
	http.HandleFunc("/health", health)
	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
	if err != nil {
		log.Errorf("Failed to start global search diagnostics on port: %s, error: %v", PORT, err)
	}
}

func health(writer http.ResponseWriter, request *http.Request) {
	retStatus := http.StatusOK
	writer.WriteHeader(retStatus)
}
