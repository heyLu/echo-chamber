package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/echo", requestLog(handleEcho))
	http.HandleFunc("/404", requestLog(handleNotFound))

	addr := "localhost:12345"
	fmt.Printf("Listening on http://%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		exit("Starting server", err)
	}
}

func handleEcho(w http.ResponseWriter, req *http.Request) {
	err := req.Write(w)
	if err != nil {
		warn("/echo", err)
	}
}

func handleNotFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

type Handler func(w http.ResponseWriter, req *http.Request)

type recordWriter struct {
	http.ResponseWriter

	StatusCode int
}

func newRecordWriter(w http.ResponseWriter) *recordWriter {
	return &recordWriter{w, 200}
}

func (rw *recordWriter) WriteHeader(s int) {
	rw.StatusCode = s
	rw.ResponseWriter.WriteHeader(s)
}

func requestLog(handler Handler) Handler {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		rw := newRecordWriter(w)
		handler(rw, req)
		took := time.Since(start)

		fmt.Printf("%s  %s %s  %d  (took %s)\n", start.UTC().Format(time.RFC3339), req.Method, req.URL.Path, rw.StatusCode, took)
	}
}

func warn(msg string, err error) {
	fmt.Fprintf(os.Stderr, "Warning: %s: %s\n", msg, err)
}

func exit(msg string, err error) {
	fmt.Fprintf(os.Stderr, "Error: %s: %s\n", msg, err)
	os.Exit(1)
}
