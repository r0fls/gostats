package stats

// TODO

// - add more distributions
// - add an initialize helper function:
//    takes a distribution and ... input and returns
//    the parameters needed for a type
//    i.e. don't do BernoulliType{discrete{bernoulli{p}.Quantile}, p}
//    each time
// - Quantile gets copied for each distribution to the type
// - Random should return an array?

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

func init() {
	Seed()
}

func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func random() float64 {
	return rand.Float64()
}

type discrete struct {
	Quantile func(p float64) int
}

func (this discrete) Random(k ...int) int {
	return this.Quantile(random())
}

type continuous struct {
	Quantile func(p float64) float64
}

func (this continuous) Random(k ...int) float64 {
	return this.Quantile(random())
}

// Bernoulli

type bernoulli struct {
	P float64
}

type BernoulliType struct {
	discrete
	bernoulli
}

func Bernoulli(p float64) BernoulliType {
	return BernoulliType{discrete{bernoulli{p}.Quantile}, bernoulli{p}}
}

func FitBernoulli(data []int) BernoulliType {
	mean := MeanInt(data)
	return Bernoulli(mean)
}

func (b bernoulli) Pmf(k int) float64 {
	if k == 1 {
		return b.P
	}
	if k == 0 {
		return 1 - b.P
	}
	return -1
}

