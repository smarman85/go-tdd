package main

import "fmt"

const german = "German"
const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const germanHelloPrefix = "Hallo, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
  if name == "" {
    name = "World"
  }

  return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

  switch language {
  case french:
    prefix = frenchHelloPrefix
  case german:
    prefix = germanHelloPrefix
  case spanish:
    prefix = spanishHelloPrefix
  default:
    prefix = englishHelloPrefix
  }

  return
}

func main() {
  fmt.Println(Hello("Sean", "German"))
}
