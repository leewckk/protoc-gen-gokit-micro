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
package v1

import (
	"fmt"

	"github.com/leewckk/protoc-gen-gokit-micro/common"

	"github.com/leewckk/protoc-gen-gokit-micro/v1/client/gin"
	"github.com/leewckk/protoc-gen-gokit-micro/v1/client/grpc"
	"github.com/leewckk/protoc-gen-gokit-micro/v1/endpoint"
	"github.com/leewckk/protoc-gen-gokit-micro/v1/service"
	"google.golang.org/protobuf/compiler/protogen"

	clientgin "github.com/leewckk/protoc-gen-gokit-micro/v1/client/gin"
	clientgrpc "github.com/leewckk/protoc-gen-gokit-micro/v1/client/grpc"
	clientproto "github.com/leewckk/protoc-gen-gokit-micro/v1/client/protocol"
)

type GeneratorV1 struct {
	generators []common.Generator
}

func NewGenerator(plugin *protogen.Plugin, options common.Options) *GeneratorV1 {

	g := &GeneratorV1{
		generators: make([]common.Generator, 0, 0),
	}

	if *options.TypeName == common.GENERATE_TYPE_NAME_CLIENT {
		g.generators = append(g.generators, clientgrpc.NewGenerator(plugin))
		g.generators = append(g.generators, clientgin.NewGenerator())
		g.generators = append(g.generators, clientproto.NewGenerator())
	} else if *options.TypeName == common.GENERATE_TYPE_NAME_SERVER {
		g.generators = append(g.generators, service.NewGenerator())
		g.generators = append(g.generators, endpoint.NewGenerator())
		g.generators = append(g.generators, grpc.NewGenerator(plugin))
		g.generators = append(g.generators, gin.NewGenerator())
	} else {
		panic(fmt.Sprintf("unsuported generate type: %v", options.TypeName))
	}
	return g
}

func (this *GeneratorV1) GenerateFile(gen *protogen.Plugin, file *protogen.File, options *common.Options) (*protogen.GeneratedFile, error) {

	for _, generator := range this.generators {
		_, err := generator.GenerateFile(gen, file, options)
		if nil != err {
			panic(err)
		}
	}
	return nil, nil
}
