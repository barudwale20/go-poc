package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Health struct {
	l *log.Logger
}

func NewHealth(l *log.Logger) *Health {
	return &Health{l}
}

func (h *Health) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// data, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	http.Error(rw, "Ohh no!", http.StatusBadRequest)
	// 	return
	// }
	fmt.Fprintf(rw, "Healthy and running!")
}
