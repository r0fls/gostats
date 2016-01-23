package main

import (
	"fmt"
	"math"
)

type Distribution interface {
	random() float64
}

// bernoulli

type bernoulli struct {
	p float64
}

func (b bernoulli) pmf(k int) float64 {
	if k == 1 {
		return b.p
	}
	if k == 0 {
		return 1 - b.p
	}
	return -1
}

func (b bernoulli) quantile(p float64) int {
	if p < 0 {
		return -1
	}
	if p < b.p {
		return 0
	}
	if b.p < 1 {
		return 1
	}
	return -1
}

// laplace

type laplace struct {
	mean, b float64
}

func (l laplace) pdf(x float64) float64 {
	return math.Exp(-math.Abs(x-l.mean)/l.b) / (2 * l.b)
}

func (l laplace) cdf(x float64) float64 {
	if x < l.mean {
		return math.Exp((x-l.mean)/l.b) / 2
	}
	if x >= l.mean {
		return 1 - math.Exp((l.mean-x)/l.b)/2
	}
}
func (l laplace) quantile(x float64) float64 {

	if x > 0 && x < .5 {
		return l.mean + l.b*math.Log(2*l.b*x)
	}

	if x > .5 && x < 1 {
		return l.mean - l.b*math.Log(2*(1-x))
	}
	throw error
}

func main() {
	b := bernoulli{.4}
	fmt.Println(b.pmf(0))
	fmt.Println(b.cdf(0))
	fmt.Println(b.quantile(.5))

	l := laplace{0, 1}
	fmt.Println(l.pdf(0))
	fmt.Println(b.cdf(0))
	fmt.Println(b.quantile(.5))

}
