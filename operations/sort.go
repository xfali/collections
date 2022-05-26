// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package operations

import "sort"

//func (s SortedOperands) Len() int {
//	return len(s)
//}
//
//func (s SortedOperands) Less(i, j int) bool {
//	return s[i].LT(s[j])
//}
//
//func (s SortedOperands) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}

func SortOperands(v OperandSet) OperandSet {
	sort.Sort(v)
	return v
}
