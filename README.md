[![Build Status](https://travis-ci.org/r0fls/gostats.png)](https://travis-ci.org/r0fls/gostats)
# gostats
Statistics for golang

### Usage

To install, do `go get github.com/r0fls/gostats`. Check out `gostats_test.go` for a working example of using each distribution.

You have to call `Seed()` initially before generating any random numbers (see
the example below).

##### Advanced

You can supply your own seed function, instead of the default:
```go
func Seed() {
    rand.Seed(time.Now().UTC().UnixNano())
}
```
To do so currently you will need to go into stats.go and modify that function.

### Example
```go
package main

import (
    "fmt"
    "github.com/r0fls/gostats"
)

func main() {
    stats.Seed()
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

And each distribution has these functions:
- Pmf or Pdf
- Cdf
- Quantile
- Random

Also there is a corresponding function named `FitDistrbution` for each distribution, as shown in the above example with the Bernoulli. That function uses the [MLE](https://en.wikipedia.org/wiki/Maximum_likelihood) for each distribution to choose the best estimation for the parameters and returns an initialized distribution with them.

### Common
Additionally there are some common functions. Most notably is LSR, which performs a least squares regression.

### Contributing
If you're interested in contributing, please submit a pull request, or raise an issue.

##### TO-DO
- add more distributions
- should `random` return an array?
- allow updating a fitted distribution with more data
