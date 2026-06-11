package main
import "fmt"
func main(){
	var n int =10
	
//for loop example
i:=0
	for i=0;i<n;i++{
		fmt.Println(i)
	}
//while loop example
j:=0
   for j < n {
	fmt.Println("This is while loop running on the ",j," time")
	j++
   }

//do-while loop
k:=0
for {
	fmt.Println("do while loop runing on the ",k," time")
	k++
	
	if k>n{
break }
 
}
}