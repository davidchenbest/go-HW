package main

import(
"fmt"
)

func main()  {
  //initialize
  // var names map[string]string
  // var names = make(map[string]string)
  var names = map[string]string {
    "0": "John",
    "1": "Jane",   // last comma is a must
}

  names["david"] = "david"

  name,exist := names["david"]

  delete(names, "1")

  fmt.Println(name,exist)

  for _, name := range names {
    fmt.Println(name)
  }


}