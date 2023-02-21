package main

import "fmt"

const a = 10

func main() {

	var b = 8

	fmt.Println("vim-go")
	fmt.Printf("Hello %s\n", "World")
	fmt.Print("Hello World!")

	fmt.Printf("2 + 5 = %v\n", 2+a)
	fmt.Printf("2 - 5 = %v\n", b-5)
	fmt.Printf("2 * 5 = %v\n", 2*5)
	fmt.Printf("2 / 5 = %v\n", 2/5)
	fmt.Printf("2 %% 5 = %v\n", 2%5)
}
