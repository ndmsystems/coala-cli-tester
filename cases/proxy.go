package cases

import (
	"fmt"
)

type ProxyClientConnection struct {
}

func (g ProxyClientConnection) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.GumAddr == "" {
		return v, msg, fmt.Errorf("missing gum address")
	}
	c.ProxySession, err = gumSession(c.CID, c.ProxyClientCID, c.GumAddr)
	if !c.ProxySession.Online {
		return c, msg, fmt.Errorf("proxy client offline")
	}
	return c, "proxy client address: " + c.ProxySession.Address + " proxy server: " + c.ProxySession.Proxy, err
}

func (g ProxyClientConnection) Title() string {
	return "Proxy connecting with client"
}

type ProxyClientInfo struct {
}

func (g ProxyClientInfo) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.ProxySession.Address == "" {
		return v, msg, fmt.Errorf("missing proxy connection")
	}
	c.ProxyClientInfo, err = getServiceInfoViaProxy(c.ProxySession)
	return c, "proxy client info: " + c.ProxyClientInfo, err
}

func (g ProxyClientInfo) Title() string {
	return "Proxy client info"
}
