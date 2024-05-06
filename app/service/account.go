package service

import (
	"crypto/rand"
	"github.com/Jangwooo/AIM_Coding_Test/app/model"
	"github.com/Jangwooo/AIM_Coding_Test/pkg"
	"io"
)

func generateAccountCode() string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, 13)
	_, _ = io.ReadAtLeast(rand.Reader, b, 13)
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}

func OpenAccount(uid string) error {
	a := model.Account{
		ID:     generateAccountCode(),
		UserID: uid,
	}

	err := a.CreateAccount()
	if err != nil {
		return err
	}

	return nil
}

func GetAccountList(uid string) ([]model.Account, error) {
	account := model.Account{
		UserID: uid,
	}
	return account.GetAccountList()
}

func GetAccountInfo(aid string) (*model.Account, error) {
	a := model.Account{
		ID: aid,
	}
	return a.GetAccountByID()
}

func Deposit(uid, aid string, amount uint) error {
	account := &model.Account{
		ID:     aid,
		UserID: uid,
	}

	account.GetAccountByID()

	t := model.Transaction{
		AccountID:    account.ID,
		Amount:       amount,
		Type:         model.TransactionTypeDeposit,
		AfterBalance: account.Balance + amount,
	}

	_ = t.CreateTransaction()

	return account.Deposit(amount)
}

func Withdraw(uid, aid string, amount uint) error {
	account := &model.Account{
		ID:     aid,
		UserID: uid,
	}

	_, _ = account.GetAccountByID()

	if account.Balance < amount {
		return pkg.ErrBalanceNotEnough
	}

	err := account.Withdraw(amount)
	if err != nil {
		return err
	}

	t := model.Transaction{
		AccountID:    aid,
		Amount:       amount,
		Type:         model.TransactionTypeWithdrawal,
		AfterBalance: account.Balance - amount,
	}

	return t.CreateTransaction()
}
