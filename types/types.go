package types

// Value is a generic interface to which all data returned by this library 
// comply.
type Value interface {
	String() string
	Equals(Value) bool
}

// Sequence is a Value interface that provides a common abstraction for 
// working with all EDN sequences (Lists, Maps, Sets, and Vectors)
type Sequence interface {
	Value

	Into(Sequence) Sequence
	Each(func(Value, int))
	Length() int
}
