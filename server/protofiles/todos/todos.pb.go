// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: todos.proto

package todos

import (
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

type FieldMask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paths []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *FieldMask) Reset() {
	*x = FieldMask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMask) ProtoMessage() {}

func (x *FieldMask) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMask.ProtoReflect.Descriptor instead.
func (*FieldMask) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{0}
}

func (x *FieldMask) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ListId   string `protobuf:"bytes,2,opt,name=listId,proto3" json:"listId,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Complete bool   `protobuf:"varint,4,opt,name=complete,proto3" json:"complete,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{1}
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todo) GetListId() string {
	if x != nil {
		return x.ListId
	}
	return ""
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

type TodoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Todos []*Todo `protobuf:"bytes,3,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *TodoList) Reset() {
	*x = TodoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoList) ProtoMessage() {}

func (x *TodoList) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoList.ProtoReflect.Descriptor instead.
func (*TodoList) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{2}
}

func (x *TodoList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodoList) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TodoList) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

type GetTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTodoRequest) Reset() {
	*x = GetTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoRequest) ProtoMessage() {}

func (x *GetTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoRequest.ProtoReflect.Descriptor instead.
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{3}
}

func (x *GetTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{4}
}

func (x *GetListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId   string `protobuf:"bytes,1,opt,name=listId,proto3" json:"listId,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Complete bool   `protobuf:"varint,2,opt,name=complete,proto3" json:"complete,omitempty"`
}

func (x *CreateTodoRequest) Reset() {
	*x = CreateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest) ProtoMessage() {}

func (x *CreateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTodoRequest) GetListId() string {
	if x != nil {
		return x.ListId
	}
	return ""
}

func (x *CreateTodoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTodoRequest) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

type UpdateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Complete bool       `protobuf:"varint,2,opt,name=complete,proto3" json:"complete,omitempty"`
	Title    string     `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Mask     *FieldMask `protobuf:"bytes,4,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *UpdateTodoRequest) Reset() {
	*x = UpdateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoRequest) ProtoMessage() {}

func (x *UpdateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoRequest.ProtoReflect.Descriptor instead.
func (*UpdateTodoRequest) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTodoRequest) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

func (x *UpdateTodoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTodoRequest) GetMask() *FieldMask {
	if x != nil {
		return x.Mask
	}
	return nil
}

type DeleteTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTodoRequest) Reset() {
	*x = DeleteTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoRequest) ProtoMessage() {}

func (x *DeleteTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoRequest.ProtoReflect.Descriptor instead.
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_todos_proto protoreflect.FileDescriptor

var file_todos_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x22, 0x60, 0x0a, 0x04, 0x54, 0x6f, 0x64,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x58, 0x0a, 0x08, 0x54,
	0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x26, 0x0a,
	0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x05,
	0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x29, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0x23, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x32, 0xc6, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x12,
	0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x3b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x64, 0x6f, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f,
	0x64, 0x6f, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todos_proto_rawDescOnce sync.Once
	file_todos_proto_rawDescData = file_todos_proto_rawDesc
)

func file_todos_proto_rawDescGZIP() []byte {
	file_todos_proto_rawDescOnce.Do(func() {
		file_todos_proto_rawDescData = protoimpl.X.CompressGZIP(file_todos_proto_rawDescData)
	})
	return file_todos_proto_rawDescData
}

var file_todos_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_todos_proto_goTypes = []interface{}{
	(*FieldMask)(nil),         // 0: protofiles.FieldMask
	(*Todo)(nil),              // 1: protofiles.Todo
	(*TodoList)(nil),          // 2: protofiles.TodoList
	(*GetTodoRequest)(nil),    // 3: protofiles.GetTodoRequest
	(*GetListRequest)(nil),    // 4: protofiles.GetListRequest
	(*CreateTodoRequest)(nil), // 5: protofiles.CreateTodoRequest
	(*UpdateTodoRequest)(nil), // 6: protofiles.UpdateTodoRequest
	(*DeleteTodoRequest)(nil), // 7: protofiles.DeleteTodoRequest
	(*emptypb.Empty)(nil),     // 8: google.protobuf.Empty
}
var file_todos_proto_depIdxs = []int32{
	1, // 0: protofiles.TodoList.todos:type_name -> protofiles.Todo
	0, // 1: protofiles.UpdateTodoRequest.mask:type_name -> protofiles.FieldMask
	5, // 2: protofiles.TodoService.CreateTodo:input_type -> protofiles.CreateTodoRequest
	3, // 3: protofiles.TodoService.GetTodo:input_type -> protofiles.GetTodoRequest
	4, // 4: protofiles.TodoService.GetList:input_type -> protofiles.GetListRequest
	6, // 5: protofiles.TodoService.UpdateTodo:input_type -> protofiles.UpdateTodoRequest
	7, // 6: protofiles.TodoService.DeleteTodo:input_type -> protofiles.DeleteTodoRequest
	1, // 7: protofiles.TodoService.CreateTodo:output_type -> protofiles.Todo
	1, // 8: protofiles.TodoService.GetTodo:output_type -> protofiles.Todo
	2, // 9: protofiles.TodoService.GetList:output_type -> protofiles.TodoList
	1, // 10: protofiles.TodoService.UpdateTodo:output_type -> protofiles.Todo
	8, // 11: protofiles.TodoService.DeleteTodo:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_todos_proto_init() }
func file_todos_proto_init() {
	if File_todos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldMask); i {
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
		file_todos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_todos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoList); i {
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
		file_todos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoRequest); i {
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
		file_todos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRequest); i {
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
		file_todos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoRequest); i {
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
		file_todos_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoRequest); i {
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
		file_todos_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodoRequest); i {
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
			RawDescriptor: file_todos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todos_proto_goTypes,
		DependencyIndexes: file_todos_proto_depIdxs,
		MessageInfos:      file_todos_proto_msgTypes,
	}.Build()
	File_todos_proto = out.File
	file_todos_proto_rawDesc = nil
	file_todos_proto_goTypes = nil
	file_todos_proto_depIdxs = nil
}
