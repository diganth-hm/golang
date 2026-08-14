package main
import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"net/url"
)


func checkerr(err error){
	if err != nil{
		panic(err)
	}
}

func PerformGet(){
	fmt.Println("Performing get request on local host")
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

func PerfornmPostJson(){
	fmt.Println("Performing  Post request on local host")
	const url="http://localhost:3000/post"
	//fake json payload
	requestBody := strings.NewReader(`
	 {
		"coursename":"Golang",
		"price":6744,
		"hours":45
    }	
	 `)
	 response,err := http.Post(url,"application/json",requestBody)
	 checkerr(err)
	 defer response.Body.Close()
	 fmt.Println("status code : ",response.StatusCode)
	 content,err := io.ReadAll(response.Body)
	 checkerr(err)
	 fmt.Println("The request sent was : \n",string(content))
}

func PerfromPostFormreq(){
	fmt.Println("Performing PostForm request on local host")
	const myurl = "http://localhost:3000/postform"

	data := url.Values{}
	data.Add("firstname","Diganth")
	data.Add("lastname","HM")
	data.Add("Email","lonny23@Ggmail.com")

	response,err := http.PostForm(myurl,data)
	checkerr(err)
	defer response.Body.Close()
	content,err := io.ReadAll(response.Body)
	checkerr(err)
	fmt.Println("Status code : ",response.StatusCode)
	fmt.Println("content in the req is : \n",string(content))
}

func main(){
	
	//PerformGet()
	//PerfornmPostJson()
	PerfromPostFormreq()
}