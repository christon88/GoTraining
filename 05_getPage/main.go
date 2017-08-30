package main

import(
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
  res, _ := http.Get("http://www.vg.no/")
  page, _ := ioutil.ReadAll(res.Body)
  res.Body.Close()
  fmt.Printf("%s", page)
}
