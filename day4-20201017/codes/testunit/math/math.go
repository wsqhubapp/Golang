package math

func Add(n1, n2 int) int {
	return n1 + n2
}

func Mul(n1, n2 int) int {
	return n1 * n2
}

func Sub(n1, n2 int) int {
	return n1 - n2
}

func Div(n1, n2 int) int {
	if n2 == 0 {
		return 0
	} else {
		return n1 / n2
	}

}
