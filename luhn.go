package luhn
import (
	"sort"
	"math/rand"
	"time"
	"strconv"
	"unicode/utf8"
)

func CheckSum(num string) string {
	products := doubled(num)
	sum := 0
	for _, n := range products {
		sum = sum + sumOf(n)
	}
	check := 10 - (sum % 10)
	res := strconv.Itoa(check)
	if check == 10 {
		res = "0"
	}
	return res
}

func sumOf(n int) int {
	sum := 0
	num := split(strconv.Itoa(n))
	for _, n := range num {
		sum += n
	}
	return sum
}

func Valid(num string) bool {
	n := split(num)
	return strconv.Itoa(n[len(n)-1]) == CheckSum(num[0:len(n)-1])
}

func doubled(num string) []int {
	s := split(num)
	keys := keys(s)
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for i, k := range keys {
		if i % 2 == 0 {
			s[i] = s[k] * 2
		} else {
			s[i] = s[k] * 1
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for i, k := range keys {
		s[i], s[k] = s[k], s[i]
	}
	return s
}

func keys(arr []int) []int {
	keys := make([]int, 0, len(arr))
	for k, _ := range arr {
		keys = append(keys, k)
	}
	return keys
}

func split(num string) []int {
	res := make([]int, 0, len(num))
	for _, d := range num {
		buf := make([]byte, 1)
		_ = utf8.EncodeRune(buf, d)
		n, _ := strconv.Atoi(string(buf))
		res = append(res, n)
	}
	return res
}

func GetRandomNumber(length int) string {
	numbers := []int{1,2,3,4,5,6,7,8,9,0}
	res := ""
	rand.Seed(time.Now().UnixNano())
	for length > 0 {
		r := rand.Intn(len(numbers))
		res = res + strconv.Itoa(numbers[r])
		length--
	}
	return res
}
