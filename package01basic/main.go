package package01basic

// A function with lowercase first letter can't be call outside the package
func calculateNumberPri(n ...int) (t int) {
	for _, v := range n {
		t += v
	}
	return t
}

// A function with uppercase first letter can be call outside the package
func CalculateNumberPub(n ...int) (t int) {
	for _, v := range n {
		t += v
	}
	return t
}
