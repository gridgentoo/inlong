//*
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: RPC.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponseHeader_Status int32

const (
	ResponseHeader_SUCCESS ResponseHeader_Status = 0
	ResponseHeader_ERROR   ResponseHeader_Status = 1
	ResponseHeader_FATAL   ResponseHeader_Status = 2
)

// Enum value maps for ResponseHeader_Status.
var (
	ResponseHeader_Status_name = map[int32]string{
		0: "SUCCESS",
		1: "ERROR",
		2: "FATAL",
	}
	ResponseHeader_Status_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   1,
		"FATAL":   2,
	}
)

func (x ResponseHeader_Status) Enum() *ResponseHeader_Status {
	p := new(ResponseHeader_Status)
	*p = x
	return p
}

func (x ResponseHeader_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseHeader_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_RPC_proto_enumTypes[0].Descriptor()
}

func (ResponseHeader_Status) Type() protoreflect.EnumType {
	return &file_RPC_proto_enumTypes[0]
}

func (x ResponseHeader_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ResponseHeader_Status) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ResponseHeader_Status(num)
	return nil
}

// Deprecated: Use ResponseHeader_Status.Descriptor instead.
func (ResponseHeader_Status) EnumDescriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{3, 0}
}

type RpcConnHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag     *int32 `protobuf:"varint,1,req,name=flag" json:"flag,omitempty"`
	TraceId  *int64 `protobuf:"varint,2,opt,name=traceId" json:"traceId,omitempty"`
	SpanId   *int64 `protobuf:"varint,3,opt,name=spanId" json:"spanId,omitempty"`
	ParentId *int64 `protobuf:"varint,4,opt,name=parentId" json:"parentId,omitempty"`
}

func (x *RpcConnHeader) Reset() {
	*x = RpcConnHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RPC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcConnHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcConnHeader) ProtoMessage() {}

func (x *RpcConnHeader) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcConnHeader.ProtoReflect.Descriptor instead.
func (*RpcConnHeader) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{0}
}

func (x *RpcConnHeader) GetFlag() int32 {
	if x != nil && x.Flag != nil {
		return *x.Flag
	}
	return 0
}

func (x *RpcConnHeader) GetTraceId() int64 {
	if x != nil && x.TraceId != nil {
		return *x.TraceId
	}
	return 0
}

func (x *RpcConnHeader) GetSpanId() int64 {
	if x != nil && x.SpanId != nil {
		return *x.SpanId
	}
	return 0
}

func (x *RpcConnHeader) GetParentId() int64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

type RequestHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceType *int32 `protobuf:"varint,1,opt,name=serviceType" json:"serviceType,omitempty"`
	ProtocolVer *int32 `protobuf:"varint,2,opt,name=protocolVer" json:"protocolVer,omitempty"`
}

func (x *RequestHeader) Reset() {
	*x = RequestHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RPC_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeader) ProtoMessage() {}

func (x *RequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeader.ProtoReflect.Descriptor instead.
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{1}
}

func (x *RequestHeader) GetServiceType() int32 {
	if x != nil && x.ServiceType != nil {
		return *x.ServiceType
	}
	return 0
}

func (x *RequestHeader) GetProtocolVer() int32 {
	if x != nil && x.ProtocolVer != nil {
		return *x.ProtocolVer
	}
	return 0
}

type RequestBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method  *int32 `protobuf:"varint,1,req,name=method" json:"method,omitempty"`
	Timeout *int64 `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
	Request []byte `protobuf:"bytes,3,opt,name=request" json:"request,omitempty"`
}

func (x *RequestBody) Reset() {
	*x = RequestBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RPC_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBody) ProtoMessage() {}

func (x *RequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBody.ProtoReflect.Descriptor instead.
func (*RequestBody) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{2}
}

func (x *RequestBody) GetMethod() int32 {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return 0
}

func (x *RequestBody) GetTimeout() int64 {
	if x != nil && x.Timeout != nil {
		return *x.Timeout
	}
	return 0
}

func (x *RequestBody) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

type ResponseHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      *ResponseHeader_Status `protobuf:"varint,1,req,name=status,enum=ResponseHeader_Status" json:"status,omitempty"`
	ServiceType *int32                 `protobuf:"varint,2,opt,name=serviceType" json:"serviceType,omitempty"`
	ProtocolVer *int32                 `protobuf:"varint,3,opt,name=protocolVer" json:"protocolVer,omitempty"`
}

func (x *ResponseHeader) Reset() {
	*x = ResponseHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RPC_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseHeader) ProtoMessage() {}

func (x *ResponseHeader) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseHeader.ProtoReflect.Descriptor instead.
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseHeader) GetStatus() ResponseHeader_Status {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ResponseHeader_SUCCESS
}

func (x *ResponseHeader) GetServiceType() int32 {
	if x != nil && x.ServiceType != nil {
		return *x.ServiceType
	}
	return 0
}

func (x *ResponseHeader) GetProtocolVer() int32 {
	if x != nil && x.ProtocolVer != nil {
		return *x.ProtocolVer
	}
	return 0
}

type RspExceptionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExceptionName *string `protobuf:"bytes,1,req,name=exceptionName" json:"exceptionName,omitempty"`
	StackTrace    *string `protobuf:"bytes,2,opt,name=stackTrace" json:"stackTrace,omitempty"`
}

func (x *RspExceptionBody) Reset() {
	*x = RspExceptionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RPC_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspExceptionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspExceptionBody) ProtoMessage() {}

func (x *RspExceptionBody) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspExceptionBody.ProtoReflect.Descriptor instead.
func (*RspExceptionBody) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{4}
}

func (x *RspExceptionBody) GetExceptionName() string {
	if x != nil && x.ExceptionName != nil {
		return *x.ExceptionName
	}
	return ""
}

func (x *RspExceptionBody) GetStackTrace() string {
	if x != nil && x.StackTrace != nil {
		return *x.StackTrace
	}
	return ""
}

type RspResponseBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method *int32 `protobuf:"varint,1,req,name=method" json:"method,omitempty"`
	Data   []byte `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
}

