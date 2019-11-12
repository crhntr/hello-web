package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	stderr := log.New(os.Stderr, "", 0)
	stdout := log.New(os.Stdout, "", 0)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		stdout.Printf("%s -> %s: %s", req.Host, req.Method, req.URL)
		res.Write([]byte("Hello, world!"))
	})); err != nil {
		stderr.Println(err)
	}
}
