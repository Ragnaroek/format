package ft

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseFormatGraph(t *testing.T) {
	tcs := []struct {
		format        string
		expectedToken *root
	}{
		{
			format:        "string only",
			expectedToken: &root{children: []ftoken{NewLiteral("string only")}},
		},
		{
			format:        "~a",
			expectedToken: &root{children: []ftoken{fdir(directive{char: rune('a')})}},
		},
		{
			format:        "~:a",
			expectedToken: &root{children: []ftoken{fdir(directive{char: rune('a'), colonMod: true})}},
		},
		{
			format:        "~:@a",
			expectedToken: &root{children: []ftoken{fdir(directive{char: rune('a'), colonMod: true, atMod: true})}},
		},
		{
			format:        "~@:a",
			expectedToken: &root{[]ftoken{fdir(directive{char: rune('a'), colonMod: true, atMod: true})}},
		},
		{
			format:        "~12d",
			expectedToken: &root{[]ftoken{fdir(directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: 12}}})}},
		},
		{
			format:        "~-12d",
			expectedToken: &root{[]ftoken{fdir(directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: -12}}})}},
		},
		{
			format:        "~+12d",
			expectedToken: &root{[]ftoken{fdir(directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: 12}}})}},
		},
		{
			format:        "~'0d",
			expectedToken: &root{[]ftoken{fdir(directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{charParam: '0'}}})}},
		},
		{
			format:        "~12,'xd",
			expectedToken: &root{[]ftoken{fdir(directive{char: rune('d'), prefixParam: []prefixParam{prefixParam{numParam: 12}, prefixParam{charParam: 'x'}}})}},
		},
		{
			format:        "~10,,,,2:@r",
			expectedToken: &root{[]ftoken{fdir(directive{char: rune('r'), atMod: true, colonMod: true, prefixParam: []prefixParam{prefixParam{numParam: 10}, prefixParam{empty: true}, prefixParam{empty: true}, prefixParam{empty: true}, prefixParam{numParam: 2}}})}},
		},
		{
			format:        "foo ~:a bar",
			expectedToken: &root{[]ftoken{NewLiteral("foo "), fdir(directive{char: rune('a'), colonMod: true}), NewLiteral(" bar")}},
		},
		{
			format:        "~{~a~^, ~}",
			expectedToken: &root{[]ftoken{&root{children: []ftoken{fdir(directive{char: 'a'}), fdir(directive{char: '^'}), NewLiteral(", ")}}}},
		},
		//err cases
		{
			format:        "~2", //missing directive
			expectedToken: expectErr("endofformat"),
		},
		{
			format:        "~{", //unbalanced }
			expectedToken: expectErr("rootpeer"),
		},
		{
			format:        "~{~}~}", //unbalanced }
			expectedToken: expectErr("nopeer"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.format, func(t *testing.T) {
			root := parseFormatGraph(tc.format)

			if len(root.children) != len(tc.expectedToken.children) {
				t.Fatalf("expected %d children, got %d", len(tc.expectedToken.children), len(root.children))
			}

			cmpToken(t, tc.expectedToken, root)
		})
	}
}

func cmpToken(t *testing.T, nodeExpected, nodeActual ftoken) {
	rootActual, ok := nodeActual.(*root)
	if ok {
		rootExpected, okExp := nodeExpected.(*root)
		if !okExp {
			t.Errorf("actual is root node, but expected is not")
		}

		for i, nodeExpected := range rootExpected.children {
			nodeActual := rootActual.children[i]
			cmpToken(t, nodeExpected, nodeActual)
		}
	} else {
		if !reflect.DeepEqual(nodeExpected, nodeActual) {
			t.Errorf("\nexpected %#v,\ngot      %#v", nodeExpected, nodeActual)
		}
	}
}

func fdir(directive directive) *directive {
	cdef := getControlDef(directive.char)
	directive.controlDef = cdef
	return &directive
}

func expectErr(code string) *root {
	return &root{children: []ftoken{NewLiteral(generalError(errors.New(code)))}}
}
