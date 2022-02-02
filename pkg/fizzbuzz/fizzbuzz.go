package fizzbuzz

import "strconv"

// UpToLength returns a list of strings with the
// fizz-buzz result for the sequence of digits from 1 to length
func UpToLength(length int) []string {
	out := []string{}
	for i := 1; i <= length; i++ {
		out = append(out, digitToFizzBuzz(i))
	}
	return out
}

func digitToFizzBuzz(input int) string {
	if input%3 == 0 && input%5 == 0 {
		return "fizzBuzz"
	} else if input%3 == 0 {
		return "fizz"
	} else if input%5 == 0 {
		return "buzz"
	} else {
		return strconv.Itoa(input)
	}
}
