package cases

import "fmt"

type GettingDataAddr struct {
}

func (g GettingDataAddr) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.GumAddr == "" {
		return v, msg, fmt.Errorf("missing gum address")
	}
	c.DataAddr, err = getService(c.CID, c.GumAddr, "data")
	return c, "data address: " + c.DataAddr, err
}

func (g GettingDataAddr) Title() string {
	return "Get Data address"
}

type GettingDataInfo struct {
}

func (g GettingDataInfo) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.DataAddr == "" {
		return v, msg, fmt.Errorf("missing data address")
	}
	info, err := getServiceInfo(c.DataAddr)
	return c, "data info: " + info, err
}

func (g GettingDataInfo) Title() string {
	return "Get Data info"
}
