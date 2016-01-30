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

func ExampleGeometricRandom() {
	g := stats.Geometric(.2)
	fmt.Println(g.Random())
}
func ExampleGeometricPMF() {
	g := stats.Geometric(.2)
	fmt.Println(g.Pmf(5))
}
func ExampleGeometricCDF() {
	g := stats.Geometric(.2)
	fmt.Println(g.Cdf(10))
}
func ExampleGeometricQuantile() {
	g := stats.Geometric(.2)
	fmt.Println(g.Quantile(.5))
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
}
func ExampleExponentialCDF() {
	e := stats.Exponential(.2)
	fmt.Println(e.Cdf(10))
}
func ExampleExponentialQuantile() {
	e := stats.Exponential(.2)
	fmt.Println(e.Quantile(.5))
}

func ExampleBinomialRandom() {
	r := stats.Binomial(10, .5)
	fmt.Println(r.Random())
}
func ExampleBinomialPMF() {
	r := stats.Binomial(10, .5)
	fmt.Println(r.Pmf(5))
}
func ExampleBinomialCDF() {
	r := stats.Binomial(10, .5)
	fmt.Println(r.Cdf(10))
}
func ExampleBinomialQuantile() {
	r := stats.Binomial(10, .5)
	fmt.Println(r.Quantile(.5))
}