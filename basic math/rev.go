package main
import "fmt"

func main(){

	var num int
	rem :=0
	rev :=0
	
	fmt.Println("Enter the number to be reversed ")
	fmt.Scan(&num)
	for num!=0 {
		rem=num%10
		rev=(rev*10)+rem
		num=num/10

	}
	fmt.Println("The reverse of the given number is = ",rev)


}