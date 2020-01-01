package format

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func applyC(arg interface{}, d *directive, _ *strings.Builder) string {
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

func applyPercent(_ interface{}, d *directive, _ *strings.Builder) string {
	param, ok := singleNumParam(d, 1)
	if !ok {
		return numParamError(d, 1)
	}
	return strings.Repeat("\n", param)
}

func applyAmp(_ interface{}, d *directive, output *strings.Builder) string {
	param, ok := singleNumParam(d, 1)
	if !ok {
		return numParamError(d, 1)
	}

	outStr := output.String()
	lenOutStr := len(outStr)
	if lenOutStr > 0 && outStr[lenOutStr-1] == '\n' {
		return strings.Repeat("\n", param-1)
	}

	return strings.Repeat("\n", param)
}

func applyVerticalBar(_ interface{}, d *directive, output *strings.Builder) string {
	param, ok := singleNumParam(d, 1)
	if !ok {
		return numParamError(d, 1)
	}
	return strings.Repeat("\x0C", param)
}

func applyTilde(_ interface{}, d *directive, output *strings.Builder) string {
	param, ok := singleNumParam(d, 1)
	if !ok {
		return numParamError(d, 1)
	}
	return strings.Repeat("~", param)
}

func applyR(arg interface{}, d *directive, _ *strings.Builder) string {
	var value int64
	switch v := arg.(type) {
	case int64:
		value = v
	case int:
		value = int64(v)
	default:
		return typeError('r', arg)
	}

	radix, ok := singleNumParam(d, 10)
	if !ok {
		return numParamError(d, 10)
	}

	return strconv.FormatInt(value, radix)
}

func applyA(arg interface{}, d *directive, _ *strings.Builder) string {
	switch v := arg.(type) {
	case string:
		return v
	default:
		return typeError('a', arg)
	}
}

func applyD(arg interface{}, d *directive, _ *strings.Builder) string {
	return "D"
}

func applyCircumflex(arg interface{}, d *directive, _ *strings.Builder) string {
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
