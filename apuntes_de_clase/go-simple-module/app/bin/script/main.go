package main

import (
	"fmt"
	"github.com/fatih/color"
	"go-bases/intro/app/lib/calculator"
)

func main() {
	color.Red("Hello world!")
	fmt.Println(calculator.Value)
	fmt.Println(calculator.String)
}
