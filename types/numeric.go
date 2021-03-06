package types

import (
	"fmt"
	"math/big"
)

// TODO: type Decimal (representing EDN fixed point numbers)

// Int represents EDN integers: 1, -5, 1000.
type Int int64

// Float represents EDN floats: 1.234, -1.2925e+9
type Float float64

// Rational represents EDN rationals: 1/100, -7/5
type Rational big.Rat

// BigInt represents EDN arbitrary precision integers.
type BigInt big.Int

// Equals compares the Float to another Value for equality.
func (this Float) Equals(other Value) bool {
	return this == other
}

// String returns the EDN string representation of the Float.
func (this Float) String() string {
	return fmt.Sprintf("%f", float64(this))
}

// Equals compares the Int to another Value for equality.
func (this Int) Equals(other Value) bool {
	return this == other
}

// String returns the EDN string representation of the Int.
func (this Int) String() string {
	return fmt.Sprintf("%d", int64(this))
}
