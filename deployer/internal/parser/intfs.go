package parser

type ProvideLine interface {
	HaveLine(lineNo int, text string)
}
