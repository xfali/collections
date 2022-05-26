// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package operations

type Int64Operand int64

type Int64Pair struct {
	Start int64
	End   int64
}

type Int64Set []Int64Pair

func MakeInt64OperandSet(size ...int) OperandSet {
	if len(size) == 2 {
		return Int64Set(make([]Int64Pair, size[0], size[1]))
	} else if len(size) == 1 {
		return Int64Set(make([]Int64Pair, size[0]))
	}
	return nil
}

func (p Int64Pair) K() Operand {
	return Int64Operand(p.Start)
}

func (p Int64Pair) V() Operand {
	return Int64Operand(p.End)
}

func (s Int64Set) Len() int {
	return len(s)
}

func (s Int64Set) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}

func (s Int64Set) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Int64Set) Copy(set OperandSet, offset int) OperandSet {
	copy(s[offset:], []Int64Pair(set.(Int64Set)))
	return s
}

func (s Int64Set) Append(k, v Operand) OperandSet {
	ret := append([]Int64Pair(s), Int64Pair{
		Start: int64(k.(Int64Operand)),
		End:   int64(v.(Int64Operand)),
	})
	return Int64Set(ret)
}

func (s Int64Set) Iterator() Iterator {
	return &Int64Iterator{
		cur:   0,
		pairs: s,
	}
}

type Int64Iterator struct {
	pairs []Int64Pair
	cur   int
}

func (it *Int64Iterator) HasNext() bool {
	return it.cur < len(it.pairs)
}

func (it *Int64Iterator) Next() Pair {
	v := it.pairs[it.cur]
	it.cur++
	return v
}

// Greater than
func (o Int64Operand) GT(other Operand) bool {
	return o > other.(Int64Operand)
}

// Greater than or equal
func (o Int64Operand) GE(other Operand) bool {
	return o >= other.(Int64Operand)
}

// Less than
func (o Int64Operand) LT(other Operand) bool {
	return o < other.(Int64Operand)
}

// Less than or equal
func (o Int64Operand) LE(other Operand) bool {
	return o <= other.(Int64Operand)
}

// Not equal
func (o Int64Operand) NE(other Operand) bool {
	return o != other.(Int64Operand)
}

// Equal
func (o Int64Operand) EQ(other Operand) bool {
	return o == other.(Int64Operand)
}

// Increase 1 step
func (o Int64Operand) SelfIncreasing() Operand {
	return o + 1
}

// Decrease 1 step
func (o Int64Operand) SelfDecreasing() Operand {
	return o - 1
}
