package mud

import (
	"regexp"
)

// Site defines the interface used by sites in the context of the
// application
type Site interface {
	Name() string
	LoginResponse() *regexp.Regexp
	MatchChat() *regexp.Regexp
}
