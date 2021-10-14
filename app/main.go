// go:build js

package main

import "syscall/js"

func main() {
	println("Hello, world!")
	js.Global().Get("document").Call("querySelector", "body").Set("innerHTML", "Hello, world!")
}
