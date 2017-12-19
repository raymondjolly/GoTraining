package main

import "fmt"

func zero(x int){
	fmt.Printf("%p\n", &x) //address in func zero
	fmt.Println(&x)  //address in func zero
	x = 0
}


func main()  {
	x:= 5
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)  //address of x in main
	zero(x)	// should be zero
	fmt.Println(x) //should be 5


}
