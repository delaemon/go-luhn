package luhn
import "testing"

func TestCheckSum(t *testing.T) {
	in := "1972096831"
	want := "6"
	out := checkSum(in)
	if out != want {
		t.Errorf("in: %s, out: %s, want: %t, out != want", in, out, want)
	}

	in = "4274801484"
	want = "0"
	out = checkSum(in)
	if out != want {
		t.Errorf("in: %s, out: %s, want: %t, out != want", in, out, want)
	}
}

func TestGetRandomNumber(t *testing.T) {
	length := 10
	number := getRandomNumber(length)
	if length != len(number) {
		t.Errorf("number:%s, len: %d, want: %d", number, len(number), length)
	}
}

func TestSumOf(t *testing.T) {
	in := 123
	want := 6
	out := sumOf(in)
	if out != want {
		t.Errorf("in: %s, out: %s, want: %t, out != want", in, out, want)
	}
}

func TestDubled(t *testing.T) {
	in := "8101304952"
	want := []int{4,5,18,4,0,0,8,18,10,4}
	out := doubled(in)
	for i := 0; i < len(in); i++ {
		if out[i] != want[i] {
			t.Errorf("in: %s, out: %s, want: %s, out != want", in, out, want)
		}
	}
}

func TestKeys(t *testing.T) {
	in := []int{9,8,7,6,5,4,3,2,1,0}
	want := []int{0,1,2,3,4,5,6,7,8,9}
	out := keys(in)
	for i := 0; i < len(in); i++ {
		if out[i] != want[i] {
			t.Errorf("in: %s, out: %s, want: %s, out != want", in, out, want)
		}
	}
}

func TestSplit(t *testing.T) {
	in := "1234567890"
	out := split(in)
	want := []int{1,2,3,4,5,6,7,8,9,0}
	for i := 0; i < len(in); i++ {
		if out[i] != want[i] {
			t.Errorf("in: %s, out: %s, want: %s, out != want", in, out, want)
		}
	}
}

func TestGenerate(t *testing.T) {
	in := 10
	want := 10
	out := Generate(in)
	if len(out) != want {
		t.Errorf("in: %s, out: %t, want: %t, len(out) != want", in, out, want)
	}
}

func TestValid(t *testing.T) {
	in := "24380106902"
	want := true
	out := Valid(in)
	if out != want {
		t.Errorf("in: %s, out: %t, want: %t, out != want", in, out, want)
	}

	in = "243801"
	want = false
	out = Valid(in)
	if out != want {
		t.Errorf("in: %s, out: %t, want: %t, out != want", in, out, want)
	}
}

