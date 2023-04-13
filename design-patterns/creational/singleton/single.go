package singleton

import (
	"errors"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func GetInstance() (*single, error) {
	var err error
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{}
		} else {
			err = errors.New("Single Instance already created")
		}
	} else {
		err = errors.New("Single Instance already created")
	}
	return singleInstance, err
}
