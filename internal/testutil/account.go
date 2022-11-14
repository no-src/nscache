package testutil

// Account a structure for tests
type Account struct {
	UserName string
	Password string
	IsValid  bool
}

// Equal the current Account is equal to target Account or not
func (a *Account) Equal(account *Account) bool {
	if a == nil || account == nil {
		return a == account
	}
	return a.UserName == account.UserName && a.Password == account.Password && a.IsValid == account.IsValid
}
