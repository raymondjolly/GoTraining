package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main(){

	res, err := http.Get("http://google.com/robots.txt")

	if err !=nil {
		log.Fatal(err)
	}

	defer res.Body.Close()  //protective programming puts Close statement first but defers to end to prevent forgetting.
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
