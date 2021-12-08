package cases

import (
	"fmt"
)

type GettingTestClientAddr struct {
}

func (g GettingTestClientAddr) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.GumAddr == "" {
		return v, msg, fmt.Errorf("missing gum address")
	}
	c.TestClientAddr, err = getService(c.CID, c.GumAddr, "test1")
	return c, "test client address: " + c.TestClientAddr, err
}

func (g GettingTestClientAddr) Title() string {
	return "Get test client address"
}

type GettingTestClientInfo struct {
}

func (g GettingTestClientInfo) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.GumAddr == "" {
		return v, msg, fmt.Errorf("missing gum address")
	}
	c.TestClientInfo, err = getServiceInfo(c.TestClientAddr)
	return c, "test client info: " + c.TestClientInfo, err
}

func (g GettingTestClientInfo) Title() string {
	return "Get test client info"
}
