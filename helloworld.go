package main

import (
	"fmt"
  "syscall/js"
)

func main() {
  setTheText()
}

func setTheText() {
  doc := js.Global().Get("document")
  container := doc.Call("getElementById", "app")
  js.ValueOf(container).Set(
    "innerHTML",
    "For no good reason, I am setting this using GO for the %d time",
  )
  fmt.Println("I SET IT!")
}

