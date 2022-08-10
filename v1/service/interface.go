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

type Interface struct {
	generators []common.Generator
}

func NewInterface() *Interface {
	g := &Interface{}
	g.generators = make([]common.Generator, 0, 0)
	return g
}

// func (this *Interface) FileName(packageName string, protoFileName string) string {
// 	return __subPath__ + packageName + fmt.Sprintf("/%v.interface.go", protoFileName)
// }

func (this *Interface) generate(gfile *protogen.GeneratedFile, svc *protogen.Service) error {

	gfile.P()
	gfile.P("type ", svc.GoName+"Service", " interface{")
	for _, method := range svc.Methods {
		gfile.P(method.Comments.Leading,
			method.GoName,
			"(", common.ContextPackage.Ident("Context"),
			", *", common.GenMessageName(method.Input),
			") (*", common.GenMessageName(method.Output), ", error)")
	}
	gfile.P("}")
	gfile.P()
	return nil
}

func (this *Interface) GenerateFile(gen *protogen.Plugin, file *protogen.File, options *common.Options) (*protogen.GeneratedFile, error) {

	packageName := string(file.GoPackageName)
	protoFileName := common.GetFileName(file)

	// fileName := common.OutPutFileName(__subPath__, packageName, protoFileName, common.LAYER_NAME_SERVICE, "", ".interface", options)
	fileName := common.OutPutFileName(__subPath__,
		packageName,
		protoFileName,
		common.LAYER_NAME_SERVICE, "", ".interface",
		options, func(layer *common.LayerConfig) bool {
			if nil != layer {
				return layer.Replace
			}
			return true
		})

	gfile := gen.NewGeneratedFile(fileName, file.GoImportPath)
	gfile, _ = common.GenerateHeader(gfile, packageName)

	for _, svc := range file.Services {
		if err := this.generate(gfile, svc); nil != err {
			return nil, err
		}
	}
	return gfile, nil
}
