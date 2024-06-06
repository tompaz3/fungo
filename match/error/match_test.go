package errmatch_test

import (
	"errors"
	"fmt"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
	errmatch "github.com/tompaz3/fungo/match/error"
)

var errTestError = errors.New("test error")

var _ = g.Describe("Switch", func() {
	g.Describe("Case", func() {
		receivedErr := PaymentGatewayRejectedError{
			ReasonCode:    "100",
			ReasonMessage: "Payment rejected",
		}

		g.When("error matches first error type", func() {
			g.Context("and matched function returns value", func() {
				g.It("should retrieve value", func() {
					executedMethods := make([]string, 0)
					res, err := errmatch.Switch(
						receivedErr,
						errmatch.Case(func(_ PaymentGatewayRejectedError) (int, error) {
							executedMethods = append(executedMethods, "PaymentGatewayRejectedError")

							return 1, nil
						}),
						errmatch.Case(func(_ InsufficientFundsError) (int, error) {
							executedMethods = append(executedMethods, "InsufficientFundsError")

							return 2, nil
						}),
						errmatch.Case(func(_ ProductNotFoundError) (int, error) {
							executedMethods = append(executedMethods, "ProductNotFoundError")

							return 3, nil
						}),
						errmatch.Case(func(_ UnsupportedCurrencyError) (int, error) {
							executedMethods = append(executedMethods, "UnsupportedCurrencyError")

							return 4, nil
						}),
					)

					o.Expect(res).To(o.Equal(1))
					o.Expect(err).ShouldNot(o.HaveOccurred())
					o.Expect(executedMethods).To(o.HaveExactElements("PaymentGatewayRejectedError"))
				})
			})

			g.Context("and matched function returns error", func() {
				g.It("should retrieve error", func() {
					executedMethods := make([]string, 0)
					res, err := errmatch.Switch(
						receivedErr,
						errmatch.Case(func(_ PaymentGatewayRejectedError) (int, error) {
							executedMethods = append(executedMethods, "PaymentGatewayRejectedError")

							return 0, errTestError
						}),
						errmatch.Case(func(_ InsufficientFundsError) (int, error) {
							executedMethods = append(executedMethods, "InsufficientFundsError")

							return 1, nil
						}),
						errmatch.Case(func(_ ProductNotFoundError) (int, error) {
							executedMethods = append(executedMethods, "ProductNotFoundError")

							return 2, nil
						}),
						errmatch.Case(func(_ UnsupportedCurrencyError) (int, error) {
							executedMethods = append(executedMethods, "UnsupportedCurrencyError")

							return 3, nil
						}),
					)

					o.Expect(res).To(o.BeZero())
					o.Expect(err).Should(o.HaveOccurred())
					o.Expect(err).To(o.Equal(errTestError))
					o.Expect(executedMethods).To(o.HaveExactElements("PaymentGatewayRejectedError"))
				})
			})
		})

		g.When("error matches middle error type", func() {
			g.Context("and matched function returns value", func() {
				g.It("should retrieve value", func() {
					executedMethods := make([]string, 0)
					res, err := errmatch.Switch(
						receivedErr,
						errmatch.Case(func(_ InsufficientFundsError) (int, error) {
							executedMethods = append(executedMethods, "InsufficientFundsError")

							return 1, nil
						}),
						errmatch.Case(func(_ PaymentGatewayRejectedError) (int, error) {
							executedMethods = append(executedMethods, "PaymentGatewayRejectedError")

							return 2, nil
						}),
						errmatch.Case(func(_ ProductNotFoundError) (int, error) {
							executedMethods = append(executedMethods, "ProductNotFoundError")

							return 3, nil
						}),
						errmatch.Case(func(_ UnsupportedCurrencyError) (int, error) {
							executedMethods = append(executedMethods, "UnsupportedCurrencyError")

							return 4, nil
						}),
					)

					o.Expect(res).To(o.Equal(2))
					o.Expect(err).ShouldNot(o.HaveOccurred())
					o.Expect(executedMethods).To(o.HaveExactElements("PaymentGatewayRejectedError"))
				})
			})

			g.Context("and matched function returns error", func() {
				g.It("should retrieve error", func() {
					executedMethods := make([]string, 0)
					res, err := errmatch.Switch(
						receivedErr,
						errmatch.Case(func(_ InsufficientFundsError) (int, error) {
							executedMethods = append(executedMethods, "InsufficientFundsError")

							return 1, nil
						}),
						errmatch.Case(func(_ PaymentGatewayRejectedError) (int, error) {
							executedMethods = append(executedMethods, "PaymentGatewayRejectedError")

							return 0, errTestError
						}),
						errmatch.Case(func(_ ProductNotFoundError) (int, error) {
							executedMethods = append(executedMethods, "ProductNotFoundError")

							return 3, nil
						}),
						errmatch.Case(func(_ UnsupportedCurrencyError) (int, error) {
							executedMethods = append(executedMethods, "UnsupportedCurrencyError")

							return 4, nil
						}),
					)

					o.Expect(res).To(o.BeZero())
					o.Expect(err).Should(o.HaveOccurred())
					o.Expect(err).To(o.Equal(errTestError))
					o.Expect(executedMethods).To(o.HaveExactElements("PaymentGatewayRejectedError"))
				})
			})
		})

		g.When("error matches last error type", func() {
			g.Context("and matched function returns value", func() {
				g.It("should retrieve value", func() {
					executedMethods := make([]string, 0)
					res, err := errmatch.Switch(
						receivedErr,
						errmatch.Case(func(_ InsufficientFundsError) (int, error) {
							executedMethods = append(executedMethods, "InsufficientFundsError")

							return 1, nil
						}),
						errmatch.Case(func(_ ProductNotFoundError) (int, error) {
							executedMethods = append(executedMethods, "ProductNotFoundError")

							return 2, nil
						}),
						errmatch.Case(func(_ UnsupportedCurrencyError) (int, error) {
							executedMethods = append(executedMethods, "UnsupportedCurrencyError")

							return 3, nil
						}),
						errmatch.Case(func(_ PaymentGatewayRejectedError) (int, error) {
							executedMethods = append(executedMethods, "PaymentGatewayRejectedError")

							return 4, nil
						}),
					)

					o.Expect(res).To(o.Equal(4))
					o.Expect(err).ShouldNot(o.HaveOccurred())
					o.Expect(executedMethods).To(o.HaveExactElements("PaymentGatewayRejectedError"))
				})
			})

			g.Context("and matched function returns error", func() {
				g.It("should retrieve error", func() {
					executedMethods := make([]string, 0)
					res, err := errmatch.Switch(
						receivedErr,
						errmatch.Case(func(_ InsufficientFundsError) (int, error) {
							executedMethods = append(executedMethods, "InsufficientFundsError")

							return 1, nil
						}),
						errmatch.Case(func(_ ProductNotFoundError) (int, error) {
							executedMethods = append(executedMethods, "ProductNotFoundError")

							return 2, nil
						}),
						errmatch.Case(func(_ UnsupportedCurrencyError) (int, error) {
							executedMethods = append(executedMethods, "UnsupportedCurrencyError")

							return 3, nil
						}),
						errmatch.Case(func(_ PaymentGatewayRejectedError) (int, error) {
							executedMethods = append(executedMethods, "PaymentGatewayRejectedError")

							return 0, errTestError
						}),
					)

					o.Expect(res).To(o.BeZero())
					o.Expect(err).Should(o.HaveOccurred())
					o.Expect(err).To(o.Equal(errTestError))
					o.Expect(executedMethods).To(o.HaveExactElements("PaymentGatewayRejectedError"))
				})
			})
		})

		g.When("no match found", func() {
			g.Context("and no default case", func() {
				g.It("should return error", func() {
					executedMethods := make([]string, 0)
					res, err := errmatch.Switch(
						receivedErr,
						errmatch.Case(func(_ InsufficientFundsError) (int, error) {
							executedMethods = append(executedMethods, "InsufficientFundsError")

							return 1, nil
						}),
						errmatch.Case(func(_ ProductNotFoundError) (int, error) {
							executedMethods = append(executedMethods, "ProductNotFoundError")

							return 2, nil
						}),
						errmatch.Case(func(_ UnsupportedCurrencyError) (int, error) {
							executedMethods = append(executedMethods, "UnsupportedCurrencyError")

							return 3, nil
						}),
					)

					o.Expect(res).To(o.BeZero())
					o.Expect(err).Should(o.HaveOccurred())
					o.Expect(err).To(o.Equal(errmatch.ErrNoMatchFound))
					o.Expect(executedMethods).To(o.BeEmpty())
				})
			})

			g.Context("and default case", func() {
				g.It("should return result", func() {
					executedMethods := make([]string, 0)
					res, err := errmatch.Switch(
						receivedErr,
						errmatch.Case(func(_ InsufficientFundsError) (int, error) {
							executedMethods = append(executedMethods, "InsufficientFundsError")

							return 1, nil
						}),
						errmatch.Case(func(_ ProductNotFoundError) (int, error) {
							executedMethods = append(executedMethods, "ProductNotFoundError")

							return 2, nil
						}),
						errmatch.Case(func(_ UnsupportedCurrencyError) (int, error) {
							executedMethods = append(executedMethods, "UnsupportedCurrencyError")

							return 3, nil
						}),
						errmatch.Default(func(_ error) (int, error) {
							executedMethods = append(executedMethods, "Default")

							return 4, nil
						}),
					)

					o.Expect(res).To(o.Equal(4))
					o.Expect(err).ShouldNot(o.HaveOccurred())
					o.Expect(executedMethods).To(o.HaveExactElements("Default"))
				})
			})
		})
	})
})

type InsufficientFundsError struct {
	Amount float64
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
	return fmt.Sprintf("insufficient amount for payment %.2f", e.Amount)
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
