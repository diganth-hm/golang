//methods 
package main

import "fmt"


type person struct {
    name string
}

func (p *person) changeName(newName string) {
    p.name = newName
}

func main() {
    a := person{name: "a"}
    
    fmt.Println("Before:", a.name)
    a.changeName("b")
    
    fmt.Println("After:", a.name)
}