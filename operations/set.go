// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package operations

import "sort"

type OperandSet interface {
	sort.Interface

	Copy(set OperandSet, offset int) OperandSet
	Append(k, v Operand) OperandSet
	Iterator() Iterator
}

type Pair interface {
	First() Operand
	Second() Operand
}

type Iterator interface {
	HasNext() bool
	Next() Pair
}
