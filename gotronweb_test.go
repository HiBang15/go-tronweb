package go_tronweb

import (
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
)

var (
	client  *Client
	tronWeb *TronWeb
)

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, errors.New("test-error")
}

func (errReader) Close() error {
	return nil
}

func setup() {
	tronWeb := TronWeb{
		FullNode:       "",
		SolidityNode:   "",
		EventServer:    "",
		TronProApiKey: "25f66928-0b70-48cd-9ac6-da6f8247c663",
	}

	client = NewClient("api.shasta.trongrid.io", "", tronWeb)

	httpmock.ActivateNonDefault(client.Client)
}

func teardown() {
	httpmock.DeactivateAndReset()
}

func loadFixture(filename string) []byte {
	f, err := ioutil.ReadFile("fixtures/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %v", filename))
	}
	return f
}


