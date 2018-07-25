package main

import (
	"fmt"
  "syscall/js"
)

var done = make(chan struct{})
var num int

func main() {
  setTheText()
  setupKillWASM()
  setupDoitAgain()
  cb := endBeforeUnload()
  defer cb.Release()

  <- done
  fmt.Println("Bye WASM!")
}

func setTheText() {
  num++
  doc := js.Global().Get("document")
  container := doc.Call("getElementById", "app")
  js.ValueOf(container).Set(
    "innerHTML",
    fmt.Sprintf("For no good reason, I am setting this using GO!! (%d)",  num),
  )
  fmt.Println("I SET IT!")
}

func endBeforeUnload() js.Callback {
  cb := js.NewEventCallback(0, beforeUnload)
  addEventListener := js.Global().Get("addEventListener")
  addEventListener.Invoke("beforeunload", cb)
  return cb;
}

func beforeUnload(event js.Value) {
  alert := js.Global().Get("alert")
  alert.Invoke("Are you sure you want to leave?")
  done <- struct{}{}
}

func setupDoitAgain() {
  again := func (_ []js.Value) {
    setTheText()
  }
  js.Global().Set("setTheText", js.NewCallback(again))
}

func setupKillWASM() {
  kill := func (_ []js.Value) {
    done <- struct{}{}
  }
  js.Global().Set("killWASM", js.NewCallback(kill))
}


