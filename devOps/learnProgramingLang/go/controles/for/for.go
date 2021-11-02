package main

import (
	"fmt"
	"time"
)

func for1() {
	for i := 0; i < 10; i++ { // variavel, condicao, incremento
		fmt.Println(i)
	}
}

func for2() {
	i := 0       // variavel
	for i < 10 { // condicao
		fmt.Println(i)
		i++ // incremento
	}
}

func for3() {
	i := 0 // variavel
	for {
		fmt.Println(i)
		i++          // incremento
		if i == 10 { // condicao
			break
		}
	}
}

func for4() {
	for {
		fmt.Println("Loop infinito")
		time.Sleep(time.Second)
	}
}

func main() {
	for1()
	for2()
	for3()
	for4()
}
