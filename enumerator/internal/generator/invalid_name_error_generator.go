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

type invalidNameErrorGenerator struct {
	enum   generationEnum
	writer *Writer
}

func newInvalidNameErrorGenerator(
	enum generationEnum,
	writer *Writer,
) *invalidNameErrorGenerator {
	return &invalidNameErrorGenerator{
		enum:   enum,
		writer: writer,
	}
}

func (g *invalidNameErrorGenerator) generateInvalidNameError() {
	w := g.writer
	e := g.enum
	w.Line("type Invalid" + e.Type + "NameError struct {")
	w.Line("\tname string")
	w.Line("}")
	w.LineBreak()
	w.Line("func (e " + e.invalidNameError + ") Error() string {")
	w.Line("\treturn \"invalid " + e.Type + " name: \" + e.name")
	w.Line("}")
	w.LineBreak()
	w.Line("func " + e.invalidNameErrorConstructor + "(name string) " + e.invalidNameError + " {")
	w.Line("\treturn " + e.invalidNameError + "{name: name}")
	w.Line("}")
	w.LineBreak()
}
