package workers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorker(t *testing.T) {
	ch := make(chan struct{})
	go Worker(ch)
	_, ok := <-ch
	assert.True(t, ok, "No value in channel")
}

func TestWorkerIncrement(t *testing.T) {
	count := 0
	handler := func() { count++ }
	testTable := []struct {
		name            string
		workersQuantity int
		handler         func()
		expect          int
	}{
		{
			name:            "first",
			workersQuantity: 10,
			handler:         handler,
			expect:          10,
		},
		{
			name:            "second",
			workersQuantity: 100,
			handler:         handler,
			expect:          100,
		},
		{
			name:            "third",
			workersQuantity: 1000,
			handler:         handler,
			expect:          1000,
		},
	}

	for _, tc := range testTable {
		count = 0
		WorkerHandler(tc.workersQuantity, tc.handler)
		assert.Equal(t, tc.expect, count)
	}
}
