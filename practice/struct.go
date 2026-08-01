package main
import "fmt"

//declaring a struct 
type Employee struct{
	 Name string
	 Id  int
	 Salary float64

}

//declaring a method
func (e Employee) Display(){

	fmt.Println("Employee Name is : ",e.Name)
	fmt.Println("Employee ID No is : ", e.Id)
	fmt.Println("Employee Salary is = ",e.Salary)

}

func (e *Employee) GiveRaise(percent float64){

	e.Salary = e.Salary + (e.Salary*(percent/100))
}

func main(){

	//initialising of struct 
	var emp1 Employee
	emp1.Name ="Sam"
	emp1.Id =43
	emp1.Salary=56000.654
	var emp2 Employee
	emp2.Name ="Jhon"
	emp2.Id =46
	emp2.Salary=53000.654
   //use of method
	emp1.Display()
      fmt.Println("\nThe salary of ",emp2.Name," Before the Raise")
	  emp2.Display()

	emp2.GiveRaise(19.8)
        fmt.Println("\nThe salary  of ",emp2.Name," After the Raise")
	emp2.Display()

}