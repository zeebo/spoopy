package spoopy

import "bytes"

// Context is the type passed into a Sink so that it can generate all of the
// output that it wants.
type Context interface {
	// Sources returns the list of available sources.
	Sources() []string

	// Lookup takes in a source name and returns the Source.
	Lookup(source string) Source

	// File indicates that the contents of the returned *bytes.Buffer should
	// be written to the path. The path is relative to the root output folder
	// and cannot escape it.
	File(path string) *bytes.Buffer
}

// Source is a type that represents some source of data.
type Source interface {
	// Name returns the name of the source.
	Name() string

	// Value is the associated data.
	Value() interface{}
}

// Sink is a type that can generate output for some Context. If an error is
// returned, no output is generated.
type Sink interface {
	Run(ctx Context) error
}

// SinkFunc is a helper type that implements Sink for a function.
type SinkFunc func(Context) error

// Run implements the Sink interface.
func (s SinkFunc) Run(ctx Context) error {
	return f(ctx)
}
