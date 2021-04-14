package go_tronweb

import "fmt"

const getAccountBasePath = "wallet/getaccount"
const createAccountBasePath = "wallet/createaccount"

//type GetAccountResponse struct {
//	AccountName     string      `json:"account_name"`
//	Address         string      `json:"address"`
//	Balance         int         `json:"balance"`
//	AccountResource interface{} `json:"account_resource"`
//}

type GetAccountResResource struct {
	Account *Account `json:"account"`
}

type Account struct {
	AccountName     string      `json:"account_name"`
	Address         string      `json:"address"`
	Balance         int         `json:"balance"`
	AccountResource interface{} `json:"account_resource"`
}

type AccountResource struct {
	Account *Account `json:"account"`
}

type GetAccountResource struct {
	Account *GetAccountRequest `json:"account"`
}

type GetAccountRequest struct {
	Address string `json:"address"`
	Visible bool   `json:"visible"`
}

type AccountService interface {
	Create(Account) (*Account, error)
	Get(GetAccountRequest) (*Account, error)
}

type AccountServiceOp struct {
	client *Client
}

func (s *AccountServiceOp) Create(account Account) (*Account, error) {
	path := fmt.Sprintf("%s", createAccountBasePath)
	wrappedData := AccountResource{Account: &account}
	resource := new(AccountResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Account, err
}

func (s *AccountServiceOp) Get(account GetAccountRequest) (*Account, error) {
	path := fmt.Sprintf("%s", getAccountBasePath)
	//wrappedData := GetAccountResource{Account: &account}
	resource := new(GetAccountResResource)
	err := s.client.Post(path, account, resource)
	return resource.Account, err
}
