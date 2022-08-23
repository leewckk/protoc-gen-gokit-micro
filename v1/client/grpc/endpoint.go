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

package grpc

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
	return "Make" + service + method + "ClientEndpoint"
}

func (this *Endpoints) generate(gfile *protogen.GeneratedFile, file *protogen.File, svc *protogen.Service, packageName string, options *common.Options) error {

	serviceName := svc.GoName
	pbImort := common.GetPBPackagePath(file, options)
	gfile.P()
	for _, method := range svc.Methods {

		endpointName := EndpointName(serviceName, method.GoName)
		encRequestName := GetEncodeRequestName(serviceName, method.GoName)
		decResponseName := GetDecodeResponseName(serviceName, method.GoName)

		gfile.P("func ", endpointName, "(serverName string, ", " client ", common.GoKitSDConsulPackage.Ident("Client"), " , opts... ", common.GoKitGRPC.Ident("ClientOption"), ") endpoint.Endpoint{")
		// gfile.P("var reply ", pbImort.Ident(method.Output.GoIdent.GoName))
		gfile.P("replyMaker := func()interface{}{i return new(", pbImort.Ident(method.Output.GoIdent.GoName), ")}")
		gfile.P("enc := ", (encRequestName))
		gfile.P("dec := ", (decResponseName))
		gfile.P("serviceName := \"", file.GoPackageName, ".", svc.GoName, "\"")
		gfile.P("methodName := \"", method.GoName, "\"")
		// gfile.P("return ", common.GokitServiceEndpointGrpc.Ident("MakeClientEndpoint"), "(client, serverName, serviceName, methodName, enc, dec, &reply, opts...)")
		gfile.P("return ", common.GokitServiceEndpointGrpc.Ident("MakeClientEndpoint"), "(client, serverName, serviceName, methodName, enc, dec, replyMaker, opts...)")
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
