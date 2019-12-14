package format

var controlTable = make(map[rune]*controlDef)

func init() {
	AddRepeatingControl('{', '}')
	AddControl(NewControlDef('a'))
	AddControl(NewControlDef('d'))
	AddControl(NewControlDef('^'))
}

type controlDef struct {
	controlChar rune
	repeatStart bool
	repeatEnd   bool
	peerChar    rune //only relevant if repeatStart or repeatEnd is true
	//TODO controlRenderFn
}

func NewControlDef(char rune) controlDef {
	return controlDef{controlChar: char}
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
