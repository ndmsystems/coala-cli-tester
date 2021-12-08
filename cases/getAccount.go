package cases

import "fmt"

type GettingAccountAddr struct {
}

func (g GettingAccountAddr) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.GumAddr == "" {
		return v, msg, fmt.Errorf("missing gum address")
	}
	c.AccountAddr, err = getService(c.CID, c.GumAddr, "account")
	return c, "account address: " + c.AccountAddr, err
}

func (g GettingAccountAddr) Title() string {
	return "Get Account address"
}

type GettingAccountInfo struct {
}

func (g GettingAccountInfo) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.AccountAddr == "" {
		return v, msg, fmt.Errorf("missing account address")
	}
	c.AccountInfo, err = getServiceInfo(c.AccountAddr)
	return c, "account info: " + c.AccountInfo, err
}

func (g GettingAccountInfo) Title() string {
	return "Get Account info"
}
