package main
import "fmt"
func main() {
	var r rune = 'a'
	fmt.Println("Rune : %c,unicode :%U\\n",r,r,r)
//converting strings into rune 

	 str := "hello"
	 runes := []rune(str)
    fmt.Println("Runes:", runes)
    fmt.Println("Number of runes:", len(runes))




	}