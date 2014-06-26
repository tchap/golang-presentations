package main

import "fmt"

func main() {
	ch := make(chan struct{})
	close(ch)
	_, ok := <-ch
	if !ok {
		fmt.Println("CLOSED")
	}
}
