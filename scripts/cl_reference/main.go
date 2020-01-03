/// Generates format test cases from sbcl output
/// Usage: go run main.go 1-100, 101-9223372036854775807->100
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

const fullRangeStart = 0
const fullRangeEnd = 100
const sampleRangeStart = 101
const sampleRangeEnd = 9223372036854775807
const sampleNum = 100
const format = "~:R"
const seed = 666

func main() {

	rng := rand.New(rand.NewSource(seed))

	format, found := os.LookupEnv("FORMAT")
	if !found {
		panic("No FORMAT given")
	}

	full := make([]int64, 0, fullRangeEnd-fullRangeStart)
	for i := fullRangeStart; i <= fullRangeEnd; i++ {
		full = append(full, int64(i))
	}

	sampled := make([]int64, 0, sampleNum+2)
	sampled = append(sampled, sampleRangeStart)
	for k := 0; k < sampleNum; k++ {
		rnd := rng.Int63n(sampleRangeEnd - sampleRangeStart)
		sampled = append(sampled, rnd)
	}
	sampled = append(sampled, sampleRangeEnd)

	all := append(full, sampled...)

	testcases := make([]testcase, 0)
	for _, s := range all {
		evalCmd := fmt.Sprintf("(format T \"%s\" %d)", format, s)
		out, err := exec.Command("sbcl", "--noinform", "--non-interactive", "--eval", evalCmd).Output()
		if err != nil {
			fmt.Printf("err = %s\n", out)
			panic(err)
		}
		testcases = append(testcases, testcase{
			input:          s,
			expectedResult: string(out),
		})
	}

	testcode := genTableTest(testcases)
	fmt.Println(testcode)
}

type testcase struct {
	input          int64
	expectedResult string
}

func genTableTest(testcases []testcase) string {

	var result strings.Builder
	result.WriteString("tcs := []formatTest{\n")
	for _, c := range testcases {
		result.WriteString(fmt.Sprintf("formatT(\"%s\", \"%s\", %d),\n", c.expectedResult, format, c.input))
	}
	result.WriteString("}\nrunTests(t, tcs)")
	return result.String()
}
