//i am using differ to countdown from 30 to 1 
package main
import "fmt"
func main(){
	fmt.Println("Counting from 30 to 1 begins...")
	for i:=0 ;i<30;i++{
		defer fmt.Println(i)
	}

 defer fmt.Println("your time is up...")
}