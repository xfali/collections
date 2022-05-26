// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/collections/operations"
	"testing"
)

func TestInt64Diff(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		ret := operations.DifferenceOperandSets(operations.MakeInt64OperandSet, operations.Int64Set(arr1), operations.Int64Set(arr2))
		checkDiff1(ret, t)
	})

	t.Run("2", func(t *testing.T) {
		ret := operations.DifferenceOperandSets(operations.MakeInt64OperandSet, operations.Int64Set(arr2), operations.Int64Set(arr1))
		checkDiff2(ret, t)
	})
}

func TestReflectDiff(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		ret := operations.DifferenceSetsResult(func(i, j int64) int {
			return int(i - j)
		}, func(a int64, i int) int64 {
			return a + int64(i)
		}, arr1, arr2).([]operations.Int64Pair)
		arr := []operations.Int64Pair{
			{1, 1}, {6, 7}, {22, 23},
		}
		if len(ret) != len(arr) {
			t.Fatal("len not match")
		}
		for i := range ret {
			t.Log(ret[i])
			if ret[i].Start != arr[i].Start || ret[i].End != arr[i].End {
				t.Fatal("Value not match")
			}
		}
	})

	t.Run("2", func(t *testing.T) {
		ret := operations.DifferenceSetsResult(func(i, j int64) int {
			return int(i - j)
		}, func(a int64, i int) int64 {
			return a + int64(i)
		}, arr2, arr1).([]operations.Int64Pair)
		arr := []operations.Int64Pair{
			{3, 4}, {10, 14}, {25, 26},
		}
		if len(ret) != len(arr) {
			t.Fatal("len not match")
		}
		for i := range ret {
			t.Log(ret[i])
			if ret[i].Start != arr[i].Start || ret[i].End != arr[i].End {
				t.Fatal("Value not match")
			}
		}
	})
}

func checkDiff1(ret operations.OperandSet, t *testing.T) {
	arr := []operations.Int64Pair{
		{1, 1}, {6, 7}, {22, 23},
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
}

func checkDiff2(ret operations.OperandSet, t *testing.T) {
	arr := []operations.Int64Pair{
		{3, 4}, {10, 14}, {25, 26},
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
}
