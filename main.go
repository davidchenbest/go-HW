package main

import(
"fmt"
"math"
)

type rect struct {
  width, height float64
}

func (r rect) area() float64 {
  return r.width * r.height
}

type cir struct {
  radius float64
}

func (c cir) area() float64 {
  return  math.Pi * c.radius * c.radius
}

type Shape interface{
  area() float64
}

func area(shape Shape) float64{
  return shape.area()
}

func main()  {

  rectangle := rect{
    width:2,
    height:3,
  }

  circle := cir{
    radius:2,
  }
  fmt.Println(area(rectangle),area(circle))

  //EMPTY INTERFACE
  var values []interface{}
  values = append(values, 42)
  values = append(values, "Hello")
  fmt.Println(values)
  
}