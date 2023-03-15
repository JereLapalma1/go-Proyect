package main

import (
	"fmt"
)

func main() {

	var varInt int = -200
	fmt.Println("my variable", varInt)

	var varUnit uint = 12
	fmt.Println("unit variable", varUnit)

	var varString string = "hello world"
	fmt.Println(varString)

	var varBool bool = true
	fmt.Println("bool variable", varBool)

	fmt.Println("variable address is ", &varString)

	//setear un valor a una variable sin definir el tipo de datoc

	varInt2 := 100
	fmt.Println("mi variable 2 is", varInt2)

	const varConst int = 1000
	fmt.Println("my constant is", varConst)

}
