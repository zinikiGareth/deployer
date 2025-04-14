package pluggable

type Definition interface {
	// Where identifies when the definition came into being
	Where() Location

	// What identifies the name of the definition
	What() SymbolType

	// ShortDescription enables clients to describe what they are pointing to in a unique way
	ShortDescription() string

	// DumpTo renders the whole of the text of the definition in a reproducible and unique, but not necessarily parseable form
	DumpTo(to IndentWriter)

	// Resolve asks the definition to examine all of its structure and ask for resolution of any unresolved names
	Resolve(r Resolver)

	// Execute the action in a given context
	Execute(runtime RuntimeStorage)
}
