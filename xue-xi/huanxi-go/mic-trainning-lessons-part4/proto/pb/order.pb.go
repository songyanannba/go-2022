// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: order.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderNo string `protobuf:"bytes,2,opt,name=OrderNo,proto3" json:"OrderNo,omitempty"`
	Status  string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderStatus) Reset() {
	*x = OrderStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatus) ProtoMessage() {}

func (x *OrderStatus) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatus.ProtoReflect.Descriptor instead.
func (*OrderStatus) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderStatus) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderStatus) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *OrderStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type OrderItemDetailRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order       *OrderItemRes      `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	ProductList []*OrderProductRes `protobuf:"bytes,2,rep,name=productList,proto3" json:"productList,omitempty"`
}

func (x *OrderItemDetailRes) Reset() {
	*x = OrderItemDetailRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemDetailRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemDetailRes) ProtoMessage() {}

func (x *OrderItemDetailRes) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemDetailRes.ProtoReflect.Descriptor instead.
func (*OrderItemDetailRes) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItemDetailRes) GetOrder() *OrderItemRes {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *OrderItemDetailRes) GetProductList() []*OrderProductRes {
	if x != nil {
		return x.ProductList
	}
	return nil
}

type OrderProductRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId     int32   `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	ProductId   int32   `protobuf:"varint,3,opt,name=productId,proto3" json:"productId,omitempty"`
	Num         int32   `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`
	ProductName string  `protobuf:"bytes,5,opt,name=productName,proto3" json:"productName,omitempty"`
	RealPrice   float32 `protobuf:"fixed32,6,opt,name=realPrice,proto3" json:"realPrice,omitempty"`
	CoverImage  string  `protobuf:"bytes,7,opt,name=coverImage,proto3" json:"coverImage,omitempty"`
}

func (x *OrderProductRes) Reset() {
	*x = OrderProductRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProductRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProductRes) ProtoMessage() {}

func (x *OrderProductRes) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProductRes.ProtoReflect.Descriptor instead.
func (*OrderProductRes) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderProductRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderProductRes) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderProductRes) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OrderProductRes) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *OrderProductRes) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *OrderProductRes) GetRealPrice() float32 {
	if x != nil {
		return x.RealPrice
	}
	return 0
}

func (x *OrderProductRes) GetCoverImage() string {
	if x != nil {
		return x.CoverImage
	}
	return ""
}

type OrderPagingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId int32 `protobuf:"varint,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	PagNo     int32 `protobuf:"varint,2,opt,name=pagNo,proto3" json:"pagNo,omitempty"`
	PageSize  int32 `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *OrderPagingReq) Reset() {
	*x = OrderPagingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPagingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPagingReq) ProtoMessage() {}

func (x *OrderPagingReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPagingReq.ProtoReflect.Descriptor instead.
func (*OrderPagingReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderPagingReq) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *OrderPagingReq) GetPagNo() int32 {
	if x != nil {
		return x.PagNo
	}
	return 0
}

func (x *OrderPagingReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type OrderListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int32           `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	ItemList []*OrderItemRes `protobuf:"bytes,2,rep,name=itemList,proto3" json:"itemList,omitempty"`
}

func (x *OrderListRes) Reset() {
	*x = OrderListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListRes) ProtoMessage() {}

func (x *OrderListRes) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListRes.ProtoReflect.Descriptor instead.
func (*OrderListRes) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderListRes) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OrderListRes) GetItemList() []*OrderItemRes {
	if x != nil {
		return x.ItemList
	}
	return nil
}

type OrderItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId int32  `protobuf:"varint,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Addr      string `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	PostCode  string `protobuf:"bytes,4,opt,name=postCode,proto3" json:"postCode,omitempty"`
	Receiver  string `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Mobile    string `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *OrderItemReq) Reset() {
	*x = OrderItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemReq) ProtoMessage() {}

func (x *OrderItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemReq.ProtoReflect.Descriptor instead.
func (*OrderItemReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderItemReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderItemReq) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *OrderItemReq) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *OrderItemReq) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *OrderItemReq) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *OrderItemReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type OrderItemRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId  int32  `protobuf:"varint,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
	PayType    string `protobuf:"bytes,3,opt,name=payType,proto3" json:"payType,omitempty"`
	OrderNo    string `protobuf:"bytes,4,opt,name=orderNo,proto3" json:"orderNo,omitempty"`
	PostCode   string `protobuf:"bytes,5,opt,name=postCode,proto3" json:"postCode,omitempty"`
	Amount     string `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Addr       string `protobuf:"bytes,7,opt,name=addr,proto3" json:"addr,omitempty"`
	Receiver   string `protobuf:"bytes,8,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Mobile     string `protobuf:"bytes,9,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Status     string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime string `protobuf:"bytes,11,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *OrderItemRes) Reset() {
	*x = OrderItemRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemRes) ProtoMessage() {}

func (x *OrderItemRes) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemRes.ProtoReflect.Descriptor instead.
func (*OrderItemRes) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *OrderItemRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderItemRes) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *OrderItemRes) GetPayType() string {
	if x != nil {
		return x.PayType
	}
	return ""
}

