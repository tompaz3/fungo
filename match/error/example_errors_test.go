package match_test

import "fmt"

type InsufficientFundsError struct {
	AccountID string
	Amount    float64
}

type ProductNotFoundError struct {
	ProductID string
}

type UnsupportedCurrencyError struct {
	Currency string
}

type PaymentGatewayRejectedError struct {
	ReasonCode    string
	ReasonMessage string
}

var (
	//nolint:errcheck // this is just a compile-time check for error implementation
	_ error = (*InsufficientFundsError)(nil)
	//nolint:errcheck // this is just a compile-time check for error implementation
	_ error = (*ProductNotFoundError)(nil)
	//nolint:errcheck // this is just a compile-time check for error implementation
	_ error = (*UnsupportedCurrencyError)(nil)
	//nolint:errcheck // this is just a compile-time check for error implementation
	_ error = (*PaymentGatewayRejectedError)(nil)
)

func (e InsufficientFundsError) Error() string {
	return fmt.Sprintf("account %q has insufficient funds to pay %.2f amount", e.AccountID, e.Amount)
}

func (e ProductNotFoundError) Error() string {
	return fmt.Sprintf("product not found %q", e.ProductID)
}

func (e UnsupportedCurrencyError) Error() string {
	return fmt.Sprintf("unsupported currency %q", e.Currency)
}

func (e PaymentGatewayRejectedError) Error() string {
	return fmt.Sprintf("payment gateway rejected with code %q and message %q", e.ReasonCode, e.ReasonMessage)
}
