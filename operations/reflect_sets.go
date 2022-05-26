// Copyright (C) 2020-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package operations

import (
	"fmt"
	"reflect"
)

var (
	reflectPairTag         = "pair"
	reflectPairTag1stValue = "first"
	reflectPairTag2ndValue = "second"
)

type reflectOperand struct {
	field     reflect.Value
	compareFn reflect.Value
	sumFn     reflect.Value
}

type reflectPair struct {
	elem        reflect.Value
	compareFn   reflect.Value
	sumFn       reflect.Value
	firstIndex  int
	secondIndex int
}

type reflectSet struct {
	slice       reflect.Value
	compareFn   reflect.Value
	sumFn       reflect.Value
	firstIndex  int
	secondIndex int
}

func DifferenceSets(compareFn, sumFunc, src, dst interface{}) OperandSet {
	s1 := NewReflectSet(src, compareFn, sumFunc)
	s2 := NewReflectSet(dst, compareFn, sumFunc)
	return DifferenceOperandSets(s1.MakeNewOperandSetFunc(), s1, s2)

}

func DifferenceSetsResult(compareFn, sumFunc, src, dst interface{}) interface{} {
	ret := DifferenceSets(compareFn, sumFunc, src, dst)
	return ret.(*reflectSet).slice.Interface()
}

func UnionSets(compareFn, sumFunc interface{}, sets ...interface{}) OperandSet {
	if len(sets) == 0 {
		return nil
	}
	params := make([]OperandSet, len(sets))
	for i := range sets {
		params[i] = NewReflectSet(sets[i], compareFn, sumFunc)
	}
	if len(sets) == 1 {
		return params[0]
	}
	return UnionOperandSets(params[0].(*reflectSet).MakeNewOperandSetFunc(), params...)
}

func UnionSetsResult(compareFn, sumFunc interface{}, sets ...interface{}) interface{} {
	ret := UnionSets(compareFn, sumFunc, sets...)
	return ret.(*reflectSet).slice.Interface()
}

func NewReflectSet(slice, compareFn, sumFunc interface{}) *reflectSet {
	t := reflect.TypeOf(slice)
	v := reflect.ValueOf(slice)
	checkSlice(t)
	elemType := t.Elem()
	fi, si, opType := checkType(elemType)
	fn := reflect.ValueOf(compareFn)
	if !checkCompareFunction(fn, opType) {
		panic("Compare function Not match, expect func(a, b Type) int")
	}
	sumFn := reflect.ValueOf(sumFunc)
	if !checkSumFunction(sumFn, opType) {
		panic("Sum function Not match, expect func(a Type, steps int) Type")
	}
	return &reflectSet{
		slice:       v,
		compareFn:   fn,
		sumFn:       sumFn,
		firstIndex:  fi,
		secondIndex: si,
	}
}

func (s *reflectSet) MakeNewOperandSetFunc() NewOperandSetFunc {
	return func(size ...int) OperandSet {
		if len(size) == 2 {
			v := reflect.MakeSlice(s.slice.Type(), size[0], size[1])
			return &reflectSet{
				slice:       v,
				compareFn:   s.compareFn,
				sumFn:       s.sumFn,
				firstIndex:  s.firstIndex,
				secondIndex: s.secondIndex,
			}
		} else if len(size) == 1 {
			v := reflect.MakeSlice(s.slice.Type(), 0, size[1])
			return &reflectSet{
				slice:       v,
				compareFn:   s.compareFn,
				sumFn:       s.sumFn,
				firstIndex:  s.firstIndex,
				secondIndex: s.secondIndex,
			}
		}
		return nil
	}
}

func (p *reflectPair) First() Operand {
	return &reflectOperand{
		field:     p.elem.Field(p.firstIndex),
		compareFn: p.compareFn,
		sumFn:     p.sumFn,
	}
}

func (p *reflectPair) Second() Operand {
	return &reflectOperand{
		field:     p.elem.Field(p.secondIndex),
		compareFn: p.compareFn,
		sumFn:     p.sumFn,
	}
}

func (s *reflectSet) Len() int {
	return s.slice.Len()
}

func (s *reflectSet) Less(i, j int) bool {
	var param [2]reflect.Value
	param[0] = s.slice.Index(i).Field(s.firstIndex)
	param[1] = s.slice.Index(j).Field(s.firstIndex)
	return s.compareFn.Call(param[:])[0].Int() < 0
}

func (s *reflectSet) Swap(i, j int) {
	reflect.Swapper(s.slice.Interface())(i, j)
}

func (s *reflectSet) Copy(set OperandSet, offset int) OperandSet {
	reflect.Copy(s.slice.Slice(offset, s.slice.Len()), set.(*reflectSet).slice)
	return s
}

