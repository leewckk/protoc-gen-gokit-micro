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

func GetDecodeRequestName(serviceName, methodName string) string {
	funcName := "Decode" + serviceName + methodName + "Request"
	return funcName
}

func GetEncodeResponseName(serviceName, methodName string) string {
	funcName := "Encode" + serviceName + methodName + "Response"
	return funcName
}

func (this *Codes) userStartBlock(serviceName string, methodName string) string {
	return `// @@USER.` + serviceName + "." + methodName
}

func (this *Codes) userEndBlock(serviceName string, methodName string) string {
	return `// $$USER.` + serviceName + "." + methodName
}

func (this *Codes) generateDecodeRequest(fileUri string, file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) error {

	parts := make(map[string]func())
	parts["GET"] = func() {
		gfile.P(`
                if err := c.ShouldBindUri(&request); err == nil {
                    return &request, nil 
                }else {
                    return nil, fmt.Errorf("Error decode request , uri: %v, err: %v", c.Request.URL, err)
                }`)
	}

	parts["JSON"] = func() {
		gfile.P(`
                if err := c.ShouldBindJSON(&request); err == nil {
                    return &request, nil 
                }else {
                    return nil, fmt.Errorf("Error decode request , uri: %v, err: %v", c.Request.URL, err)
                }`)
	}

	gfile.P()
	for _, method := range svc.Methods {

		for part, out := range parts {
			funcName := GetDecodeRequestName(svc.GoName, method.GoName) + part
			gfile.P("//// THIS IS FOR SERVICE DEOCDE REQUEST")
			gfile.P("//// CLIENT.REQUEST -> SERVICE.REQUEST")

			gfile.P(method.Comments.Leading,
				"func ",
				funcName,
				"(c* ", common.GinPackage.Ident("Context"), ")",
				"(interface{} , error){",
			)

			serviceName := svc.GoName
			methodName := funcName
			// gfile.P(this.userStartBlock(serviceName, methodName))

			userblock := common.UserBlock(fileUri, serviceName, methodName,
				this.userStartBlock(serviceName, methodName), this.userEndBlock(serviceName, methodName))
			if "" == userblock {
				gfile.P("var request ", common.GetPackagePath("endpoint", string(file.GoPackageName), options).Ident(method.Input.GoIdent.GoName))
				out()
			} else {
				gfile.P(userblock)
			}
			// gfile.P(this.userEndBlock(serviceName, methodName))
			gfile.P("}")
		}
	}
	gfile.P("//// ")
	return nil
}

func (this *Codes) generateEncodeResponse(fileUri string, file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) error {

	gfile.P()
	for _, method := range svc.Methods {
		funcName := "Encode" + svc.GoName + method.GoName + "Response"
		gfile.P("//// THIS IS FOR SERVICE ENCODE RESPONSE")
		gfile.P("//// SERVICE.RESPONSE -> CLIENT.RESPONSE")

		gfile.P(method.Comments.Leading,
			"func ",
			funcName,
			"(c *", common.GinPackage.Ident("Context"), " , resp interface{} )",
			" error{",
		)
		serviceName := svc.GoName
		methodName := funcName
		// gfile.P(this.userStartBlock(serviceName, methodName))

		userblock := common.UserBlock(fileUri, serviceName, methodName,
			this.userStartBlock(serviceName, methodName), this.userEndBlock(serviceName, methodName))
		if "" == userblock {
			gfile.P(`c.JSON(`, common.HttpPackage.Ident("StatusOK"), `,resp)
        return nil `)

		} else {
			gfile.P(userblock)
		}
		// gfile.P(this.userEndBlock(serviceName, methodName))

		gfile.P("}")
	}
	gfile.P("//// ")
	return nil
}

func (this *Codes) GenerateFile(fileUri string, gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) (*protogen.GeneratedFile, error) {
	this.generateDecodeRequest(fileUri, file, gfile, svc, options)
	this.generateEncodeResponse(fileUri, file, gfile, svc, options)
	return gfile, nil
}
