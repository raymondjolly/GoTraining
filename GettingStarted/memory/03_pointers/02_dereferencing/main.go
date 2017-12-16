package main

import "fmt"

func main(){

	a:=45

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b) //will provide address in memory
	fmt.Println(*b) //gives value at memory address

	*b = 42
	fmt.Println(a)
	//fmt.Println(*b)

	//	the code above makes b a pointer to the memory address where an int is stored
	//  b is of type "int" pointer
	// b points to the memory where address where an int is stored
	// to see the value in the memory address, add * in front of the variable
	// this is know as dereferencing
	// the * is an operator in this case.

}
