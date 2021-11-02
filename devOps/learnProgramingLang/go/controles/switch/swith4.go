package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case func():
		return "func"
	}
	return "desconocido"
}

func main() {
	fmt.Println(tipo(1))
	fmt.Println(tipo("hola"))
	fmt.Println(tipo(true))
	fmt.Println(tipo(func() {}))
}
