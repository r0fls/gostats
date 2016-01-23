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

func (b bernoulli) cdf(k int) float64 {
	if k < 0 {
		return 0
	}

	if k < 1 {
		return 1 - b.p
	}
	if k >= 1 {
		return 1
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
	if b.p <= 1 {
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
	return -1
}

func (l laplace) quantile(p float64) float64 {

	if p > 0 && p <= .5 {
		return l.mean + l.b*math.Log(2*p)
	}

	if p > .5 && p < 1 {
		return l.mean - l.b*math.Log(2*(1-p))
	}
	panic("wrong domain")
	return -1
}

func main() {
	b := bernoulli{.5}
	fmt.Println(b.pmf(0))
	fmt.Println(b.cdf(0))
	fmt.Println(b.quantile(.4))

	l := laplace{0, 1}
	fmt.Println(l.pdf(0))
	fmt.Println(l.cdf(0))
	fmt.Println(l.quantile(.5))
}
