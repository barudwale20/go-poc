package handlers

import (
	"log"
	"net/http"
	"time"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	time.Sleep(10 * time.Second)
	rw.Write([]byte("Bye!"))
}
