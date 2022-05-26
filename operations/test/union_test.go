// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/collections/operations"
	"testing"
)

func TestInt64Union(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		ret := operations.UnionOperandSets(operations.MakeInt64OperandSet, operations.Int64Set(arr1), operations.Int64Set(arr2))
		arr := []operations.Int64Pair{
			{1, 20}, {22, 23}, {25, 26},
		}
		if ret.Len() != len(arr) {
			t.Fatal("len not match")
		}
		it := ret.Iterator()
		for _, v := range arr {
			x := it.Next()
			t.Log(x)
			if v.First().NE(x.First()) || v.Second().NE(x.Second()) {
				t.Fatal("Value not match")
			}
		}
	})

	t.Run("2", func(t *testing.T) {
		ret := operations.UnionOperandSets(operations.MakeInt64OperandSet, operations.Int64Set(arr2), operations.Int64Set(arr1))
		arr := []operations.Int64Pair{
			{1, 20}, {22, 23}, {25, 26},
		}
		if ret.Len() != len(arr) {
			t.Fatal("len not match")
		}
		it := ret.Iterator()
		for _, v := range arr {
			x := it.Next()
			t.Log(x)
			if v.First().NE(x.First()) || v.Second().NE(x.Second()) {
				t.Fatal("Value not match")
			}
		}
	})
}
