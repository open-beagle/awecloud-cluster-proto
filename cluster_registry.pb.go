// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: cluster_registry.proto

package awecloud_cluster_proto

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

// LoginRequest 集群信息
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID, 客户端集群ID，aliyun
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Secret, 客户端集群的密钥, Bp#q3sM6uxK4
	Secret string `protobuf:"bytes,2,opt,name=Secret,proto3" json:"Secret,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_cluster_registry_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *LoginRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

// LoginResponse 集群信息
type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path, 集群的URL的访问路径, /awecloud/cluster/access
	Path string `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	// Secret, 集群的密钥, Bp#q3sM6uxK4
	Secret string `protobuf:"bytes,2,opt,name=Secret,proto3" json:"Secret,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_cluster_registry_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *LoginResponse) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

// ListenRequest ， 服务信息
type ListenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID , 客户端集群ID ， aliyun
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ListenRequest) Reset() {
	*x = ListenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenRequest) ProtoMessage() {}

func (x *ListenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenRequest.ProtoReflect.Descriptor instead.
func (*ListenRequest) Descriptor() ([]byte, []int) {
	return file_cluster_registry_proto_rawDescGZIP(), []int{2}
}

func (x *ListenRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// ListenResponse ， 服务信息
type ListenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type , 服务类型 ， 如service
	Type string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	// Service , 服务 ， 如k8s
	Service string `protobuf:"bytes,2,opt,name=Service,proto3" json:"Service,omitempty"`
	// Action , 服务动作 ， 如create,update,delete
	Action string `protobuf:"bytes,3,opt,name=Action,proto3" json:"Action,omitempty"`
	// Data , 服务数据 ， 如Json
	Data string `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ListenResponse) Reset() {
	*x = ListenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenResponse) ProtoMessage() {}

func (x *ListenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenResponse.ProtoReflect.Descriptor instead.
func (*ListenResponse) Descriptor() ([]byte, []int) {
	return file_cluster_registry_proto_rawDescGZIP(), []int{3}
}

func (x *ListenResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ListenResponse) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ListenResponse) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ListenResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// PublishServiceRequest 发布服务信息请求
type PublishServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID, 集群的ID，aliyun
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Service, 注册的服务信息
	Service *RegisterService `protobuf:"bytes,2,opt,name=Service,proto3" json:"Service,omitempty"`
}

func (x *PublishServiceRequest) Reset() {
	*x = PublishServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishServiceRequest) ProtoMessage() {}

func (x *PublishServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishServiceRequest.ProtoReflect.Descriptor instead.
func (*PublishServiceRequest) Descriptor() ([]byte, []int) {
	return file_cluster_registry_proto_rawDescGZIP(), []int{4}
}

func (x *PublishServiceRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *PublishServiceRequest) GetService() *RegisterService {
	if x != nil {
		return x.Service
	}
	return nil
}

// RegisterService ， 服务信息
type RegisterService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name , 服务名称 ， 如k8s
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// Type , 服务类型 ， 如tcp
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	// Host , 服务主机 ， 如kubernetes.default
	Host string `protobuf:"bytes,3,opt,name=Host,proto3" json:"Host,omitempty"`
	// Port , 服务端口 ， 如6443
	Port int32 `protobuf:"varint,4,opt,name=Port,proto3" json:"Port,omitempty"`
	// Remote , 服务映射端口 ， 如7000
	Remote int32 `protobuf:"varint,5,opt,name=Remote,proto3" json:"Remote,omitempty"`
	// ApproveState, 集群状态 1 代接入  2 待同意  3 已同意 4 已拒绝
	ApproveState int32 `protobuf:"varint,6,opt,name=ApproveState,proto3" json:"ApproveState,omitempty"`
}

func (x *RegisterService) Reset() {
	*x = RegisterService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_registry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterService) ProtoMessage() {}

func (x *RegisterService) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_registry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterService.ProtoReflect.Descriptor instead.
func (*RegisterService) Descriptor() ([]byte, []int) {
	return file_cluster_registry_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterService) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RegisterService) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RegisterService) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RegisterService) GetRemote() int32 {
	if x != nil {
		return x.Remote
	}
	return 0
}

func (x *RegisterService) GetApproveState() int32 {
	if x != nil {
		return x.ApproveState
	}
	return 0
}

var File_cluster_registry_proto protoreflect.FileDescriptor

var file_cluster_registry_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x77, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x22, 0x36, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x3b, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x6a, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x5c, 0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x33, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x61, 0x77, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x32, 0xde, 0x01, 0x0a, 0x0f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x16, 0x2e, 0x61, 0x77, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x77, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x61,
	0x77, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x77, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x77, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x77, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x62, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x77,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cluster_registry_proto_rawDescOnce sync.Once
	file_cluster_registry_proto_rawDescData = file_cluster_registry_proto_rawDesc
)

func file_cluster_registry_proto_rawDescGZIP() []byte {
	file_cluster_registry_proto_rawDescOnce.Do(func() {
		file_cluster_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_cluster_registry_proto_rawDescData)
	})
	return file_cluster_registry_proto_rawDescData
}

var file_cluster_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cluster_registry_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),          // 0: awecloud.LoginRequest
	(*LoginResponse)(nil),         // 1: awecloud.LoginResponse
	(*ListenRequest)(nil),         // 2: awecloud.ListenRequest
	(*ListenResponse)(nil),        // 3: awecloud.ListenResponse
	(*PublishServiceRequest)(nil), // 4: awecloud.PublishServiceRequest
	(*RegisterService)(nil),       // 5: awecloud.RegisterService
}
var file_cluster_registry_proto_depIdxs = []int32{
	5, // 0: awecloud.PublishServiceRequest.Service:type_name -> awecloud.RegisterService
	0, // 1: awecloud.ClusterRegistry.Login:input_type -> awecloud.LoginRequest
	2, // 2: awecloud.ClusterRegistry.Listen:input_type -> awecloud.ListenRequest
	4, // 3: awecloud.ClusterRegistry.PublishService:input_type -> awecloud.PublishServiceRequest
	1, // 4: awecloud.ClusterRegistry.Login:output_type -> awecloud.LoginResponse
	3, // 5: awecloud.ClusterRegistry.Listen:output_type -> awecloud.ListenResponse
	5, // 6: awecloud.ClusterRegistry.PublishService:output_type -> awecloud.RegisterService
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cluster_registry_proto_init() }
func file_cluster_registry_proto_init() {
	if File_cluster_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cluster_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_cluster_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_cluster_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenRequest); i {
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
		file_cluster_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenResponse); i {
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
		file_cluster_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishServiceRequest); i {
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
		file_cluster_registry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterService); i {
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
			RawDescriptor: file_cluster_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cluster_registry_proto_goTypes,
		DependencyIndexes: file_cluster_registry_proto_depIdxs,
		MessageInfos:      file_cluster_registry_proto_msgTypes,
	}.Build()
	File_cluster_registry_proto = out.File
	file_cluster_registry_proto_rawDesc = nil
	file_cluster_registry_proto_goTypes = nil
	file_cluster_registry_proto_depIdxs = nil
}
