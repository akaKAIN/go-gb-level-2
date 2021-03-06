package myatomic

import (
	"github.com/akaKAIN/go-gb-level-2/myerrors"
	"sync"
)

var ErrorIndexOutOfRange = myerrors.NewErrorTime("Index is out of range")

type IntArray struct {
	ArrayBody []int
	lock      sync.RWMutex
}

func NewMutexIntArray(arr ...int) *IntArray {
	if arr == nil {
		arr = make([]int, 0)
	}
	return &IntArray{
		ArrayBody: arr,
	}
}

func (i *IntArray) Push(newInt int) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.ArrayBody = append(i.ArrayBody, newInt)
}

func (i *IntArray) Replace(ind, newInt int) (int, error) {
	if ind < 0 || ind >= len(i.ArrayBody) {
		return 0, ErrorIndexOutOfRange
	}
	i.lock.Lock()
	defer i.lock.Unlock()
	oldInt := i.ArrayBody[ind]
	i.ArrayBody[ind] = newInt
	return oldInt, nil
}

func (i *IntArray) GetByIndex(ind int) (int, error) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	if ind < 0 || ind >= len(i.ArrayBody) {
		return 0, ErrorIndexOutOfRange
	}
	return i.ArrayBody[ind], nil
}
