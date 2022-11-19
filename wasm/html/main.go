package main

import (
	"syscall/js"
    // "time"
    // "fmt"
    "nuxui.org/nuxui/log"
    "nuxui.org/nuxui/nux"
    // _ "nuxui.org/nuxui/ui"
)

func init(){
    println("golang init") // expecting 5
}

type dd struct {

}

// This calls a JS function from Go.
func main() {
    log.I("nuxui", "hhhh = %d", 32)

    doc := js.Global().Get("document")
    c := doc.Call("getElementById", "myCanvas")
    canvas := c.Call("getContext", "2d")
    canvas.Set("font", "30px Arial")

    nux.App().Init(manifest)
	nux.App().Run()
}