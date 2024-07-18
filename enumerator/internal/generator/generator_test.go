// Copyright (c) 2024-2024 Tomasz Paździurek
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package generator_test

import (
	"os"

	"github.com/tompaz3/fungo/enumerator/internal/generator"

	_ "embed" // embed package is imported for go:embed directive

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
	colorplain "github.com/tompaz3/fungo/enumerator/internal/generator/colorplain"
	colorjson "github.com/tompaz3/fungo/enumerator/internal/generator/colorwithjsonmarshalling"
	colorundef "github.com/tompaz3/fungo/enumerator/internal/generator/colorwithundefined"
	colorundefjson "github.com/tompaz3/fungo/enumerator/internal/generator/colorwithundefinedandjsonmarshalling"
	colorundefjsonnil "github.com/tompaz3/fungo/enumerator/internal/generator/colorwithundefinedandjsonmarshallingniltoundefined"
)

//go:embed colorplain/expected_color.txt
var expectedColor []byte

var _ = g.Describe("generate plain", func() {
	destination := "./colorplain/color.go"

	g.It("should generate", func() {
		enum := generator.Enum{
			Destination:   &destination,
			CopyrightFile: "../../../LICENSE",
			Package:       "color",
			Type:          "Color",
			Values:        []string{"Undefined", "Red", "Green", "Blue"},
		}
		err := generator.Generate(enum)

		o.Expect(err).ShouldNot(o.HaveOccurred())
	})

	g.It("should be equal to expected", func() {
		colorContent, err := os.ReadFile(destination)

		o.Expect(err).ShouldNot(o.HaveOccurred())
		o.Expect(colorContent).Should(o.Equal(expectedColor))
	})

	g.Describe("Of", func() {
		g.When("Undefined", func() {
			g.It("then Undefined", func() {
				value, err := colorplain.Of("Undefined")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorplain.Undefined))
			})
		})
		g.When("Red", func() {
			g.It("then Red", func() {
				value, err := colorplain.Of("Red")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorplain.Red))
			})
		})
		g.When("Green", func() {
			g.It("then Green", func() {
				value, err := colorplain.Of("Green")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorplain.Green))
			})
		})
		g.When("Blue", func() {
			g.It("then Blue", func() {
				value, err := colorplain.Of("Blue")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorplain.Blue))
			})
		})
		g.When("empty string", func() {
			g.It("then error", func() {
				value, err := colorplain.Of("")
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(value).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("invalid Color name: "))
			})
		})
		g.When("invalid value", func() {
			g.It("then error", func() {
				value, err := colorplain.Of("invalid-color")
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(value).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("invalid Color name: invalid-color"))
			})
		})
	})
})

//go:embed colorwithundefined/expected_color.txt
var expectedColorWithUndefined []byte

var _ = g.Describe("generate with undefined", func() {
	destination := "./colorwithundefined/color.go"

	g.It("should generate", func() {
		enum := generator.Enum{
			Destination:    &destination,
			CopyrightFile:  "../../../LICENSE",
			Package:        "color",
			Type:           "Color",
			Values:         []string{"Undefined", "Red", "Green", "Blue"},
			UndefinedValue: "Undefined",
		}
		err := generator.Generate(enum)

		o.Expect(err).ShouldNot(o.HaveOccurred())
	})

	g.It("should be equal to expected", func() {
		colorContent, err := os.ReadFile(destination)

		o.Expect(err).ShouldNot(o.HaveOccurred())
		o.Expect(colorContent).Should(o.Equal(expectedColorWithUndefined))
	})

	g.Describe("Of", func() {
		g.When("Undefined", func() {
			g.It("then Undefined", func() {
				value, err := colorundef.Of("Undefined")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorundef.Undefined))
			})
		})
		g.When("Red", func() {
			g.It("then Red", func() {
				value, err := colorundef.Of("Red")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorundef.Red))
			})
		})
		g.When("Green", func() {
			g.It("then Green", func() {
				value, err := colorundef.Of("Green")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorundef.Green))
			})
		})
		g.When("Blue", func() {
			g.It("then Blue", func() {
				value, err := colorundef.Of("Blue")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorundef.Blue))
			})
		})
		g.When("empty string", func() {
			g.It("then error", func() {
				value, err := colorundef.Of("")
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(value).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("invalid Color name: "))
			})
		})
		g.When("invalid value", func() {
			g.It("then error", func() {
				value, err := colorundef.Of("invalid-color")
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(value).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("invalid Color name: invalid-color"))
			})
		})
	})

	g.Describe("OfOrUndefined", func() {
		g.When("Undefined", func() {
			g.It("then Undefined", func() {
				value := colorundef.OfOrUndefined("Undefined")
				o.Expect(value).Should(o.Equal(colorundef.Undefined))
			})
		})
		g.When("Red", func() {
			g.It("then Red", func() {
				value := colorundef.OfOrUndefined("Red")
				o.Expect(value).Should(o.Equal(colorundef.Red))
			})
		})
		g.When("Green", func() {
			g.It("then Green", func() {
				value := colorundef.OfOrUndefined("Green")
				o.Expect(value).Should(o.Equal(colorundef.Green))
			})
		})
		g.When("Blue", func() {
			g.It("then Blue", func() {
				value := colorundef.OfOrUndefined("Blue")
				o.Expect(value).Should(o.Equal(colorundef.Blue))
			})
		})
		g.When("empty string", func() {
			g.It("then Undefined", func() {
				value := colorundef.OfOrUndefined("")
				o.Expect(value).Should(o.Equal(colorundef.Undefined))
			})
		})
		g.When("invalid value", func() {
			g.It("then Undefined", func() {
				value := colorundef.OfOrUndefined("invalid-color")
				o.Expect(value).Should(o.Equal(colorundef.Undefined))
			})
		})
	})
})

