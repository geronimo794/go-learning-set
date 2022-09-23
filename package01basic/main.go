package package01basic

var name string

// Int function will be always call first when the package
func init() {
	name = "Package name : package01basic"
}

func GetName() string {
	return name
}

// PRIVATE FUNCTION
// A function with lowercase first letter can't be call outside the package
func calculateNumberPri(n ...int) (t int) {
	for _, v := range n {
		t += v
	}
	return t
}

// PUBLIC FUNCTION
// A function with uppercase first letter can be call outside the package
func CalculateNumberPub(n ...int) (t int) {
	for _, v := range n {
		t += v
	}
	return t
}
