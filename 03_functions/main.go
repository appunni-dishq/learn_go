package main

import "fmt"

func greeting(name string) string {
	return "Gello " + name
}

func main() {
    fmt.Println(greeting("Appunni"))
}