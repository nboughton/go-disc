package sites

import (
	"regexp"
)

// Discworld is a discworld mud site definition
type Discworld struct {
	name       string
	host       string
	port       int
	loginRegex *regexp.Regexp
	chatRegex  *regexp.Regexp
}

// Name satisfies the Name func for the Site interface
func (dw *Discworld) Name() string {
	return dw.name
}

// Host satisfies the Host func for the Site interface
func (dw *Discworld) Host() string {
	return dw.host
}

// Port satisfies the Port func for the Site interface
func (dw *Discworld) Port() int {
	return dw.port
}

// LoginSuccess satisfies the LoginRespone func for the Site interface
func (dw *Discworld) LoginSuccess(line string) bool {
	return dw.loginRegex.MatchString(line)
}

// IsChat satisfies the IsChat func for the Site interface
func (dw *Discworld) IsChat(line string) bool {
	return dw.chatRegex.MatchString(line)
}

func init() {
	dw := &Discworld{
		name:       "Discworld",
		host:       "discworld.atuin.net",
		port:       4242,
		loginRegex: regexp.MustCompile(`You (last logged in from|are already playing)`),
		chatRegex: regexp.MustCompile(`\[1m\[32m`),
	}

	Supported[dw.Name()] = dw
}