//go:embed colorwithjsonmarshalling/expected_color.txt
var expectedColorWithJSONMarshalling []byte

var _ = g.Describe("generate with JSON marshalling", func() {
	destination := "./colorwithjsonmarshalling/color.go"

	g.It("should generate", func() {
		enum := generator.Enum{
			Destination:   &destination,
			CopyrightFile: "../../../LICENSE",
			Package:       "color",
			Type:          "Color",
			Values:        []string{"Red", "Green", "Blue"},
			Marshalling: generator.MarshalOptions{
				JSONOptions: generator.JSONMarshalOptions{
					Generate: true,
				},
			},
		}
		err := generator.Generate(enum)

		o.Expect(err).ShouldNot(o.HaveOccurred())
	})

	g.It("should be equal to expected", func() {
		colorContent, err := os.ReadFile(destination)

		o.Expect(err).ShouldNot(o.HaveOccurred())
		o.Expect(colorContent).Should(o.Equal(expectedColorWithJSONMarshalling))
	})

	g.Describe("Of", func() {
		g.When("Red", func() {
			g.It("then Red", func() {
				value, err := colorjson.Of("Red")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorjson.Red))
			})
		})
		g.When("Green", func() {
			g.It("then Green", func() {
				value, err := colorjson.Of("Green")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorjson.Green))
			})
		})
		g.When("Blue", func() {
			g.It("then Blue", func() {
				value, err := colorjson.Of("Blue")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorjson.Blue))
			})
		})
		g.When("empty string", func() {
			g.It("then error", func() {
				value, err := colorjson.Of("")
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(value).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("invalid Color name: "))
			})
		})
		g.When("invalid value", func() {
			g.It("then error", func() {
				value, err := colorjson.Of("invalid-color")
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(value).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("invalid Color name: invalid-color"))
			})
		})
	})

	g.Describe("MarshalJSON", func() {
		g.When("Red", func() {
			g.It("then marshalled", func() {
				json, err := colorjson.Red.ToMarshallable().MarshalJSON()
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(json).Should(o.Equal([]byte(`"Red"`)))
			})
		})
		g.When("Green", func() {
			g.It("then marshalled", func() {
				json, err := colorjson.Green.ToMarshallable().MarshalJSON()
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(json).Should(o.Equal([]byte(`"Green"`)))
			})
		})
		g.When("Blue", func() {
			g.It("then marshalled", func() {
				json, err := colorjson.Blue.ToMarshallable().MarshalJSON()
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(json).Should(o.Equal([]byte(`"Blue"`)))
			})
		})
	})

	g.Describe("UnmarshalJSON", func() {
		g.When("Red", func() {
			g.It("then unmarshalled", func() {
				mColor := &colorjson.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"Red"`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorjson.Red))
			})
		})
		g.When("Green", func() {
			g.It("then unmarshalled", func() {
				mColor := &colorjson.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"Green"`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorjson.Green))
			})
		})
		g.When("Blue", func() {
			g.It("then unmarshalled", func() {
				mColor := &colorjson.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"Blue"`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorjson.Blue))
			})
		})
		g.When("empty json", func() {
			g.It("then nil", func() {
				mColor := &colorjson.MarshallableColor{}
				err := mColor.UnmarshalJSON(make([]byte, 0))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.BeNil())
			})
		})
		g.When("null json", func() {
			g.It("then nil", func() {
				mColor := &colorjson.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`null`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.BeNil())
			})
		})
		g.When("invalid-color json", func() {
			g.It("then error", func() {
				mColor := &colorjson.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"invalid-color"`))
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("could not unmarshal Color from JSON\ninvalid Color name: invalid-color"))
			})
		})
	})
})

