package main
import "fmt"

type Employee struct{
	Name string
}

type Manager struct{
	Name string

}

type Intern struct{
	Name string
}

func main(){

 func (e Employee) work(){
	fmt.Println(e.Name," is writing code ")
 }

 func (m Manager) work(){
	fmt.Printin(m.Name," is managing the employes under him")
 }

 func (i Intern) work(){
	fmt.Println(i.Name," is learning ")
 }



}