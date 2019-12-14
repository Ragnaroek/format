package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFormatGraph(t *testing.T) {
	tcs := []struct {
		format            string
		expectedTokens    []ftoken
		expectedRepeation []repeatDef
	}{
		{
			format:         "string only",
			expectedTokens: []ftoken{NewLiteral("string only")},
		},
		{
			format:         "~a",
			expectedTokens: []ftoken{&directive{char: rune('a')}},
		},
		{
			format:         "~:a",
			expectedTokens: []ftoken{&directive{char: rune('a'), colonMod: true}},
		},
		{
			format:         "~:@a",
			expectedTokens: []ftoken{&directive{char: rune('a'), colonMod: true, atMod: true}},
		},
		{
			format:         "~@:a",
			expectedTokens: []ftoken{&directive{char: rune('a'), colonMod: true, atMod: true}},
		},
		{
			format:         "~12d",
			expectedTokens: []ftoken{&directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: 12}}}},
		},
		{
			format:         "~-12d",
			expectedTokens: []ftoken{&directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: -12}}}},
		},
		{
			format:         "~+12d",
			expectedTokens: []ftoken{&directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: 12}}}},
		},
		{
			format:         "~'0d",
			expectedTokens: []ftoken{&directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{charParam: '0'}}}},
		},
		{
			format:         "~12,'xd",
			expectedTokens: []ftoken{&directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: 12}, prefixParam{charParam: 'x'}}}},
		},
		{
			format:         "foo ~:a bar",
			expectedTokens: []ftoken{NewLiteral("foo "), &directive{char: rune('a'), colonMod: true}, NewLiteral(" bar")},
		},
		{
			format:            "~{~a~^, ~}",
			expectedTokens:    []ftoken{NewCharDir('{'), NewCharDir('a'), NewCharDir('^'), NewLiteral(", "), NewCharDir('}')},
			expectedRepeation: []repeatDef{repeatDef{index: 4, linksTo: 0}},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.format, func(t *testing.T) {
			node := parseFormatGraph(tc.format)
			var nodeExpected ftoken
			graph := buildGraph(tc.expectedTokens, tc.expectedRepeation)
			nodeExpected = &graph
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

func buildGraph(tokens []ftoken, repeatDef []repeatDef) root {
	rootDir := root{}
	var curDir ftoken
	curDir = &rootDir
	for ix, token := range tokens {
		dir, ok := token.(*directive)
		if ok {
			dir.controlDef = getControlDef(dir.char)
		}
		curDir.SetNext(token)
		curDir = token

		for _, repeat := range repeatDef {
			if repeat.index == ix {
				token.SetRepeatRef(tokens[repeat.linksTo])
			}
		}
	}

	return rootDir
}
