package main
import (
	"fmt"

	"github.com/delaemon/luhn"
)

func main() {
	length := 10
	luhnNumber := luhn.Generate(length)
	valid := luhn.Valid(luhnNumber)
	fmt.Printf("len: %d, luhn: %s, valid: %t\n", length, luhnNumber, valid)
}