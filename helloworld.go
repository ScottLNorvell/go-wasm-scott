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

var colors = getColors()

func setTheText() {
  doc := js.Global().Get("document")
  container := js.ValueOf(doc.Call("getElementById", "app"))
  container.Set(
    "innerHTML",
    fmt.Sprintf("For no good reason, I am setting this using GO!! (%d)",  num),
  )
  doc.Get("body").Get("style").Set("backgroundColor", colors[num % len(colors)])
  num++
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


func getColors() (colors [7]string) {
  colors[0] = "red"
  colors[1] = "orange"
  colors[2] = "yellow"
  colors[3] = "green"
  colors[4] = "blue"
  colors[5] = "indigo"
  colors[6] = "violet"
  return
}


