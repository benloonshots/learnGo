package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount is Constructor creating Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
