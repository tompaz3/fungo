package maybe_test

import (
	"testing"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
)

func TestMaybe(t *testing.T) {
	t.Parallel()
	o.RegisterFailHandler(g.Fail)
	g.RunSpecs(t, "Maybe Suite")
}
