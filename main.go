package main

import (
	"fmt"
	"syscall/js"
)

func exposedFunction(this js.Value, inputs []js.Value) interface{} {
	fmt.Println("Exposed Function Executed!")
	return nil
}

func main() {
	// channel to keep the wasm running
	c := make(chan int)
	fmt.Println("Go WebAssembly")

	js.Global().Set("exposedFunction", js.FuncOf(exposedFunction))

	document := js.Global().Get("document")

	hello := document.Call("createElement", "h1")
	hello.Set("innerText", "Go WebAssembly")
	document.Get("body").Call("appendChild", hello)
	// value for c is never send
	<-c
}
