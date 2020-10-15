package semaphore

import "errors"

const (
	errorAcquire = "error acquire semaphore"
	errorRelease = "error release semaphore"
)

type Semaphore interface {
	Acquire() error
	Release() error
}

type semaphore struct {
	channel chan struct{}
}

func (s *semaphore) Acquire() error {
	select {
	case s.channel <- struct{}{}:
		return nil
	default:
		return errors.New(errorAcquire)
	}
}

func (s *semaphore) Release() error {
	select {
	case <-s.channel:
		return nil
	default:
		return errors.New(errorRelease)
	}
}

func New(bufferSize int) Semaphore {
	return &semaphore{channel: make(chan struct{}, bufferSize)}
}
