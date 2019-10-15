## An implementation of iso6346 check-digit algorithm

[![Go Report Card](https://goreportcard.com/badge/github.com/kolkov/iso6346)](https://goreportcard.com/report/github.com/kolkov/iso6346)
[![Coverage Status](https://coveralls.io/repos/github/kolkov/iso6346/badge.svg?branch=master)](https://coveralls.io/github/kolkov/iso6346?branch=master)
[![Build Status](https://travis-ci.com/kolkov/iso6346.svg?branch=master)](https://travis-ci.com/kolkov/iso6346)

Algorithm will detect wrong container numbers.

It is not intended to be a cryptographically secure hash function. It is mostly used for preflight container numbers.

Compatible with the ozzo-validation package.

### Usage ###

```
import "github.com/kolkov/iso6346"

err := iso6346.Validate("CMAU5110875")

signed := iso6346.Generate("CMAU511087")
```

test on your own by running `make benchmark`

### Resources ###

* [Wikipedia - iso6346 algorithm](https://en.wikipedia.org/wiki/ISO_6346)
