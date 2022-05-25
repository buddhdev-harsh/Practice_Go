package basics

func Switch_result(name string) (value int) {
	switch name {
	case "harsh":
		value = 5
	case "noname":
		value = 0
	default:
		value = 1000
	}
	return
}
