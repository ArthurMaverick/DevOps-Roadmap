package main

import "fmt"

func main() {
	i := 1
	var p *int = &i // p recebe o endereço de i
	fmt.Println(p)  // endereço de i
	fmt.Println(*p) // valor de i
	*p = 2          // valor de i agora é 2
	fmt.Println(i)  // 2

	// Go não suporta a passagem de ponteiros por valor.
	u := 2
	var x *int = nil
	x = &u // p recebe o endereço de
	*x++
	u++
	fmt.Println(x)  // endereço de x
	fmt.Println(*x) // valor de x
	fmt.Println(u)  // valor de u é  3
	fmt.Println(&u) // endereço de u é

}
