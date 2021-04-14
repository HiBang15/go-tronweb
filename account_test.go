package go_tronweb

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"log"
	"testing"
)

func Test_CreateAccount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprint("https://api.shasta.trongrid.io"),
		httpmock.NewBytesResponder(
			200,
			loadFixture("get_account.json"),
		))
	getReq := GetAccountRequest{
		Address: "41E552F6487585C2B58BC2C9BB4492BC1F17132CD0",
		Visible: true,
	}
	account, err := client.Account.Get(getReq)
	if err != nil {
		log.Println(err)
		//t.Errorf("DraftOrder.Get returned error: %v", err)
	}
	log.Println(account)
}
