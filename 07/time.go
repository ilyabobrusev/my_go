package main

import (
	"fmt"
	"time"
)

func main() {
	//timeout :=3000
	var timeout time.Duration = 3000
	fmt.Println("before ")
	//time.Sleep(time.Duration(timeout) * time.Millisecond)
	time.Sleep(timeout * time.Millisecond)
	fmt.Println("after")
}
