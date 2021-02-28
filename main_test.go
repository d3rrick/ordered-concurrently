package orderedconcurrently

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// The work that needs to be performed
func work(val interface{}) interface{} {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	return val
}

func Test(t *testing.T) {
	max := 10
	inputChan := make(chan *OrderedInput)
	wg := &sync.WaitGroup{}
	go func() {
		outChan := Process(inputChan, work, 10)
		for out := range outChan {
			t.Log(out.Value)
			wg.Done()
		}
	}()

	// Create work and the associated order
	for work, order := 0, 0; work < max; work, order = work+1, order+1 {
		wg.Add(1)
		input := &OrderedInput{work, order}
		inputChan <- input
	}
	wg.Wait()
}
