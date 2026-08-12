package main
import "fmt"

func main(){
	students := make(map[string]int)
	students["samay"]=98
	students["karan"]=99
	fmt.Println(students["karan"])
   for key,value := range students{

	fmt.Printf("The marks obtained by %s is = %d percent \n",key,value)
   }

}
