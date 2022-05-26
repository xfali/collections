// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package operations

func DifferenceOperandSets(newFunc NewOperandSetFunc, src, dst OperandSet) OperandSet {
	retMap := make(map[Operand]Operand, dst.Len())
	it := dst.Iterator()
	for it.HasNext() {
		v := it.Next()
		retMap[v.K()] = v.V()
	}
	it = src.Iterator()
	for it.HasNext() {
		v := it.Next()
		st := v.K()
		ed := v.V()

		for s, e := range retMap {
			if st.LT(s) && ed.GE(s) && ed.LT(e) {
				// 起点早于区间，终点属于区间
				delete(retMap, s)
				ed := ed.SelfIncreasing()
				retMap[ed] = e
			} else if st.LE(s) && ed.GE(e) {
				// 起点早于区间，终点晚于区间
				delete(retMap, s)
			} else if ed.GE(e) && st.GE(s) && st.LE(e) {
				// 起点属于区间，终点晚于区间
				delete(retMap, s)
				st := st.SelfDecreasing()
				retMap[s] = st
			} else if ed.LT(s) || st.GT(e) {
				// 无重合区间
			} else {
				delete(retMap, s)
			}
		}
	}
	ret := newFunc(0, len(retMap))
	for k, v := range retMap {
		ret = ret.Append(k, v)
	}
	return SortOperands(ret)
}
