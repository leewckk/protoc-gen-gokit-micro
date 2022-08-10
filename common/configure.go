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
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	DEFAULT_CONFIG_FILE  = "gokit-gen.yaml"
	LAYER_NAME_TRANSPORT = "transport"
	LAYER_NAME_ENDPOINT  = "endpoint"
	LAYER_NAME_CLIENT    = "client"
	LAYER_NAME_SERVICE   = "service"
)

type LayerSubConfig struct {
	Enable  bool `json:"enable" yaml:"enable"`
	Replace bool `json:"replace" yaml:"replace"`
}

type LayerConfig struct {
	Name    string         `json:"name" yaml:"name"`
	Enable  bool           `json:"enable" yaml:"enable"`
	Replace bool           `json:"replace" yaml:"replace"`
	Grpc    LayerSubConfig `json:"grpc" yaml:"grpc"`
	Http    LayerSubConfig `json:"http" yaml:"http"`
}

type GenConfig struct {
	Version string        `json:"version" yaml:"version"`
	Project string        `json:"project" yaml:"project"`
	Layers  []LayerConfig `json:"layers" yaml:"layers"`
}

type Configure struct {
	Uri string
	cfg *viper.Viper
	gen GenConfig
}

func NewConfigure(uri string) *Configure {
	c := Configure{
		Uri: uri,
		cfg: viper.New(),
	}
	c.Load()
	return &c
}

func DefaultConfig() *Configure {
	return NewConfigure(DEFAULT_CONFIG_FILE)
}

func (c *Configure) Load() error {

	c.cfg.SetConfigFile(c.Uri)

	c.cfg.SetDefault("version", "v1")
	c.cfg.SetDefault("project", "unknown")

	if FileExist(c.Uri) != true {
		return nil
	}

	err := c.cfg.ReadInConfig()
	if nil != err {
		panic(err)
	}

	err = c.cfg.Unmarshal(&c.gen)
	if nil != err {
		panic(err)
	}
	return nil
}

func (c *Configure) GetGen() *GenConfig {
	return &c.gen
}

func (c *Configure) Get() *viper.Viper {
	return c.cfg
}

func (c *Configure) GetLayer(layer string) *LayerConfig {
	for _, l := range c.gen.Layers {
		if l.Name == layer {
			return &l
		}
	}
	return nil
}

func (c *Configure) DumpJson() string {
	js, _ := json.MarshalIndent(c.gen, "", "  ")
	return string(js)
}
