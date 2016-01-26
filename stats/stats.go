package stats

// TODO

// - refactor all distributions before Geometric to match
// 	  it's style
// - add more distributions
// - add an initialize helper function:
//    takes a distribution and ... input and returns
//    the parameters needed for a type
//    i.e. don't do BernoulliType{Discrete{bernoulli{p}.Quantile}, p}
//    each time
// - Quantile gets copied for each distribution to the type
// - get laid

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

func (l LaplaceType) Quantile(p float64) float64 {
	return laplace{l.Mean, l.B}.Quantile(p)
}

// Poisson

type poisson struct {
	Mean float64
}

type PoissonType struct {
	Discrete
	Mean float64
}

func Poisson(m float64) PoissonType {
	return PoissonType{Discrete{poisson{m}.Quantile}, m}
}

func (p PoissonType) Pmf(k int) float64 {
	return math.Pow(p.Mean, float64(k)) * math.Exp(-p.Mean) / math.Gamma(float64(k+1))
}

func (p PoissonType) Cdf(k int) float64 {
	total := 0.0
	for i := 0; i <= k; i++ {
		total += p.Pmf(i)
	}
	return total
}

// unfortunately duplication of Pmf/Pdf
// is required if the Quantile uses Pmf/Pdf
// Could be refactored to use the incomplete gamma function

func (p poisson) Pmf(k int) float64 {
	return math.Pow(p.Mean, float64(k)) * math.Exp(-p.Mean) / math.Gamma(float64(k+1))
}

func (p poisson) Quantile(x float64) int {
	total := 0.0
	j := 0
	for total < x {
		j += 1
		total += p.Pmf(j)
	}
	return j
}

// Geometric

type geometric struct {
	P float64
}

type GeometricType struct {
	Discrete
	geometric
}

func Geometric(p float64) GeometricType {
	return GeometricType{Discrete{geometric{p}.Quantile}, geometric{p}}
}

func (g geometric) Pmf(k int) float64 {
	return math.Pow(1-g.P, float64(k-1)) * g.P
}

func (g geometric) Cdf(k int) float64 {
	return 1 - math.Pow(1-g.P, float64(k))
}

func (g geometric) Quantile(p float64) int {
	return int(math.Ceil(math.Log(1-p) / math.Log(1-g.P)))
}

func (g GeometricType) Quantile(p float64) int {
	return geometric{g.P}.Quantile(p)
}

// Weibull

type weibull struct {
	L, K float64
}

type WeibullType struct {
	Continuous
	weibull
}

func Weibull(l float64, k float64) WeibullType {
	return WeibullType{Continuous{weibull{l, k}.Quantile}, weibull{l, k}}
}

func (w weibull) Pdf(x float64) float64 {
	if x >= 0 {
		return (w.K / w.L) * math.Pow(x/w.L, w.K-1) * math.Exp(-math.Pow(x/w.L, w.K))
	}
	if x < 0 {
		return 0
	}
	return -1
}

func (w weibull) Cdf(x float64) float64 {
	if x >= 0 {
		return 1 - math.Exp(-math.Pow(x/w.L, w.K))
	}
	if x < 0 {
		return 0
	}
	return -1
}

func (w weibull) Quantile(p float64) float64 {
	return w.L * (math.Pow(-math.Log(1-p), 1/w.K))
}

func (w WeibullType) Quantile(p float64) float64 {
	return weibull{w.L, w.K}.Quantile(p)
}

// Exponential
// Pareto
