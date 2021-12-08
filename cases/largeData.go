package cases

import "fmt"

type GettingLargeData struct {
}

func (g GettingLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.TestClientAddr == "" {
		return v, msg, fmt.Errorf("missing test client address")
	}
	dur, err := downloadData(c.TestClientAddr, c.DataSize)
	return c, fmt.Sprintf("downloading ⬇ %d Bytes duration: %d ms", c.DataSize, dur), err
}

func (g GettingLargeData) Title() string {
	return "Get large data"
}

type SendingLargeData struct {
}

func (g SendingLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.TestClientAddr == "" {
		return v, msg, fmt.Errorf("missing test client address")
	}
	dur, err := sendData(c.TestClientAddr, c.DataSize)
	return c, fmt.Sprintf("sending ⬆ %d Bytes duration: %d ms", c.DataSize, dur), err
}

func (g SendingLargeData) Title() string {
	return "Send large data"
}

type MirrorLargeData struct {
}

func (g MirrorLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.TestClientAddr == "" {
		return v, msg, fmt.Errorf("missing test client address")
	}
	dur, err := sendMirror(c.TestClientAddr, c.DataSize)
	return c, fmt.Sprintf("full proccessing ⬆ %d Bytes ⬇ %d Bytes duration: %d ms", c.DataSize, c.DataSize, dur), err
}

func (g MirrorLargeData) Title() string {
	return "Mirror testing of large data"
}
