package main

import "fmt"

func foo() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func sum(x, y int, c chan int) {
	c <- x + y
}

func main() {
	foo()
	c := make(chan int)
	go sum(24, 18, c)
	fmt.Println(<-c)
}
