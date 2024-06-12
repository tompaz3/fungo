package match_test

import (
	match "github.com/tompaz3/fungo/match/error"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
)

var _ = g.Describe("ErrorType", func() {
	receivedErr := PaymentGatewayRejectedError{
		ReasonCode:    "100",
		ReasonMessage: "Payment rejected",
	}

	g.When("error is of the expected type", func() {
		g.It("should return the typed error and true", func() {
			typedErr, ok := match.ErrorType[PaymentGatewayRejectedError](receivedErr)
			o.Expect(ok).To(o.BeTrue())
			o.Expect(typedErr).To(o.Equal(receivedErr))
		})
	})

	g.When("error is of a different type", func() {
		g.It("should not match error", func() {
			typedErr, ok := match.ErrorType[InsufficientFundsError](receivedErr)
			o.Expect(ok).To(o.BeFalse())
			o.Expect(typedErr).To(o.Equal(InsufficientFundsError{}))
		})
	})
})

var _ = g.Describe("ErrorTypeMatches", func() {
	receivedErr := PaymentGatewayRejectedError{
		ReasonCode:    "100",
		ReasonMessage: "Payment rejected",
	}

	g.When("error is of expected type", func() {
		g.Describe("and matches the predicate", func() {
			g.It("should return the typed error and true", func() {
				typedErr, ok := match.ErrorTypeMatches[PaymentGatewayRejectedError](
					receivedErr,
					func(err PaymentGatewayRejectedError) bool {
						return err.ReasonCode == "100"
					},
				)
				o.Expect(ok).To(o.BeTrue())
				o.Expect(typedErr).To(o.Equal(receivedErr))
			})
		})
		g.Describe("and doesn't match the predicate", func() {
			g.It("should return the typed error and true", func() {
				typedErr, ok := match.ErrorTypeMatches[PaymentGatewayRejectedError](
					receivedErr,
					func(err PaymentGatewayRejectedError) bool {
						return err.ReasonCode == "200"
					},
				)
				o.Expect(ok).To(o.BeFalse())
				o.Expect(typedErr).To(o.Equal(PaymentGatewayRejectedError{}))
			})
		})
	})

	g.When("error is of a different type", func() {
		g.Describe("and matches the predicate", func() {
			g.It("should not match and do not invoke the predicate", func() {
				predicateExecuted := false
				typedErr, ok := match.ErrorTypeMatches[InsufficientFundsError](receivedErr, func(_ InsufficientFundsError) bool {
					predicateExecuted = true

					return true
				})
				o.Expect(predicateExecuted).To(o.BeFalse())
				o.Expect(ok).To(o.BeFalse())
				o.Expect(typedErr).To(o.Equal(InsufficientFundsError{}))
			})
		})
		g.Describe("and doesn't match the predicate", func() {
			g.It("should not match and do not invoke the predicate", func() {
				predicateExecuted := false
				typedErr, ok := match.ErrorTypeMatches[InsufficientFundsError](receivedErr, func(_ InsufficientFundsError) bool {
					predicateExecuted = true
					return false
				})
				o.Expect(predicateExecuted).To(o.BeFalse())
				o.Expect(ok).To(o.BeFalse())
				o.Expect(typedErr).To(o.Equal(InsufficientFundsError{}))
			})
		})
	})
})

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
