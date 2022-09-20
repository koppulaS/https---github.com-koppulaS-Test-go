package main

import (
	"fmt"
	"sync"
)

func main() {
	letter := make(chan bool)
	number:=make(chan bool)

	wg := sync.WaitGroup{}
	go func() {

		for ch := 'A'; ch <= 'Z'; ch += 1 {
			letter <- true
			fmt.Print(string(ch))
			<-number

		}
		close(letter)
	}()

	wg.Add(1)
	go func() {
		start := 1
		for {
			number <- true

			fmt.Print(start)
			start += 1

			_, ok := <-letter
			if ok == false {
				break
			}
		}
		wg.Done()
	}()

	<-number

	wg.Wait()
	fmt.Print("\n")

}
