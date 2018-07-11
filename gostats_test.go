package stats_test

import (
	"fmt"
	"github.com/r0fls/gostats"
)

func ExampleBernoulliRandom() {
	stats.Seed()
	b := stats.Bernoulli(.6)
	fmt.Println(b.Random())
}

func ExampleBernoulliPMF() {
	b := stats.Bernoulli(.6)
	fmt.Println(b.Pmf(0))
	// Output: 0.4
}

func ExampleFitBernoulli() {
	data := []int{1, 1, 1, 0, 0}
	b := stats.FitBernoulli(data)
	fmt.Println(b.Cdf(0))
	// Output: 0.4
}

func ExampleBernoulliCDF() {
	b := stats.Bernoulli(.6)
	fmt.Println(b.Cdf(0))
	// Output: 0.4
}

func ExampleBernoulliQuantile() {
	b := stats.Bernoulli(.6)
	fmt.Println(b.Quantile(.3))
	// Output: 0
}

func ExampleLaplaceRandom() {
	l := stats.Laplace(0, 1)
	fmt.Println(l.Random())
}

func ExampleLaplacePDF() {
	l := stats.Laplace(0, 1)
	fmt.Println(l.Pdf(0))
	// Output: 0.5
}
func ExampleLaplaceCDF() {
	l := stats.Laplace(0, 1)
	fmt.Println(l.Cdf(0))
	// Output: 0.5
}
func ExampleLaplaceQuantile() {
	l := stats.Laplace(0, 1)
	fmt.Println(l.Quantile(.5))
	// Output: 0

}

func ExampleFitLaplace() {
	data := []float64{1, 1, 1, 0, 0}
	l := stats.FitLaplace(data)
	fmt.Println(l.Quantile(.5))
	// Output: 1
}

func ExamplePoissonRandom() {
	p := stats.Poisson(5)
	fmt.Println(p.Random())
}
func ExamplePoissonPMF() {
	p := stats.Poisson(5)
	fmt.Println(p.Pmf(5))
	// Output: 0.1754673697678507
}
func ExamplePoissonCDF() {
	p := stats.Poisson(5)
	fmt.Println(p.Cdf(10))
	// Output: 0.9863047314016171
}
func ExamplePoissonQuantile() {
	p := stats.Poisson(5)
	fmt.Println(p.Quantile(.5))
	// Output: 5
}

func ExampleFitPoisson() {
	data := []int{10, 3, 3, 4, 5}
	p := stats.FitPoisson(data)
	fmt.Println(p.Quantile(.5))
	// Output: 5
}

func ExampleGeometricRandom() {
	g := stats.Geometric(.2)
	fmt.Println(g.Random())
}
func ExampleGeometricPMF() {
	g := stats.Geometric(.2)
	fmt.Println(g.Pmf(5))
	// Output: 0.08192000000000005
}
func ExampleGeometricCDF() {
	g := stats.Geometric(.2)
	fmt.Println(g.Cdf(10))
	// Output: 0.8926258175999999
}
func ExampleGeometricQuantile() {
	g := stats.Geometric(.2)
	fmt.Println(g.Quantile(.5))
	// Output: 4
}

func ExampleFitGeometric() {
	data := []int{10, 3, 3, 4, 5}
	g := stats.FitGeometric(data)
	fmt.Println(g.Quantile(.5))
	// Output: 4
}

func ExampleWeibullRandom() {
	w := stats.Weibull(.2, 1)
	fmt.Println(w.Random())
}
func ExampleWeibullPDF() {
	w := stats.Weibull(.2, 1)
	fmt.Println(w.Pdf(5))
}
func ExampleWeibullCDF() {
	w := stats.Weibull(.2, 1)
	fmt.Println(w.Cdf(10))
}
func ExampleWeibullQuantile() {
	w := stats.Weibull(.2, 1)
	fmt.Println(w.Quantile(.5))
}

func ExampleExponentialRandom() {
	e := stats.Exponential(.2)
	fmt.Println(e.Random())
}
func ExampleExponentialPDF() {
	e := stats.Exponential(.2)
	fmt.Println(e.Pdf(5))
	// Output: 0.07357588823428847
}
func ExampleExponentialCDF() {
	e := stats.Exponential(.2)
	fmt.Println(e.Cdf(10))
	// Output: 0.8646647167633873
}
func ExampleExponentialQuantile() {
	e := stats.Exponential(.2)
	fmt.Println(e.Quantile(.5))
	// Output: 3.465735902799726
}

func ExampleFitExponential() {
	data := []float64{10, 3, 3, 4, 5}
	g := stats.FitExponential(data)
	fmt.Println(g.Quantile(.5))
	// Output: 3.465735902799726
}

func ExampleBinomialRandom() {
	r := stats.Binomial(10, .5)
	fmt.Println(r.Random())
}

func ExampleBinomialPMF() {
	r := stats.Binomial(10, .5)
	fmt.Println(r.Pmf(5))
	// Output: 0.24609375
}
func ExampleBinomialCDF() {
	r := stats.Binomial(10, .5)
	fmt.Println(r.Cdf(10))
	// Output: 1
}
func ExampleBinomialQuantile() {
	r := stats.Binomial(10, .5)
	fmt.Println(r.Quantile(.5))
	// Output: 5
}

func ExampleNegativeBinomialRandom() {
	r := stats.NegativeBinomial(10, .5)
	fmt.Println(r.Random())
}

func ExampleNegativeBinomialPMF() {
	r := stats.NegativeBinomial(10, .5)
	fmt.Println(r.Pmf(5))
	// Output: 0.06109619140625
}
func ExampleNegativeBinomialCDF() {
	r := stats.NegativeBinomial(10, .5)
	fmt.Println(r.Cdf(10))
	// Output: 0.5880985260009766
}

func ExampleNegativeBinomialQuantile() {
	r := stats.NegativeBinomial(10, .5)
	fmt.Println(r.Quantile(.5))
	// Output: 9
}

// test common functions

func ExampleSumInt() {
	s := []int{2, 3, 5, 7, 11, 13}
	sum := stats.SumInt(s)
	fmt.Println(sum)
	// Output: 41
}

func ExampleSumFloat64() {
	s := []float64{2, 3, 5, 7, 11, 13}
	sum := stats.SumFloat64(s)
	fmt.Println(sum)
	// Output: 41
}

func ExampleFactorial() {
	fmt.Println(stats.Factorial(4))
	// Output: 24
}
func ExampleChoose() {
	fmt.Println(stats.Choose(4, 2))
	// Output: 6
}

func ExampleMedianInt() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(stats.MedianInt(s))
	// Output: 6
}

func ExampleLSR() {
	data := [][]float64{{60.0, 3.1}, {61.0, 3.6}, {62.0, 3.8}, {63, 4}, {65.0, 4.1}}
	fmt.Println(stats.LSR(data))
	// Output: [-7.963513513513208 0.18783783783783292]
}
