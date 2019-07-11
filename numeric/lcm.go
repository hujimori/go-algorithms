package numeric

func lcmBase(a, b int) int {
	return a * (b / gcdBase(a, b))
}

// Lcm return Least common multiple
func Lcm(args ...int) int {
	b := args[0]

	for _, a := range args {
		b = lcmBase(a, b)
	}

	return b
}
