/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalNodeInfo writes a value of the 'node_info' type to the given writer.
func MarshalNodeInfo(object *NodeInfo, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeNodeInfo(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// writeNodeInfo writes a value of the 'node_info' type to the given stream.
func writeNodeInfo(object *NodeInfo, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = object.bitmap_&1 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("amount")
		stream.WriteInt(object.amount)
		count++
	}
	present_ = object.bitmap_&2 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("type")
		stream.WriteString(string(object.type_))
	}
	stream.WriteObjectEnd()
}

// UnmarshalNodeInfo reads a value of the 'node_info' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalNodeInfo(source interface{}) (object *NodeInfo, err error) {
	if source == http.NoBody {
		return
	}
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readNodeInfo(iterator)
	err = iterator.Error
	return
}

// readNodeInfo reads a value of the 'node_info' type from the given iterator.
func readNodeInfo(iterator *jsoniter.Iterator) *NodeInfo {
	object := &NodeInfo{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "amount":
			value := iterator.ReadInt()
			object.amount = value
			object.bitmap_ |= 1
		case "type":
			text := iterator.ReadString()
			value := NodeType(text)
			object.type_ = value
			object.bitmap_ |= 2
		default:
			iterator.ReadAny()
		}
	}
	return object
}
