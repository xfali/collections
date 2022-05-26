// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package operations

type NewOperandSetFunc func(size ...int) OperandSet

func MergeOperands(newFunc NewOperandSetFunc, sort bool, datas ...OperandSet) OperandSet {
	if len(datas) == 0 {
		return nil
	} else if len(datas) == 1 {
		return datas[0]
	}

	size := 0
	for _, v := range datas {
		size += v.Len()
	}
	all := newFunc(size)
	size = 0
	for _, v := range datas {
		all = all.Copy(v, size)
		size += v.Len()
	}
	if !sort {
		return all
	}
	return SortOperands(all)
}
