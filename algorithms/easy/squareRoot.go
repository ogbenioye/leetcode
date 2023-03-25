package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	for i := 1; i <= x; i++ {
		if i*i == x {
			return i
		} else if i*i > x {
			return i - 1
		}
	}
	return 0
}

func main() {
	mySqrt(8)
}
