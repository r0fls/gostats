package main

import (
	"./stats"
	"fmt"
)

func main() {
	//b := new(stats.Bernoulli)
	//b.P = .5
	b := stats.New(.5)
	//fmt.Println(b.Random(1))
	fmt.Println(b.Pmf(0))
	fmt.Println(b.Cdf(0))
	fmt.Println(b.Quantile(.6))
	//	fmt.Println(b.Random(5))

	l := stats.Laplace{0, 1}
	fmt.Println(l.Pdf(0))
	fmt.Println(l.Cdf(0))
	fmt.Println(l.Quantile(.5))

}
