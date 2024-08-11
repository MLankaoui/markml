package main

import (
	"fmt"

	"marouane.com/markml/utils"
)

func main() {
	// headers tests
	fmt.Println(utils.HeaderParser("# WELCOME TO OUR PAGE"))
	fmt.Println(utils.HeaderParser("## WELCOME TO OUR PAGE"))
	fmt.Println(utils.HeaderParser("### WELCOME TO OUR PAGE"))
	fmt.Println(utils.HeaderParser("#### WELCOME TO OUR PAGE"))
	fmt.Println(utils.HeaderParser("##### WELCOME TO OUR PAGE"))
	fmt.Println(utils.HeaderParser("###### WELCOME TO OUR PAGE"))

	// fmt.Print
}
