package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	port := ""
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
	} else {
		port = ":8080"
	}
	log.Println("Server started at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(pwd))))
}
