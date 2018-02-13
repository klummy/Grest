package grest

// ContextKey type to create a new context key
type ContextKey string

func (c ContextKey) String() string {
	return "grest custom key " + string(c)
}
