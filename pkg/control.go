package format

import (
	"strings"
)

var controlTable = make(map[rune]*controlDef)

func init() {
	//Tested
	AddControl(NewControlDef('c', applyC))
	AddControl(NewNoArgControlDef('%', applyPercent))
	AddControl(NewNoArgControlDef('&', applyAmp))

	//Untested yet
	AddRepeatingControl('{', '}')
	AddControl(NewControlDef('a', applyA))
	AddControl(NewControlDef('d', applyD))
	AddControl(NewControlDef('^', applyCircumflex))

}

type ApplyFn func(interface{}, *directive, *strings.Builder) string

type controlDef struct {
	controlChar rune
	consumesArg bool
	repeatStart bool
	repeatEnd   bool
	peerChar    rune //only relevant if repeatStart or repeatEnd is true
	applyFn     ApplyFn
}

func NewControlDef(char rune, fn ApplyFn) controlDef {
	return controlDef{controlChar: char, applyFn: fn, consumesArg: true}
}

//Create a control that consumes no args
func NewNoArgControlDef(char rune, fn ApplyFn) controlDef {
	return controlDef{controlChar: char, applyFn: fn, consumesArg: false}
}

func AddRepeatingControl(startChar rune, endChar rune) {
	AddControl(controlDef{controlChar: startChar, repeatStart: true, peerChar: endChar, consumesArg: true})
	AddControl(controlDef{controlChar: endChar, repeatEnd: true, peerChar: startChar, consumesArg: true})
}

func AddControl(def controlDef) {
	//TOOD Check if def not already defined
	controlTable[def.controlChar] = &def
}

func getControlDef(char rune) *controlDef {
	return controlTable[char]
}
