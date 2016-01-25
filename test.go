package main

import (
	"./stats"
	"fmt"
)

func main() {
	//b := new(stats.Bernoulli)
	//b.P = .5
	stats.Seed()
	b := stats.Bernoulli(.5)
	fmt.Println(b.Random())
	fmt.Println(b.Pmf(0))
	fmt.Println(b.Cdf(0))
	fmt.Println(b.Quantile(.4))

	l := stats.Laplace(0, 1)
	fmt.Println(l.Random())
	fmt.Println(l.Pdf(0))
	fmt.Println(l.Cdf(0))
	fmt.Println(l.Quantile(.5))

}
