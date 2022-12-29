package main

import "fmt"

//600851475143

func dels() []int {
	a := 600851475143
	var z []int
	for g := 2; len(z) < 4; {
		a = a / g
		z = append(z, a)
		g *= 2
	}
	return z
}

func min(z []int) {
	var a int
	var c []int
	for _, value := range z {
		for i := 2; i < value; i++ {
			a = i % value
			c = append(c, a)
			fmt.Println(c)
		}
	}
}

func main() {
	a := dels()
	min(a)
}
