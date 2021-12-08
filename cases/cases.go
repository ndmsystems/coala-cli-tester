package cases

type CaseRunner interface {
	Run(c CommonVariables) (v CommonVariables, msg string, err error)
	Title() string
}

type CommonVariables struct {
	CID           string
	GumAddr       string
	GumInfo       string
	DataAddr      string
	DataInfo      string
	AccountAddr   string
	AccountInfo   string
	P2PClientInfo string
	DataSize      int
	P2PSession    SessionData
}

type SessionData struct {
	Online  bool   `json:"online"`
	Address string `json:"address"`
	Proxy   string `json:"proxy"`
}
