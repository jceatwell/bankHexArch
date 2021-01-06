package domain

import (
	"github.com/jceatwell/bankHexArch/dto"
	"github.com/jceatwell/bankHexArch/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpenDate    string
	AccountType string
	Amount      string
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
