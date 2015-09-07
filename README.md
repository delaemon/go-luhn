# go-luhn [![Build Status](https://travis-ci.org/delaemon/go-luhn.svg?branch=master)](https://travis-ci.org/delaemon/go-luhn) [![GoCover](http://gocover.io/_badge/github.com/delaemon/go-luhn)](http://gocover.io/github.com/delaemon/go-luhn) [![GoDoc](https://godoc.org/github.com/delaemon/go-luhn?status.png)](https://godoc.org/github.com/delaemon/go-luhn)

##Description
Very simple library to calculate and validate Luhn numbers.

##Installation
This package can be installed with the go get command:
```sh
go get github.com/delaemon/go-luhn
```

###Usage
```
package main

import (
	"github.com/delaemon/go-luhn"
)

func main() {
	length := 10
	luhnNumber := luhn.Generate(length)
	valid := luhn.Valid(luhnNumber)
}
```
Examples can be found under the ./_example directory
