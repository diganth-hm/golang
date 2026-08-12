package main
import(
	"fmt"
	"os"
	"io"
	
)

func main(){
   
    fmt.Println("welcome to files in golang ")
	content := "This needs to go in the new file that will be created "

	//now creating a file using os pkg

	file,err := os.Create("./newfile.txt")
	if err != nil {
		panic(err)
	}
   //io pkg is useed to write in to the file
   // i am using length to find the len of the file ,by this i would know it worked or not
    lenght,err := io.WriteString(file,content)
     if err != nil{

		panic(err)
	 }
     fmt.Println("Lenght of the file is : ",lenght)
	 defer file.Close()
	 fmt.Println("\nnow reading a file ")
	 readFile("myfile.txt")
}


func readFile(filename string){
	data,err := os.ReadFile(filename)
    if err != nil{
		panic(err)
	}
	fmt.Println("The text data inside the file is \n",string(data))
}