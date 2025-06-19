package pluggable

type IndentWriter interface {
	Intro(format string, args ...any)
	AttrsWhere(at Locatable)
	TextAttr(field string, value string)
	ListAttr(field string)
	EndList()
	NestedAttr(field string, obj Describable)
	EndAttrs()

	// And to cope with everything else
	Indent()
	UnIndent()
	IndPrintf(format string, args ...any)
}
