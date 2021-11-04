package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}   //array
	s1 := []int{1, 2, 3, 4} //slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5} //array
	s2 := a2[1:4]               //slice of array a2
	fmt.Println(a2, s2)

	s3 := a2[:2] //slice of array a2
	fmt.Println(s3)

	a4 := []bool{true, false, true, false}
	s4 := a4[1:3]
	fmt.Println(s4)

	//? ps: slice nao é array

}
