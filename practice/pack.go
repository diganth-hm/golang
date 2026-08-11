package main
import (
	"fmt"
	"os"
	"bufio"
)

func main(){

	fmt.Println("Please rate the pizza out of 10")
	//using this to take whole sentances as input
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us a := ",input)
}