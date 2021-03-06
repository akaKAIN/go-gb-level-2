package myatomic

import (
	"reflect"
	"testing"
)

func TestNewMutexIntArray(t *testing.T) {
	testTable := []struct {
		name   string
		inArr  []int
		expect int
	}{
		{
			name:   "test length for 4",
			inArr:  []int{1, 2, 3, 4},
			expect: 4,
		},
		{
			name:   "test length for 0",
			inArr:  []int{},
			expect: 0,
		},
		{
			name:   "test length for 2",
			inArr:  []int{0, 3},
			expect: 2,
		},
		{
			name:   "test arr is nil",
			inArr:  nil,
			expect: 0,
		},
	}

	for _, tc := range testTable {
		ia := NewMutexIntArray(tc.inArr...)
		if tc.expect != len(ia.ArrayBody) {
			t.Fatalf("%s: expect: %d, got: %d\n", tc.name, tc.expect, len(ia.ArrayBody))
		}
	}
}

func TestIntArray_Push(t *testing.T) {
	testTable := []struct {
		name      string
		inject    int
		inArr     []int
		expectArr []int
	}{
		{
			name:      "0",
			inject:    10,
			inArr:     []int{},
			expectArr: []int{10},
		},
		{
			name:      "-3, 0, 10",
			inject:    -15,
			inArr:     []int{-3, 0, 10},
			expectArr: []int{-3, 0, 10, -15},
		},
		{
			name:      "1,2,3",
			inject:    4,
			inArr:     []int{1, 2, 3},
			expectArr: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testTable {
		ai := NewMutexIntArray(tc.inArr...)
		ai.Push(tc.inject)
		if !reflect.DeepEqual(ai.ArrayBody, tc.expectArr) {
			t.Fatalf("%s: must by equal - expect: %v, got %v", tc.name, tc.expectArr, ai.ArrayBody)
		}
	}
}
