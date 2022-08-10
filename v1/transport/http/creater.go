// The MIT License (MIT)
//
// Copyright (c) 2022 leewckk@126.com
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
//

package http

import (
	"github.com/leewckk/protoc-gen-gokit-micro/common"

	"google.golang.org/protobuf/compiler/protogen"
)

type CreaterFunc struct {
	generators []Generator
}

func NewCreaterFunc() *CreaterFunc {
	p := &CreaterFunc{
		generators: make([]Generator, 0, 0),
	}
	return p
}

func GetCreateTransportFuncName(serviceName string) string {
	funcName := "New" + serviceName + "Transport"
	return funcName
}

func (this *CreaterFunc) GenerateFile(gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) (*protogen.GeneratedFile, error) {
	serviceName := svc.GoName + "ServiceTransport"
	funcName := GetCreateTransportFuncName(svc.GoName)
	gfile.P()
	gfile.P("func ", funcName, "(opts ...", common.GoKitHTTP.Ident("ServerOption"), ") *", serviceName, "{")
	gfile.P()
	gfile.P("var t ", serviceName)

	gfile.P(`if len(opts) > 0 {
        t.options = append(t.options, opts...)
    }`)
	gfile.P("t.options = append(t.options,", CommonImporPath(options).Ident("ErrorServerOption"), "())")
	gfile.P("return &t")
	gfile.P("}")
	gfile.P()
	return gfile, nil
}
