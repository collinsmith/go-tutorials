package main

import "fmt"

const enprefix = "Hello, "

func Hello(name string) string {
  return enprefix + name
}

func main() {
  fmt.Println(Hello("Collin"))
}