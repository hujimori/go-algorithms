package numeric

func gcdBase(a, b int) int {
	if b == 0 {
		return a
	}

	return gcdBase(b, a%b)
}

// Gcd return Greatest common divisor
func Gcd(args ...int) int {
	b := args[0]

	for _, a := range args {
		b = gcdBase(a, b)
	}
	return b
}