func (x *OrderItemRes) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *OrderItemRes) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *OrderItemRes) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *OrderItemRes) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *OrderItemRes) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *OrderItemRes) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *OrderItemRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderItemRes) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6d, 0x0a, 0x12, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xcb, 0x01, 0x0a, 0x0f, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x72,
	0x65, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x60, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x4e,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x4e, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x4f, 0x0a, 0x0c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x29, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xa0, 0x01, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0xa4,
	0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x4e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xd5, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x0f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x1a, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x12, 0x31, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x05, 0x5a,
	0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_order_proto_goTypes = []interface{}{
	(*OrderStatus)(nil),        // 0: OrderStatus
	(*OrderItemDetailRes)(nil), // 1: OrderItemDetailRes
	(*OrderProductRes)(nil),    // 2: OrderProductRes
	(*OrderPagingReq)(nil),     // 3: OrderPagingReq
	(*OrderListRes)(nil),       // 4: OrderListRes
	(*OrderItemReq)(nil),       // 5: OrderItemReq
	(*OrderItemRes)(nil),       // 6: OrderItemRes
	(*emptypb.Empty)(nil),      // 7: google.protobuf.Empty
}
var file_order_proto_depIdxs = []int32{
	6, // 0: OrderItemDetailRes.order:type_name -> OrderItemRes
	2, // 1: OrderItemDetailRes.productList:type_name -> OrderProductRes
	6, // 2: OrderListRes.itemList:type_name -> OrderItemRes
	5, // 3: OrderService.CreateOrder:input_type -> OrderItemReq
	3, // 4: OrderService.OrderList:input_type -> OrderPagingReq
	5, // 5: OrderService.OrderDetail:input_type -> OrderItemReq
	0, // 6: OrderService.ChangOrderStatus:input_type -> OrderStatus
	6, // 7: OrderService.CreateOrder:output_type -> OrderItemRes
	4, // 8: OrderService.OrderList:output_type -> OrderListRes
	1, // 9: OrderService.OrderDetail:output_type -> OrderItemDetailRes
	7, // 10: OrderService.ChangOrderStatus:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStatus); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemDetailRes); i {
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
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderProductRes); i {
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
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPagingReq); i {
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
		file_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListRes); i {
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
		file_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemReq); i {
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
		file_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemRes); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	//新建订单
	CreateOrder(ctx context.Context, in *OrderItemReq, opts ...grpc.CallOption) (*OrderItemRes, error)
	//订单列表
	OrderList(ctx context.Context, in *OrderPagingReq, opts ...grpc.CallOption) (*OrderListRes, error)
	//订单详情
	OrderDetail(ctx context.Context, in *OrderItemReq, opts ...grpc.CallOption) (*OrderItemDetailRes, error)
	//更改状态
	ChangOrderStatus(ctx context.Context, in *OrderStatus, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *OrderItemReq, opts ...grpc.CallOption) (*OrderItemRes, error) {
	out := new(OrderItemRes)
	err := c.cc.Invoke(ctx, "/OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) OrderList(ctx context.Context, in *OrderPagingReq, opts ...grpc.CallOption) (*OrderListRes, error) {
	out := new(OrderListRes)
	err := c.cc.Invoke(ctx, "/OrderService/OrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) OrderDetail(ctx context.Context, in *OrderItemReq, opts ...grpc.CallOption) (*OrderItemDetailRes, error) {
	out := new(OrderItemDetailRes)
	err := c.cc.Invoke(ctx, "/OrderService/OrderDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ChangOrderStatus(ctx context.Context, in *OrderStatus, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/OrderService/ChangOrderStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	//新建订单
	CreateOrder(context.Context, *OrderItemReq) (*OrderItemRes, error)
	//订单列表
	OrderList(context.Context, *OrderPagingReq) (*OrderListRes, error)
	//订单详情
	OrderDetail(context.Context, *OrderItemReq) (*OrderItemDetailRes, error)
	//更改状态
	ChangOrderStatus(context.Context, *OrderStatus) (*emptypb.Empty, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) CreateOrder(context.Context, *OrderItemReq) (*OrderItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (*UnimplementedOrderServiceServer) OrderList(context.Context, *OrderPagingReq) (*OrderListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (*UnimplementedOrderServiceServer) OrderDetail(context.Context, *OrderItemReq) (*OrderItemDetailRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDetail not implemented")
}
func (*UnimplementedOrderServiceServer) ChangOrderStatus(context.Context, *OrderStatus) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangOrderStatus not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*OrderItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderPagingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/OrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderList(ctx, req.(*OrderPagingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_OrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/OrderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderDetail(ctx, req.(*OrderItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ChangOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ChangOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/ChangOrderStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ChangOrderStatus(ctx, req.(*OrderStatus))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "OrderList",
			Handler:    _OrderService_OrderList_Handler,
		},
		{
			MethodName: "OrderDetail",
			Handler:    _OrderService_OrderDetail_Handler,
		},
		{
			MethodName: "ChangOrderStatus",
			Handler:    _OrderService_ChangOrderStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
