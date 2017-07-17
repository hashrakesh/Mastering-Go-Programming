package main

import "fmt"

func main() {
	fmt.Println("Counting")
	for i := 1; i <= 4; i++ {
		fmt.Println(i) // print normaly i.e. FIFO
		defer fmt.Println(i)  // output in reverse i.e. LIFO
	}
	fmt.Prinln("Done")
}
