package edn

type Value interface {
	String() string
	Equals(Value) bool
}
