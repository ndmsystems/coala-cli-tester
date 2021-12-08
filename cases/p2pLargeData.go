package cases

import "fmt"

type P2PGettingLargeData struct {
}

func (g P2PGettingLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.P2PSession.Address == "" {
		return v, msg, fmt.Errorf("missing p2p connection")
	}
	dur, err := downloadData(c.P2PSession.Address, c.DataSize)
	return c, fmt.Sprintf("downloading ⬇ %d Bytes duration: %d ms", c.DataSize, dur), err
}

func (g P2PGettingLargeData) Title() string {
	return "P2P Get large data"
}

type P2PSendingLargeData struct {
}

func (g P2PSendingLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.P2PSession.Address == "" {
		return v, msg, fmt.Errorf("missing p2p connection")
	}
	dur, err := sendData(c.P2PSession.Address, c.DataSize)
	return c, fmt.Sprintf("sending ⬆ %d Bytes duration: %d ms", c.DataSize, dur), err
}

func (g P2PSendingLargeData) Title() string {
	return "P2P Send large data"
}

type P2PMirrorLargeData struct {
}

func (g P2PMirrorLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.P2PSession.Address == "" {
		return v, msg, fmt.Errorf("missing p2p connection")
	}
	dur, err := sendMirror(c.P2PSession.Address, c.DataSize)
	return c, fmt.Sprintf("full proccessing ⬆ %d Bytes ⬇ %d Bytes duration: %d ms", c.DataSize, c.DataSize, dur), err
}

func (g P2PMirrorLargeData) Title() string {
	return "P2P Mirror testing of large data"
}
