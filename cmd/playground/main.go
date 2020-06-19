package main

import (
	"fmt"
	"math"
	"syscall/js"

	"github.com/ragnaroek/format/pkg"
)

func registerCallbacks() {
	js.Global().Set("Sformat", js.FuncOf(jsSformat))
}

func jsSformat(this js.Value, args []js.Value) interface{} {

	if len(args) < 1 {
		panic(fmt.Errorf("Sformat: at least one arg expected"))
	}

	if args[0].Type() != js.TypeString {
		panic(fmt.Errorf("Sformat: first arg must be a string"))
	}

	goVals := make([]interface{}, 0, len(args)-1)
	for _, valArg := range args[1:] {
		goVals = append(goVals, toGoVal(valArg))
	}

	return format.Sformat(args[0].String(), goVals...)
}

func toGoVal(val js.Value) interface{} {

	switch val.Type() {
	case js.TypeUndefined:
		return nil
	case js.TypeNull:
		return nil
	case js.TypeBoolean:
		return val.Bool()
	case js.TypeNumber:
		return toGoNum(val.Float())
	case js.TypeString:
		return val.String()
	default:
		//TODO Support array and map types
		panic(fmt.Sprintf("unsupported type: %s", val.Type()))
	}
}

func toGoNum(f float64) interface{} {
	if math.Mod(f, 1.0) == 0 {
		return int64(f)
	}
	return f
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
