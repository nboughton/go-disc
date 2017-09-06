package sites

// Site defines the interface used by sites
type Site interface {
	Name() string
	Host() string
	Port() int
	LoginSuccess(string) bool
	IsChat(string) bool
}

// Supported contains all currently supported sites as registered so far
var Supported = make(map[string]Site)
