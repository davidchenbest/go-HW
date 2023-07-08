package main

import(
  "fmt"
)

func main()  {
  names := []string{"susan","david"}
  names = append(names,"gina")
  l:=len(names)
  fmt.Println(l)
  fmt.Println(names[1:])
}