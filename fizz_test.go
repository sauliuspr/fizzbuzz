package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(15)
	want :=
		`1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
`
	if got != want {
		t.Errorf("FizzBuzz(15) \n got: \n%v \n want: \n%v", got, want)
	}
}
