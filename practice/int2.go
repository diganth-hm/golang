package main
import "fmt"

type Email struct{
	Address string
}

type Sms struct{
	PhoneNumber string
}

func (e Email) send() string{
	return "Sending  Email to "+ e.Address
}

func (s Sms) send() string{
	return "Sending SMS to "+ s.PhoneNumber
}

type Notifier interface{
	send() string
}

func Alertuser(n Notifier){
	fmt.Println(n.send())
}

func main(){
 myemail:= Email{Address : "user234@gmail.com" }
 myph:= Sms{ PhoneNumber :"2345645645"}
    Alertuser(myemail)
	Alertuser(myph)

}