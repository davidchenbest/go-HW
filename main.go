package main

import(
  "time"
  "fmt"
  "sync"

)

var wg sync.WaitGroup

func main()  {
  wg.Add(1)
  Announce("hey", 1)
  wg.Wait()
}

func Announce(message string, delay time.Duration) {
  go func() {
      time.Sleep(delay)
      fmt.Println(message)
      wg.Done()
  }()  // Note the parentheses - must call the function.
}