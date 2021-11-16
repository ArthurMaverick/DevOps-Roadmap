package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10
	b := 5
	c := 13.5
	fmt.Println("soma:", a+b)
	fmt.Println("sub:", a-b)
	fmt.Println("mul:", a*b)
	fmt.Println("div:", a/b)
	fmt.Println("pow", a>>b) // div por 2 b vezes
	fmt.Println("pow", a<<b) // mult por 2 b vezes

	//bitwise
	fmt.Println("bitwise and:", a&b)      //10 & 5 = 0
	fmt.Println("bitwise or:", a|b)       //10 | 5 = 15
	fmt.Println("bitwise xor:", a^b)      //10 ^ 5 = 15
	fmt.Println("bitwise not:", ^a)       //not 10 = -11
	fmt.Println("bitwise and not:", a&^b) //10 &^ 5 = 15
	fmt.Println("bitwise or not:", a|^b)  //10 |^ 5 = 11

	// usando math
	fmt.Println("mul:", a*b)
	fmt.Println("math:", math.Sqrt(c))   // raiz quadrada
	fmt.Println("math:", math.Pow(c, 2)) // c elevado a 2
	fmt.Println("math:", math.Max(c, 4)) // maior valor
	fmt.Println("math:", math.Min(c, 4)) // menor valor
	fmt.Println("math:", math.Mod(c, 4)) // resto da divisão

}
