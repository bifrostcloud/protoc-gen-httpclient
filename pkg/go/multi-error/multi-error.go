package multierror

import (
	"errors"
	"strings"
	"sync"
)

// MultiError  -
type MultiError struct {
	sync.Mutex
	Errs []error
}

// Add -
func (m *MultiError) Add(err string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.Errs = append(m.Errs, errors.New(err))
}

// HasError -
func (m *MultiError) HasError() error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	if len(m.Errs) == 0 {
		return nil
	}

	return m
}

// Error -
func (m *MultiError) Error() string {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	formattedError := make([]string, len(m.Errs))
	for i, e := range m.Errs {
		formattedError[i] = e.Error()
	}

	return strings.Join(formattedError, ", ")
}
