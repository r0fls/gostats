# gostats
Statistics for go

###Usage
To install, clone the repo and put it in your project. See the `example.go` for a working example of using each distribution.

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
If you're interested in contributing, please submit a pull request.

####TO-DO
- add more distributions
- make a `fit` method using the MLE for each distribution
- make `random` return an array?