func (s *reflectSet) Append(k, v Operand) OperandSet {
	t := s.slice.Type().Elem()
	elem := reflect.New(t).Elem()
	elem.Field(s.firstIndex).Set(k.(*reflectOperand).field)
	elem.Field(s.secondIndex).Set(v.(*reflectOperand).field)
	s.slice = reflect.Append(s.slice, elem)
	return s
}

func (s *reflectSet) Iterator() Iterator {
	return &ReflectIterator{
		cur: 0,
		set: s,
	}
}

type ReflectIterator struct {
	set *reflectSet
	cur int
}

func (it *ReflectIterator) HasNext() bool {
	return it.cur < it.set.slice.Len()
}

func (it *ReflectIterator) Next() Pair {
	v := it.set.slice.Index(it.cur)
	it.cur++
	return &reflectPair{
		elem:        v,
		compareFn:   it.set.compareFn,
		sumFn:       it.set.sumFn,
		firstIndex:  it.set.firstIndex,
		secondIndex: it.set.secondIndex,
	}
}

// Greater than
func (o *reflectOperand) GT(other Operand) bool {
	var param [2]reflect.Value
	param[0] = o.field
	param[1] = other.(*reflectOperand).field
	return o.compareFn.Call(param[:])[0].Int() > 0
}

// Greater than or equal
func (o *reflectOperand) GE(other Operand) bool {
	var param [2]reflect.Value
	param[0] = o.field
	param[1] = other.(*reflectOperand).field
	return o.compareFn.Call(param[:])[0].Int() >= 0
}

// Less than
func (o *reflectOperand) LT(other Operand) bool {
	var param [2]reflect.Value
	param[0] = o.field
	param[1] = other.(*reflectOperand).field
	return o.compareFn.Call(param[:])[0].Int() < 0
}

// Less than or equal
func (o *reflectOperand) LE(other Operand) bool {
	var param [2]reflect.Value
	param[0] = o.field
	param[1] = other.(*reflectOperand).field
	return o.compareFn.Call(param[:])[0].Int() <= 0
}

// Not equal
func (o *reflectOperand) NE(other Operand) bool {
	var param [2]reflect.Value
	param[0] = o.field
	param[1] = other.(*reflectOperand).field
	return o.compareFn.Call(param[:])[0].Int() != 0
}

// Equal
func (o *reflectOperand) EQ(other Operand) bool {
	var param [2]reflect.Value
	param[0] = o.field
	param[1] = other.(*reflectOperand).field
	return o.compareFn.Call(param[:])[0].Int() == 0
}

// Increase 1 step
func (o *reflectOperand) SelfIncreasing() Operand {
	var param [2]reflect.Value
	param[0] = o.field
	param[1] = reflect.ValueOf(1)
	o.field = o.sumFn.Call(param[:])[0]
	return o
}

// Decrease 1 step
func (o *reflectOperand) SelfDecreasing() Operand {
	var param [2]reflect.Value
	param[0] = o.field
	param[1] = reflect.ValueOf(-1)
	o.field = o.sumFn.Call(param[:])[0]
	return o
}

func checkSlice(t reflect.Type) {
	if t.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Type %t is not slice!", t))
	}
}

func checkType(t reflect.Type) (int, int, reflect.Type) {
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Type %t is not struct!", t))
	}
	s := t.NumField()
	found1st := -1
	found2nd := -1
	var t1, t2 reflect.Type
	for i := 0; i < s; i++ {
		f := t.Field(i)
		if v, ok := f.Tag.Lookup(reflectPairTag); ok {
			if found1st == -1 && v == reflectPairTag1stValue {
				found1st = i
				t1 = f.Type
			} else if found2nd == -1 && v == reflectPairTag2ndValue {
				found2nd = i
				t2 = f.Type
			}
			if found1st != -1 && found2nd != -1 {
				if t1 != t2 {
					panic("Struct's field with tag \"pair\" is different type. ")
				}
				return found1st, found2nd, t1
			}
		}
	}
	panic(`Tag pair:"first" or pair:"second" not found.`)
}

func checkCompareFunction(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 2 || fn.Type().NumOut() != 1 {
		return false
	}
	if elemType != fn.Type().In(0) || elemType != fn.Type().In(1) || reflect.Int != fn.Type().Out(0).Kind() {
		return false
	}
	return true
}

func checkSumFunction(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 2 || fn.Type().NumOut() != 1 {
		return false
	}
	fmt.Println(elemType, fn.Type().In(0), fn.Type().In(1).Kind(), fn.Type().Out(0))
	if elemType != fn.Type().In(0) || reflect.Int != fn.Type().In(1).Kind() || elemType != fn.Type().Out(0) {
		return false
	}
	return true
}
