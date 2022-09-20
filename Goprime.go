package main

import (
	"fmt"
)

// Generate natural seri number: 2,3,4,...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter: delete the number which is divisible by a prime number to find prime number
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural()
	for i := 2; i < 10; i++ {
		prime := <-ch

		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)

	}
}
