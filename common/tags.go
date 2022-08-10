// The MIT License (MIT)

// Copyright (c) 2022 leewckk@126.com

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package common

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func omitEmpty(field *protogen.Field) string {

	omitempty := ",omitempty"
	if field.Desc.Cardinality() == protoreflect.Repeated {
		if field.Desc.Kind() == protoreflect.MessageKind {
			return omitempty
		}
		return omitempty
	}

	if field.Desc.HasOptionalKeyword() {
		return omitempty
	}

	if field.Desc.Kind() == protoreflect.MessageKind {
		return omitempty
	}
	return ""
}

func GenerateTags(field *protogen.Field) string {
	return "json:\"" + field.Desc.JSONName() + omitEmpty(field) + "\"" + ` uri:"` + field.Desc.JSONName() + `"` + ` form:"` + field.Desc.JSONName() + `"`
}
