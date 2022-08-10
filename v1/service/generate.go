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
	"log"

	"github.com/leewckk/protoc-gen-gokit-micro/common"

	"google.golang.org/protobuf/compiler/protogen"
)

type Generate struct {
	generators []common.Generator
}

const (
	__subPath__ = "service/"
)

func NewGenerator() *Generate {
	g := &Generate{
		generators: make([]common.Generator, 0, 0),
	}
	g.generators = append(g.generators, NewImplement())
	g.generators = append(g.generators, NewInterface())
	g.generators = append(g.generators, NewProtocol())
	log.Printf("create service generator, output path: %v", __subPath__)
	return g
}

func (this *Generate) GenerateFile(gen *protogen.Plugin, file *protogen.File, options *common.Options) (*protogen.GeneratedFile, error) {
	for _, generator := range this.generators {
		generator.GenerateFile(gen, file, options)
	}

	return nil, nil
}
