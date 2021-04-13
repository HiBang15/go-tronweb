package go_tronweb

import "fmt"

const accountBasePath = "accounts"

type Account struct {
	OwnerAddress   string `json:"owner_address"`
	AccountAddress string `json:"account_address"`
	Visible        bool   `json:"visible"`
	PermissionId   int32  `json:"permission_id"`
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

func (s *AccountServiceOp)Create(account Account) (*Account, error)  {
	path := fmt.Sprintf("%s", accountBasePath)
	wrappedData := AccountResource{Account: &account}
	resource := new(AccountResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Account, err
}

func (s *AccountServiceOp)Get(account GetAccountRequest) (*Account, error)  {
	path := fmt.Sprintf("%s", accountBasePath)
	wrappedData := GetAccountResource{Account: &account}
	resource := new(AccountResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Account, err
}
