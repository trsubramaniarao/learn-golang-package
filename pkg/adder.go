package pkg

func Adder(a int, b ...int) int {
	result := a
	for _, x := range b {
		result += x
	}
	return result
}
