package main
import (
	"fmt"
	"net/http"
	"io"
)

const url ="https://google.com"

func checkerr(err error){
	if err != nil {
	 panic(err)
	}
}

func main(){
	fmt.Println("Get web Request")
	response,err := http.Get(url)
    checkerr(err)
	//u should alway close the body at the last never forget it 
	defer  response.Body.Close()
	databytes,err := io.ReadAll(response.Body)
	checkerr(err)
	data:= string(databytes)
	fmt.Println("The data in the url is \n",data)

}

