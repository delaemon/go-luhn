package main
import (
	"fmt"

	"github.com/delaemon/luhn"
)

func main() {
	l := 10
	number := luhn.GetRandomNumber(l)
	checkSum := luhn.CheckSum(number)
	luhnNumber := number + checkSum
	valid := luhn.Valid(luhnNumber)
	fmt.Printf("len: %d, num: %s, checkSum: %s, luhn: %s, valid: %t\n",
		l, number, checkSum, luhnNumber, valid)
}