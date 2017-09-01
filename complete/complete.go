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
	index map[string]bool
}

// New returns a new completion object.
func New() *Data {
	idx := make(map[string]bool)
	return &Data{idx}
}

// Add adds a string to the dictionary
func (d *Data) Add(str string) {
	d.index[str] = true
}

// Tab pseudo randomly attempts to get a string
// from the dictionary
func (d *Data) Tab(str string) (string, error) {
	for key := range d.index {
		if strings.HasPrefix(key, str) {
			ing := strings.TrimPrefix(key, str)
			return ing, nil
		}
	}

	return "", ErrNoMatch
}
