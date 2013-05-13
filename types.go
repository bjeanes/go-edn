package edn

type Value interface {
	String() string
	Equals(Value) bool
}

type Sequence interface {
	Value

	Into(Sequence) Sequence
	Each(func(Value, int)) // TODO: consider a Value() method that returns a channel instead
	Length() int
}
