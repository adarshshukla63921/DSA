package main

import "fmt"

func myPow(x float64, n int) float64 {
	ans := 1.0
	for n > 0 {
		if n%2 == 1 {
			ans = ans * x
			n = n - 1
		}
		x = x * x
		n = n / 2
	}
	return ans
}
func main() {
	fmt.Println(myPow(2.00000, 10))
}