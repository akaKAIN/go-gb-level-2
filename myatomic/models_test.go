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

func TestIntArray_GetByIndex(t *testing.T) {
	testTable := []struct {
		name          string
		inArr         []int
		currentInd    int
		expected      int
		expectedError bool
	}{
		{
			name:          "Must return error1",
			inArr:         nil,
			currentInd:    0,
			expected:      0,
			expectedError: true,
		},
		{
			name:          "Must return error2",
			inArr:         []int{1, 2, 3},
			currentInd:    5,
			expected:      0,
			expectedError: true,
		},
		{
			name:          "Must return 2 and no error",
			inArr:         []int{1, 2, 3},
			currentInd:    1,
			expected:      2,
			expectedError: false,
		},
	}

	for _, tc := range testTable {
		ia := NewMutexIntArray(tc.inArr...)
		num, err := ia.GetByIndex(tc.currentInd)
		if num != tc.expected || (err != nil) != tc.expectedError {
			t.Fatalf(
				"%s: expect: %d (error: %t), got: %d (error: %t\n)",
				tc.name,
				tc.expected,
				tc.expectedError,
				num,
				err != nil,
			)
		}
	}
}

func TestIntArray_Replace(t *testing.T) {
	testTable := []struct {
		name          string
		inArr         []int
		currentInd    int
		oldNum        int
		newNum        int
		expectNum     int
		expectArr     []int
		expectedError bool
	}{
		{
			name:          "Must return error1",
			inArr:         nil,
			currentInd:    0,
			oldNum:        0,
			newNum:        1,
			expectNum:     0,
			expectArr:     []int{},
			expectedError: true,
		},
		{
			name:          "Must return error2",
			inArr:         []int{1, 2, 3},
			currentInd:    4,
			oldNum:        0,
			newNum:        4,
			expectNum:     0,
			expectArr:     []int{1, 2, 3},
			expectedError: true,
		},
		{
			name:          "Replacing",
			inArr:         []int{1, 2, 3, 4},
			currentInd:    2,
			oldNum:        3,
			newNum:        5,
			expectNum:     0,
			expectArr:     []int{1, 2, 5, 4},
			expectedError: false,
		},
	}

	for _, tc := range testTable {
		ia := NewMutexIntArray(tc.inArr...)
		oldNum, err := ia.Replace(tc.currentInd, tc.newNum)
		if tc.expectedError != (err != nil) {
			t.Fatalf(
				"%s: expect error: %t, but got error: %v",
				tc.name,
				tc.expectedError,
				err,
			)
		}

		if oldNum != tc.oldNum {
			t.Fatalf(
				"%s: wrong old number: expect %d, got %d\n",
				tc.name,
				tc.expectNum,
				oldNum,
			)
		}

		if !reflect.DeepEqual(tc.expectArr, ia.ArrayBody) {
			t.Fatalf(
				"%s: expect %v, got %v",
				tc.name, tc.expectArr,
				ia.ArrayBody,
			)
		}
	}
}

func BenchmarkIntArray_Push(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	injectNum := 0
	ia := NewMutexIntArray(arr...)
	b.Run("cpu", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ia.Push(injectNum)
				injectNum++
			}
		})
	})
}

func BenchmarkCase90_10(b *testing.B) {
	b.Run("cpu", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				StartReadAndWrite(900, 100)
			}
		})
	})
}
func BenchmarkCase50_50(b *testing.B) {
	b.Run("cpu", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				StartReadAndWrite(500, 500)
			}
		})
	})
}
func BenchmarkCase10_90(b *testing.B) {
	b.Run("cpu", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				StartReadAndWrite(100, 900)
			}
		})
	})
}
