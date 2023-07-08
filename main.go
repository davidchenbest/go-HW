package main

import(
"fmt"
)

func main()  {
  names := []string {"david","susan"}
  fmt.Println(names)

  for i, name :=range names{
    fmt.Println(name,i)
  }
}