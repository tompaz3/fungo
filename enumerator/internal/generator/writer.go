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

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strings"
)

type Writer struct {
	writer *bufio.Writer
	errors []error
}

func NewWriter(writer *bufio.Writer) *Writer {
	return &Writer{
		writer: writer,
		errors: make([]error, 0),
	}
}

func (w *Writer) Line(value string) *Writer {
	return w.String(value).
		String("\n")
}

func (w *Writer) LineBreak() *Writer {
	return w.String("\n")
}

func (w *Writer) String(value string) *Writer {
	_, err := w.writer.WriteString(value)
	if err != nil {
		w.errors = append(w.errors, err)
	}
	return w
}

func (w *Writer) FileContentCommented(filePath string) *Writer {
	content, err := os.ReadFile(filePath)
	if err != nil {
		w.errors = append(w.errors, err)
		return w
	}

	copyrightString := bytes.NewBuffer(content).String()
	for _, line := range strings.Split(copyrightString, "\n") {
		w.Line("// " + line)
	}

	return w
}

func (w *Writer) Flush() error {
	if err := w.writer.Flush(); err != nil {
		w.errors = append(w.errors, err)
	}
	return errors.Join(w.errors...)
}
