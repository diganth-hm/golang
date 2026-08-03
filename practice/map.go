package main
import "fmt"

func main(){
	//declareing map map([key] value)
	student:=make(map[string]int)
    fmt.Println(student)
	student["samay"]=28
    fmt.Println(student)

	//another way
	students:=map[string]int{
		"sam":27,
		"samay":28
	}
}