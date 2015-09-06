package luhn
import "fmt"

func test_luhn() {
	len := 10
	num := getRandomNumber(len)
	checkSum := checkSum(num)
	luhn := string(num) + string(checkSum)
	valid := valid(luhn)
	fmt.Printf("len: %d, num: %d, checkSum: %d, luhn: %s, valid: %t", len, num, checkSum, luhn, valid)
}
