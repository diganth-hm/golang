package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diganth-hm/golang/buildapi/routes"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("server is runnig .....")
	fmt.Println("Server runing on http://localhost:5454")
	//assing the route
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":5454", r))

}
