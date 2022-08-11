package grpc

import (
	"github.com/leewckk/protoc-gen-gokit-micro/common"
	"github.com/valyala/fasttemplate"
	"google.golang.org/protobuf/compiler/protogen"
)

type Register struct {
	generators []Generator
}

func NewRegister() *Register {
	p := &Register{
		generators: make([]Generator, 0, 0),
	}
	return p
}

func GetRegisterTransportFuncName(serviceName string) string {
	funcName := "Register" + serviceName + "Proc"
	return funcName
}

func GetRegisterGrpcFuncName(serviceName string) string {
	funcName := "Register" + serviceName + "Server"
	return funcName
}

func (this *Register) GenerateFile(fileUri string, gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, svc *protogen.Service, options *common.Options) (*protogen.GeneratedFile, error) {
	funcName := GetRegisterTransportFuncName(svc.GoName)
	regName := GetRegisterGrpcFuncName(svc.GoName)
	pb := common.GetPBPackagePath(file, options)

	gfile.P()
	gfile.P("func ", funcName, "(",
		"s *", common.GrpcPackage.Ident("Server"), ", opts...", common.GoKitGRPC.Ident("ServerOption"), "){")

	template := `(s, {{creater}}(opts...))`
	tmpl := fasttemplate.New(template, "{{", "}}")
	gfile.P(pb.Ident(regName),
		tmpl.ExecuteString(map[string]interface{}{
			"creater": GetCreateTransportFuncName(svc.GoName),
		}))

	gfile.P("}")
	gfile.P()
	return gfile, nil

}
