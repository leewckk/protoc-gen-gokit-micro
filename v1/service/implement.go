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
package service

import (
	"github.com/leewckk/protoc-gen-gokit-micro/common"
	"google.golang.org/protobuf/compiler/protogen"
)

type Implement struct {
	generators []common.Generator
}

func NewImplement() *Implement {
	g := &Implement{}
	g.generators = make([]common.Generator, 0, 0)
	return g
}

func (this *Implement) userStartBlock(serviceName string, methodName string) string {
	return `// @@USER.` + serviceName + "." + methodName
}

func (this *Implement) userEndBlock(serviceName string, methodName string) string {
	return `// $$USER.` + serviceName + "." + methodName
}

func (this *Implement) generateMethods(fileUri, serviceName string, gfile *protogen.GeneratedFile, methods []*protogen.Method) error {

	for _, method := range methods {
		gfile.P(method.Comments.Leading, "func (this* "+serviceName,
			")",
			method.GoName,
			"(ctx ", common.ContextPackage.Ident("Context"),
			",request *", common.GenMessageName(method.Input),
			") (*", common.GenMessageName(method.Output), ", error) {")
		gfile.P(this.userStartBlock(serviceName, method.GoName))

		userblock := common.UserBlock(fileUri, serviceName, method.GoName,
			this.userStartBlock(serviceName, method.GoName), this.userEndBlock(serviceName, method.GoName))
		if "" == userblock {
			gfile.P("resp := &", common.GenMessageName(method.Output), "{}")
			gfile.P("return resp, nil")
		} else {
			gfile.P(userblock)
		}
		gfile.P(this.userEndBlock(serviceName, method.GoName))
		gfile.P("}")
	}
	gfile.P()
	return nil
}

func (this *Implement) generate(fileUri string, gfile *protogen.GeneratedFile, svc *protogen.Service) error {

	serviceName := svc.GoName + "ServiceImplement"
	gfile.P()
	gfile.P("type ", serviceName, " struct {")
	gfile.P("}")

	gfile.P()
	gfile.P("func New", svc.GoName+"Service() ", svc.GoName+"Service {")

	methodName := "New"

	gfile.P(this.userStartBlock(serviceName, methodName))
	userblock := common.UserBlock(fileUri, serviceName, methodName,
		this.userStartBlock(serviceName, methodName), this.userEndBlock(serviceName, methodName))

	if "" == userblock {
		gfile.P("s := &", serviceName+"{}")
		gfile.P("return s")
	} else {
		gfile.P(userblock)
	}
	gfile.P(this.userEndBlock(serviceName, methodName))

	gfile.P("}")
	gfile.P()
	return this.generateMethods(fileUri, serviceName, gfile, svc.Methods)
}

func (this *Implement) GenerateFile(gen *protogen.Plugin, file *protogen.File, options *common.Options) (*protogen.GeneratedFile, error) {

	packageName := string(file.GoPackageName)
	protoFileName := common.GetFileName(file)

	fileName := common.OutPutFileName(__subPath__,
		packageName,
		protoFileName,
		common.LAYER_NAME_SERVICE, "", "",
		options, func(layer *common.LayerConfig) bool {
			if nil != layer {
				return layer.Replace
			}
			return true
		})
	gfile := gen.NewGeneratedFile(fileName, file.GoImportPath)
	gfile, _ = common.GenerateHeader(gfile, packageName)

	for _, svc := range file.Services {
		if err := this.generate(fileName, gfile, svc); nil != err {
			return nil, err
		}
	}
	return gfile, nil
}
