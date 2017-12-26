package main

import (
	"fmt"

)

func main(){

	fmt.Println("Start")
	defer fmt.Println("this was deferred")  //Defer will always execute action last or LIFO if using multiple defers
	panic("Something Bad happened") // not that defered will execute BEFORE the panic so that resources can be closed
	fmt.Println("End")

}
