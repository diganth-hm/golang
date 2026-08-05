package main
import "fmt"

type Creditcard struct{
	 
	Cardnumber string
}

type Cryptowallet struct{

	Walletaddress string
}

func (c Creditcard ) propayment(amount float64) string{
	return fmt.Sprintf("Processing %.2f  using creditcard ending in %s ",amount,c.Cardnumber)
}

func (w Cryptowallet) propayment(amount float64) string{
	return fmt.Sprintf("Processing %.2f using a Walletaddress ending with %s ",amount,w.Walletaddress)
}

type paymethod interface{
	propayment(amount float64) string
}

func  checkout(p paymethod ,amount float64){
	result:= p.propayment(amount)
	fmt.Println(result)
}

func main() {
	
	myCard := Creditcard{Cardnumber: "4564-4334-7886-2386"}
	myWallet := Cryptowallet{Walletaddress: "7vttc3475t3v4785tc3473r4by"}

	checkout(myCard, 99.99)
	checkout(myWallet, 45.50)
}