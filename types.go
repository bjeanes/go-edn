package edn

type Value interface {
	String() string
	Equals(Value) bool
}

type Sequence interface {
	Value

	Into(Sequence) Sequence
	Each(func(Value, int))
	Length() int
}
