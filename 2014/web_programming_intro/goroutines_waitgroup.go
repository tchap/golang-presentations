package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i != 10; i++ {
		go func(v interface{}) {
			defer wg.Done(); fmt.Println(v)
		}(i)
	}
	wg.Wait()
}
