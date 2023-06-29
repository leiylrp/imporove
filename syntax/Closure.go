package syntax

import (
	"fmt"
	"time"
)

func newCounter() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}

type funcSign func(int)

func timer(f func(int)) func(n int) {
	return func(n int) {
		start := time.Now()
		f(n)
		end := time.Now()
		fmt.Println("this operation take ", end.Sub(start))
	}
}

func printN(n int) {
	fmt.Println("-----", n)
}

