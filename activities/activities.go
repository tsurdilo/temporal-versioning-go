package activities

import (
	"context"
	"go.temporal.io/sdk/activity"
	"temporal/demo/versioning/model"
)

func CheckCustomerAccount(ctx context.Context, customer model.Customer) (bool, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("CheckCustomerAccount", "customer", customer.Name)
	return true, nil
}

func GetCustomerAccount(ctx context.Context, customer model.Customer) (*model.Account, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("GetCustomerAccount", "customer", customer.Name)

	account := &model.Account {
		AccountNum:  customer.AccountNum,
		CustomerNum: customer.AccountNum,
		Message:     "Customer Account",
		Customer:    customer,
		Amount:      0,
	}

	return account, nil
}

func UpdateCustomerAccount(ctx context.Context, account model.Account, amount int) (model.Account, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("UpdateCustomerAccount", "customer", account.Customer.Name)
	account.Amount += amount

	return account, nil
}

func SendBonusEmail(ctx context.Context, customer model.Customer, message string) (bool, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("SendBonusEmail", "customer", customer.Name, "message", message)
	return true, nil
}
