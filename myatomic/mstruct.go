package myatomic

import "sync"

type IntMapM struct {
	lock sync.Mutex
	Map  map[int]int
}

func (i *IntMapM) Add(key, val int) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.Map[key] = val
	return nil
}

func (i *IntMapM) Get(key int) (int, error) {
	i.lock.Lock()
	defer i.lock.Unlock()

	num, ok := i.Map[key]
	if ok {
		return num, nil
	}

	return 0, ErrorNoKeyInMap
}
