package mud

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/nboughton/go-disc/history"
	"github.com/stesla/gotelnet"
)

// Site defines the interface used by sites in the context of the
// application
type Site interface {
	Name() string
	LoginResponse() *regexp.Regexp
	MatchChat() *regexp.Regexp
}

var loginSuccess = regexp.MustCompile(`You (last logged in from|are already playing)`)

// Client wraps the telnet connection and provides some extra functionality
type Client struct {
	r             chan string      // Receiver channel for server text
	Cmds          *history.History // Command history
	loggedIn      bool             // Logged in or not
	gotelnet.Conn                  // Wrap Conn interface for reading/writing data
}

// NewClient attempts to connect to the host and return a working client connection
func NewClient(host string, port int) (*Client, error) {
	c := new(Client)

	// Connect
	var err error
	c.Conn, err = gotelnet.Dial(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return c, err
	}

	// Start listening for return data
	go c.listen()

	// initialise command history
	c.Cmds = history.New()

	return c, err
}

func (c *Client) listen() {
	c.r = make(chan string)
	b := bufio.NewReader(c.Conn)

	for {
		l, _, _ := b.ReadLine()
		line := string(l)

		if !c.LoggedIn() && loginSuccess.MatchString(line) {
			c.Cmds.SetLogging(true)
			c.SetLoggedIn(true)
		}
		c.r <- string(l)
	}
}

// LoggedIn returns whether or not the client thinks a successful login
// has occurred
func (c *Client) LoggedIn() bool {
	return c.loggedIn
}

// SetLoggedIn allows one to set whether or not a successful login has
// occurred
func (c *Client) SetLoggedIn(b bool) {
	c.loggedIn = b
}

// Receive creates a listener channel for server response text
func (c *Client) Receive() chan string {
	return c.r
}

// Send attempts to write a line to the server and returns the error result
func (c *Client) Send(line string) error {
	if _, err := c.Conn.Write([]byte(line + "\n")); err != nil {
		return err
	}

	c.Cmds.Log(line)

	return nil
}
