package main

import(
	"fmt"
)

// Fizz Buzz
func main(){
	for i := 1; i <= 100; i++{
		printNum := true
		if i % 3 == 0{
			fmt.Printf("Fizz")
			printNum = false
		}
		if i % 5 == 0{
			fmt.Printf("Buzz")
			printNum = false
		}
		if printNum {
			fmt.Printf("%v", i)
		}
		fmt.Printf("\n")
	}
}