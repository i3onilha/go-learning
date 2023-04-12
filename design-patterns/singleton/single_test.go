package singleton

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

var errs = map[string]int{}

func TestSingleton(t *testing.T) {
	instances := 30
	wg.Add(instances)
	for i := 0; i < instances; i++ {
		go func() {
			defer wg.Done()
			_, err := GetInstance()
			if err != nil {
				errs[err.Error()]++
			} else {
				errs["nil"]++
			}
		}()
	}
	wg.Wait()
	if errs["nil"] > 1 {
		t.Error("More than one instance created")
	}
}
