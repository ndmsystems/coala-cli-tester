package cases

import (
	"fmt"
)

type P2PClientConnection struct {
}

func (g P2PClientConnection) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.GumAddr == "" {
		return v, msg, fmt.Errorf("missing gum address")
	}
	c.P2PSession, err = gumSession(c.CID, c.P2PClientCID, c.GumAddr)
	if !c.P2PSession.Online {
		return c, msg, fmt.Errorf("p2p client offline")
	}
	return c, "p2p client address: " + c.P2PSession.Address, err
}

func (g P2PClientConnection) Title() string {
	return "P2P connecting with client"
}

type P2PClientInfo struct {
}

func (g P2PClientInfo) Run(c CommonVariables) (v CommonVariables, msg string, err error) {
	if c.P2PSession.Address == "" {
		return v, msg, fmt.Errorf("missing p2p connection")
	}
	info, err := getServiceInfo(c.P2PSession.Address)
	return c, "p2p client info: " + info, err
}

func (g P2PClientInfo) Title() string {
	return "P2P client info"
}
