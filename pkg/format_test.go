package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFormatGraph(t *testing.T) {
	tcs := []struct {
		format              string
		expectedFormatGraph root
	}{
		{
			"string only",
			buildGraph([]ftoken{NewLiteral("string only")}),
		},
		{
			"~a",
			buildGraph([]ftoken{&directive{char: rune('a')}}),
		},
		{
			"~:a",
			buildGraph([]ftoken{&directive{char: rune('a'), colonMod: true}}),
		},
		{
			"~:@a",
			buildGraph([]ftoken{&directive{char: rune('a'), colonMod: true, atMod: true}}),
		},
		{
			"~@:a",
			buildGraph([]ftoken{&directive{char: rune('a'), colonMod: true, atMod: true}}),
		},
		{
			"~12d",
			buildGraph([]ftoken{&directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: 12}}}}),
		},
		{
			"~-12d",
			buildGraph([]ftoken{&directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: -12}}}}),
		},
		{
			"~+12d",
			buildGraph([]ftoken{&directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: 12}}}}),
		},
		{
			"~'0d",
			buildGraph([]ftoken{&directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{charParam: '0'}}}}),
		},
		{
			"~12,'xd",
			buildGraph([]ftoken{&directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: 12}, prefixParam{charParam: 'x'}}}}),
		},
		{
			"foo ~:a bar",
			buildGraph([]ftoken{NewLiteral("foo "), &directive{char: rune('a'), colonMod: true}, NewLiteral(" bar")}),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.format, func(t *testing.T) {
			node := parseFormatGraph(tc.format)
			var nodeExpected ftoken
			nodeExpected = &tc.expectedFormatGraph
			assert.NotNil(t, node)

			var i = 0
			for node.Next() != nil {
				assert.Equal(t, nodeExpected, node, "expected %#v, actual %#v", nodeExpected, node)
				node = node.Next()
				nodeExpected = nodeExpected.Next()
				i++
			}
		})
	}
}

/*
func Test_A(t *testing.T) {
	tcs := []formatTest{
		formatT("~a", "Hello"),
	}

	for _, tc := range tcs {
		t.Run(tc.format, func(t *testing.T) {
			result := Sformat(tc.format, tc.args...)
			if result != tc.expected {
				t.Errorf("expected `%s`, got `%s`", tc.expected, result)
			}
		})
	}
}
*/

//helper

type formatTest struct {
	format   string
	args     []interface{}
	expected string
}

func formatT(expected string, format string, args ...interface{}) formatTest {
	return formatTest{format, args, expected}
}

func buildGraph(tokens []ftoken) root {
	rootDir := root{}
	var curDir ftoken
	curDir = &rootDir
	for _, token := range tokens {
		curDir.SetNext(token)
		curDir = token
	}
	return rootDir
}
