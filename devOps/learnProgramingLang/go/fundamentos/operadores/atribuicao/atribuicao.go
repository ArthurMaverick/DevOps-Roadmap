package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i *= 3
	i /= 3
	i %= 3
	i <<= 3 // i = i * 2^3
	i >>= 3 // i = i / 2^3
	i &= 3  // i = i mod 2
	i ^= 3  // i = i xor 2
	i |= 3  // i = i or 2
	fmt.Println(i)

	c, e := 3, 2
	fmt.Println(c, e)
	c, e = e, c
	fmt.Println(c, e)
}
