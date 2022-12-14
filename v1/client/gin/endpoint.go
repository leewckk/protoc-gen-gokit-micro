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

package gin

import (
	"github.com/leewckk/protoc-gen-gokit-micro/common"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
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
	return "Make" + service + method + "ClientEndpoint"
}

func (this *Endpoints) generate(gfile *protogen.GeneratedFile, file *protogen.File, svc *protogen.Service, packageName string, options *common.Options) error {

	serviceName := svc.GoName
	// transportCommonImport := common.GetPackagePath("invoker/http", "common", options)
	transportCommonImport := common.GokitServiceEndpointHttp

	gfile.P()
	for _, method := range svc.Methods {

		endpointName := EndpointName(serviceName, method.GoName)
		encRequestName := GetEncodeRequestName(serviceName, method.GoName)
		decResponseName := GetDecodeResponseName(serviceName, method.GoName)

		gfile.P("func ", endpointName, "(serverName string, client ", common.GoKitSDConsulPackage.Ident("Client"), " , opts ...", common.GoKitHTTP.Ident("ClientOption"), ") endpoint.Endpoint{")
		gfile.P("enc := ", (encRequestName))
		gfile.P("dec := ", (decResponseName))

		if opt, ok := method.Desc.Options().(*descriptorpb.MethodOptions); ok {
			if rule, ok := proto.GetExtension(opt, annotations.E_Http).(*annotations.HttpRule); ok {

				out := func(methodName, pattern string) {
					gfile.P(`api := "`, pattern, `"`)
					gfile.P(`httpMethod := "`, methodName, `"`)
				}

				m, p, err := common.ParseHttpOption(rule)
				if nil == err {
					out(m, p)
				}
			}
		}

		gfile.P("return ", transportCommonImport.Ident("MakeClientEndpoint"), "(client, serverName, httpMethod, api, enc, dec, opts...)")
		gfile.P("}")
	}
	gfile.P()
	return nil
}

func (this *Endpoints) GenerateFile(fileUri string, gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, options *common.Options) (*protogen.GeneratedFile, error) {

	packageName := string(file.GoPackageName)
	for _, svc := range file.Services {
		if err := this.generate(gfile, file, svc, packageName, options); nil != err {
			return nil, err
		}
	}
	return gfile, nil
}
