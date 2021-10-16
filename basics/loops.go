package basics

import "fmt"

func Loopf() {
	fmt.Println("lets print loops")

	var sum int = 0
	fmt.Println("Enter number ")
	var numb int
	fmt.Scanf("%d", &numb)
	for i := 0; i < numb; i++ {
		sum += i
	}
	fmt.Print(sum)
}
