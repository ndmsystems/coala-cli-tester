package cases

import "fmt"

type GettingLargeData struct {
}

func (g GettingLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.TestClientAddr == "" {
		return v, msg, fmt.Errorf("missing test client address")
	}
	dur, err := downloadData(c.TestClientAddr)
	return c, fmt.Sprintf("downloading duration: %d ms", dur), err
}

func (g GettingLargeData) Title() string {
	return "Get large data ⬇ 512 KB"
}

type SendingLargeData struct {
}

func (g SendingLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.TestClientAddr == "" {
		return v, msg, fmt.Errorf("missing test client address")
	}
	dur, err := sendData(c.TestClientAddr)
	return c, fmt.Sprintf("sending duration: %d ms", dur), err
}

func (g SendingLargeData) Title() string {
	return "Send large data ⬆ 512 KB "
}

type MirrorLargeData struct {
}

func (g MirrorLargeData) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.TestClientAddr == "" {
		return v, msg, fmt.Errorf("missing test client address")
	}
	dur, err := sendMirror(c.TestClientAddr)
	return c, fmt.Sprintf("full proccessing duration: %d ms", dur), err
}

func (g MirrorLargeData) Title() string {
	return "Mirror testing of large data ⬆ 512 KB ⬇ 512 KB"
}
