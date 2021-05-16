package main

import "strings"

type lexicalAnalyzer struct {
	input    string
	inputArr []string
	current  int
}

func (l *lexicalAnalyzer) Lex(s *yySymType) int {
	if !l.hasNext() {
		return 0
	}
	str := l.next()
	switch n := strings.ToLower(str); n {
	case "harkishen":
		s.s = str
		return HARKISHEN
	case "singh":
		s.s = str
		return SINGH
	default:
		return RANDOM
	}
}

func (l *lexicalAnalyzer) Error(err string) {
	panic(err)
}

func (l *lexicalAnalyzer) next() string {
	l.current++
	return l.inputArr[l.current-1]
}

func (l *lexicalAnalyzer) hasNext() bool {
	return l.current < len(l.inputArr)
}

func main() {
	str := "some string some other string "
	arr := strings.Split(str, " ")
	yyParse(&lexicalAnalyzer{input: str, inputArr: arr})
}
