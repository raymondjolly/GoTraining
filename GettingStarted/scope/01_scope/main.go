package _1_scope

import "fmt"

var x = 42

func main(){

	foo()
}

//function calling another function.
func foo(){
	fmt.Println(x)
}
