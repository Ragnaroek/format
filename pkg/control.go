package format

var controlTable = make(map[rune]*controlDef)

func init() {
	AddRepeatingControl('{', '}')
	AddControl(NewControlDef('a', applyA))
	AddControl(NewControlDef('d', applyD))
	AddControl(NewControlDef('^', applyCircumflex))
}

type ApplyFn func(interface{}) string

type controlDef struct {
	controlChar rune
	repeatStart bool
	repeatEnd   bool
	peerChar    rune //only relevant if repeatStart or repeatEnd is true
	applyFn     ApplyFn
}

func NewControlDef(char rune, fn ApplyFn) controlDef {
	return controlDef{controlChar: char, applyFn: fn}
}

func AddRepeatingControl(startChar rune, endChar rune) {
	AddControl(controlDef{controlChar: startChar, repeatStart: true, peerChar: endChar})
	AddControl(controlDef{controlChar: endChar, repeatEnd: true, peerChar: startChar})
}

func AddControl(def controlDef) {
	//TOOD Check if def not already defined
	controlTable[def.controlChar] = &def
}

func getControlDef(char rune) *controlDef {
	return controlTable[char]
}
