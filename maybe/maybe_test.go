package maybe_test

import (
	"errors"
	"strconv"

	"github.com/tompaz3/fungo/maybe"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
)

var errTestError = errors.New("test error")

var _ = g.Describe("Maybe", func() {
	g.Describe("IsDefined", func() {
		g.When("some value present", func() {
			g.It("should return true", func() {
				o.Expect(maybe.Some(1).IsDefined()).To(o.BeTrue())
			})
		})

		g.When("no value present", func() {
			g.It("should return false", func() {
				o.Expect(maybe.None[int]().IsDefined()).To(o.BeFalse())
			})
		})
	})

	g.Describe("IsEmpty", func() {
		g.When("some value present", func() {
			g.It("should return false", func() {
				o.Expect(maybe.Some(1).IsEmpty()).To(o.BeFalse())
			})
		})

		g.When("no value present", func() {
			g.It("should return true", func() {
				o.Expect(maybe.None[int]().IsEmpty()).To(o.BeTrue())
			})
		})
	})

	g.Describe("Get", func() {
		g.When("some value present", func() {
			g.It("should return the value", func() {
				value, err := maybe.Some(1).Get()
				o.Expect(value).To(o.Equal(1))
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})

		g.When("no value present", func() {
			g.It("should return an error", func() {
				value, err := maybe.None[int]().Get()
				o.Expect(value).To(o.Equal(0))
				o.Expect(err).Should(o.HaveOccurred())
			})
		})
	})

	g.Describe("MustGet", func() {
		g.When("some value present", func() {
			g.It("should return the value", func() {
				value := maybe.Some(1).MustGet()
				o.Expect(value).To(o.Equal(1))
			})
		})

		g.When("when no value present", func() {
			g.It("should panic", func() {
				o.Expect(func() {
					maybe.None[int]().MustGet()
				}).Should(o.PanicWith(maybe.ErrEmptyMaybe))
			})
		})
	})

	g.Describe("OrElse", func() {
		g.When("some value present", func() {
			g.It("should return the value", func() {
				value := maybe.Some(1).OrElse(2)
				o.Expect(value).To(o.Equal(1))
			})
		})

		g.When("when no value present", func() {
			g.It("should return the other value", func() {
				value := maybe.None[int]().OrElse(2)
				o.Expect(value).To(o.Equal(2))
			})
		})
	})

	g.Describe("OrElseGet", func() {
		g.When("some value present", func() {
			g.It("should return the value", func() {
				value := maybe.Some(1).OrElseGet(func() int { return 2 })
				o.Expect(value).To(o.Equal(1))
			})
		})

		g.When("no value present", func() {
			g.It("should return the other value", func() {
				value := maybe.None[int]().OrElseGet(func() int { return 2 })
				o.Expect(value).To(o.Equal(2))
			})
		})
	})

	g.Describe("OrElseTryGet", func() {
		g.When("some value present", func() {
			g.Context("and other function returns value", func() {
				g.It("should return the value", func() {
					value, err := maybe.Some(1).
						OrElseTryGet(func() (int, error) { return 2, nil })
					o.Expect(value).To(o.Equal(1))
					o.Expect(err).ShouldNot(o.HaveOccurred())
				})
			})

			g.Context("and other function returns error", func() {
				g.It("should return the value", func() {
					value, err := maybe.Some(1).
						OrElseTryGet(func() (int, error) { return 0, errTestError })
					o.Expect(value).To(o.Equal(1))
					o.Expect(err).ShouldNot(o.HaveOccurred())
				})
			})
		})

		g.When("no value present", func() {
			g.Context("and other function returns value", func() {
				g.It("should return the other value", func() {
					value, err := maybe.None[int]().
						OrElseTryGet(func() (int, error) { return 2, nil })
					o.Expect(value).To(o.Equal(2))
					o.Expect(err).ShouldNot(o.HaveOccurred())
				})
			})

			g.Context("and other function returns error", func() {
				g.It("should return the other error", func() {
					value, err := maybe.None[int]().
						OrElseTryGet(func() (int, error) { return 0, errTestError })
					o.Expect(value).To(o.BeZero())
					o.Expect(err).Should(o.HaveOccurred())
					o.Expect(err).To(o.Equal(errTestError))
				})
			})
		})
	})
})

var _ = g.Describe("Map", func() {
	g.When("some value present", func() {
		g.It("should return the mapped value", func() {
			mb := maybe.Some(1)
			mapped := maybe.Map(mb, strconv.Itoa)
			o.Expect(mapped.IsDefined()).To(o.BeTrue())
			o.Expect(mapped.Get()).To(o.Equal("1"))
		})
	})

	g.When("no value present", func() {
		g.It("should return None", func() {
			mb := maybe.None[int]()
			mapped := maybe.Map(mb, strconv.Itoa)
			o.Expect(mapped.IsEmpty()).To(o.BeTrue())
		})
	})
})

var _ = g.Describe("TryMap", func() {
	g.Describe("some value present", func() {
		g.When("mapping function returns value", func() {
			g.It("should return the mapped value", func() {
				mb := maybe.Some(1)
				mapped, err := maybe.TryMap(mb, func(it int) (string, error) {
					return strconv.Itoa(it), nil
				})
				o.Expect(mapped.IsDefined()).To(o.BeTrue())
				o.Expect(mapped.Get()).To(o.Equal("1"))
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})

		g.When("mapping function returns error", func() {
			g.It("should return the error", func() {
				mb := maybe.Some(1)
				mapped, err := maybe.TryMap(mb, func(_ int) (string, error) {
					return "", errTestError
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(err).To(o.Equal(errTestError))
			})
		})
	})

	g.Describe("no value present", func() {
		g.When("mapping function returns value", func() {
			g.It("should return None", func() {
				mb := maybe.None[int]()
				mapped, err := maybe.TryMap(mb, func(it int) (string, error) {
					return strconv.Itoa(it), nil
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})

		g.When("mapping function returns error", func() {
			g.It("should return None", func() {
				mb := maybe.None[int]()
				mapped, err := maybe.TryMap(mb, func(_ int) (string, error) {
					return "", errTestError
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})
	})
})

var _ = g.Describe("FlatMap", func() {
	g.When("some value present", func() {
		g.It("should return the mapped value", func() {
			mb := maybe.Some(1)
			mapped := maybe.FlatMap(mb, func(it int) maybe.Maybe[string] {
				return maybe.Some(strconv.Itoa(it))
			})
			o.Expect(mapped.IsDefined()).To(o.BeTrue())
			o.Expect(mapped.Get()).To(o.Equal("1"))
		})
	})

	g.When("no value present", func() {
		g.It("should return None", func() {
			mb := maybe.None[int]()
			mapped := maybe.FlatMap(mb, func(it int) maybe.Maybe[string] {
				return maybe.Some(strconv.Itoa(it))
			})
			o.Expect(mapped.IsEmpty()).To(o.BeTrue())
		})
	})
})

var _ = g.Describe("TryFlatMap", func() {
	g.Describe("some value present", func() {
		g.When("mapping function returns value", func() {
			g.It("should return the mapped value", func() {
				mb := maybe.Some(1)
				mapped, err := maybe.TryFlatMap(mb, func(it int) (maybe.Maybe[string], error) {
					return maybe.Some(strconv.Itoa(it)), nil
				})
				o.Expect(mapped.IsDefined()).To(o.BeTrue())
				o.Expect(mapped.Get()).To(o.Equal("1"))
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})

		g.When("mapping function returns error", func() {
			g.It("should return the error", func() {
				mb := maybe.Some(1)
				mapped, err := maybe.TryFlatMap(mb, func(_ int) (maybe.Maybe[string], error) {
					return maybe.None[string](), errTestError
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(err).To(o.Equal(errTestError))
			})
		})
	})

	g.Describe("no value present", func() {
		g.When("mapping function returns value", func() {
			g.It("should return None", func() {
				mb := maybe.None[int]()
				mapped, err := maybe.TryFlatMap(mb, func(it int) (maybe.Maybe[string], error) {
					return maybe.Some(strconv.Itoa(it)), nil
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})

		g.When("mapping function returns error", func() {
			g.It("should return None", func() {
				mb := maybe.None[int]()
				mapped, err := maybe.TryFlatMap(mb, func(_ int) (maybe.Maybe[string], error) {
					return maybe.None[string](), errTestError
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})
	})
})
