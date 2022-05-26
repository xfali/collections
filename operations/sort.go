// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package operations

import "sort"

func SortOperands(v OperandSet) OperandSet {
	sort.Sort(v)
	return v
}
