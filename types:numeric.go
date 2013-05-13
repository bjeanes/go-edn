package edn

import "fmt"
import "math/big"

type Int int64
type Float float64
type Rational big.Rat
type BigInt big.Int

func (f Float) Equals(v Value) bool {
	return f == v
}

func (f Float) String() string {
	return fmt.Sprintf("%f", float64(f))
}

func (i Int) Equals(v Value) bool {
	return i == v
}

func (i Int) String() string {
	return fmt.Sprintf("%d", int64(i))
}
