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
	"net/http"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
)

type Options struct {
	ModName    *string `json:"modName"`
	TypeName   *string `json:"typeName"`
	EnableGrpc bool    `json:"enableGrpc"`
	Cfg        *Configure
}

const (
	GENERATE_TYPE_NAME_CLIENT = "client"
	GENERATE_TYPE_NAME_SERVER = "server"
)

func (this *Options) GetModName() string {
	if nil != this.ModName {
		return *this.ModName
	}
	return this.Cfg.GetGen().Project
}

func (this *Options) LayerConfig(layer string) *LayerConfig {
	return this.Cfg.GetLayer(layer)
}

func ParseHttpOption(httpRule *annotations.HttpRule) (httpMethod string, pattern string, err error) {

	err = nil
	switch httpRule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		httpMethod = http.MethodGet
		pattern = httpRule.GetGet()
	case *annotations.HttpRule_Put:
		httpMethod = http.MethodPut
		pattern = httpRule.GetPut()
	case *annotations.HttpRule_Post:
		httpMethod = http.MethodPost
		pattern = httpRule.GetPost()
	case *annotations.HttpRule_Delete:
		httpMethod = http.MethodDelete
		pattern = httpRule.GetDelete()
	case *annotations.HttpRule_Patch:
		httpMethod = http.MethodPatch
		pattern = httpRule.GetPatch()
	default:
		err = fmt.Errorf("Error decode http rule")
	}
	pattern = strings.Replace(pattern, "}", "", -1)
	pattern = strings.Replace(pattern, "{", ":", -1)
	return
}