//go:embed colorwithundefinedandjsonmarshalling/expected_color.txt
var expectedColorWithUndefinedAndJSONMarshalling []byte

var _ = g.Describe("generate with undefined and JSON marshalling", func() {
	destination := "./colorwithundefinedandjsonmarshalling/color.go"

	g.It("should generate", func() {
		enum := generator.Enum{
			Destination:    &destination,
			CopyrightFile:  "../../../LICENSE",
			Package:        "color",
			Type:           "Color",
			Values:         []string{"Undefined", "Red", "Green", "Blue"},
			UndefinedValue: "Undefined",
			Marshalling: generator.MarshalOptions{
				JSONOptions: generator.JSONMarshalOptions{
					Generate: true,
				},
			},
		}
		err := generator.Generate(enum)

		o.Expect(err).ShouldNot(o.HaveOccurred())
	})

	g.It("should be equal to expected", func() {
		colorContent, err := os.ReadFile(destination)

		o.Expect(err).ShouldNot(o.HaveOccurred())
		o.Expect(colorContent).Should(o.Equal(expectedColorWithUndefinedAndJSONMarshalling))
	})

	g.Describe("Of", func() {
		g.When("Red", func() {
			g.It("then Red", func() {
				value, err := colorundefjson.Of("Red")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorundefjson.Red))
			})
		})
		g.When("Green", func() {
			g.It("then Green", func() {
				value, err := colorundefjson.Of("Green")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorundefjson.Green))
			})
		})
		g.When("Blue", func() {
			g.It("then Blue", func() {
				value, err := colorundefjson.Of("Blue")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorundefjson.Blue))
			})
		})
		g.When("empty string", func() {
			g.It("then error", func() {
				value, err := colorundefjson.Of("")
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(value).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("invalid Color name: "))
			})
		})
		g.When("invalid value", func() {
			g.It("then error", func() {
				value, err := colorundefjson.Of("invalid-color")
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(value).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("invalid Color name: invalid-color"))
			})
		})
	})

	g.Describe("OfOrUndefined", func() {
		g.When("Undefined", func() {
			g.It("then Undefined", func() {
				value := colorundefjson.OfOrUndefined("Undefined")
				o.Expect(value).Should(o.Equal(colorundefjson.Undefined))
			})
		})
		g.When("Red", func() {
			g.It("then Red", func() {
				value := colorundefjson.OfOrUndefined("Red")
				o.Expect(value).Should(o.Equal(colorundefjson.Red))
			})
		})
		g.When("Green", func() {
			g.It("then Green", func() {
				value := colorundefjson.OfOrUndefined("Green")
				o.Expect(value).Should(o.Equal(colorundefjson.Green))
			})
		})
		g.When("Blue", func() {
			g.It("then Blue", func() {
				value := colorundefjson.OfOrUndefined("Blue")
				o.Expect(value).Should(o.Equal(colorundefjson.Blue))
			})
		})
		g.When("empty string", func() {
			g.It("then Undefined", func() {
				value := colorundefjson.OfOrUndefined("")
				o.Expect(value).Should(o.Equal(colorundefjson.Undefined))
			})
		})
		g.When("invalid value", func() {
			g.It("then Undefined", func() {
				value := colorundefjson.OfOrUndefined("invalid-color")
				o.Expect(value).Should(o.Equal(colorundefjson.Undefined))
			})
		})
	})

	g.Describe("MarshalJSON", func() {
		g.When("Red", func() {
			g.It("then marshalled", func() {
				json, err := colorundefjson.Red.ToMarshallable().MarshalJSON()
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(json).Should(o.Equal([]byte(`"Red"`)))
			})
		})
		g.When("Green", func() {
			g.It("then marshalled", func() {
				json, err := colorundefjson.Green.ToMarshallable().MarshalJSON()
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(json).Should(o.Equal([]byte(`"Green"`)))
			})
		})
		g.When("Blue", func() {
			g.It("then marshalled", func() {
				json, err := colorundefjson.Blue.ToMarshallable().MarshalJSON()
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(json).Should(o.Equal([]byte(`"Blue"`)))
			})
		})
	})

	g.Describe("UnmarshalJSON", func() {
		g.When("Red", func() {
			g.It("then unmarshalled", func() {
				mColor := &colorundefjson.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"Red"`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorundefjson.Red))
			})
		})
		g.When("Green", func() {
			g.It("then unmarshalled", func() {
				mColor := &colorundefjson.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"Green"`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorundefjson.Green))
			})
		})
		g.When("Blue", func() {
			g.It("then unmarshalled", func() {
				mColor := &colorundefjson.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"Blue"`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorundefjson.Blue))
			})
		})
		g.When("empty json", func() {
			g.It("then nil", func() {
				mColor := &colorundefjson.MarshallableColor{}
				err := mColor.UnmarshalJSON(make([]byte, 0))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.BeNil())
			})
		})
		g.When("null json", func() {
			g.It("then nil", func() {
				mColor := &colorundefjson.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`null`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.BeNil())
			})
		})
		g.When("invalid-color json", func() {
			g.It("then error", func() {
				mColor := &colorundefjson.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"invalid-color"`))
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("could not unmarshal Color from JSON\ninvalid Color name: invalid-color"))
			})
		})
	})
})

