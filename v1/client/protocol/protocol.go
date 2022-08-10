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

package protocol

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
