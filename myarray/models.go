package myarray

import (
	"bytes"
	"errors"
	"sync"
)

var (
	ErrorIndexOutOfRange = errors.New("index out of range")
	ErrorNoMatchFound = errors.New("no match found")
)

type ByteArray struct {
	lock  sync.Mutex
	Array []byte
}

func (b *ByteArray) Lock() {
	b.lock.Lock()
}

func (b *ByteArray) Unlock() {
	b.lock.Unlock()
}

func (b *ByteArray) ReplaceSubString(old, new string) error {
	// Заменяет исходном массиве на переданные данные
	b.Lock()
	defer b.Unlock()

	oldByte := []byte(old)
	if !bytes.Contains(b.Array, oldByte) {
		return ErrorNoMatchFound
	}

	b.Array = bytes.ReplaceAll(b.Array, []byte(old), []byte(new))
	return nil
}

func (b *ByteArray) ReplaceByIndex(ind int, s string) error {
	if ind < 0 {
		return ErrorIndexOutOfRange
	}
	b.Lock()
	defer b.Unlock()

	for {
		if ind < len(b.Array) {
			break
		}
		ind %= len(b.Array)
	}
	b.Array[ind] = []byte(s)[0]
	return nil
}

func (b *ByteArray) GetByteByInd(ind int) (byte, error) {
	if ind < 0 {
		return 0, ErrorIndexOutOfRange
	}
	b.Lock()
	defer b.Unlock()

	for {
		if ind < len(b.Array) {
			break
		}
		ind %= len(b.Array)
	}
	return b.Array[ind], nil
}

func (b *ByteArray) String() string {
	return string(b.Array)
}
