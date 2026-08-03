//Arrays and structs inside arrary and use of loops and range

package main
import "fmt"

//decelaring a struct
type Employee struct{
	Name string
	Id   int
	Salary float64

}

func main(){
	var high_sal float64 = 0
	var high_sal_emp string
	
	employee := [...] Employee{
		{Name:"Samay" ,Id:322,Salary:57345.646 },
		{Name:"Karan" ,Id:563,Salary:60246.326 },
		{Name:"Tanmay" ,Id:347,Salary:54675.245 },
		
	}
	fmt.Println("\nThe details of the Employees are")
	for i, emp := range employee{
		
		fmt.Println("\nEmployee no:",i+1)
		fmt.Println("Name :",emp.Name)
		fmt.Println("ID : ",emp.Id)
		fmt.Println("Salary : ",emp.Salary)
		if emp.Salary > high_sal{
			high_sal=emp.Salary
			high_sal_emp = emp.Name
		}
	}
    fmt.Println("\nThe Employee with highest salary is : ",high_sal_emp)
}