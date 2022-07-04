// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/protos/v1/simpleservice.proto

package simpleservicev1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListTransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blockchain string `protobuf:"bytes,1,opt,name=blockchain,proto3" json:"blockchain,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Network    string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	PageSize   *int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
	PageToken  *int32 `protobuf:"varint,5,opt,name=page_token,json=pageToken,proto3,oneof" json:"page_token,omitempty"`
}

func (x *ListTransactionsRequest) Reset() {
	*x = ListTransactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_simpleservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactionsRequest) ProtoMessage() {}

func (x *ListTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_simpleservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransactionsRequest.ProtoReflect.Descriptor instead.
func (*ListTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_simpleservice_proto_rawDescGZIP(), []int{0}
}

func (x *ListTransactionsRequest) GetBlockchain() string {
	if x != nil {
		return x.Blockchain
	}
	return ""
}

func (x *ListTransactionsRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ListTransactionsRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *ListTransactionsRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListTransactionsRequest) GetPageToken() int32 {
	if x != nil && x.PageToken != nil {
		return *x.PageToken
	}
	return 0
}

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blockchain string `protobuf:"bytes,1,opt,name=blockchain,proto3" json:"blockchain,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Network    string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_simpleservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_simpleservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_simpleservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetBalanceRequest) GetBlockchain() string {
	if x != nil {
		return x.Blockchain
	}
	return ""
}

func (x *GetBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetBalanceRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

type ListTransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions  []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	NextPageToken int32          `protobuf:"varint,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTransactionsResponse) Reset() {
	*x = ListTransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_simpleservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactionsResponse) ProtoMessage() {}

func (x *ListTransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_simpleservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransactionsResponse.ProtoReflect.Descriptor instead.
func (*ListTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_simpleservice_proto_rawDescGZIP(), []int{2}
}

func (x *ListTransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *ListTransactionsResponse) GetNextPageToken() int32 {
	if x != nil {
		return x.NextPageToken
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_simpleservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_simpleservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_simpleservice_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinalBalance float32 `protobuf:"fixed32,1,opt,name=final_balance,json=finalBalance,proto3" json:"final_balance,omitempty"`
	Token        string  `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_simpleservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_simpleservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_simpleservice_proto_rawDescGZIP(), []int{4}
}

func (x *Balance) GetFinalBalance() float32 {
	if x != nil {
		return x.FinalBalance
	}
	return 0
}

func (x *Balance) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_api_protos_v1_simpleservice_proto protoreflect.FileDescriptor

var file_api_protos_v1_simpleservice_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x67, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x22, 0x88, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x06, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x21, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x22, 0x44, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x8a, 0x02, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x3d, 0x2a, 0x7d, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x2a, 0x7d, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x3a, 0x01, 0x2a, 0x42, 0x87, 0x01, 0x0a, 0x11, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x12, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6b, 0x69,
	0x72, 0x6f, 0x2d, 0x6a, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31,
	0x92, 0x41, 0x1a, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x30, 0x2a, 0x01, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protos_v1_simpleservice_proto_rawDescOnce sync.Once
	file_api_protos_v1_simpleservice_proto_rawDescData = file_api_protos_v1_simpleservice_proto_rawDesc
)

func file_api_protos_v1_simpleservice_proto_rawDescGZIP() []byte {
	file_api_protos_v1_simpleservice_proto_rawDescOnce.Do(func() {
		file_api_protos_v1_simpleservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protos_v1_simpleservice_proto_rawDescData)
	})
	return file_api_protos_v1_simpleservice_proto_rawDescData
}

var file_api_protos_v1_simpleservice_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_protos_v1_simpleservice_proto_goTypes = []interface{}{
	(*ListTransactionsRequest)(nil),  // 0: api.protos.v1.ListTransactionsRequest
	(*GetBalanceRequest)(nil),        // 1: api.protos.v1.GetBalanceRequest
	(*ListTransactionsResponse)(nil), // 2: api.protos.v1.ListTransactionsResponse
	(*Transaction)(nil),              // 3: api.protos.v1.Transaction
	(*Balance)(nil),                  // 4: api.protos.v1.Balance
}
var file_api_protos_v1_simpleservice_proto_depIdxs = []int32{
	3, // 0: api.protos.v1.ListTransactionsResponse.transactions:type_name -> api.protos.v1.Transaction
	0, // 1: api.protos.v1.SimpleService.ListTransactions:input_type -> api.protos.v1.ListTransactionsRequest
	1, // 2: api.protos.v1.SimpleService.GetBalance:input_type -> api.protos.v1.GetBalanceRequest
	2, // 3: api.protos.v1.SimpleService.ListTransactions:output_type -> api.protos.v1.ListTransactionsResponse
	4, // 4: api.protos.v1.SimpleService.GetBalance:output_type -> api.protos.v1.Balance
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_protos_v1_simpleservice_proto_init() }
func file_api_protos_v1_simpleservice_proto_init() {
	if File_api_protos_v1_simpleservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protos_v1_simpleservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransactionsRequest); i {
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
		file_api_protos_v1_simpleservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
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
		file_api_protos_v1_simpleservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransactionsResponse); i {
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
		file_api_protos_v1_simpleservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_api_protos_v1_simpleservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
	file_api_protos_v1_simpleservice_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_protos_v1_simpleservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_protos_v1_simpleservice_proto_goTypes,
		DependencyIndexes: file_api_protos_v1_simpleservice_proto_depIdxs,
		MessageInfos:      file_api_protos_v1_simpleservice_proto_msgTypes,
	}.Build()
	File_api_protos_v1_simpleservice_proto = out.File
	file_api_protos_v1_simpleservice_proto_rawDesc = nil
	file_api_protos_v1_simpleservice_proto_goTypes = nil
	file_api_protos_v1_simpleservice_proto_depIdxs = nil
}
