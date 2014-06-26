package main

import "fmt"

func main() {
	ch := make(chan struct{})
	go func() { fmt.Println("Hello"); close(ch) }()
	<-ch
}
