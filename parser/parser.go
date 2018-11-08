package parser

type ParserInput interface {
	CurrentCodePoint() rune
	RemainingInput() ParserInput
}

type RuneArrayInput struct {
	slice  []rune
	cursor int
}

type ParserResult struct {
	Result         interface{}
	RemainingInput ParserInput
}

func (t RuneArrayInput) CurrentCodePoint() rune {
	return t.slice[t.cursor]
}

func (t RuneArrayInput) RemainingInput() ParserInput {
	if t.cursor+1 >= len(t.slice) {
		return nil
	}
	return NewRuneArrayInput(t.slice, t.cursor+1)
}

func NewRuneArrayInput(slice []rune, cursor int) RuneArrayInput {
	return RuneArrayInput{slice, cursor}
}

func stringToInput(s string) ParserInput {
	return NewRuneArrayInput([]rune(s), 0)
}

func parserA(input ParserInput) ParserResult {
	if input != nil && input.CurrentCodePoint() == 'A' {
		return ParserResult{'A', input.RemainingInput()}
	} else {
		return ParserResult{nil, input}
	}
}
