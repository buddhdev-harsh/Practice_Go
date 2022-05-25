package basics

import (
	"fmt"
)

func TakeInput() string {
	var name string
	fmt.Scanln(&name)
	return name
}
