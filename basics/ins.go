package basics

import "fmt"

func Ins_Of_all_Type() {
	var i int
	var s string
	var f float32
	// var b bool
	fmt.Println("enter int: ")
	fmt.Scanf("%d", &i)
	fmt.Println("enter string: ")
	fmt.Scanf("%s", &s)
	fmt.Println("enter float: ")
	fmt.Scanf("%f", &f)

	Get_type(i)
	Get_type(s)
	Get_type(f)

	fmt.Printf("%v loves number %v and his age is: %v", s, i, f)
}
