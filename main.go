package main

import (
	"coala-cli-tester/cases"
	"flag"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	var (
		err  error
		msg  string
		vars cases.CommonVariables

		casesList = []cases.CaseRunner{
			cases.GettingGUMAddr{},
			cases.GettingGUMInfo{},
			cases.GettingDataAddr{},
			cases.GettingDataInfo{},
			cases.GettingAccountAddr{},
			cases.GettingAccountInfo{},
			cases.P2PClientConnection{},
			cases.P2PClientInfo{},
			cases.P2PGettingLargeData{},
			cases.P2PSendingLargeData{},
			cases.P2PMirrorLargeData{},
		}
	)

	flag.StringVar(&vars.CID, "cid", "", "CID for this testing client. If empty then will be setted randomly CID.")
	flag.IntVar(&vars.DataSize, "size", 512*1024, "Size of payload in bytes for large data tests.")

	flag.Parse()

	if vars.CID == "" {
		vars.CID = uuid.New().String()
	}

	fmt.Printf("\nKeenetic test page starting with CID: %s\n\n", vars.CID)

	for i, caseItem := range casesList {
		space := "  "
		if i > 9 {
			space = " "
		}

		fmt.Printf("%d.%s%s:\n", i, space, caseItem.Title())
		vars, msg, err = caseItem.Run(vars)

		printResult(msg, err)
		fmt.Println("")
	}
}

func printResult(msg string, err error) {
	if err == nil {
		fmt.Println("    🟢 " + msg)
	} else {
		fmt.Println("    🔴 error: " + err.Error())
	}
}
