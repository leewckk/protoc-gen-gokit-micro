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
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var __fieldTypeDict__ map[string]string

func init() {
	__fieldTypeDict__ = make(map[string]string)
	__fieldTypeDict__["float"] = "float32"
	__fieldTypeDict__["double"] = "float64"
	__fieldTypeDict__["bytes"] = "[]byte"
	__fieldTypeDict__["sint32"] = "int32"
	__fieldTypeDict__["sint64"] = "int64"
	__fieldTypeDict__["fixed32"] = "uint32"
	__fieldTypeDict__["fixed64"] = "uint64"
	__fieldTypeDict__["sfixed32"] = "int32"
	__fieldTypeDict__["sfixed64"] = "int64"
	__fieldTypeDict__["float"] = "float32"
	__fieldTypeDict__["float"] = "float32"

}

func IsBaseType(desc protoreflect.FieldDescriptor) bool {

	if _, ok := __fieldTypeDict__[desc.Kind().String()]; ok {
		return true
	}
	return false
}

func fieldTypePrefix(field *protogen.Field) string {

	if field.Desc.Cardinality() == protoreflect.Repeated {
		if field.Desc.Kind() == protoreflect.MessageKind {
			return "[]*"
		}
		return "[]"
	}

	if field.Desc.HasOptionalKeyword() {
		return "*"
	}

	if field.Desc.Kind() == protoreflect.MessageKind {
		return "*"
	}
	return ""
}

func FieldMapKeyTypeName(field *protogen.Field) string {
	return field.Desc.MapKey().Kind().String()
}

func FieldMapValTypeName(field *protogen.Field) string {
	kindString := field.Desc.MapValue().Kind().String()
	if v, ok := __fieldTypeDict__[kindString]; ok {
		return v
	}
	return "*" + string(field.Desc.MapValue().Message().Name())
}

// func FieldMapValTypeFullName(module protogen.GoImportPath, field *protogen.Field) string {
// 	kindString := field.Desc.MapValue().Kind().String()
// 	if v, ok := __fieldTypeDict__[kindString]; ok {
// 		return v
// 	}
// 	return "*" + module.Ident(string("*" + field.Desc.MapValue().Message().Name())).String()
// }

func FieldMapValTypeNamePure(field *protogen.Field) string {
	kindString := field.Desc.MapValue().Kind().String()
	if v, ok := __fieldTypeDict__[kindString]; ok {
		return v
	}
	return string(field.Desc.MapValue().Message().Name())
}

func GenFieldType(field *protogen.Field) string {

	kindString := field.Desc.Kind().String()
	if v, ok := __fieldTypeDict__[kindString]; ok {
		return fieldTypePrefix(field) + v
	}

	if field.Desc.IsMap() {
		return fmt.Sprintf("map[%v]%v", FieldMapKeyTypeName(field), FieldMapValTypeName(field))
	}

	if field.Message != nil {
		return fieldTypePrefix(field) + field.Message.GoIdent.GoName
	}

	return fieldTypePrefix(field) + field.Desc.Kind().String()
}

// func FieldTypeName(desc protoreflect.FieldDescriptor) string {
// 	kindString := desc.Kind().String()
// 	if v, ok := __fieldTypeDict__[kindString]; ok {
// 		return v
// 	}

// 	// if field.Desc.IsMap() {
// 	// 	return fmt.Sprintf("map[%v]%v", fieldMapKey(field), fieldMapVal(field))
// 	// }

// 	return desc.Kind().String()
// }
