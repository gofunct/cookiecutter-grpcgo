package main

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

var _ descriptor.Message
var _ gogoproto.EnableFunc
var _ proto.Buffer