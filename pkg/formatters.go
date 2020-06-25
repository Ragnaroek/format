package ft

import (
	"fmt"
	"math"
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

func applyD(arg interface{}, d *directive, _ *strings.Builder) string {
	value, ok := valueInt64(arg, d)
	if !ok {
		return typeError(d.char, arg)
	}
	return intFormat(d, value, 10, 0)
}

func applyB(arg interface{}, d *directive, _ *strings.Builder) string {
	value, ok := valueInt64(arg, d)
	if !ok {
		return typeError(d.char, arg)
	}
	return intFormat(d, value, 2, 0)
}

func applyO(arg interface{}, d *directive, _ *strings.Builder) string {
	value, ok := valueInt64(arg, d)
	if !ok {
		return typeError(d.char, arg)
	}
	return intFormat(d, value, 8, 0)
}

func applyX(arg interface{}, d *directive, _ *strings.Builder) string {
	value, ok := valueInt64(arg, d)
	if !ok {
		return typeError(d.char, arg)
	}
	return intFormat(d, value, 16, 0)
}

func applyR(arg interface{}, d *directive, _ *strings.Builder) string {
	value, ok := valueInt64(arg, d)
	if !ok {
		return typeError(d.char, arg)
	}

	if len(d.prefixParam) == 0 {
		if !d.atMod {
			return wordR(value, d.colonMod)
		} else {
			return romanR(value, d.colonMod)
		}
	}

	radix, ok := numParam(0, d, 10)
	if !ok {
		return numParamError(d, 10)
	}

	return intFormat(d, value, radix, 1)
}

func intFormat(d *directive, value int64, radix int, argOffset int) string {
	mincol, ok := numParam(argOffset, d, 0)
	if !ok {
		return numParamError(d, 0)
	}
	padchar, ok := charParam(argOffset+1, d, ' ')
	if !ok {
		return charParamError(d, ' ')
	}
	commaChar, ok := charParam(argOffset+2, d, ',')
	if !ok {
		return charParamError(d, ',')
	}
	commaInterval, ok := numParam(argOffset+3, d, 3)
	if !ok {
		return numParamError(d, 3)
	}

	formatted := strconv.FormatInt(value, radix)
	if d.atMod && value >= 0 {
		formatted = "+" + formatted
	}
	if d.colonMod {
		formatted = formatSeparator(formatted, commaInterval, commaChar)
	}
	return padLeft(formatted, mincol, padchar)
}

func valueInt64(arg interface{}, d *directive) (int64, bool) {
	var value int64
	switch v := arg.(type) {
	case int64:
		value = v
	case int:
		value = int64(v)
	default:
		return 0, false
	}
	return value, true
}

func padLeft(num string, mincol int, padchar rune) string {
	pad := mincol - len([]rune(num))
	if pad > 0 {
		return strings.Repeat(string(padchar), pad) + num
	}
	return num
}

func formatSeparator(num string, interval int, sepChar rune) string {
	if interval == 0 {
		return num
	}
	var builder strings.Builder
	numDigits := []rune(num)

	sepNum := len(numDigits)
	if numDigits[0] == '+' {
		builder.WriteRune('+')
		sepNum--
		numDigits = numDigits[1:]
	} else if numDigits[0] == '-' {
		builder.WriteRune('-')
		sepNum--
		numDigits = numDigits[1:]
	}
	maxSep := (sepNum - 1) / interval
	sep := 0
	for i, c := range numDigits {
		if i > 0 && (sepNum-i)%interval == 0 && sep < maxSep {
			builder.WriteRune(sepChar)
			sep++
		}
		builder.WriteRune(c)
	}
	return builder.String()
}

func romanR(value int64, colonMod bool) string {
	if value <= 0 || value > 3999 {
		return romanError('r')
	}

	strRep := []rune(strconv.FormatInt(value, 10))

	pow := len(strRep)
	var builder strings.Builder
	if colonMod {
		for _, digit := range strRep {
			switch pow {
			case 4:
				builder.WriteString(romanThousand(digit))
			case 3:
				builder.WriteString(oldRomanHundred(digit))
			case 2:
				builder.WriteString(oldRomanTens(digit))
			case 1:
				builder.WriteString(oldRomanSingle(digit))
			default:
			}
			pow--
		}
	} else {
		for _, digit := range strRep {
			switch pow {
			case 4:
				builder.WriteString(romanThousand(digit))
			case 3:
				builder.WriteString(romanHundred(digit))
			case 2:
				builder.WriteString(romanTens(digit))
			case 1:
				builder.WriteString(romanSingle(digit))
			default:
			}
			pow--
		}
	}

	return builder.String()
}

func romanThousand(digit rune) string {
	switch digit {
	case '1':
		return "M"
	case '2':
		return "MM"
	case '3':
		return "MMM"
	default:
		return romanError('r')
	}
}

func romanHundred(digit rune) string {
	switch digit {
	case '0':
		return ""
	case '1':
		return "C"
	case '2':
		return "CC"
	case '3':
		return "CCC"
	case '4':
		return "CD"
	case '5':
		return "D"
	case '6':
		return "DC"
	case '7':
		return "DCC"
	case '8':
		return "DCCC"
	case '9':
		return "CM"
	default:
		return romanError('r')
	}
}

func oldRomanHundred(digit rune) string {
	switch digit {
	case '0':
		return ""
	case '1':
		return "C"
	case '2':
		return "CC"
	case '3':
		return "CCC"
	case '4':
		return "CCCC"
	case '5':
		return "D"
	case '6':
		return "DC"
	case '7':
		return "DCC"
	case '8':
		return "DCCC"
	case '9':
		return "DCCCC"
	default:
		return romanError('r')
	}
}

func romanTens(digit rune) string {
	switch digit {
	case '0':
		return ""
	case '1':
		return "X"
	case '2':
		return "XX"
	case '3':
		return "XXX"
	case '4':
		return "XL"
	case '5':
		return "L"
	case '6':
		return "LX"
	case '7':
		return "LXX"
	case '8':
		return "LXXX"
	case '9':
		return "XC"
	default:
		return romanError('r')
	}
}

func oldRomanTens(digit rune) string {
	switch digit {
	case '0':
		return ""
	case '1':
		return "X"
	case '2':
		return "XX"
	case '3':
		return "XXX"
	case '4':
		return "XXXX"
	case '5':
		return "L"
	case '6':
		return "LX"
	case '7':
		return "LXX"
	case '8':
		return "LXXX"
	case '9':
		return "LXXXX"
	default:
		return romanError('r')
	}
}

func romanSingle(digit rune) string {
	switch digit {
	case '0':
		return ""
	case '1':
		return "I"
	case '2':
		return "II"
	case '3':
		return "III"
	case '4':
		return "IV"
	case '5':
		return "V"
	case '6':
		return "VI"
	case '7':
		return "VII"
	case '8':
		return "VIII"
	case '9':
		return "IX"
	default:
		return romanError('r')
	}
}

func oldRomanSingle(digit rune) string {
	switch digit {
	case '0':
		return ""
	case '1':
		return "I"
	case '2':
		return "II"
	case '3':
		return "III"
	case '4':
		return "IIII"
	case '5':
		return "V"
	case '6':
		return "VI"
	case '7':
		return "VII"
	case '8':
		return "VIII"
	case '9':
		return "VIIII"
	default:
		return romanError('r')
	}
}

func wordR(valueIn int64, colonMod bool) string {
	negative := valueIn < 0
	value := valueIn
	if negative {
		value = -1 * valueIn
	}

	parts := cardinalRRecu([]rune(strconv.FormatInt(value, 10)), 0, colonMod)
	var builder strings.Builder
	if negative {
		builder.WriteString("negative ")
	}
	for i, part := range parts {
		if i != 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(part)
	}
	return builder.String()
}

func cardinalRRecu(num []rune, pow int, colonMod bool) []string {

	if pow >= len(num) {
		return []string{}
	}

	end := len(num) - pow
	start := maxi(0, end-3)
	pack := num[start:end]

	var tens []rune
	var powNum *rune
	if len(pack) == 3 {
		tens = pack[1:]
		pack0 := pack[0]
		if pack0 != '0' {
			powNum = &pack0
		}
	} else {
		tens = pack
		powNum = nil
	}

	var builder strings.Builder
	if powNum != nil {
		builder.WriteString(nameTenCardinal(string(*powNum)))
		builder.WriteString(" ")

		if string(tens) != "00" {
			builder.WriteString("hundred")
			builder.WriteString(" ")
			if colonMod && pow == 0 {
				builder.WriteString(nameTenOrdinal(string(tens)))
			} else {
				builder.WriteString(nameTenCardinal(string(tens)))
			}
		} else {
			if colonMod && pow == 0 {
				builder.WriteString("hundredth")
			} else {
				builder.WriteString("hundred")
			}
		}
	} else {
		if colonMod && pow == 0 {
			builder.WriteString(nameTenOrdinal(string(tens)))
		} else {
			builder.WriteString(nameTenCardinal(string(tens)))
		}
	}

	if pow > 0 {
		builder.WriteString(" ")
		builder.WriteString(namePow(pow))
	}
	recuResult := cardinalRRecu(num, pow+3, colonMod)
	return append(recuResult, builder.String())
}

func namePow(pow int) string {
	switch pow {
	case 3:
		return "thousand"
	case 6:
		return "million"
	case 9:
		return "billion"
	case 12:
		return "trillion"
	case 15:
		return "quadrillion"
	case 18:
		return "quintillion"
	default:
		return ""
	}
}

func nameTenOrdinal(tenIn string) string {
	if len(tenIn) > 2 {
		return "<err ten ordinal name>"
	}
	ten := normaliseTen(tenIn)

	switch ten {
	case "0":
		return "zeroth"
	case "1":
		return "first"
	case "2":
		return "second"
	case "3":
		return "third"
	case "4":
		return "fourth"
	case "5":
		return "fifth"
	case "6":
		return "sixth"
	case "7":
		return "seventh"
	case "8":
		return "eighth"
	case "9":
		return "ninth"
	case "10":
		return "tenth"
	case "11":
		return "eleventh"
	case "12":
		return "twelfth"
	case "13":
		return "thirteenth"
	case "14":
		return "fourteenth"
	case "15":
		return "fifteenth"
	case "16":
		return "sixteenth"
	case "17":
		return "seventeenth"
	case "18":
		return "eighteenth"
	case "19":
		return "nineteenth"
	case "20":
		return "twentieth"
	case "21":
		return "twenty-first"
	case "22":
		return "twenty-second"
	case "23":
		return "twenty-third"
	case "24":
		return "twenty-fourth"
	case "25":
		return "twenty-fifth"
	case "26":
		return "twenty-sixth"
	case "27":
		return "twenty-seventh"
	case "28":
		return "twenty-eighth"
	case "29":
		return "twenty-ninth"
	case "30":
		return "thirtieth"
	case "31":
		return "thirty-first"
	case "32":
		return "thirty-second"
	case "33":
		return "thirty-third"
	case "34":
		return "thirty-fourth"
	case "35":
		return "thirty-fifth"
	case "36":
		return "thirty-sixth"
	case "37":
		return "thirty-seventh"
	case "38":
		return "thirty-eighth"
	case "39":
		return "thirty-ninth"
	case "40":
		return "fortieth"
	case "41":
		return "forty-first"
	case "42":
		return "forty-second"
	case "43":
		return "forty-third"
	case "44":
		return "forty-fourth"
	case "45":
		return "forty-fifth"
	case "46":
		return "forty-sixth"
	case "47":
		return "forty-seventh"
	case "48":
		return "forty-eighth"
	case "49":
		return "forty-ninth"
	case "50":
		return "fiftieth"
	case "51":
		return "fifty-first"
	case "52":
		return "fifty-second"
	case "53":
		return "fifty-third"
	case "54":
		return "fifty-fourth"
	case "55":
		return "fifty-fifth"
	case "56":
		return "fifty-sixth"
	case "57":
		return "fifty-seventh"
	case "58":
		return "fifty-eighth"
	case "59":
		return "fifty-ninth"
	case "60":
		return "sixtieth"
	case "61":
		return "sixty-first"
	case "62":
		return "sixty-second"
	case "63":
		return "sixty-third"
	case "64":
		return "sixty-fourth"
	case "65":
		return "sixty-fifth"
	case "66":
		return "sixty-sixth"
	case "67":
		return "sixty-seventh"
	case "68":
		return "sixty-eighth"
	case "69":
		return "sixty-ninth"
	case "70":
		return "seventieth"
	case "71":
		return "seventy-first"
	case "72":
		return "seventy-second"
	case "73":
		return "seventy-third"
	case "74":
		return "seventy-fourth"
	case "75":
		return "seventy-fifth"
	case "76":
		return "seventy-sixth"
	case "77":
		return "seventy-seventh"
	case "78":
		return "seventy-eighth"
	case "79":
		return "seventy-ninth"
	case "80":
		return "eightieth"
	case "81":
		return "eighty-first"
	case "82":
		return "eighty-second"
	case "83":
		return "eighty-third"
	case "84":
		return "eighty-fourth"
	case "85":
		return "eighty-fifth"
	case "86":
		return "eighty-sixth"
	case "87":
		return "eighty-seventh"
	case "88":
		return "eighty-eighth"
	case "89":
		return "eighty-ninth"
	case "90":
		return "ninetieth"
	case "91":
		return "ninety-first"
	case "92":
		return "ninety-second"
	case "93":
		return "ninety-third"
	case "94":
		return "ninety-fourth"
	case "95":
		return "ninety-fifth"
	case "96":
		return "ninety-sixth"
	case "97":
		return "ninety-seventh"
	case "98":
		return "ninety-eighth"
	case "99":
		return "ninety-ninth"
	default:
		return "<err ten ordinal name>"
	}
}

func normaliseTen(tenIn string) string {
	if len(tenIn) == 2 && tenIn[0] == '0' {
		return string(tenIn[1])
	}
	return tenIn
}

func nameTenCardinal(tenIn string) string {
	if len(tenIn) > 2 {
		return "<err cardinal name>"
	}

	ten := normaliseTen(tenIn)

	switch ten {
	case "0":
		return "zero"
	case "1":
		return "one"
	case "2":
		return "two"
	case "3":
		return "three"
	case "4":
		return "four"
	case "5":
		return "five"
	case "6":
		return "six"
	case "7":
		return "seven"
	case "8":
		return "eight"
	case "9":
		return "nine"
	case "10":
		return "ten"
	case "11":
		return "eleven"
	case "12":
		return "twelve"
	case "13":
		return "thirteen"
	case "14":
		return "fourteen"
	case "15":
		return "fifteen"
	case "16":
		return "sixteen"
	case "17":
		return "seventeen"
	case "18":
		return "eighteen"
	case "19":
		return "nineteen"
	case "20":
		return "twenty"
	case "21":
		return "twenty-one"
	case "22":
		return "twenty-two"
	case "23":
		return "twenty-three"
	case "24":
		return "twenty-four"
	case "25":
		return "twenty-five"
	case "26":
		return "twenty-six"
	case "27":
		return "twenty-seven"
	case "28":
		return "twenty-eight"
	case "29":
		return "twenty-nine"
	case "30":
		return "thirty"
	case "31":
		return "thirty-one"
	case "32":
		return "thirty-two"
	case "33":
		return "thirty-three"
	case "34":
		return "thirty-four"
	case "35":
		return "thirty-five"
	case "36":
		return "thirty-six"
	case "37":
		return "thirty-seven"
	case "38":
		return "thirty-eight"
	case "39":
		return "thirty-nine"
	case "40":
		return "forty"
	case "41":
		return "forty-one"
	case "42":
		return "forty-two"
	case "43":
		return "forty-three"
	case "44":
		return "forty-four"
	case "45":
		return "forty-five"
	case "46":
		return "forty-six"
	case "47":
		return "forty-seven"
	case "48":
		return "forty-eight"
	case "49":
		return "forty-nine"
	case "50":
		return "fifty"
	case "51":
		return "fifty-one"
	case "52":
		return "fifty-two"
	case "53":
		return "fifty-three"
	case "54":
		return "fifty-four"
	case "55":
		return "fifty-five"
	case "56":
		return "fifty-six"
	case "57":
		return "fifty-seven"
	case "58":
		return "fifty-eight"
	case "59":
		return "fifty-nine"
	case "60":
		return "sixty"
	case "61":
		return "sixty-one"
	case "62":
		return "sixty-two"
	case "63":
		return "sixty-three"
	case "64":
		return "sixty-four"
	case "65":
		return "sixty-five"
	case "66":
		return "sixty-six"
	case "67":
		return "sixty-seven"
	case "68":
		return "sixty-eight"
	case "69":
		return "sixty-nine"
	case "70":
		return "seventy"
	case "71":
		return "seventy-one"
	case "72":
		return "seventy-two"
	case "73":
		return "seventy-three"
	case "74":
		return "seventy-four"
	case "75":
		return "seventy-five"
	case "76":
		return "seventy-six"
	case "77":
		return "seventy-seven"
	case "78":
		return "seventy-eight"
	case "79":
		return "seventy-nine"
	case "80":
		return "eighty"
	case "81":
		return "eighty-one"
	case "82":
		return "eighty-two"
	case "83":
		return "eighty-three"
	case "84":
		return "eighty-four"
	case "85":
		return "eighty-five"
	case "86":
		return "eighty-six"
	case "87":
		return "eighty-seven"
	case "88":
		return "eighty-eight"
	case "89":
		return "eighty-nine"
	case "90":
		return "ninety"
	case "91":
		return "ninety-one"
	case "92":
		return "ninety-two"
	case "93":
		return "ninety-three"
	case "94":
		return "ninety-four"
	case "95":
		return "ninety-five"
	case "96":
		return "ninety-six"
	case "97":
		return "ninety-seven"
	case "98":
		return "ninety-eight"
	case "99":
		return "ninety-nine"
	default:
		return "<err cardinal name>"
	}
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func applyF(arg interface{}, d *directive, _ *strings.Builder) string {
	w, ok := numParam(0, d, -1)
	if !ok {
		return numParamError(d, 0)
	}
	de, ok := numParam(1, d, -1)
	if !ok {
		return numParamError(d, 0)
	}
	k, ok := numParam(2, d, 0)
	if !ok {
		return numParamError(d, 0)
	}
	overflowchar, ok := charParamNoDefault(3, d)
	if !ok {
		return charParamError(d, ' ')
	}
	padchar, ok := charParam(4, d, ' ')
	if !ok {
		return charParamError(d, ' ')
	}

	fmt.Printf("%#v, %#v, %#v, %#v\n", de, k, overflowchar, padchar)

	var formatted string
	switch v := arg.(type) {
	case int64:
		formatted = formatInt(v, w)
	case int:
		formatted = formatInt(int64(v), w)
	case float64:
		formatted = formatFloat(v, w)
	case float32:
		formatted = formatFloat(float64(v), w)
	default:
		return typeError('f', arg)
	}

	return formatted
}

func formatInt(v int64, w int) string {
	fomt := strconv.FormatInt(v, 10)
	if w == -1 || len(fomt) <= w-2 {
		return fomt + ".0"
	}
	return fomt + "."
}

func formatFloat(f float64, w int) string {
	if math.Mod(f, 1.0) == 0 {
		return formatInt(int64(f), w)
	}
	fomt := strconv.FormatFloat(f, 'f', -1, 64)

	if w == -1 {
		return fomt
	} else if w < 2 {
		if f < 1.0 {
			return fomt[1:]
		}
		return fomt
	} else {
		if len(fomt) > w {
			if f < 1.0 {
				rounded := strconv.FormatFloat(round(f, w-1), 'f', -1, 64)
				return rounded[1:]
			}
			lenInt := int(math.Log10(f) + 1)
			rounded := strconv.FormatFloat(round(f, w-1-lenInt), 'f', -1, 64)
			return rounded
		}
		return fomt
	}
}

func round(f float64, precision int) float64 {
	p := math.Pow10(precision)
	s := f * p
	return math.Round(s) / p
}

func applyA(arg interface{}, d *directive, _ *strings.Builder) string {
	switch v := arg.(type) {
	case string:
		return v
	default:
		return typeError('a', arg)
	}
}

func applyCircumflex(arg interface{}, d *directive, _ *strings.Builder) string {
	return "^"
}
