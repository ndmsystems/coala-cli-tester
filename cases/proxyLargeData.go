package cases

import "fmt"

type ProxyGettingLargeData struct {
}

func (g ProxyGettingLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.ProxySession.Address == "" {
		return v, msg, fmt.Errorf("missing proxy connection")
	}
	dur, err := downloadDataViaProxy(c.ProxySession, c.DataSize)
	return c, fmt.Sprintf("downloading ⬇ %d Bytes duration: %d ms", c.DataSize, dur), err
}

func (g ProxyGettingLargeData) Title() string {
	return "Proxy Get large data"
}

type ProxySendingLargeData struct {
}

func (g ProxySendingLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.ProxySession.Address == "" {
		return v, msg, fmt.Errorf("missing proxy connection")
	}
	dur, err := sendDataViaProxy(c.ProxySession, c.DataSize)
	return c, fmt.Sprintf("sending ⬆ %d Bytes duration: %d ms", c.DataSize, dur), err
}

func (g ProxySendingLargeData) Title() string {
	return "Proxy Send large data"
}

type ProxyMirrorLargeData struct {
}

func (g ProxyMirrorLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.ProxySession.Address == "" {
		return v, msg, fmt.Errorf("missing proxy connection")
	}
	dur, err := sendMirrorViaProxy(c.ProxySession, c.DataSize)
	return c, fmt.Sprintf("full proccessing ⬆ %d Bytes ⬇ %d Bytes duration: %d ms", c.DataSize, c.DataSize, dur), err
}

func (g ProxyMirrorLargeData) Title() string {
	return "Proxy Mirror testing of large data"
}
