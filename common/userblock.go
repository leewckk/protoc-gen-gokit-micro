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
	"bufio"
	"io"
	"os"
	"strings"
)

func UserBlock(fileUri, serviceName string, methodName string, start, end string) string {

	if FileExist(fileUri) != true {
		return ""
	}

	/// FIND USER BOLCK
	fi, err := os.Open(fileUri)
	if nil != err {
		panic(err)
	}

	matched := make([]string, 0, 0)

	/// scan line by line
	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if len(matched) == 0 {
			/// FIND startBolck
			if strings.Contains(string(line), start) {
				matched = append(matched, "")
			}
		} else {
			//// FIND endBlock
			if strings.Contains(string(line), end) {
				break
			}
			matched = append(matched, string(line)+"\n")
		}
	}

	userblock := ""
	for idx, line := range matched {
		if idx == len(matched)-1 {
			userblock += strings.Replace(line, "\n", "", -1)
		} else {
			userblock += (line)
		}
	}
	return userblock
}
