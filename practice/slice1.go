package main
import "fmt"

type Employee struct{
	Name string
	Salary float64
}

func main(){
	//created empty slcice 
	employee := [] Employee{}
	employee=append(employee,Employee{Name:"Karan",Salary:63579.856},)
	employee=append(employee,Employee{Name:"Samay",Salary:62239.546},)
	employee=append(employee,Employee{Name:"Tanmay",Salary:57453.345},)
	employee=append(employee,Employee{Name:"carry",Salary:54867.346},)
	employee=append(employee,Employee{Name:"balraj",Salary:60534.677},)
	fmt.Println("The lsit of all the employees ")		
	var high_sal float64 = 0	
	var high_nam string
	for _,emp:= range employee{
		fmt.Println("\nName : ",emp.Name,
	                 "\nSalary : ",emp.Salary)
		if emp.Salary > high_sal{
			high_sal=emp.Salary
            high_nam=emp.Name
		}
	
		
	}
	fmt.Println("\nThe highest paid employee is : ",high_nam)

	for i := range employee{

		//giving a raise of 10% to all
		employee[i].Salary=employee[i].Salary+(employee[i].Salary*(10.00/100.00))
	}

	fmt.Println("\nThe updated lsit after the Raise was given ")
		for _,emp:= range employee{
		fmt.Println("\nName : ",emp.Name,
	                 "\nSalary : ",emp.Salary)
		}
		

}
