# go-luhn [![Build Status](https://travis-ci.org/delaemon/go-luhn.svg?branch=master)](https://travis-ci.org/delaemon/go-luhn)

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
	"github.com/delaemon/luhn"
)

func main() {
	l := 10
	number := luhn.GetRandomNumber(l)
	checkSum := luhn.CheckSum(number)
	luhnNumber := number + checkSum
	valid := luhn.Valid(luhnNumber)
}
```
Examples can be found under the ./_example directory
