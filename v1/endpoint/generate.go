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
	"log"

	"github.com/leewckk/protoc-gen-gokit-micro/common"
	"google.golang.org/protobuf/compiler/protogen"
)

type Generator interface {
	GenerateFile(gen *protogen.Plugin, file *protogen.File, gfile *protogen.GeneratedFile, options *common.Options) (*protogen.GeneratedFile, error)
}

type Generate struct {
	generators []Generator
}

const (
	__subPath__ = "endpoint/"
)

func NewGenerator() *Generate {
	g := &Generate{
		generators: make([]Generator, 0, 0),
	}
	g.generators = append(g.generators, NewProtocol())
	g.generators = append(g.generators, NewPrototype())
	g.generators = append(g.generators, NewEndpoints())
	log.Printf("create endpoint generator, output path: %v", __subPath__)
	return g
}

func (this *Generate) GenerateFile(gen *protogen.Plugin, file *protogen.File, options *common.Options) (*protogen.GeneratedFile, error) {

	packageName := string(file.GoPackageName)
	protoFileName := common.GetFileName(file)

	fileName := common.OutPutFileName(__subPath__,
		packageName,
		protoFileName,
		common.LAYER_NAME_ENDPOINT, "", "",
		options, func(layer *common.LayerConfig) bool {
			if nil != layer {
				return layer.Replace
			}
			return true
		})

	gfile := gen.NewGeneratedFile(fileName, file.GoImportPath)
	gfile, _ = common.GenerateHeader(gfile, packageName)

	for _, generator := range this.generators {
		gfile, _ = generator.GenerateFile(gen, file, gfile, options)
	}

	return gfile, nil
}
