package bajik

import (
	"fmt"
	"reflect"
	"strings"
)

func String(s string) {
	// s := "Hello this is string program. thanks for running it!"
	for i, p := range s {
		fmt.Printf("i：%d,s：%c == %d\n", i, p, p)
	}
}

func Do_Join() {
	s := strings.Join([]string{"hello", "world", "this", "is", "harsh"}, ", ")
	fmt.Print(s)
}

func Get_type(x interface{}) {
	fmt.Println(reflect.TypeOf(x).String())
}
