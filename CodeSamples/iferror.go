package main

import (
	"fmt"
	"os"
)

func main() {

	homePath := os.Getenv("HOME")
	// fmt.Println("Home Path:", homePath)

	_, err := os.Open(homePath + "/update_dns_on_pip.rb")

	if err != nil {
		fmt.Println("Error returned:", err)
	}

	fmt.Println("Error returned:", err)
}
