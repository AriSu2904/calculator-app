package main

import (
	"fmt"
	calculator_module "github.com/AriSu2904/calculator-module"
)

func main() {
	fmt.Println(calculator_module.Sum(5, 5))

	result, err := calculator_module.Divide(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
