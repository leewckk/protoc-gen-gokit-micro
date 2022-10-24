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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leewckk/protoc-gen-gokit-micro/common"
	"google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
)

func init() {

	version := flag.Bool("version", false, "display version")
	flag.Parse()

	if *version == true {
		vers := FullVersion()
		for _, v := range vers {
			fmt.Println(v)

		}
		os.Exit(0)

	}
}

func main() {

	var fs flag.FlagSet
	typeName := fs.String("gentype", "server", "server / client mode")
	modName := fs.String("mod", "unknown", "project mod name")
	enableGrpc := fs.Bool("enable_grpc", false, "enable grpc supoort")

	opts := &protogen.Options{
		ParamFunc: fs.Set,
	}
	opts.Run(func(p *protogen.Plugin) error {

		options := common.Options{
			ModName:    modName,
			TypeName:   typeName,
			EnableGrpc: *enableGrpc,
			Cfg:        common.DefaultConfig(),
		}
		generator := NewGenerator(p, options)
		for _, f := range p.Files {
			if f.Generate {
				generator.GenerateFile(p, f, &options)
			}
		}
		p.SupportedFeatures = internal_gengo.SupportedFeatures
		return nil
	})
}
