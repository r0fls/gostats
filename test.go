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

	p := stats.Poisson(5)
	fmt.Println(p.Random())
	fmt.Println(p.Pmf(5))
	fmt.Println(p.Cdf(10))
	fmt.Println(p.Quantile(.5))

	g := stats.Geometric(.2)
	fmt.Println(g.Random())
	fmt.Println(g.Pmf(5))
	fmt.Println(g.Cdf(10))
	fmt.Println(g.Quantile(.5))

	w := stats.Weibull(.2, 1)
	fmt.Println(w.Random())
	fmt.Println(w.Pdf(5))
	fmt.Println(w.Cdf(10))
	fmt.Println(w.Quantile(.5))

}
