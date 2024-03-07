package main

import (
	addtwonumbers "Day1/add_two_numbers"
	evenodd "Day1/even_odd"
	helloworld "Day1/hello_world"
	"fmt"
)

func main() {
	sum := addtwonumbers.Add(5, 6)
	fmt.Println(helloworld.Hello_world())
	check := evenodd.Even_odd(5)
	if check {
		fmt.Println("Given number is Even")
	} else {
		fmt.Println("Given number is odd")
	}
	fmt.Printf("Sum of 5 and 6 is %d", sum)
}
