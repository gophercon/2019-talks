package main

import "syscall/js"

func setDemo() {
	doc := js.Global().Get("document")
	d := doc.Call("getElementById", "demo")
	d.Set("innerHTML", "Hello <b>Go!</b>")
}
