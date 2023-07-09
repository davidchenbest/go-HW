package main

import(
"fmt"
)

type User struct{
  Name string
  Age int
}

func main()  {
  david := User{
    Name:"david",
    Age : 2,
  }

  users := []User{
    david,
  }

  var p *User
  p = &david

  p.Name = "s"
  p.Age =4

  fmt.Println(p, david,users)
}