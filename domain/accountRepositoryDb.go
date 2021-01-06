package domain

import (
	"strconv"

	"github.com/jceatwell/bankHexArch/errs"
	"github.com/jceatwell/bankHexArch/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRespositoryDb struct {
	client *sqlx.DB
}

func (d AccountRespositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := `Insert INTO accounts (customer_id, opening_date, account_type, amount, status)
				  values (?, ?, ?, ?, ?)`
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpenDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from Database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting laste insert id for new account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from Database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRespositoryDb(dbClient *sqlx.DB) AccountRespositoryDb {
	return AccountRespositoryDb{client: dbClient}
}
