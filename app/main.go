// go:build js

package main

import "github.com/crhntr/window"

func main() {
	println("Hello, world!")
	window.Document.Body().SetInnerHTML( /* html */ "<h1>Hello, world!</h1>")
}
