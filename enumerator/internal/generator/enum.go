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

import "errors"

var (
	ErrEmptyPackage                           = errors.New("package name is empty")
	ErrEmptyType                              = errors.New("type name is empty")
	ErrEmptyValues                            = errors.New("values are empty")
	ErrUndefinedValueNotFound                 = errors.New("undefined value not found in values")
	ErrUndefinedValueForUnmarshallingNotFound = errors.New("undefined value for unmarshalling not found")
)

type Enum struct {
	Destination *string

	CopyrightFile string

	Package string
	Type    string
	Values  []string

	UndefinedValue string
	Marshalling    MarshalOptions
}

type MarshalOptions struct {
	JSONOptions JSONMarshalOptions
}

type JSONMarshalOptions struct {
	Generate       bool
	NilToUndefined bool
}

func (e Enum) validate() error {
	if e.Package == "" {
		return ErrEmptyPackage
	}
	if e.Type == "" {
		return ErrEmptyType
	}
	if len(e.Values) == 0 {
		return ErrEmptyValues
	}

	return e.validateUndefined()
}

func (e Enum) validateUndefined() error {
	if e.UndefinedValue != "" {
		found := false
		for _, value := range e.Values {
			if value == e.UndefinedValue {
				found = true
				break
			}
		}
		if !found {
			return ErrUndefinedValueNotFound
		}
	}

	if e.Marshalling.JSONOptions.Generate &&
		e.Marshalling.JSONOptions.NilToUndefined &&
		e.UndefinedValue == "" {
		return ErrUndefinedValueForUnmarshallingNotFound
	}
	return nil
}
