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
//

package generator

type jsonMarshallerGenerator struct {
	enum   generationEnum
	writer *Writer
}

func newJSONMarshallerGenerator(
	enum generationEnum,
	writer *Writer,
) *jsonMarshallerGenerator {
	return &jsonMarshallerGenerator{
		enum:   enum,
		writer: writer,
	}
}

func (g *jsonMarshallerGenerator) generateImports() {
	if !g.enum.Marshalling.JSONOptions.Generate {
		return
	}
	if !g.enum.Marshalling.JSONOptions.NilToUndefined {
		g.writer.Line("\t\"errors\"")
	}

	g.writer.Line("\t\"strings\"")
}

func (g *jsonMarshallerGenerator) generateToMarshallerDeclaration() {
	if !g.enum.Marshalling.JSONOptions.Generate {
		return
	}
	g.writer.Line("\tToMarshallable() " + g.enum.marshallableStruct)
}

func (g *jsonMarshallerGenerator) generateJSONMarshalling() {
	if !g.enum.Marshalling.JSONOptions.Generate {
		return
	}
	g.generateMarshallableStruct()
	g.generateMarshalJSON()
	g.generateUnmarshalJSON()
	g.generateToMarshallable()
	g.generateFromMarshallable()
}

func (g *jsonMarshallerGenerator) generateMarshallableStruct() {
	w := g.writer
	e := g.enum
	w.Line("type " + e.marshallableStruct + " struct {")
	w.Line("\ten " + e.Type)
	w.Line("}")
	w.LineBreak()
}

func (g *jsonMarshallerGenerator) generateMarshalJSON() {
	w := g.writer
	e := g.enum
	w.Line("func (b " + e.marshallableStruct + ") MarshalJSON() ([]byte, error) {")
	w.Line("\treturn []byte(\"\\\"\" + b.en.String() + \"\\\"\"), nil")
	w.Line("}")
	w.LineBreak()
}

func (g *jsonMarshallerGenerator) generateUnmarshalJSON() {
	w := g.writer
	e := g.enum
	w.Line("func (b *" + e.marshallableStruct + ") UnmarshalJSON(jsonBytes []byte) error {")

	g.generateUnmarshalFromEmptyBytes()

	w.Line("\tjsonString := bytes.NewBuffer(jsonBytes).String();")

	g.generateUnmarshalFromNull()
	g.generateUnmarshalFromString()

	w.Line("\treturn nil")

	w.Line("}")
	w.LineBreak()
}

func (g *jsonMarshallerGenerator) generateUnmarshalFromEmptyBytes() {
	w := g.writer
	e := g.enum

	// if len(jsonBytes) == 0 {
	w.Line("\tif len(jsonBytes) == 0 {")

	// b = Undefined
	if e.Marshalling.JSONOptions.NilToUndefined {
		w.Line("\t\tb.en = " + e.UndefinedValue)
	} else { // or just return
		w.Line("\t\treturn nil")
	}

	// } end if len(jsonBytes) == 0
	w.Line("\t}")

	w.LineBreak()
}

func (g *jsonMarshallerGenerator) generateUnmarshalFromNull() {
	w := g.writer
	e := g.enum

	// if jsonString == "null" {
	w.Line("\tif jsonString == \"null\" {")

	// b = Undefined
	if e.Marshalling.JSONOptions.NilToUndefined {
		w.Line("\t\tb.en = " + e.UndefinedValue)
	} else { // or just return
		w.Line("\t\treturn nil")
	}

	// } end if jsonString == "null"
	w.Line("\t}")

	w.LineBreak()
}

func (g *jsonMarshallerGenerator) generateUnmarshalFromString() {
	w := g.writer
	e := g.enum

	w.Line("\ttrimmedString := strings.Trim(jsonString, \"\\\"\")")

	// OfOrUndefined
	if e.Marshalling.JSONOptions.NilToUndefined {
		w.Line("\torUndefined := OfOrUndefined(trimmedString)")
		w.Line("\tb.en = orUndefined")
	} else { // or fail
		w.Line("\tvalue, err := Of(trimmedString)")
		w.Line("\tif err != nil {")
		w.Line("\t\treturn errors.Join(errors.New(\"could not unmarshal " + e.Type + " from JSON\"), err)")
		w.Line("\t}")
		w.Line("\tb.en = value")
	}

	w.LineBreak()
}

func (g *jsonMarshallerGenerator) generateToMarshallable() {
	w := g.writer
	e := g.enum

	w.Line("func (b " + e.baseStruct + ") ToMarshallable() " + e.marshallableStruct + " {")
	w.Line("\treturn " + e.marshallableStruct + "{en: b}")
	w.Line("}")
	w.LineBreak()
}

func (g *jsonMarshallerGenerator) generateFromMarshallable() {
	w := g.writer
	e := g.enum

	w.Line("func (m " + e.marshallableStruct + ") FromMarshallable() " + e.Type + " {")
	w.Line("\treturn m.en")
	w.Line("}")
	w.LineBreak()
}
