package binarysearch

func LowerBound(t []int, k int) int {
	lb := -1
	ub := len(t) - 1
	for ub-lb > 1 {
		mid := (lb + ub) / 2

		if t[mid] >= k {
			ub = mid
		} else {
			lb = mid
		}
	}
	return ub
}

func UpperBound(t []int, k int) int {
	lb := -1
	ub := len(t) - 1
	for ub-lb > 1 {
		mid := (lb + ub) / 2

		if t[mid] <= k {
			lb = mid
		} else {
			ub = mid
		}
	}
	return ub
}
