package main

import (
	"fmt"
	"log"
)

func main(){

	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker(){
	fmt.Println("about to panic")
	defer func(){
		if err:= recover(); err != nil {
			log.Println("Error:", err)
			panic(err)
		}
	}()

	panic("Something Happened")
	fmt.Println("done panicking")
}