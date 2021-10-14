package main

import (
	"bytes"
	"embed"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"text/template"
)

//go:embed assets
var assets embed.FS

func main() {
	port := "8080"
	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
	log.Println("starting server on port", port)
	bp := NewBufferPool()

	mux := http.NewServeMux()
	mux.Handle("/assets/", wasmCT(http.FileServer(http.FS(assets))))
	mux.Handle("/", indexHandler(bp))

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

//go:embed assets/index.gohtml
var indexHTML string

func indexHandler(bp *BufferPool) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		t := template.Must(template.New("index.html").Parse(indexHTML))
		res.Header().Set("content-type", "text/html")
		buf := bp.Get()
		defer bp.Release(buf)
		err := t.Execute(buf, struct{}{})
		if err != nil {
			http.Error(res, "failed to render page", http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusOK)
		_, _ = io.Copy(res, buf)
	}
}

func wasmCT(handler http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			res.Header().Set("content-type", "application/wasm")
		}
		handler.ServeHTTP(res, req)
	}
}

type BufferPool struct {
	pool sync.Pool
}

func NewBufferPool() *BufferPool {
	var bp BufferPool
	bp.pool.New = func() interface{} {
		return new(bytes.Buffer)
	}
	return &bp
}

func (bp *BufferPool) Get() *bytes.Buffer {
	return bp.pool.Get().(*bytes.Buffer)
}

func (bp *BufferPool) Release(buf *bytes.Buffer) {
	buf.Reset()
	bp.pool.Put(buf)
}
