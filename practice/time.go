package main
import(
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to my pratice session")

	presenttime := time.Now()
	fmt.Println("The time now is = ",presenttime)
//format of the date and time same every time you use it anywhere
    fmt.Println(presenttime.Format("01-02-2006 Monday"))
}