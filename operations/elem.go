// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package operations

type GreaterOperator interface {
	// Greater than
	GT(other Operand) bool

	// Greater than or equal
	GE(other Operand) bool
}

type LessOperator interface {
	// Less than
	LT(other Operand) bool

	// Less than or equal
	LE(other Operand) bool
}

type EqualOperator interface {
	// Not equal
	NE(other Operand) bool

	// Equal
	EQ(other Operand) bool
}

type Operand interface {
	GreaterOperator

	LessOperator

	EqualOperator

	// Increase 1 step
	SelfIncreasing() Operand

	// Decrease 1 step
	SelfDecreasing() Operand
}
