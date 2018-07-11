[![Build Status](https://travis-ci.org/r0fls/gostats.png)](https://travis-ci.org/r0fls/gostats)
# gostats
Statistics for golang

### Usage

To install, do `go get github.com/r0fls/gostats`. Check out `gostats_test.go` for a working example of using each distribution.

### Example
```go
package main

import (
    "fmt"
    "github.com/r0fls/gostats"
)

func main() {
    b := stats.Bernoulli(.5)
    fmt.Println(b.Random())

    // or fit a distribution from data...
    b = stats.FitBernoulli([0,1,1,1])
    fmt.Println(b.Random())
}
```
### Distributions
Currently the following distributions are included:
- Bernoulli
- Laplace
- Poisson
- Geometric
- Weibull
- Exponential
- Binomial
- NegativeBinomial \*

>\*note the negative binomial takes parameters `r`, `p` where `r` is the number allowed success before termination of the trials and `p` is the success of a given trial. The random variable produced by `NegativeBinomial` is the number of failures.


And each distribution has these functions:
- Pmf or Pdf
- Cdf
- Quantile
- Random

Also there is a corresponding function named `FitDistrbution` for each distribution, as shown in the above example with the Bernoulli. That function uses the [MLE](https://en.wikipedia.org/wiki/Maximum_likelihood) for each distribution to choose the best estimation for the parameters and returns an initialized distribution with them.

### Common
Additionally there are some common functions. Most notably is LSR, which performs a least squares regression.

### Additional Info

The default seed function is `time.Now().UTC().UnixNano()`. However you can override that by calling `rand.Seed(uniqueInteger)` yourself before generating random numbers.

### Contributing
If you're interested in contributing, please submit a pull request, or raise an issue.

##### TO-DO
- add more distributions
- should `random` return an array?
- allow updating a fitted distribution with more data
