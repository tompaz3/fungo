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

type ofStringGenerator struct {
	enum   generationEnum
	writer *Writer
}

func newOfStringGenerator(
	enum generationEnum,
	writer *Writer,
) *ofStringGenerator {
	return &ofStringGenerator{
		enum:   enum,
		writer: writer,
	}
}

func (g *ofStringGenerator) generateOfStringMethods() {
	g.generateOfString()
	g.generateOfOrUndefined()
}

func (g *ofStringGenerator) generateOfString() {
	w := g.writer
	e := g.enum
	w.Line("func Of(name string) (" + e.Type + ", error) {")
	w.Line("\tif value, ok := allValuesByString[name]; ok {")
	w.Line("\t\treturn value, nil")
	w.Line("\t}")
	w.Line("\treturn nil, " + e.invalidNameErrorConstructor + "(name)")
	w.Line("}")
	w.LineBreak()
}

func (g *ofStringGenerator) generateOfOrUndefined() {
	if g.enum.UndefinedValue == "" {
		return
	}

	w := g.writer
	e := g.enum
	w.Line("func OfOrUndefined(name string) " + e.Type + " {")
	w.Line("\tif value, ok := allValuesByString[name]; ok {")
	w.Line("\t\treturn value")
	w.Line("\t}")
	w.Line("\treturn " + e.UndefinedValue)
	w.Line("}")
	w.LineBreak()
}
