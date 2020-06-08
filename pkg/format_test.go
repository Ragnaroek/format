package format

import (
	"reflect"
	"testing"
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
			format:         "~10,,,,2:@r",
			expectedTokens: []ftoken{&directive{char: rune('r'), atMod: true, colonMod: true, prefixParam: []prefixParam{prefixParam{numParam: 10}, prefixParam{empty: true}, prefixParam{empty: true}, prefixParam{empty: true}, prefixParam{numParam: 2}}}},
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
			root := parseFormatGraph(tc.format)
			var node ftoken = &root
			var nodeExpected ftoken
			graph := buildGraph(tc.expectedTokens, tc.expectedRepeation)
			nodeExpected = &graph
			if node == nil {
				t.Errorf("node nil")
			}

			var i = 0
			for node.Next() != nil {
				if !reflect.DeepEqual(nodeExpected, node) {
					t.Errorf("expected %#v, actual %#v", nodeExpected, node)
				}
				node = node.Next()
				nodeExpected = nodeExpected.Next()
				i++
			}
		})
	}
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
