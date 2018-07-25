package main

import (
	"fmt"
  "syscall/js"
  "net/http"
  "io/ioutil"
)

var done = make(chan struct{})
var num int

func main() {
  setTheText()
  setupKillWASM()
  setupDoitAgain()
  cb := endBeforeUnload()
  defer cb.Release()

  getTheData()

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

func getTheData() {
  resp, err := http.Get("https://www.jsonstore.io/33abd1336ed1cb35f157ef70ef0958f53eae56696bb7be46377364fb356d20b8/")

  if err != nil {
    fmt.Printf("%s", err)
  } else {
    defer resp.Body.Close()
    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      fmt.Printf("%s", err)
    }
    result := string(contents)
    fmt.Printf("RESULT -> %s\n", result)
    js.Global().Set("scott", result)
  }
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


