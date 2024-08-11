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

	// test decoration
	fmt.Println(utils.Decoration("**bold**"))
	fmt.Println(utils.Decoration("*italic*"))
	fmt.Println(utils.Decoration("word"))
	fmt.Println(utils.Decoration("~~strikethrough~~"))
}
