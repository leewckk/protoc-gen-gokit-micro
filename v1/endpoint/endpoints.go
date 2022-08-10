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
package endpoint

import (
	"github.com/leewckk/protoc-gen-gokit-micro/common"
	"google.golang.org/protobuf/compiler/protogen"
)

type Endpoints struct {
	generators []Generator
}

func NewEndpoints() *Endpoints {
	g := &Endpoints{}
	g.generators = make([]Generator, 0, 0)
	return g
}

func EndpointName(service, method string) string {
	return "Make" + service + method + "Endpoint"
}

func (this *Endpoints) generate(gfile *protogen.GeneratedFile, svc *protogen.Service, packageName string, options *common.Options) error {

	serviceName := svc.GoName
	gfile.P()
	for _, method := range svc.Methods {

		endpointName := EndpointName(serviceName, method.GoName)
		gfile.P("func ", endpointName,
			" () ", common.GoKitEndpoint.Ident("Endpoint"), "{")
		gfile.P("return ", MiddleWareImporPath(options).Ident(""), "InjectMiddleWares(func(ctx ", common.ContextPackage.Ident("Context"), " , request interface{}) (interface{} , error){")

		gfile.P("if r, ok := request.(*", common.GenMessageName(method.Input), ");ok {")
		gfile.P("return ", PrototypeProcName(svc, method), "(ctx, r)")
		gfile.P("}")

		gfile.P("return nil, ", common.FmtPackage.Ident("Errorf"), `("Error convert interface : %v", `, common.ReflectPackage.Ident(""), "TypeOf(request))")
		gfile.P("})")
		gfile.P("}")
	}

	gfile.P()

	return nil
}

func (this *Endpoints) GenerateFile(gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, options *common.Options) (*protogen.GeneratedFile, error) {

	gfile.P("/// ENDPOINT MAKER")
	packageName := string(file.GoPackageName)
	for _, svc := range file.Services {
		if err := this.generate(gfile, svc, packageName, options); nil != err {
			return nil, err
		}
	}
	return gfile, nil
}