func (x *RspResponseBody) Reset() {
	*x = RspResponseBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RPC_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspResponseBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspResponseBody) ProtoMessage() {}

func (x *RspResponseBody) ProtoReflect() protoreflect.Message {
	mi := &file_RPC_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspResponseBody.ProtoReflect.Descriptor instead.
func (*RspResponseBody) Descriptor() ([]byte, []int) {
	return file_RPC_proto_rawDescGZIP(), []int{5}
}

func (x *RspResponseBody) GetMethod() int32 {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return 0
}

func (x *RspResponseBody) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_RPC_proto protoreflect.FileDescriptor

var file_RPC_proto_rawDesc = []byte{
	0x0a, 0x09, 0x52, 0x50, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0d, 0x52,
	0x70, 0x63, 0x43, 0x6f, 0x6e, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70,
	0x61, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x53,
	0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x56, 0x65, 0x72, 0x22, 0x59, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb1,
	0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x56, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c,
	0x10, 0x02, 0x22, 0x58, 0x0a, 0x10, 0x52, 0x73, 0x70, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0d, 0x65,
	0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x22, 0x3d, 0x0a, 0x0f,
	0x52, 0x73, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x51, 0x0a, 0x34, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x69, 0x6e, 0x6c, 0x6f, 0x6e, 0x67,
	0x2e, 0x74, 0x75, 0x62, 0x65, 0x6d, 0x71, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x09, 0x52, 0x50, 0x43, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01,
	0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0xa0, 0x01, 0x01,
}

var (
	file_RPC_proto_rawDescOnce sync.Once
	file_RPC_proto_rawDescData = file_RPC_proto_rawDesc
)

func file_RPC_proto_rawDescGZIP() []byte {
	file_RPC_proto_rawDescOnce.Do(func() {
		file_RPC_proto_rawDescData = protoimpl.X.CompressGZIP(file_RPC_proto_rawDescData)
	})
	return file_RPC_proto_rawDescData
}

var file_RPC_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RPC_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_RPC_proto_goTypes = []interface{}{
	(ResponseHeader_Status)(0), // 0: ResponseHeader.Status
	(*RpcConnHeader)(nil),      // 1: RpcConnHeader
	(*RequestHeader)(nil),      // 2: RequestHeader
	(*RequestBody)(nil),        // 3: RequestBody
	(*ResponseHeader)(nil),     // 4: ResponseHeader
	(*RspExceptionBody)(nil),   // 5: RspExceptionBody
	(*RspResponseBody)(nil),    // 6: RspResponseBody
}
var file_RPC_proto_depIdxs = []int32{
	0, // 0: ResponseHeader.status:type_name -> ResponseHeader.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RPC_proto_init() }
func file_RPC_proto_init() {
	if File_RPC_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RPC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcConnHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_RPC_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_RPC_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBody); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_RPC_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_RPC_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspExceptionBody); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_RPC_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspResponseBody); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RPC_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RPC_proto_goTypes,
		DependencyIndexes: file_RPC_proto_depIdxs,
		EnumInfos:         file_RPC_proto_enumTypes,
		MessageInfos:      file_RPC_proto_msgTypes,
	}.Build()
	File_RPC_proto = out.File
	file_RPC_proto_rawDesc = nil
	file_RPC_proto_goTypes = nil
	file_RPC_proto_depIdxs = nil
}
