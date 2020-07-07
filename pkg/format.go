package ft

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

var fgCache = make(map[string]*root, 0)

func Format(format string, a ...interface{}) {
	print(Sformat(format, a...))
}

func Sformat(format string, a ...interface{}) string {
	fg, found := fgCache[format]
	if !found {
		fg = parseFormatGraph(format)
		fgCache[format] = fg
	}
	return applyFormat(fg, a...)
}

func applyFormat(rootNode *root, a ...interface{}) string {
	var result strings.Builder

	argPtr := 0
	for _, node := range rootNode.children {
		root, isRoot := node.(*root)
		if isRoot {
			arg := a[argPtr]
			v := reflect.ValueOf(arg)
			if v.Kind() == reflect.Slice {
				for i := 0; i < v.Len(); i++ {
					result.WriteString(applyFormat(root, v.Index(i).Interface()))
				}
			} else {
				result.WriteString(applyFormat(root, arg))
			}
			argPtr++
		} else {
			var str string
			if node.ConsumesArg() {
				str = node.Format(a[argPtr], &result)
				argPtr++
			} else {
				str = node.Format(nil, &result)
			}
			result.WriteString(str)
		}
	}

	return result.String()
}

type recu struct {
	dir  *directive //nil on toplevel
	root *root
}

func parseFormatGraph(format string) *root {
	runes := []rune(format)

	recuStack := []*recu{&recu{dir: nil, root: &root{}}}
	headRecu := recuStack[0]

	var literalBuf bytes.Buffer

	for i := 0; i < len(runes); i++ {
		if runes[i] == '~' {
			if literalBuf.Len() != 0 {
				ltoken := literal{literal: string(literalBuf.Bytes())}
				literalBuf.Reset()
				headRecu.root.children = append(headRecu.root.children, &ltoken)
			}
			directive, skip, err := parseDirective(i, runes)
			if err != nil {
				return errDir(err)
			}
			i += skip

			if directive.controlDef.repeatStart {
				newRecu := &recu{dir: &directive, root: &root{}}
				recuStack = append(recuStack, newRecu)
				headRecu = newRecu
			} else if directive.controlDef.repeatEnd {
				n := len(recuStack) - 1
				if n <= 0 { //there should alway be the top root node (that can't be popped)
					return errDir(errors.New("nopeer"))
				}
				popped := recuStack[n]
				recuStack = recuStack[:n] //pop
				headRecu = recuStack[n-1]
				if directive.controlDef.peerChar != popped.dir.char {
					return errDir(errors.New("balancepeer"))
				}
				headRecu.root.children = append(headRecu.root.children, popped.root)
			} else {
				headRecu.root.children = append(headRecu.root.children, &directive)
			}
		} else {
			literalBuf.WriteRune(runes[i])
		}
	}

	if literalBuf.Len() != 0 {
		ltoken := literal{literal: string(literalBuf.Bytes())}
		headRecu.root.children = append(headRecu.root.children, &ltoken)
	}

	if headRecu.dir != nil {
		return errDir(errors.New("rootpeer"))
	}

	return headRecu.root
}

//~num|char_t[,num|char_t][:][@]char
//char_t = 'char
func parseDirective(start int, format []rune) (directive, int, error) {

	i := start //i always points to the character currently parsed
	if format[i] != '~' {
		return directive{}, 0, fmt.Errorf("startofformat")
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
		if i >= len(format) {
			return 0, 0, errors.New("endofformat")
		}
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

func errDir(err error) *root {
	r := root{}
	r.children = []ftoken{NewLiteral(generalError(err))}
	return &r
}

type ftoken interface {
	ConsumesArg() bool
	Format(interface{}, *strings.Builder) string
}

type root struct {
	children []ftoken
}

func (l *root) ConsumesArg() bool {
	return false
}

func (l *root) Format(_ interface{}, _ *strings.Builder) string {
	return ""
}

type literal struct {
	literal string
}

func NewLiteral(lit string) ftoken {
	return &literal{literal: lit}
}

func (l *literal) ConsumesArg() bool {
	return false
}

func (l *literal) Format(_ interface{}, _ *strings.Builder) string {
	return l.literal
}

type directive struct {
	prefixParam []prefixParam
	colonMod    bool
	atMod       bool
	char        rune
	next        ftoken
	controlDef  *controlDef
}

type prefixParam struct {
	numParam  int
	charParam rune
	empty     bool
}

func NewCharDir(char rune) ftoken {
	return &directive{char: char}
}

func (l *directive) ConsumesArg() bool {
	return l.controlDef.consumesArg
}

func (l *directive) Format(arg interface{}, input *strings.Builder) string {
	return l.controlDef.applyFn(arg, l, input)
}
