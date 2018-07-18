package fizzbuzz

import "fmt"

func fizz(n int) (r string) {
	if n%3 == 0 {
		r += "Fizz"
	}
	if n%5 == 0 {
		r += "Buzz"
	}
	if r != "" {
		return r
	}
	return fmt.Sprintf("%d", n)
}

// FizzBuzz Gerenrates String from 1 to amount
func FizzBuzz(amount int) (res string) {
	for i := 1; i <= amount; i++ {
		res += fizz(i) + "\n"
	}
	return res
}
