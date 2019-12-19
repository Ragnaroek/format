package format

import (
	"testing"
)

//~c
func Test_C(t *testing.T) {
	tcs := []formatTest{
		formatT("k", "~c", 'k'),
		formatT("🥭", "~c", '🥭'),
		formatT(" ", "~c", ' '),
		formatT("\n", "~c", '\n'),
		formatT("\x00", "~c", '\x00'),
		formatT("~!c(string=foo)", "~c", "foo"),
		//modifier @
		formatT("'k'", "~@c", 'k'),
		formatT("'🥭'", "~@c", '🥭'),
	}
	runTests(t, tcs)
}

//~%
func Test_Percent(t *testing.T) {
	tcs := []formatTest{
		formatT("", "~0%"),
		formatT("\n", "~%"),
		formatT("\n", "~1%"),
		formatT("\n\n\n\n\n\n", "~6%"),
		//errs
		formatT("~!%(prefix.num!=1)", "~6,8%"),
	}
	runTests(t, tcs)
}

func Test_A(t *testing.T) {
	tcs := []formatTest{
		formatT("Hello", "~a", "Hello"),
	}
	runTests(t, tcs)
}

//helper

func runTests(t *testing.T, tests []formatTest) {
	for _, tc := range tests {
		t.Run(tc.format, func(t *testing.T) {
			result := Sformat(tc.format, tc.args...)
			if result != tc.expected {
				t.Errorf("expected `%s`, got `%s`", tc.expected, result)
			}
		})
	}
}

type repeatDef struct {
	index   int
	linksTo int
}

type formatTest struct {
	format   string
	args     []interface{}
	expected string
}

func formatT(expected string, format string, args ...interface{}) formatTest {
	return formatTest{format, args, expected}
}
