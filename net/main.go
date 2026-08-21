package main

import (
	"fmt"
	"log"
)

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	fmt.Println()
}
