// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package operations

func UnionOperandSets(newFunc NewOperandSetFunc, sets ...OperandSet) OperandSet {
	all := MergeOperands(newFunc, true, sets...)
	if all == nil || all.Len() == 0 {
		return all
	}

	retMap := make(map[Operand]Operand, all.Len())
	it := all.Iterator()
	for it.HasNext() {
		v := it.Next()
		st := v.First()
		ed := v.Second()

		if len(retMap) == 0 {
			retMap[st] = ed
		} else {
			i := len(retMap)
			for s, e := range retMap {
				if st.LE(s) && ed.GE(s) && ed.LE(e) {
					// 起点早于区间，终点属于区间
					delete(retMap, s)
					retMap[st] = e
				} else if st.LE(s) && ed.GE(e) {
					// 起点早于区间，终点晚于区间
					delete(retMap, s)
					retMap[st] = ed
				} else if ed.GE(e) && st.GE(s) && st.LE(e) {
					// 起点属于区间，终点晚于区间
					retMap[s] = ed
				} else if ed.LT(s) || st.GT(e) {
					i--
					// 无重合区间
					if i == 0 {
						retMap[st] = ed
					}
				}
			}
		}
	}
	ret := newFunc(0, len(retMap))
	for k, v := range retMap {
		ret = ret.Append(k, v)
	}
	return SortOperands(ret)
}
