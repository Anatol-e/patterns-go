package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	var counter1 Singleton
	counter1 = GetInstance()
	if counter1 == nil {
		t.Fatalf("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	counter1.AddOne()
	currentCount := counter1.GetCount()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	var counter2, expectedCounter Singleton
	expectedCounter = counter1
	counter2 = GetInstance()
	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but it got a different instance")
	}

	counter2.AddOne()
	currentCount = counter2.GetCount()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}
}

func TestParallel(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 5000

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			singleton.AddOne()
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			singleton2.AddOne()
		}()
	}

	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())

	wg.Wait()

	fmt.Printf("Current count is %d\n", singleton.GetCount())

	currentCount1 := singleton.GetCount()
	currentCount2 := singleton2.GetCount()
	if currentCount1 != currentCount2 {
		t.Errorf("Counts not match\nCurrentCount1=%d\nCurrentCount2=%d", currentCount1, currentCount2)
	}

	if currentCount1 != n*2 {
		t.Errorf("Counts not match\nCurrentCount1=%d\nN*2=%d", currentCount1, n*2)
	}
}
