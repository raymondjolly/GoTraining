package main

import "fmt"

func main(){

	for i, j :=0, 0; i<5; i, j = i+1 , j+1{ //note that using i++ or j++ will not work in this case
		fmt.Println(i, j)
	}
}
