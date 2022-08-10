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

type Codes struct {
	generators []Generator
}

func NewCodes() *Codes {
	p := &Codes{
		generators: make([]Generator, 0, 0),
	}
	return p
}

func GetEncodeRequestName(serviceName, methodName string) string {
	funcName := "Encode" + serviceName + methodName + "Request"
	return funcName
}

func GetDecodeRequestName(serviceName, methodName string) string {
	funcName := "Decode" + serviceName + methodName + "Request"
	return funcName
}

func GetEncodeResponseName(serviceName, methodName string) string {
	funcName := "Encode" + serviceName + methodName + "Response"
	return funcName
}

func GetDecodeResponseName(serviceName, methodName string) string {
	funcName := "Decode" + serviceName + methodName + "Response"
	return funcName
}

func (this *Codes) generateEncodeRequest(file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) error {

	gfile.P()
	for _, method := range svc.Methods {
		// funcName := "Encode" + svc.GoName + method.GoName + "Request"
		funcName := GetEncodeRequestName(svc.GoName, method.GoName)
		gfile.P("//// THIS IS FOR CLIENT INVKE SERVICE REQUEST ENCODE")
		gfile.P("//// CLIENT.REQUEST - > SERVICE.REQUEST")

		gfile.P(method.Comments.Leading,
			"func ",
			funcName,
			"(ctx ", common.ContextPackage.Ident("Context"), ", req *", common.HttpPackage.Ident("Request"), " , request interface{} )",
			"error {",
		)
		gfile.P("/// TODO")
		gfile.P("return ", CommonImporPath(options).Ident("EncodeJSONRequest"), "(ctx, req, request)")
		gfile.P("}")
	}
	gfile.P("//// ")

	return nil
}

func (this *Codes) generateDecodeRequest(file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) error {

	gfile.P()
	for _, method := range svc.Methods {
		funcName := GetDecodeRequestName(svc.GoName, method.GoName)
		gfile.P("//// THIS IS FOR SERVICE DEOCDE REQUEST")
		gfile.P("//// CLIENT.REQUEST -> SERVICE.REQUEST")

		gfile.P(method.Comments.Leading,
			"func ",
			funcName,
			"(ctx ", common.ContextPackage.Ident("Context"), " , r *", common.HttpPackage.Ident("Request"), " )",
			"(interface{} , error){",
		)

		gfile.P("var request ", common.GetPackagePath("service", string(file.GoPackageName), options).Ident(method.Input.GoIdent.GoName))
		gfile.P(`
            err := json.NewDecoder(r.Body).Decode(&request)
            if nil != err {
                panic("Error decode request: " + err.Error())
                            
            }
        `)
		gfile.P("return nil, ", common.FmtPackage.Ident("Errorf"), `("Error convert interface : %v", `, common.ReflectPackage.Ident(""), "TypeOf(request))")
		gfile.P("}")
	}

	gfile.P("//// ")
	return nil
}

func (this *Codes) generateEncodeResponse(file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) error {

	gfile.P()
	for _, method := range svc.Methods {
		funcName := "Encode" + svc.GoName + method.GoName + "Response"
		gfile.P("//// THIS IS FOR SERVICE ENCODE RESPONSE")
		gfile.P("//// SERVICE.RESPONSE -> CLIENT.RESPONSE")

		gfile.P(method.Comments.Leading,
			"func ",
			funcName,
			"(ctx ", common.ContextPackage.Ident("Context"), " , w ", common.HttpPackage.Ident("ResponseWriter"), " , resp interface{} )",
			" error{",
		)
		gfile.P("//// TODO")
		gfile.P(`w.Header().Set("Content-Type", "application/json")
            return json.NewEncoder(w).Encode(resp)`)
		gfile.P("}")
	}
	gfile.P("//// ")
	return nil
}

func (this *Codes) generateDecodeResponse(file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) error {

	gfile.P()
	for _, method := range svc.Methods {
		funcName := "Decode" + svc.GoName + method.GoName + "Response"
		gfile.P("//// THIS IS FOR CLIENT DECODE RESPONSE FROM SERVICE")
		gfile.P("//// SERVICE.RESPONSE -> CLIENT.RESPONSE")

		gfile.P(method.Comments.Leading,
			"func ",
			funcName,
			"(ctx ", common.ContextPackage.Ident("Context"), " , resp *", common.HttpPackage.Ident("Response"), " )",
			"(interface{} , error){",
		)

		gfile.P("//// TODO")
		gfile.P("var r ", common.GetPackagePath("service", string(file.GoPackageName), options).Ident(method.Output.GoIdent.GoName))
		gfile.P(`
            if err := json.NewDecoder(resp.Body).Decode(&r);nil != err {
                return nil, err 
            }
            return &r, nil`)
		gfile.P("}")
	}
	gfile.P("//// ")
	return nil
}

func (this *Codes) GenerateFile(gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) (*protogen.GeneratedFile, error) {
	this.generateEncodeRequest(file, gfile, svc, options)
	this.generateDecodeRequest(file, gfile, svc, options)
	this.generateEncodeResponse(file, gfile, svc, options)
	this.generateDecodeResponse(file, gfile, svc, options)
	return gfile, nil
}
