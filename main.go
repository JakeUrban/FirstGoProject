package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"os"
)

func main() {
	kp, kpErr := keypair.Random()
	if kpErr != nil {
		fmt.Println("Unable to generate Stellar keypair")
		os.Exit(1)
	}

	client := horizonclient.DefaultTestNetClient
	_, fundErr := client.Fund(kp.Address())
	if fundErr != nil {
		fmt.Println("Unable to fund Stellar account")
		os.Exit(1)
	}

	account, accountErr := client.AccountDetail(horizonclient.AccountRequest{
		AccountID: kp.Address(),
	})
	if accountErr != nil {
		fmt.Println(accountErr)
		os.Exit(1)
	}

	accountJSON, _ := json.Marshal(account)
	buf := bytes.Buffer{}
	err := json.Indent(&buf, accountJSON, "", "   ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(buf.String())
}