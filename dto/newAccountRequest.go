package dto

import (
	"github.com/fprogress17/banking/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"cusomter_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("account type should be checking or saving")
	}

	return nil
}
