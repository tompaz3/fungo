package patternmatch_test

import (
	"testing"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
)

func TestPatternMatch(t *testing.T) {
	t.Parallel()
	o.RegisterFailHandler(g.Fail)
	g.RunSpecs(t, "PatternMatch Suite")
}
