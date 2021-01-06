package service

import (
	"fmt"
	"time"

	"github.com/jceatwell/bankHexArch/domain"
	"github.com/jceatwell/bankHexArch/dto"
	"github.com/jceatwell/bankHexArch/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccoungRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpenDate:    time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      fmt.Sprintf("%f", req.Amount),
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

func NewAccountService(repo domain.AccoungRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
