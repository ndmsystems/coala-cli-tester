package cases

import (
	"fmt"
	"io"
	"net/http"
)

type GettingGUMAddr struct {
}

func (g GettingGUMAddr) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	c.GumAddr, err = getGumAddrFromMaster(c.MasterAddr, c.CID)
	return c, "gum address: " + c.GumAddr, err
}

func (g GettingGUMAddr) Title() string {
	return "Get GUM address"
}

type GettingGUMInfo struct {
}

func (g GettingGUMInfo) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.GumAddr == "" {
		return v, msg, fmt.Errorf("missing gum address")
	}
	info, err := getServiceInfo(c.GumAddr)
	return c, "gum info: " + info, err
}

func (g GettingGUMInfo) Title() string {
	return "Get GUM info"
}

func getGumAddrFromMaster(addr, cid string) (gumAddr string, err error) {
	resp, err := http.Get(addr + "/?cid=" + cid)
	if err != nil {
		return "", fmt.Errorf("get request by HTTP: %s", err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read an HTTP response: %s", err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("invalid response code %s with payload: %s", resp.Status, string(b))
	}

	return string(b), nil
}
