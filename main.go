package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	log.Println("starting server on port", port)
	err := http.ListenAndServe(":"+port, http.HandlerFunc(hello))
	if err != nil {
		log.Fatal(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	log.Printf("%s | %s -> %s: %s", time.Now(), req.Host, req.Method, req.URL)
	_, _ = res.Write([]byte("Hello, world!"))
}
