package match_test

import (
	"fmt"
	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
	match "github.com/tompaz3/fungo/match/error"
	"net/http"
)

var _ = g.Describe("ErrorMatches", func() {
	receivedErr := PaymentGatewayRejectedError{
		ReasonCode:    "100",
		ReasonMessage: "Payment rejected",
	}

	g.When("error matches the predicate", func() {
		g.It("should return the error and true", func() {
			matchedErr, ok := match.ErrorMatches(receivedErr, func(err error) bool {
				return err.Error() == `payment gateway rejected with code "100" and message "Payment rejected"`
			})
			o.Expect(ok).To(o.BeTrue())
			o.Expect(matchedErr).To(o.Equal(receivedErr))
		})
	})

	g.When("error doesn't match the predicate", func() {
		g.It("should return the error and false", func() {
			matchedErr, ok := match.ErrorMatches(receivedErr, func(err error) bool {
				return err.Error() == `payment gateway rejected with code "200" and message "Payment rejected"`
			})
			o.Expect(ok).To(o.BeFalse())
			o.Expect(matchedErr).To(o.BeZero())
		})
	})
})

type Response struct {
	StatusCode   int
	ErrorCode    ErrorCode
	ErrorMessage string
}

type ErrorCode string

const (
	ErrorCodeProductNotFound     ErrorCode = "product_not_found"
	ErrorCodeUnexpectedError     ErrorCode = "unexpected_error"
	ErrorCodeUnsupportedCurrency ErrorCode = "unsupported_currency"
	ErrorPaymentError            ErrorCode = "payment_error"
)

func HandleErrorByType(recErr error) Response {
	if err, ok := match.ErrorType[ProductNotFoundError](recErr); ok {
		return Response{
			StatusCode:   http.StatusBadRequest,
			ErrorCode:    ErrorCodeProductNotFound,
			ErrorMessage: fmt.Sprintf("product with id %q not found", err.ProductID),
		}
	}
	if err, ok := match.ErrorType[UnsupportedCurrencyError](recErr); ok {
		return Response{
			StatusCode:   http.StatusBadRequest,
			ErrorCode:    ErrorCodeUnsupportedCurrency,
			ErrorMessage: fmt.Sprintf("unuspported currency %q", err.Currency),
		}
	}
	if err, ok := match.ErrorType[InsufficientFundsError](recErr); ok {
		fmt.Println(err.Error())
		return Response{
			StatusCode:   http.StatusInternalServerError,
			ErrorCode:    ErrorPaymentError,
			ErrorMessage: "payment error",
		}
	}
	if err, ok := match.ErrorType[PaymentGatewayRejectedError](recErr); ok {
		return Response{
			StatusCode: http.StatusInternalServerError,
			ErrorCode:  ErrorPaymentError,
			ErrorMessage: fmt.Sprintf(
				"payment gateway rejected with code %q and message %q",
				err.ReasonCode,
				err.ReasonMessage,
			),
		}
	}

	return Response{
		StatusCode:   http.StatusInternalServerError,
		ErrorCode:    ErrorCodeUnexpectedError,
		ErrorMessage: "unexpected error occurred",
	}
}
