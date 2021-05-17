package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

type analyzer struct {
	reader *bufio.Reader
	pos    int
	len    int
}

func (a *analyzer) hasNext() bool {
	return a.pos < a.len
}

func (a *analyzer) next() rune {
	r, _, err := a.reader.ReadRune()
	if err != nil {
		panic(err)
	}
	a.pos++
	return r
}

func (a *analyzer) Lex(s *yySymType) int {
	if !a.hasNext() {
		return 0
	}

	switch char := a.next(); {
	case char == '+':
		s.sym = byte(char)
		return SYMBOL
	case unicode.IsDigit(char):
		s.digit = string(char)
		return DIGIT
	default:
		return RANDOM_STRING
	}
}

func (a *analyzer) Error(err string) {
	panic(err)
}

func concat(items ...string) string {
	var s string
	for i := range items {
		s += fmt.Sprintf("%s%s", s, items[i])
	}
	return s
}

func main() {
	str := "+91797867207"
	anaz := &analyzer{
		reader: bufio.NewReader(strings.NewReader(str)),
		len:    len(str),
		pos:    0,
	}
	yyParse(anaz)
}
