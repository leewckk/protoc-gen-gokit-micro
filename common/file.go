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

package common

import (
	"fmt"
	"os"
	"path"

	"google.golang.org/protobuf/compiler/protogen"
)

func FileExist(uri string) bool {
	_, err := os.Stat(uri)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func GetFileName(file *protogen.File) string {
	protoName := file.Proto.GetName()
	fileNameFull := path.Base(protoName)
	fileSuffix := path.Ext(protoName)
	protoFileName := fileNameFull[0 : len(fileNameFull)-len(fileSuffix)]
	return protoFileName
}

type IsReplace func(layer *LayerConfig) bool

func OutPutFileName(subPath string, packageName string, protoFileName string, layerName string, prefix, postfix string, options *Options, isReplace IsReplace) string {

	uri := subPath + packageName + fmt.Sprintf("/%v%v%v.go", prefix, protoFileName, postfix)
	layer := options.LayerConfig(layerName)

	replace := true
	if nil != isReplace {
		replace = isReplace(layer)
	}

	if replace {
		return uri
	}

	if FileExist(uri) {
		uri = subPath + packageName + fmt.Sprintf("/1/%v%v%v.go", prefix, protoFileName, postfix)
	}
	return uri
}

func findMessage(file *protogen.File, name string) *protogen.Message {

	for _, msg := range file.Messages {
		if string(msg.Desc.Name()) == name {
			return msg
		}
	}
	for _, ext := range file.Extensions {
		if string(ext.Desc.Name()) == name {
			return ext.Message
		}
	}
	return nil
}

func FindMessage(plugin *protogen.Plugin, file *protogen.File, name string) *protogen.Message {

	message := findMessage(file, name)
	if nil != message {
		return message
	}

	//// 从全局搜索
	for _, f := range plugin.Files {
		for i := 0; i < file.Desc.Imports().Len(); i++ {
			importName := string(file.Desc.Imports().Get(i).Path())
			if importName == f.Proto.GetName() {
				message := findMessage(f, name)
				if nil != message {
					return message
				}
			}
		}
	}
	return nil
}
