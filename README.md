# gostats
Statistics for go

###Usage
To install, clone the repo and put it in your project. See the `example.go` for a working example of using each distribution.

You have to call `Seed()` initially before generating any random numbers.

#####Advanced
You can supply your own seed function, instead of the default:

    func Seed() {
	    rand.Seed(time.Now().UTC().UnixNano())
    }

To do so currently you will need to go into stats.go and modify that function.

###Distributions
Thus far, the following distributions are included: 
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

###Contributing
If you're interested in contributing, please submit a pull request, or raise an issue.

#####TO-DO
- add more distributions
- make a `fit` method using the MLE for each distribution
- make `random` return an array?
