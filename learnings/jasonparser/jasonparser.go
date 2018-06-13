package main

import (
        "fmt"
        "encoding/json"
        "io/ioutil"
        "os"
)

type Name struct {
  Name string //`json:"name"`
}

func (p Name) toString() string {
    return toJson(p)
}

func toJson(p interface{}) string {
    bytes, err := json.Marshal(p)
    if err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
    }

    return string(bytes)
}

func main() {

    names := getNames()
    for _, p:= range names {
      fmt.Println(p.toString())
    }
    fmt.Println(toJson(names))
}

func getNames() []Name {
  raw, err := ioutil.ReadFile("./raw.json")
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
  var c []Name
  json.Unmarshal(raw, &c)
  return c
}
