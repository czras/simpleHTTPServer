package main

import (
	"log"
	"net/http"
	"os"
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
	log.Println("Server started at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(pwd))))
}
