package main

import "fmt"

//note that this is for familiarity and not a best practice for naming packages.
//goal is to become familiar with variable declaration

func main(){

	a := 10
	b := "golang"
	c := 3.14
	d := "Detroit, Mi"
	e := false

	//

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)

	fmt.Printf("%T \n", c)

	j, k, l, m := 0, "backpack", false,  "sports"

	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)

	var myName = "Ray"

	fmt.Println("Hello, ", myName)





}
