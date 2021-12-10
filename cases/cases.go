package cases

type CaseRunner interface {
	Run(c CommonVariables) (v CommonVariables, msg string, err error)
	Title() string
}

type CommonVariables struct {
	CID            string
	ProxyClientCID string
	P2PClientCID   string

	DataSize int

	MasterAddr  string
	GumAddr     string
	DataAddr    string
	AccountAddr string

	P2PSession   SessionData
	ProxySession SessionData
}

type SessionData struct {
	Online  bool   `json:"online"`
	Address string `json:"address"`
	Proxy   string `json:"proxy"`
}