func (b bernoulli) Cdf(k int) float64 {
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

func (b bernoulli) Quantile(P float64) int {
	if P < 0 {
		return -1
	} else if P < 1-b.P {
		return 0
	} else if P <= 1 {
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
	continuous
	laplace
}

func Laplace(mean float64, b float64) LaplaceType {
	return LaplaceType{continuous{laplace{mean, b}.Quantile}, laplace{mean, b}}
}

func FitLaplace(data []float64) LaplaceType {
	mean := MedianFloat64(data)
	b := 0.0
	for _, value := range data {
		b += math.Abs(value - mean)
	}
	return Laplace(mean, b)
}

func (l laplace) Pdf(x float64) float64 {
	return math.Exp(-math.Abs(x-l.Mean)/l.B) / (2 * l.B)
}

func (l laplace) Cdf(x float64) float64 {
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
	discrete
	poisson
}

func Poisson(m float64) PoissonType {
	return PoissonType{discrete{poisson{m}.Quantile}, poisson{m}}
}

func FitPoisson(data []int) PoissonType {
	return Poisson(MeanInt(data))
}

func (p poisson) Pmf(k int) float64 {
	return math.Pow(p.Mean, float64(k)) * math.Exp(-p.Mean) / math.Gamma(float64(k+1))
}

func (p poisson) Cdf(k int) float64 {
	total := 0.0
	for i := 0; i <= k; i++ {
		total += p.Pmf(i)
	}
	return total
}

func (p poisson) Quantile(x float64) int {
	j := 0
	total := p.Pmf(0)
	for total < x {
		j += 1
		total += p.Pmf(j)
	}
	return j
}

func (p PoissonType) Quantile(x float64) int {
	return poisson{p.Mean}.Quantile(x)
}

// Geometric

type geometric struct {
	P float64
}

type GeometricType struct {
	discrete
	geometric
}

func Geometric(p float64) GeometricType {
	return GeometricType{discrete{geometric{p}.Quantile}, geometric{p}}
}

func FitGeometric(data []int) GeometricType {
	return Geometric(1 / MeanInt(data))
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
	continuous
	weibull
}

func Weibull(l float64, k float64) WeibullType {
	return WeibullType{continuous{weibull{l, k}.Quantile}, weibull{l, k}}
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

type exponential struct {
	L float64
}

type ExponentialType struct {
	continuous
	exponential
}

func Exponential(l float64) ExponentialType {
	return ExponentialType{continuous{exponential{l}.Quantile}, exponential{l}}
}

func FitExponential(data []float64) ExponentialType {
	return Exponential(1 / MeanFloat64(data))
}

func (e exponential) Pdf(x float64) float64 {
	return e.L * math.Exp(-e.L*x)
}

func (e exponential) Cdf(x float64) float64 {
	return 1 - math.Exp(-e.L*x)
}

func (e exponential) Quantile(p float64) float64 {
	return -math.Log(1-p) / e.L
}

func (e ExponentialType) Quantile(p float64) float64 {
	return exponential{e.L}.Quantile(p)
}

// Binomial
type binomial struct {
	N int
	P float64
}

type BinomialType struct {
	discrete
	binomial
}

func Binomial(n int, p float64) BinomialType {
	return BinomialType{discrete{binomial{n, p}.Quantile}, binomial{n, p}}
}

func (b binomial) Pmf(k int) float64 {
	r := float64(k)
	return float64(Choose(b.N, k)) *
		math.Pow(b.P, r) * math.Pow(1-b.P, float64(b.N-k))
}

func (b binomial) Cdf(k int) float64 {
	total := 0.0
	for i := 0; i <= k; i++ {
		total += b.Pmf(i)
	}
	return total
}

func (b binomial) Quantile(x float64) int {
	j := 0
	total := b.Pmf(0)
	for total < x {
		j += 1
		total += b.Pmf(j)
	}
	return j
}

func (b BinomialType) Quantile(x float64) int {
	return binomial{b.N, b.P}.Quantile(x)
}

// NegativeBinomial
type negativeBinomial struct {
	K int
	P float64
}

type NegativeBinomialType struct {
	discrete
	negativeBinomial
}

func NegativeBinomial(k int, p float64) NegativeBinomialType {
	return NegativeBinomialType{discrete{negativeBinomial{k, p}.Quantile}, negativeBinomial{k, p}}
}

func (b negativeBinomial) Pmf(r int) float64 {
	return float64(Choose(b.K+r-1, r)) * math.Pow(b.P, float64(b.K)) * math.Pow(1-b.P, float64(r))
}

func (b negativeBinomial) Cdf(r int) float64 {
	total := 0.0
	for i := 0; i <= r; i++ {
		total += b.Pmf(i)
	}
	return total
}

func (b negativeBinomial) Quantile(x float64) int {
	j := 0
	total := b.Pmf(0)
	for total < x {
		j += 1
		total += b.Pmf(j)
	}
	return j
}

func (b NegativeBinomialType) Quantile(x float64) int {
	return negativeBinomial{b.K, b.P}.Quantile(x)
}

// Common functions
func Factorial(n int) int {
	return int(math.Gamma(float64(n) + 1))
}

func Choose(n int, k int) int {
	return int(Factorial(n) / (Factorial(k) * Factorial(n-k)))
}

func MeanInt(data []int) float64 {
	return float64(SumInt(data)) / float64(len(data))
}

func SumInt(data []int) int {
	total := 0
	for _, value := range data {
		total += value
	}
	return total
}

func MeanFloat64(data []float64) float64 {
	return SumFloat64(data) / float64(len(data))
}

func SumFloat64(data []float64) float64 {
	total := 0.0
	for _, value := range data {
		total += value
	}
	return total
}

func MedianInt(data []int) float64 {
	if sort.IntsAreSorted(data) {
		if len(data)%2 == 1 {
			return float64(data[len(data)/2])
		} else {
			return float64(data[len(data)/2]+data[len(data)/2-1]) / 2.0
		}
	} else {
		sort.Ints(data)
		if len(data)%2 == 1 {
			return float64(data[len(data)/2])
		} else {
			return float64(data[len(data)/2]+data[len(data)/2-1]) / 2.0
		}
	}
	panic("Error in MedianInt")
	return -1
}

func MedianFloat64(data []float64) float64 {
	if sort.Float64sAreSorted(data) {
		if len(data)%2 == 1 {
			return data[len(data)/2]
		} else {
			return (data[len(data)/2] + data[len(data)/2-1]) / 2.0
		}
	} else {
		sort.Float64s(data)
		if len(data)%2 == 1 {
			return data[len(data)/2]
		} else {
			return (data[len(data)/2] + data[len(data)/2-1]) / 2.0
		}
	}
	panic("Error in MedianInt")
	return -1
}

func LSR(data [][]float64) []float64 {
	total_x := 0.0
	total_xy := 0.0
	total_y := 0.0
	total_x2 := 0.0

	for _, value := range data {
		total_x += value[0]
		total_y += value[1]
		total_xy += value[0] * value[1]
		total_x2 += math.Pow(value[0], 2)
	}
	N := float64(len(data))
	b := (N*total_xy - total_x*total_y) / (N*total_x2 - math.Pow(total_x, 2))
	a := (total_y - b*total_x) / N
	return []float64{a, b}
}
