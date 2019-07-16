package main

import "fmt"

const en = "en"
const es = "es"
const fr = "fr"

const enprefix = "Hello, "
const esprefix = "Hola, "
const frprefix = "Bonjour, "

func prefix(lang string) (prefix string) {
  switch lang {
  case en: prefix = enprefix
  case es: prefix = esprefix
  case fr: prefix = frprefix
  default: prefix = enprefix
  }
  return
}

func Hello(name string, lang string) string {
  if name == "" {
    name = "world"
  }

  return prefix(lang) + name
}

func main() {
  fmt.Println(Hello("Collin", en))
}