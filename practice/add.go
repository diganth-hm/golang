package main 
import "fmt"

func main(){
//declaring variables 
	var a int =10
	var b int =155
	var c int = a+b
	fmt.Println("The sum of two numbers is =",c)

//another way of  declaring variables 
   
    var d,e,f int =14,14,25
	var sum int = d+f+e
	fmt.Println("The sum of the 3 numbers is =",sum)
 //There is also an SHORT DECLARATION way 

    name := "Diganth"
	age := 19
	fmt.Println("The name of the student is =",name)
	fmt.Println("The age of the student is =",age)
//constants 
    const num int = 1542
	


}