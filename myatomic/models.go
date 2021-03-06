package myatomic

import (
	"github.com/akaKAIN/go-gb-level-2/myerrors"
	"sync"
)

var (
	ErrorIndexOutOfRange = myerrors.NewErrorTime("Index is out of range")
	ErrorNoKeyInMap = myerrors.NewErrorTime("Key did not exist in map")
)

type Buffer interface {
	Add(int, int) error
	Get(int) (int, error)
}

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

func (i *IntArray) Add(ind, num int) error {
	if _, err := i.Replace(ind, num); err != nil {
		return err
	}
	return nil
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

type IntMap struct {
	lock sync.RWMutex
	Map map[int]int
}

func (i *IntMap) Add(key, val int) error {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.Map[key] = val
	return nil
}

func (i *IntMap) Get(key int) (int, error) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	num, ok := i.Map[key]
	if ok {
		return num, nil
	}
	return 0, ErrorNoKeyInMap
}