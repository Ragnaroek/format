package format

import (
	"fmt"
	"reflect"
	"strings"
)

func applyC(arg interface{}, d *directive) string {
	switch v := arg.(type) {
	case rune:
		if d.atMod {
			return "'" + string(v) + "'"
		}
		return string(v)
	default:
		return typeError('c', arg)
	}
}

func applyPercent(_ interface{}, d *directive) string {
	param, ok := singleNumParam(d, 1)
	if !ok {
		return numParamError(d, 1)
	}
	return strings.Repeat("\n", param)
}

func applyA(arg interface{}, d *directive) string {
	switch v := arg.(type) {
	case string:
		return v
	default:
		return typeError('a', arg)
	}
}

func applyD(arg interface{}, d *directive) string {
	return "D"
}

func applyCircumflex(arg interface{}, d *directive) string {
	return "^"
}

//Helpers

func typeError(dirChar rune, arg interface{}) string {
	argType := reflect.TypeOf(arg)
	return fmt.Sprintf("~!%c(%s=%+v)", dirChar, argType.Name(), arg)
}

func numParamError(d *directive, i int) string {
	return fmt.Sprintf("~!%c(prefix.num!=%d)", d.char, i)
}

func singleNumParam(d *directive, defaultValue int) (int, bool) {
	l := len(d.prefixParam)
	if l > 1 {
		return 0, false
	}
	if l == 0 {
		return defaultValue, true
	}
	param := d.prefixParam[0]
	if param.charParam != 0 {
		return 0, false
	}
	return param.numParam, true
}
