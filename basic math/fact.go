package main
import "fmt"

func main(){
	var fac float64 = 1
	var i float64
	var n float64
	fmt.Println("Enter the number to obtain its factorial ")
	fmt.Scan(&n)
	if n<1 {
		fmt.Println("The number should be greater than 0")
	}else{
		for i=2;i<=n;i++{
			fac=fac*i
		}
		fmt.Println("The factorial of ",n,"is = ",fac)
	}

}