package ft

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var fgCache = make(map[string]root, 0)

func Format(format string, a ...interface{}) {
	print(Sformat(format, a...))
}

func Sformat(format string, a ...interface{}) string {
	//TODO cache format graph
	fg, found := fgCache[format]
	if !found {
		fg = parseFormatGraph(format)
		fgCache[format] = fg
	}
	return applyFormat(fg, a...)
}

func applyFormat(root root, a ...interface{}) string {
	var result strings.Builder

	node := root.Next()
	argPtr := 0
	for node != nil {
		var str string
		if node.ConsumesArg() {
			str = node.Format(a[argPtr], &result)
			argPtr++
		} else {
			str = node.Format(nil, &result)
		}
		result.WriteString(str)

		node = node.Next()
	}

	return result.String()
}

func parseFormatGraph(format string) root {
	runes := []rune(format)
	var parenStack []*directive
	rootDir := root{}
	var curDirective ftoken = &rootDir
	literalBuf := make([]rune, 0)
	for i := 0; i < len(runes); i++ {
		if runes[i] == '~' {
			if len(literalBuf) != 0 {
				ltoken := literal{literal: string(literalBuf)}
				literalBuf = make([]rune, 0)
				curDirective.SetNext(&ltoken)
				curDirective = &ltoken
			}
			directive, skip, err := parseDirective(i, runes)
			if err != nil {
				panic(err) //TODO add err directive to graph, this will result in an error in the oupt
			}
			i += skip
			curDirective.SetNext(&directive)
			curDirective = &directive

			if directive.controlDef.repeatStart {
				parenStack = append(parenStack, &directive)
			}
			if directive.controlDef.repeatEnd {
				n := len(parenStack) - 1
				if n < 0 {
					panic("no peer for repeat control found") //TODO add error directive
				}
				popped := parenStack[n]
				parenStack = parenStack[:n] //pop
				if directive.controlDef.peerChar != popped.char {
					panic(fmt.Sprintf("unbalanced nested controls, expected %c, got %c", directive.controlDef.peerChar, popped.char)) //TODO add error directive
				}
				directive.SetRepeatRef(popped)
			}
		} else {
			literalBuf = append(literalBuf, runes[i]) //OPT: propably too slow
		}
	}

	if len(literalBuf) != 0 {
		ltoken := literal{literal: string(literalBuf)}
		curDirective.SetNext(&ltoken)
		curDirective = &ltoken
	}
	return rootDir
}

//~num|char_t[,num|char_t][:][@]char
//char_t = 'char
func parseDirective(start int, format []rune) (directive, int, error) {

	i := start //i always points to the character currently parsed
	if format[i] != '~' {
		return directive{}, 0, fmt.Errorf("unexpected start of format")
	}
	i++

	next, err := nextChar(i, format)
	if err != nil {
		return directive{}, 0, err
	}

	atMod := false
	colonMod := false
	var prefixParams []prefixParam = nil
	empty := true
	for {
		if unicode.IsDigit(next) {
			numParam, l, err := parseNum(i, format)
			if err != nil {
				return directive{}, 0, err
			}
			empty = false
			prefixParams = append(prefixParams, prefixParam{
				numParam: numParam,
			})
			i += l
		} else if next == '-' {
			numParam, l, err := parseNum(i+1, format)
			if err != nil {
				return directive{}, 0, err
			}
			empty = false
			prefixParams = append(prefixParams, prefixParam{
				numParam: -1 * numParam,
			})
			i += 1 + l
		} else if next == '+' {
			numParam, l, err := parseNum(i+1, format)
			if err != nil {
				return directive{}, 0, err
			}
			empty = false
			prefixParams = append(prefixParams, prefixParam{
				numParam: numParam,
			})
			i += 1 + l
		} else if next == '\'' {
			i++
			charParam, err := nextChar(i, format)
			if err != nil {
				return directive{}, 0, err
			}
			empty = false
			prefixParams = append(prefixParams, prefixParam{
				charParam: charParam,
			})
			i++
		} else if next == ':' {
			colonMod = true
			i++
		} else if next == '@' {
			atMod = true
			i++
		} else if next == ',' {
			if empty {
				prefixParams = append(prefixParams, prefixParam{empty: true})
			}
			empty = true
			i++
		} else {

			controlChar := format[i]
			controlDef := getControlDef(controlChar)
			if controlDef == nil {
				return directive{}, 0, fmt.Errorf("control char %c not found", controlChar) //TODO add error directive instead
			}

			dir := directive{
				atMod:       atMod,
				colonMod:    colonMod,
				prefixParam: prefixParams,
				char:        controlChar,
				controlDef:  controlDef,
			}
			return dir, (i - start), nil
		}

		next, err = nextChar(i, format)
		if err != nil {
			return directive{}, 0, err
		}
	}
}

func parseNum(start int, format []rune) (int, int, error) {
	i := start
	numStr := make([]rune, 0)
	for {
		if !unicode.IsDigit(format[i]) {
			n, err := strconv.Atoi(string(numStr))
			if err != nil {
				return 0, 0, err
			}
			return n, len(numStr), nil
		} else {
			numStr = append(numStr, format[i])
		}
		i++
	}
}

func nextChar(i int, format []rune) (rune, error) {
	if i >= len(format) {
		return '0', fmt.Errorf("nextChar not found")
	}
	return format[i], nil
}

type ftoken interface {
	ConsumesArg() bool

	Repeats() bool
	RepeatRef() ftoken
	SetRepeatRef(ftoken)

	Next() ftoken
	SetNext(ftoken)

	Format(interface{}, *strings.Builder) string
}

type root struct {
	next ftoken
}

func (l *root) ConsumesArg() bool {
	return false
}

func (l *root) Next() ftoken {
	return l.next
}

func (l *root) SetNext(token ftoken) {
	l.next = token
}

func (l *root) Format(_ interface{}, _ *strings.Builder) string {
	return ""
}

func (l *root) Repeats() bool {
	return false
}

func (l *root) RepeatRef() ftoken {
	return nil
}

func (l *root) SetRepeatRef(_ ftoken) {
}

type literal struct {
	literal string
	next    ftoken
}

func NewLiteral(lit string) ftoken {
	return &literal{literal: lit}
}

func (l *literal) ConsumesArg() bool {
	return false
}

func (l *literal) Next() ftoken {
	return l.next
}

func (l *literal) SetNext(token ftoken) {
	l.next = token
}

func (l *literal) Format(_ interface{}, _ *strings.Builder) string {
	return l.literal
}

func (l *literal) Repeats() bool {
	return false
}

func (l *literal) RepeatRef() ftoken {
	return nil
}

func (l *literal) SetRepeatRef(_ ftoken) {
}

type directive struct {
	prefixParam []prefixParam
	colonMod    bool
	atMod       bool
	char        rune
	next        ftoken
	repeats     bool
	repeatRef   ftoken
	controlDef  *controlDef
}

func NewCharDir(char rune) ftoken {
	return &directive{char: char}
}

type prefixParam struct {
	numParam  int
	charParam rune
	empty     bool
}

func (l *directive) ConsumesArg() bool {
	return l.controlDef.consumesArg
}

func (l *directive) Next() ftoken {
	return l.next
}

func (l *directive) SetNext(token ftoken) {
	l.next = token
}

func (l *directive) Format(arg interface{}, input *strings.Builder) string {
	return l.controlDef.applyFn(arg, l, input)
}

func (l *directive) Repeats() bool {
	return l.repeats
}

func (l *directive) RepeatRef() ftoken {
	return l.repeatRef
}

func (l *directive) SetRepeatRef(token ftoken) {
	l.repeatRef = token
}
