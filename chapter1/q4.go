package main

//fmt is short for format? That would make sense!
import(
	"fmt"
	"unicode/utf8"
)


func main(){
	fmt.Printf("Part One:\n")
	partOne()
	fmt.Printf("Part Two:\n")
	partTwo()
	fmt.Printf("Part Three:\n")
	partThree()
	fmt.Printf("Part Four:\n")
}

//make a traingle out of the letter A that goes from 1 to 100 characters across.
func partOne(){
	baseString := "A"
	for len(baseString) <= 100{
		fmt.Printf(baseString + "\n")
		baseString += "A"
	}
}

//count length and bytes
func partTwo(){
	check := []byte("asSASA ddd dsjkdsjs dk")
	fmt.Printf("Runes: %v\n", utf8.RuneCount(check))
	fmt.Printf("Bytes: %v\n", len(check))
}

func partThree(){
	check := []byte("asSASA ddd dsjkdsjs dk")
	copy(check[4:], "abc")
	fmt.Printf(string(check))
}

func partFour(){
	
}

