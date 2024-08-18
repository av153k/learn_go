package app

import "fmt"

func hello(name string) string {
	return "Hello, " + name + "!"
}

func PrintHelloWorld() {
	fmt.Println(hello("World"))
}
