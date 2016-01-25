package stats

import (
	"math"
)

type Discrete struct {
	Quantile func(p float64) int
}

func (this Discrete) Random(k int) int {
	return k
}

//func (d Distribution) random(k int) []float64 {
//	return float64(k)
//}

// Bernoulli

type bernoulli struct {
	P float64
}

type BernoulliType struct {
	Discrete
	P float64
}

func Bernoulli(p float64) BernoulliType {
	return BernoulliType{Discrete{bernoulli{p}.Quantile}, p}
}

func (b bernoulli) Quantile(P float64) int {
	if P < 0 {
		return -1
	}
	if P < b.P {
		return 0
	}
	if b.P <= 1 {
		return 1
	}
	return -1
}

func (b BernoulliType) Pmf(k int) float64 {
	if k == 1 {
		return b.P
	}
	if k == 0 {
		return 1 - b.P
	}
	return -1
}

func (b BernoulliType) Cdf(k int) float64 {
	if k < 0 {
		return 0
	}

	if k < 1 {
		return 1 - b.P
	}
	if k >= 1 {
		return 1
	}
	return -1
}

func (b BernoulliType) Quantile(p float64) int {
	return bernoulli{b.P}.Quantile(p)
}

/*
{
	if P < 0 {
		return -1
	}
	if P < b.P {
		return 0
	}
	if b.P <= 1 {
		return 1
	}
	return -1
}
*/

// Laplace

type Laplace struct {
	Mean, B float64
}

func (l Laplace) Pdf(x float64) float64 {
	return math.Exp(-math.Abs(x-l.Mean)/l.B) / (2 * l.B)
}

func (l Laplace) Cdf(x float64) float64 {
	if x < l.Mean {
		return math.Exp((x-l.Mean)/l.B) / 2
	}
	if x >= l.Mean {
		return 1 - math.Exp((l.Mean-x)/l.B)/2
	}
	return -1
}

func (l Laplace) Quantile(p float64) float64 {

	if p > 0 && p <= .5 {
		return l.Mean + l.B*math.Log(2*p)
	}

	if p > .5 && p < 1 {
		return l.Mean - l.B*math.Log(2*(1-p))
	}
	panic("wrong domain")
	return -1
}
