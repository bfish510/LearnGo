package main

import(
	"fmt"
)

func main(){
	fmt.Printf("Part One\n")
	partOne()
	fmt.Printf("Part Two\n")
	partTwo()
	fmt.Printf("Part Three\n")
	partThree()
	fmt.Printf("Complete")
}

func partOne(){
	for i := 0; i < 10; i++{
		fmt.Printf("i = %v\n", i)
	}
}

func partTwo(){
	i := 0
Loop:
	fmt.Printf("i = %v\n", i)
	i++
	if i < 10{
		goto Loop
	}
}

func partThree(){
	var array [10]int
	for i := 0; i < 10; i++{
		array[i] = i
	}
	fmt.Printf("Array: %v\n", array)
}