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
	"github.com/valyala/fasttemplate"
	"google.golang.org/protobuf/compiler/protogen"
)

type Prototype struct {
	generators []Generator
}

func NewPrototype() *Prototype {
	g := &Prototype{}
	g.generators = make([]Generator, 0, 0)
	return g
}

func PrototypeProcName(svc *protogen.Service, method *protogen.Method) string {
	return "__" + svc.GoName + method.GoName + "Proc__"
}

func (this *Prototype) generate(gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) error {

	template := `type {{prototypeName}} func(context.Context, *{{input}})(*{{output}}, error)
                 var {{prototypeProc}} {{prototypeName}}
                 func Implement{{prototypeName}} (impl func() {{prototypeName}}) {
                     __{{prototypeName}}Proc__ = impl()
                 }
    `
	serviceName := svc.GoName
	gfile.P()
	for _, method := range svc.Methods {
		// prototypeName := serviceName + method.GoName
		gfile.P("//// PROTOTYPE OF SERVICE: ", svc.GoName, " , METHOD: ", method.GoName)
		tmpl := fasttemplate.New(template, "{{", "}}")
		gfile.P(tmpl.ExecuteString(map[string]interface{}{
			"prototypeName": serviceName + method.GoName,
			"prototypeProc": PrototypeProcName(svc, method),
			"input":         method.Input.GoIdent.GoName,
			"output":        method.Output.GoIdent.GoName,
		}))
		// gfile.P(fmt.Sprintf("type %v func(context.Context, *%v)(*%v , error)", prototypeName, common.GenMessageName(method.Input), common.GenMessageName(method.Output)))
		gfile.P()

	}
	gfile.P()

	return nil
}

func (this *Prototype) GenerateFile(gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, options *common.Options) (*protogen.GeneratedFile, error) {

	for _, svc := range file.Services {
		if err := this.generate(gfile, svc, options); nil != err {
			return nil, err
		}
	}
	return gfile, nil
}
