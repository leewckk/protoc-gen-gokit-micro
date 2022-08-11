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

func GetDecodeResponseName(serviceName, methodName string) string {
	funcName := "Decode" + serviceName + methodName + "Response"
	return funcName
}

func CommonImporPath(options *common.Options) protogen.GoImportPath {
	return common.GokitServiceTransportHttp
	// return protogen.GoImportPath(options.GetModName() + "/transport/gin/common")
}

func GinHttpMiddlewareImportPath(options *common.Options) protogen.GoImportPath {
	return common.GokitServiceTransportHttp
	// return protogen.GoImportPath(options.GetModName() + "/middlewares/transport/http/gin")
}

func (this *Codes) userStartBlock(serviceName string, methodName string) string {
	return `// @@USER.` + serviceName + "." + methodName
}

func (this *Codes) userEndBlock(serviceName string, methodName string) string {
	return `// $$USER.` + serviceName + "." + methodName
}

func (this *Codes) generateEncodeRequest(fileUri string, file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) error {

	gfile.P()
	for _, method := range svc.Methods {

		funcName := GetEncodeRequestName(svc.GoName, method.GoName)
		gfile.P("//// THIS IS FOR CLIENT INVKE SERVICE REQUEST ENCODE")
		gfile.P("//// CLIENT.REQUEST - > SERVICE.REQUEST")

		gfile.P(method.Comments.Leading,
			"func ",
			funcName,
			"(ctx ", common.ContextPackage.Ident("Context"), ", req *", common.HttpPackage.Ident("Request"), " , request interface{} )",
			"error {",
		)

		serviceName := svc.GoName
		// gfile.P(this.userStartBlock(serviceName, method.GoName))

		userblock := common.UserBlock(fileUri, serviceName, method.GoName,
			this.userStartBlock(serviceName, method.GoName), this.userEndBlock(serviceName, method.GoName))
		if "" == userblock {
			gfile.P("/// TODO")
			gfile.P(`
            pattern := req.URL.Path
            pattern = `, GinHttpMiddlewareImportPath(options).Ident("MarshalPattern"), `(pattern, request)
            req.URL.Path = pattern`)
			gfile.P("return ", common.GokitServiceTransportHttp.Ident("EncodeJSONRequest"), "(ctx, req, request)")

		} else {
			gfile.P(userblock)
		}
		gfile.P("}")
	}
	gfile.P("//// ")

	return nil
}

func (this *Codes) generateDecodeResponse(fileUri string, file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) error {

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

		serviceName := svc.GoName
		methodName := funcName

		userblock := common.UserBlock(fileUri, serviceName, methodName,
			this.userStartBlock(serviceName, methodName), this.userEndBlock(serviceName, methodName))
		if "" == userblock {
			gfile.P("//// TODO")
			gfile.P("var r ", common.GetPackagePath("invoker/protocol", string(file.GoPackageName), options).Ident(method.Output.GoIdent.GoName))
			gfile.P(`
            if err := json.NewDecoder(resp.Body).Decode(&r);nil != err {
                return nil, err 
            }
            return &r, nil`)
		} else {
			gfile.P(userblock)
		}
		gfile.P("}")
	}
	gfile.P("//// ")
	return nil
}

func (this *Codes) GenerateFile(fileUri string, gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, options *common.Options) (*protogen.GeneratedFile, error) {

	for _, svc := range file.Services {
		this.generateEncodeRequest(fileUri, file, gfile, svc, options)
		this.generateDecodeResponse(fileUri, file, gfile, svc, options)
	}
	return gfile, nil
}
