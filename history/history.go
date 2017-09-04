// Package history is used to manage the command history of the application
package history

// History wraps the buffer and its current index.
type History struct {
	buffer []string
	idx    int
	log    bool
}

// New returns an empty, zero value'd History struct
func New() *History {
	return &History{}
}

// Logging returns whether or not logging is currently enabled.
func (h *History) Logging() bool {
	return h.log
}

// SetLogging can be used to turn logging on or off
func (h *History) SetLogging(b bool) {
	h.log = b
}

// Log adds a new line to the buffer and sets the new 0 value of the
// index
func (h *History) Log(str string) {
	if h.log {
		h.buffer = append(h.buffer, str)
		h.idx = len(h.buffer)
	}
}

// Prev returns the previous item in the log buffer
func (h *History) Prev() string {
	i := h.idx - 1
	if i >= 0 {
		h.idx = i
		return h.buffer[i]
	}

	return ""
}

// Next returns the next line in the log buffer
func (h *History) Next() string {
	i := h.idx + 1
	if i <= len(h.buffer) {
		h.idx = i
		return h.buffer[i]
	}

	return ""
}
