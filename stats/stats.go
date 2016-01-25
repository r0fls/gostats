package stats

import (
	"math"
	"math/rand"
	"time"
)

func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func random() float64 {
	return rand.Float64()
}

type Discrete struct {
	Quantile func(p float64) int
}

func (this Discrete) Random(k ...int) int {
	return this.Quantile(random())
}

type Continuous struct {
	Quantile func(p float64) float64
}

func (this Continuous) Random(k ...int) float64 {
	return this.Quantile(random())
}

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

// Laplace

type laplace struct {
	Mean, B float64
}

type LaplaceType struct {
	Continuous
	Mean, B float64
}

func Laplace(mean float64, b float64) LaplaceType {
	return LaplaceType{Continuous{laplace{mean, b}.Quantile}, mean, b}
}

func (l LaplaceType) Pdf(x float64) float64 {
	return math.Exp(-math.Abs(x-l.Mean)/l.B) / (2 * l.B)
}

func (l LaplaceType) Cdf(x float64) float64 {
	if x < l.Mean {
		return math.Exp((x-l.Mean)/l.B) / 2
	}
	if x >= l.Mean {
		return 1 - math.Exp((l.Mean-x)/l.B)/2
	}
	return -1
}

func (l laplace) Quantile(p float64) float64 {

	if p > 0 && p <= .5 {
		return l.Mean + l.B*math.Log(2*p)
	}

	if p > .5 && p < 1 {
		return l.Mean - l.B*math.Log(2*(1-p))
	}
	panic("wrong domain")
	return -1
}
