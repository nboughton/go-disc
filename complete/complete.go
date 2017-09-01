package complete

import (
	"errors"
	"strings"
)

// Exported vars
var (
	ErrNoMatch = errors.New("No matching string")
)

// Data contains the reference data for a tab complete
type Data struct {
	index map[string]int
}

// New returns a new completion object.
func New() *Data {
	idx := make(map[string]int)
	return &Data{idx}
}

// Add adds a string to the dictionary
func (d *Data) Add(str string) {
	d.index[str]++
}

// Tab pseudo randomly attempts to get a string
// from the dictionary
func (d *Data) Tab(str string) (string, error) {
	for word := range d.index {
		if strings.HasPrefix(word, str) {
			s := strings.TrimPrefix(word, str)
			return s, nil
		}
	}

	return "", ErrNoMatch
}
