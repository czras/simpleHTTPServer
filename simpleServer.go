package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	port := ""
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
	} else {
		port = ":8080"
	}
	log.Printf("Server started at http://localhost%s serving files from %s\n", port, pwd)
	log.Fatal(http.ListenAndServe(port, logHandler(http.FileServer(http.Dir(pwd)))))
}

type ResponseWriter struct {
	http.ResponseWriter
	status   int
	bodySize int
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.bodySize += len(data)
	return w.ResponseWriter.Write(data)
}

func logHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &ResponseWriter{ResponseWriter: w, status: 200}
		h.ServeHTTP(rw, r)
		bodySize := rw.Header().Get("Content-Length")
		if len(bodySize) == 0 {
			bodySize = strconv.Itoa(rw.bodySize)
		}
		log.Printf("Request{url: %s, method: %s, remote: %s} Response{status: %d, content-type: %s, content-length: %s, time: %s}",
			r.URL.RequestURI(), r.Method, r.RemoteAddr, rw.status, rw.Header().Get("Content-Type"), bodySize, time.Since(start))
	}
}
