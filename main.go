package main

import (
	"Practice_Go/basics"
	"fmt"
)

func main() {

	//take input for name
	fmt.Println("Enter your name:")
	name := basics.TakeInput()
	fmt.Println("your name is: " + name)

	fmt.Println("enter any vlaue")
	var value int
	fmt.Scanln(&value)
	//converting string to int
	// fmt.Println(basics.Change_Value(name))

	//call switch statement
	// value := basics.Switch_result(name)
	// fmt.Println(value)

	//update or pass by reference in go
	// basics.Update(&value, &name)
	// fmt.Println(value, name)

	//anonymous functions
	fmt.Println(basics.Call(3, 2))

	func(l int, b int) {
		fmt.Println(l + b)
	}(10, 10)

	fmt.Println("this will do multiplication", basics.Return_sum(5)(5)(10))

	basics.More_Value_FUnction("hello", "shello", "krte", "jao")

	basics.Pass_Anything(1, "red", true)
}
