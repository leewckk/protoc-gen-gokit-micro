package endpoint

import (
	"strings"

	"github.com/leewckk/protoc-gen-gokit-micro/common"
	"google.golang.org/protobuf/compiler/protogen"
)

type Protocol struct {
	generators []common.Generator
}

func NewProtocol() *Protocol {
	g := &Protocol{}
	g.generators = make([]common.Generator, 0, 0)
	return g
}

func (this *Protocol) generateFileds(gfile *protogen.GeneratedFile, msg *protogen.Message) error {

	for _, field := range msg.Fields {
		gfile.P(
			field.Comments.Leading.String(),
			field.GoName, " ",
			common.GenFieldType(field),
			"`",
			common.GenerateTags(field),
			"`",
			strings.Replace(field.Comments.Trailing.String(), "\n", "", -1),
		)
	}
	return nil
}

func (this *Protocol) generate(gfile *protogen.GeneratedFile, msg *protogen.Message) error {

	gfile.P()
	gfile.P("/// PROTOCOL DEFINITION")
	gfile.P(msg.Comments.Leading.String(), "type ", msg.GoIdent, " struct {")
	/// 生成消息字段
	this.generateFileds(gfile, msg)
	gfile.P("}", msg.Comments.Trailing.String())
	gfile.P()
	return nil
}

func (this *Protocol) GenerateFile(gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, options *common.Options) (*protogen.GeneratedFile, error) {

	for _, msg := range file.Messages {
		if err := this.generate(gfile, msg); nil != err {
			return nil, err
		}
	}
	return gfile, nil
}
