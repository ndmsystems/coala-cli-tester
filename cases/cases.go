package cases

type CaseRunner interface {
	Run(c CommonVariables) (v CommonVariables, msg string, err error)
	Title() string
}

type CommonVariables struct {
	CID            string
	GumAddr        string
	GumInfo        string
	DataAddr       string
	DataInfo       string
	AccountAddr    string
	AccountInfo    string
	TestClientAddr string
	TestClientInfo string
}
