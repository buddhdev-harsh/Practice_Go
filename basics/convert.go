package basics

import (
	"strconv"
)

func Change_Value(receive string) (sender int, err error) {
	sender, err = strconv.Atoi(receive)
	return
}
