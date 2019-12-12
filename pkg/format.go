package format

import (
	"fmt"
	"strconv"
	"unicode"
)

func Format(format string, a ...interface{}) {
	print(Sformat(format, a...))
}

func Sformat(format string, a ...interface{}) string {
	fg := parseFormatGraph(format)
	fmt.Printf("fg = %#v", fg)
	return ""
}

func parseFormatGraph(format string) ftoken {
	runes := []rune(format)
	var rootDir ftoken
	rootDir = &root{}
	curDirective := rootDir
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
		} else {
			literalBuf = append(literalBuf, runes[i]) //OPT: propably too slow
			//TODO collect runes for literal format directive
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
	for {
		if unicode.IsDigit(next) {
			numParam, l, err := parseNum(i, format)
			if err != nil {
				return directive{}, 0, err
			}
			prefixParams = append(prefixParams, prefixParam{
				numParam: numParam,
			})
			i += l
		} else if next == '-' {
			numParam, l, err := parseNum(i+1, format)
			if err != nil {
				return directive{}, 0, err
			}
			prefixParams = append(prefixParams, prefixParam{
				numParam: -1 * numParam,
			})
			i += 1 + l
		} else if next == '+' {
			numParam, l, err := parseNum(i+1, format)
			if err != nil {
				return directive{}, 0, err
			}
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
			i++
		} else {
			return directive{
				atMod:       atMod,
				colonMod:    colonMod,
				prefixParam: prefixParams,
				char:        format[i],
			}, (i - start), nil
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
	Next() ftoken
	SetNext(ftoken)
	Format(interface{}) string
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

func (l *root) Format(_ interface{}) string {
	return ""
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

func (l *literal) Format(_ interface{}) string {
	return l.literal
}

type directive struct {
	prefixParam []prefixParam
	colonMod    bool
	atMod       bool
	char        rune
	next        ftoken
}

type prefixParam struct {
	numParam  int
	charParam rune
}

func (l *directive) ConsumesArg() bool {
	return true
}

func (l *directive) Next() ftoken {
	return l.next
}

func (l *directive) SetNext(token ftoken) {
	l.next = token
}

func (l *directive) Format(_ interface{}) string {
	return "error"
}
