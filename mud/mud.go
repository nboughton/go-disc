package mud

import (
	"bufio"
	"fmt"

	"github.com/nboughton/go-disc/history"
	"github.com/nboughton/go-disc/mud/sites"
	"github.com/stesla/gotelnet"
)

// Client wraps the telnet connection and provides some extra functionality
type Client struct {
	r             chan string      // Receiver channel for server text
	Cmds          *history.History // Command history
	Site          sites.Site       // Supported site
	loggedIn      bool             // Logged in or not
	gotelnet.Conn                  // Wrap Conn interface for reading/writing data
}

// NewClient attempts to connect to the host and return a working client connection
func NewClient(site string) (*Client, error) {
	c := new(Client)

	// Check site
	s, ok := sites.Supported[site]
	if !ok {
		fmt.Println("Supported sites:")
		for sName := range sites.Supported {
			fmt.Println(sName)
		}
		return c, fmt.Errorf("Unsupported site [%s]", site)
	}

	// Connect
	var err error
	c.Conn, err = gotelnet.Dial(fmt.Sprintf("%s:%d", s.Host(), s.Port()))
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

		if !c.LoggedIn() && c.Site.LoginSuccess(line) {
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
