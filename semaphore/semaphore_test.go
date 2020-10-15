package semaphore

import (
	"fmt"
	"sync"
	"testing"
)

func TestSemaphoreCreating(t *testing.T) {
	tikets := 5
	s := New(tikets)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		i := i
		if err := s.Acquire(); err != nil {
			t.Errorf("error -> %v\n", err)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("val=%v\n", i)
			if err := s.Release(); err != nil {
				t.Errorf("error -> %v\n", err)
			}
		}()
	}
	wg.Wait()
}
