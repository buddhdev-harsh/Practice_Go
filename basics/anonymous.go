package basics

import (
	"fmt"
	"reflect"
)

var Call = func(a int, c int) int {
	return a * c
}

func Return_sum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x * y * z
		}
	}
}

func More_Value_FUnction(p ...string) {
	fmt.Println(p[0], p[1])
}

func Pass_Anything(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "---", reflect.ValueOf(v).Kind())
	}
}
