package main

import "fmt"

//Not Allowed z := 900. You can specify the type as well
var z int = 65
type new_type int
var a new_type = 90

func main() {
	//n bytes, it is necessary to use the variables declared
	//; are optional. Compiler add for us.
	n, err := fmt.Println("Hello, here is what you return")
	fmt.Println(n)
	fmt.Println(err)

	//_ to get discarded. Avoid code pollution
	nBytes, _ := fmt.Println("Hello, return type ")
	fmt.Println(nBytes)

	// to declare a variable := (Short Declaration)
	x := 900
	fmt.Println("x := ", x)
	x=9100
	fmt.Println("x = ", x)
	fmt.Println("var z =  ",z)
	//This returns a String
	output := fmt.Sprintf("Hello")
	fmt.Println("output from Sprinng is", output)
	fmt.Printf("Type of a is %T",a)
}
