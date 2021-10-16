package basics

import "fmt"

func Arr_make() {
	var size int

	fmt.Println("create Array of your number")
	fmt.Scanf("%d", &size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %vth array value: ", i+1)
		fmt.Scanf("%d", &arr[i])
	}
	print_arr(arr, size)

	var position int
	fmt.Println("enterr position value: ")
	fmt.Scanf("%d", &position)

	arr = ret_arr(arr, position)

	print_arr(arr, size)
}

func print_arr(p []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf(" %vth array value: \n", p[i])
	}
}

func ret_arr(re []int, position int) []int {
	var update int
	fmt.Println("Enter value you need to update: ")
	fmt.Scanf("%d", &update)
	re[position-1] = update
	return re
}
