package main

import "fmt"

func main(){

Loop:
	for i:=1; i<=10; i++{
		for j:=1; j<=10; j++{
			fmt.Println(i*j)
			if i*j>=10{
				break Loop
			}
		}
	}
}
