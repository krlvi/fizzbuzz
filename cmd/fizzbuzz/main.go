package main

import (
	"fizzbuzz/pkg/fizzbuzz"
	"fmt"
)

func main() {
	for _, v := range fizzbuzz.UpToLength(20) {
		fmt.Println(v)
	}
}
