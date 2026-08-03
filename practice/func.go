package main
import "fmt"

func add(x,y int) int{
	return x+y
}
//the above func can also be written as 
var multi =func(x,y int)int {
     return x*y
}

func main(){
	var x,y int
	fmt.Println("Enter the values for addition of X and y ")
	fmt.Scan(&x,&y)
	fmt.Println("The addtion of x and y is = ",add(x,y))
	fmt.Println("The multiplication of x and y is = ",multi(x,y))

}