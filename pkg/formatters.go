package format

import (
	"fmt"
	"reflect"
)

func applyA(arg interface{}, d *directive) string {
	switch v := arg.(type) {
	case string:
		return v
	default:
		return typeError('a', arg)
	}
}

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
