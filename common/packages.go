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
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

/// STD
var (
	BytesPackage   = protogen.GoImportPath("bytes")
	ContextPackage = protogen.GoImportPath("context")
	Base64Package  = protogen.GoImportPath("encoding/base64")
	JSONPackage    = protogen.GoImportPath("encoding/json")
	FmtPackage     = protogen.GoImportPath("fmt")
	IoPackage      = protogen.GoImportPath("io")
	IoutilPackage  = protogen.GoImportPath("io/ioutil")
	MimePackage    = protogen.GoImportPath("mime")
	HttpPackage    = protogen.GoImportPath("net/http")
	StrconvPackage = protogen.GoImportPath("strconv")
	StringsPackage = protogen.GoImportPath("strings")
	ReflectPackage = protogen.GoImportPath("reflect")
)

/// FOR GO-KIT
var (
	GoKitEndpoint        = protogen.GoImportPath("github.com/go-kit/kit/endpoint")
	GoKitGRPC            = protogen.GoImportPath("github.com/go-kit/kit/transport/grpc")
	GoKitHTTP            = protogen.GoImportPath("github.com/go-kit/kit/transport/http")
	GoKitSDPackage       = protogen.GoImportPath("github.com/go-kit/kit/sd")
	GoKitSDConsulPackage = protogen.GoImportPath("github.com/go-kit/kit/sd/consul")
	GoKitSDLBPackage     = protogen.GoImportPath("github.com/go-kit/kit/sd/lb")
	ConsulAPIPackage     = protogen.GoImportPath("github.com/hashicorp/consul/api")
)

/// FOR GIN
var (
	GinPackage = protogen.GoImportPath("github.com/gin-gonic/gin")
)

/// FOR PB
var (
	ProtoPackage           = protogen.GoImportPath("google.golang.org/protobuf/proto")
	ProtojsonPackage       = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	GrpcPackage            = protogen.GoImportPath("google.golang.org/grpc")
	CodesPackage           = protogen.GoImportPath("google.golang.org/grpc/codes")
	StatusPackage          = protogen.GoImportPath("google.golang.org/grpc/status")
	AnypbPackage           = protogen.GoImportPath("google.golang.org/protobuf/types/known/anypb")
	ApipbPackage           = protogen.GoImportPath("google.golang.org/protobuf/types/known/apipb")
	DurationpbPackage      = protogen.GoImportPath("google.golang.org/protobuf/types/known/durationpb")
	EmptypbPackage         = protogen.GoImportPath("google.golang.org/protobuf/types/known/emptypb")
	FieldmaskpbPackage     = protogen.GoImportPath("google.golang.org/protobuf/types/known/fieldmaskpb")
	SourcecontextpbPackage = protogen.GoImportPath("google.golang.org/protobuf/types/known/sourcecontextpb")
	StructpbPackage        = protogen.GoImportPath("google.golang.org/protobuf/types/known/structpb")
	TimestamppbPackage     = protogen.GoImportPath("google.golang.org/protobuf/types/known/timestamppb")
	TypepbPackage          = protogen.GoImportPath("google.golang.org/protobuf/types/known/typepb")
	WrapperspbPackage      = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
)

func GenMessageName(msg *protogen.Message) protogen.GoIdent {
	switch msg.Location.SourceFile {
	case "google/protobuf/any.proto":
		return AnypbPackage.Ident(msg.GoIdent.GoName)
	case "google/protobuf/api.proto":
		return ApipbPackage.Ident(msg.GoIdent.GoName)
	case "google/protobuf/duration.proto":
		return DurationpbPackage.Ident(msg.GoIdent.GoName)
	case "google/protobuf/empty.proto":
		return EmptypbPackage.Ident(msg.GoIdent.GoName)
	case "google/protobuf/field_mask.proto":
		return FieldmaskpbPackage.Ident(msg.GoIdent.GoName)
	case "google/protobuf/source_context.proto":
		return SourcecontextpbPackage.Ident(msg.GoIdent.GoName)
	case "google/protobuf/struct.proto":
		return StatusPackage.Ident(msg.GoIdent.GoName)
	case "google/protobuf/timestamp.proto":
		return TimestamppbPackage.Ident(msg.GoIdent.GoName)
	case "google/protobuf/type.proto":
		return TypepbPackage.Ident(msg.GoIdent.GoName)
	case "google/protobuf/wrappers.proto":
		return WrapperspbPackage.Ident(msg.GoIdent.GoName)
	default:
		return msg.GoIdent
	}
}

func GetPackagePath(subModuleName, packageName string, options *Options) protogen.GoImportPath {
	return protogen.GoImportPath(options.GetModName() + "/" + subModuleName + "/" + packageName)
}

func GetPBPackagePath(file *protogen.File, option *Options) protogen.GoImportPath {
	path := option.GetModName() + "/" + strings.Replace(file.GoImportPath.String(), "\"", "", -1)
	return protogen.GoImportPath(path)
}
