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
	"fmt"

	"github.com/leewckk/protoc-gen-gokit-micro/common"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Codes struct {
	plugin     *protogen.Plugin
	generators []Generator
}

func NewCodes(plugin *protogen.Plugin) *Codes {
	p := &Codes{
		plugin:     plugin,
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

func (this *Codes) userStartBlock(serviceName string, methodName string) string {
	return `// @@USER.` + serviceName + "." + methodName
}

func (this *Codes) userEndBlock(serviceName string, methodName string) string {
	return `// $$USER.` + serviceName + "." + methodName
}

func (this *Codes) assigin(file *protogen.File, gfile *protogen.GeneratedFile, prefixSrc, prefixDest string, destFieldImport protogen.GoImportPath, field *protogen.Field, options *common.Options, deep int) error {

	varname := fmt.Sprintf("s%v", deep)
	if field.Desc.Kind() == protoreflect.MessageKind {
		if field.Desc.IsList() {
			gfile.P(`for _, sub := range `, prefixSrc+field.GoName, `{`)
			gfile.P(`var `, varname, " ", destFieldImport.Ident(field.Message.GoIdent.GoName))
			for _, sub := range field.Message.Fields {
				this.assigin(file, gfile, "sub.", varname+".", destFieldImport, sub, options, deep+1)
			}

			gfile.P(prefixDest, field.GoName, ` = append(`, prefixDest, field.GoName, ` , &`, varname, ` )`)
			gfile.P(`}`)
		} else if field.Desc.IsMap() {

			if common.IsBaseType(field.Desc.MapValue()) {
				gfile.P(prefixDest+field.GoName, ` = make(map[`, common.FieldMapKeyTypeName(field), `]`, (common.FieldMapValTypeNamePure(field)), `)`)
			} else {
				gfile.P(prefixDest+field.GoName, ` = make(map[`, common.FieldMapKeyTypeName(field), `]*`, destFieldImport.Ident(common.FieldMapValTypeNamePure(field)), `)`)
			}
			gfile.P(`for key, sub := range `, prefixSrc+field.GoName, `{`)

			if common.IsBaseType(field.Desc.MapValue()) {
				gfile.P(prefixDest, field.GoName, `[key] = sub`)
			} else {
				gfile.P(`var `, varname, " ", destFieldImport.Ident(common.FieldMapValTypeNamePure(field)))
				msg := common.FindMessage(this.plugin, file, common.FieldMapValTypeNamePure(field))
				if nil != msg {
					for _, sub := range msg.Fields {
						this.assigin(file, gfile, "sub.", varname+".", destFieldImport, sub, options, deep+1)
					}

					gfile.P(prefixDest, field.GoName, `[key] = &`, varname)
				} else {
					panic(fmt.Sprintf("failed find message : %v ", common.FieldMapValTypeNamePure(field)))
				}
			}
			gfile.P(`}`)

		} else {
			gfile.P(`if nil != `, prefixSrc+field.GoName, `{`)
			gfile.P(`var `, varname+" ", destFieldImport.Ident(field.Message.GoIdent.GoName))
			gfile.P(prefixDest, field.GoName, ` = &`, varname)
			for _, sub := range field.Message.Fields {
				this.assigin(file, gfile, prefixSrc+field.GoName+".", prefixDest+field.GoName+".", destFieldImport, sub, options, deep+1)
			}
			gfile.P(`}`)
		}
	} else {
		gfile.P(prefixDest, field.GoName, " = "+prefixSrc, field.GoName)
	}
	return nil
}

func (this *Codes) generateEncodeRequest(fileUri string, file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) error {

	gfile.P()
	for _, method := range svc.Methods {
		// funcName := "Encode" + svc.GoName + method.GoName + "Request"
		funcName := GetEncodeRequestName(svc.GoName, method.GoName)
		gfile.P("//// THIS IS FOR CLIENT INVKE SERVICE REQUEST ENCODE")
		gfile.P("//// SERVICE.REQUEST - > PB.REQUEST")

		gfile.P(method.Comments.Leading,
			"func ",
			funcName,
			"(ctx ", common.ContextPackage.Ident("Context"), " , request interface{} )",
			"(interface{} , error){",
		)

		serviceName := svc.GoName
		methodName := funcName
		// gfile.P(this.userStartBlock(serviceName, methodName))

		userblock := common.UserBlock(fileUri, serviceName, methodName,
			this.userStartBlock(serviceName, methodName), this.userEndBlock(serviceName, methodName))
		if "" == userblock {
			if len(method.Input.Fields) == 0 {
				gfile.P("return &", common.GetPBPackagePath(file, options).Ident(method.Input.GoIdent.GoName), "{}, nil")
			} else {

				gfile.P("if r, ok := request.(*", common.GetPackagePath("invoker/protocol", string(file.GoPackageName), options).Ident(method.Input.GoIdent.GoName), "); ok {")
				gfile.P("req := ", common.GetPBPackagePath(file, options).Ident(method.Input.GoIdent.GoName), "{}")

				/// TODO
				gfile.P("/// TODO")
				for _, field := range method.Input.Fields {
					this.assigin(file, gfile, "r.", "req.", common.GetPBPackagePath(file, options), field, options, 0)
				}
				gfile.P("return &req, nil")
				gfile.P("}")

				gfile.P("return nil, ", common.FmtPackage.Ident("Errorf"), `("Error convert interface : %v -> *%v", `,
					common.ReflectPackage.Ident("TypeOf"), "(request), \"",
					common.GetPackagePath("invoker/protocol", string(file.GoPackageName), options).Ident(method.Input.GoIdent.GoName), "\")")
			}

		} else {
			gfile.P(userblock)
		}
		// gfile.P(this.userEndBlock(serviceName, methodName))

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
		gfile.P("//// PB.RESPONSE -> SERVICE.RESPONSE")

		gfile.P(method.Comments.Leading,
			"func ",
			funcName,
			"(ctx ", common.ContextPackage.Ident("Context"), " , resp interface{} )",
			"(interface{} , error){",
		)

		serviceName := svc.GoName
		methodName := funcName
		// gfile.P(this.userStartBlock(serviceName, methodName))

		userblock := common.UserBlock(fileUri, serviceName, methodName,
			this.userStartBlock(serviceName, methodName), this.userEndBlock(serviceName, methodName))
		if "" == userblock {
			if len(method.Output.Fields) == 0 {
				gfile.P("return &", common.GetPackagePath("invoker/protocol", string(file.GoPackageName), options).Ident(method.Output.GoIdent.GoName), "{}, nil")
			} else {

				destTypeImport := common.GetPBPackagePath(file, options)
				gfile.P("if response, ok := resp.(*", destTypeImport.Ident(method.Output.GoIdent.GoName), "); ok {")

				gfile.P("r :=", common.GetPackagePath("invoker/protocol", string(file.GoPackageName), options).Ident(method.Output.GoIdent.GoName), "{}")
				gfile.P("/// TODO")
				for _, field := range method.Output.Fields {
					// gfile.P("r.", field.GoName, " = response.", field.GoName)
					this.assigin(file, gfile, "response.", "r.", common.GetPackagePath("invoker/protocol", string(file.GoPackageName), options), field, options, 0)
				}
				gfile.P("return &r, nil")
				gfile.P("}")
				gfile.P("return nil, ", common.FmtPackage.Ident("Errorf"), `("Error convert interface : %v -> *%v", `, common.ReflectPackage.Ident("TypeOf"), "(resp) , \"",
					destTypeImport.Ident(method.Output.GoIdent.GoName), "\")")
			}

		} else {
			gfile.P(userblock)
		}
		// gfile.P(this.userEndBlock(serviceName, methodName))
		gfile.P("}")
	}
	gfile.P("//// ")
	return nil
}

// func (this *Codes) GenerateFile(fileUri string, gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) (*protogen.GeneratedFile, error) {

// 	this.generateEncodeRequest(fileUri, file, gfile, svc, options)
// 	this.generateDecodeResponse(fileUri, file, gfile, svc, options)
// 	return gfile, nil
// }

func (this *Codes) GenerateFile(fileUri string, gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, options *common.Options) (*protogen.GeneratedFile, error) {

	for _, svc := range file.Services {
		this.generateEncodeRequest(fileUri, file, gfile, svc, options)
		this.generateDecodeResponse(fileUri, file, gfile, svc, options)
	}
	return gfile, nil
}