//go:embed colorwithundefinedandjsonmarshallingniltoundefined/expected_color.txt
var expectedColorWithUndefinedAndJSONMarshallingNilToUndefined []byte

var _ = g.Describe("generate with undefined and JSON marshalling and nil to undefined", func() {
	destination := "./colorwithundefinedandjsonmarshallingniltoundefined/color.go"

	g.It("should generate", func() {
		enum := generator.Enum{
			Destination:    &destination,
			CopyrightFile:  "../../../LICENSE",
			Package:        "color",
			Type:           "Color",
			Values:         []string{"Undefined", "Red", "Green", "Blue"},
			UndefinedValue: "Undefined",
			Marshalling: generator.MarshalOptions{
				JSONOptions: generator.JSONMarshalOptions{
					Generate:       true,
					NilToUndefined: true,
				},
			},
		}
		err := generator.Generate(enum)

		o.Expect(err).ShouldNot(o.HaveOccurred())
	})

	g.It("should be equal to expected", func() {
		colorContent, err := os.ReadFile(destination)

		o.Expect(err).ShouldNot(o.HaveOccurred())
		o.Expect(colorContent).Should(o.Equal(expectedColorWithUndefinedAndJSONMarshallingNilToUndefined))
	})

	g.Describe("Of", func() {
		g.When("Red", func() {
			g.It("then Red", func() {
				value, err := colorundefjsonnil.Of("Red")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorundefjsonnil.Red))
			})
		})
		g.When("Green", func() {
			g.It("then Green", func() {
				value, err := colorundefjsonnil.Of("Green")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorundefjsonnil.Green))
			})
		})
		g.When("Blue", func() {
			g.It("then Blue", func() {
				value, err := colorundefjsonnil.Of("Blue")
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(value).Should(o.Equal(colorundefjsonnil.Blue))
			})
		})
		g.When("empty string", func() {
			g.It("then error", func() {
				value, err := colorundefjsonnil.Of("")
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(value).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("invalid Color name: "))
			})
		})
		g.When("invalid value", func() {
			g.It("then error", func() {
				value, err := colorundefjsonnil.Of("invalid-color")
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(value).Should(o.BeNil())
				o.Expect(err.Error()).Should(o.Equal("invalid Color name: invalid-color"))
			})
		})
	})

	g.Describe("OfOrUndefined", func() {
		g.When("Undefined", func() {
			g.It("then Undefined", func() {
				value := colorundefjsonnil.OfOrUndefined("Undefined")
				o.Expect(value).Should(o.Equal(colorundefjsonnil.Undefined))
			})
		})
		g.When("Red", func() {
			g.It("then Red", func() {
				value := colorundefjsonnil.OfOrUndefined("Red")
				o.Expect(value).Should(o.Equal(colorundefjsonnil.Red))
			})
		})
		g.When("Green", func() {
			g.It("then Green", func() {
				value := colorundefjsonnil.OfOrUndefined("Green")
				o.Expect(value).Should(o.Equal(colorundefjsonnil.Green))
			})
		})
		g.When("Blue", func() {
			g.It("then Blue", func() {
				value := colorundefjsonnil.OfOrUndefined("Blue")
				o.Expect(value).Should(o.Equal(colorundefjsonnil.Blue))
			})
		})
		g.When("empty string", func() {
			g.It("then Undefined", func() {
				value := colorundefjsonnil.OfOrUndefined("")
				o.Expect(value).Should(o.Equal(colorundefjsonnil.Undefined))
			})
		})
		g.When("invalid value", func() {
			g.It("then Undefined", func() {
				value := colorundefjsonnil.OfOrUndefined("invalid-color")
				o.Expect(value).Should(o.Equal(colorundefjsonnil.Undefined))
			})
		})
	})

	g.Describe("MarshalJSON", func() {
		g.When("Red", func() {
			g.It("then marshalled", func() {
				json, err := colorundefjsonnil.Red.ToMarshallable().MarshalJSON()
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(json).Should(o.Equal([]byte(`"Red"`)))
			})
		})
		g.When("Green", func() {
			g.It("then marshalled", func() {
				json, err := colorundefjsonnil.Green.ToMarshallable().MarshalJSON()
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(json).Should(o.Equal([]byte(`"Green"`)))
			})
		})
		g.When("Blue", func() {
			g.It("then marshalled", func() {
				json, err := colorundefjsonnil.Blue.ToMarshallable().MarshalJSON()
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(json).Should(o.Equal([]byte(`"Blue"`)))
			})
		})
	})

	g.Describe("UnmarshalJSON", func() {
		g.When("Red", func() {
			g.It("then unmarshalled", func() {
				mColor := &colorundefjsonnil.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"Red"`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorundefjsonnil.Red))
			})
		})
		g.When("Green", func() {
			g.It("then unmarshalled", func() {
				mColor := &colorundefjsonnil.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"Green"`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorundefjsonnil.Green))
			})
		})
		g.When("Blue", func() {
			g.It("then unmarshalled", func() {
				mColor := &colorundefjsonnil.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"Blue"`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorundefjsonnil.Blue))
			})
		})
		g.When("empty json", func() {
			g.It("then Undefined", func() {
				mColor := &colorundefjsonnil.MarshallableColor{}
				err := mColor.UnmarshalJSON(make([]byte, 0))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorundefjsonnil.Undefined))
			})
		})
		g.When("null json", func() {
			g.It("then Undefined", func() {
				mColor := &colorundefjsonnil.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`null`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorundefjsonnil.Undefined))
			})
		})
		g.When("invalid-color json", func() {
			g.It("then Undefined", func() {
				mColor := &colorundefjsonnil.MarshallableColor{}
				err := mColor.UnmarshalJSON([]byte(`"invalid-color"`))
				o.Expect(err).ShouldNot(o.HaveOccurred())
				o.Expect(mColor.FromMarshallable()).Should(o.Equal(colorundefjsonnil.Undefined))
			})
		})
	})
})

//go:embed colorwithoutcopyright/expected_color.txt
var expectedColorWithoutCopyrightClause []byte

var _ = g.Describe("generate without copyright clause", func() {
	destination := "./colorwithoutcopyright/color.go"

	g.It("should generate", func() {
		enum := generator.Enum{
			Destination:    &destination,
			Package:        "color",
			Type:           "Color",
			Values:         []string{"Undefined", "Red", "Green", "Blue"},
			UndefinedValue: "Undefined",
			Marshalling: generator.MarshalOptions{
				JSONOptions: generator.JSONMarshalOptions{
					Generate:       true,
					NilToUndefined: true,
				},
			},
		}
		err := generator.Generate(enum)

		o.Expect(err).ShouldNot(o.HaveOccurred())
	})

	g.It("should be equal to expected", func() {
		colorContent, err := os.ReadFile(destination)

		o.Expect(err).ShouldNot(o.HaveOccurred())
		o.Expect(colorContent).Should(o.Equal(expectedColorWithoutCopyrightClause))
	})
})
