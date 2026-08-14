package main
import (
	"fmt"
	"io"
	"net/http"
	"strings"
)


func checkerr(err error){
	if err != nil{
		panic(err)
	}
}

func PerformGet(){
	const url="http://localhost:3000/get"
	response,err := http.Get(url)
	checkerr(err)
	defer response.Body.Close()
	fmt.Println("Status code :",response.StatusCode)
	fmt.Println("Content length : ",response.ContentLength)
	fmt.Println("Header : ",response.Header)
	content,err := io.ReadAll(response.Body)
	checkerr(err)
	fmt.Println("conversion method \n The content is : \n",string(content))
//we can aslo use strings pkg to convert the response to string using a builder
   var responseString strings.Builder
   _,err = responseString.Write(content)
   checkerr(err)
   fmt.Println("using strings pkg  and builder \n The content is ",responseString.String())
}

func main(){
	fmt.Println("Performing get request on local host")
	PerformGet()
}