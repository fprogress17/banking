package domain

import (
	"github.com/fprogress17/banking/errs"
	"github.com/fprogress17/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRespositoryDb struct {
	client *sqlx.DB
}

func (d AccountRespositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("error while new account:" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last inserted id for new account:" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRespositoryDb {
	return AccountRespositoryDb{dbClient}
}
