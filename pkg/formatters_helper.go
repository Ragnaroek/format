package ft

import (
	"fmt"
	"reflect"
	"strings"
)

func romanError(dirChar rune) string {
	return fmt.Sprintf("~!%c(roman range=[1,3999])", dirChar)
}

func typeError(dirChar rune, arg interface{}) string {
	argType := reflect.TypeOf(arg)
	return fmt.Sprintf("~!%c(%s=%+v)", dirChar, argType.Name(), arg)
}

func numParamError(d *directive, i int) string {
	return fmt.Sprintf("~!%c(prefix.num!=%d)", d.char, i)
}

func charParamError(d *directive, r rune) string {
	return fmt.Sprintf("~!%c(prefix.char!=%c)", d.char, r)
}

func singleNumParam(d *directive, defaultValue int) (int, bool) {
	l := len(d.prefixParam)
	if l > 1 {
		return 0, false
	}
	if l == 0 {
		return defaultValue, true
	}
	return extractNumParam(0, d, defaultValue)
}

func numParam(ix int, d *directive, defaultValue int) (int, bool) {
	if ix > len(d.prefixParam)-1 {
		return defaultValue, true
	}
	return extractNumParam(ix, d, defaultValue)
}

func extractNumParam(ix int, d *directive, defaultValue int) (int, bool) {
	param := d.prefixParam[ix]
	if param.empty {
		return defaultValue, true
	}
	if param.charParam != 0 {
		return 0, false
	}
	return param.numParam, true
}

func charParam(ix int, d *directive, defaultValue rune) (rune, bool) {
	if ix > len(d.prefixParam)-1 {
		return defaultValue, true
	}
	r, ok := extractCharParam(ix, d)
	if r == nil {
		return defaultValue, ok
	}
	return *r, ok
}

func charParamNoDefault(ix int, d *directive) (*rune, bool) {
	if ix > len(d.prefixParam)-1 {
		return nil, true
	}
	return extractCharParam(ix, d)
}

func extractCharParam(ix int, d *directive) (*rune, bool) {
	param := d.prefixParam[ix]
	if param.empty {
		return nil, true
	}
	if param.numParam != 0 {
		return nil, false
	}
	return &param.charParam, true
}

func ptrRune(r rune) *rune {
	return &r
}

func padLeft(num string, mincol int, padchar rune) string {
	pad := mincol - len([]rune(num))
	if pad > 0 {
		return strings.Repeat(string(padchar), pad) + num
	}
	return num
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}
