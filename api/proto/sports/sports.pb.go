// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: sports/sports.proto

package sports

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//Types of sports
type SportType int32

const (
	SportType_RUGBY   SportType = 0
	SportType_CRICKET SportType = 1
	SportType_AFL     SportType = 2
	SportType_SOCCER  SportType = 3
)

// Enum value maps for SportType.
var (
	SportType_name = map[int32]string{
		0: "RUGBY",
		1: "CRICKET",
		2: "AFL",
		3: "SOCCER",
	}
	SportType_value = map[string]int32{
		"RUGBY":   0,
		"CRICKET": 1,
		"AFL":     2,
		"SOCCER":  3,
	}
)

func (x SportType) Enum() *SportType {
	p := new(SportType)
	*p = x
	return p
}

func (x SportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SportType) Descriptor() protoreflect.EnumDescriptor {
	return file_sports_sports_proto_enumTypes[0].Descriptor()
}

func (SportType) Type() protoreflect.EnumType {
	return &file_sports_sports_proto_enumTypes[0]
}

func (x SportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SportType.Descriptor instead.
func (SportType) EnumDescriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{0}
}

// Request for GetEvent call.
type GetEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{0}
}

func (x *GetEventRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response to GetEvent call.
type GetEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *GetEventResponse) Reset() {
	*x = GetEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventResponse) ProtoMessage() {}

func (x *GetEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventResponse.ProtoReflect.Descriptor instead.
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{1}
}

func (x *GetEventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

// Request for ListEvent call.
type ListEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *ListEventsRequestFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	SortBy *ListEventsRequestSortBy `protobuf:"bytes,2,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
}

func (x *ListEventsRequest) Reset() {
	*x = ListEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRequest) ProtoMessage() {}

func (x *ListEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsRequest.ProtoReflect.Descriptor instead.
func (*ListEventsRequest) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{2}
}

func (x *ListEventsRequest) GetFilter() *ListEventsRequestFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListEventsRequest) GetSortBy() *ListEventsRequestSortBy {
	if x != nil {
		return x.SortBy
	}
	return nil
}

// Response to ListEvents call.
type ListEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ListEventsResponse) Reset() {
	*x = ListEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsResponse) ProtoMessage() {}

func (x *ListEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsResponse.ProtoReflect.Descriptor instead.
func (*ListEventsResponse) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{3}
}

func (x *ListEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

// Filter for listing events.
type ListEventsRequestFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    []SportType `protobuf:"varint,1,rep,packed,name=type,proto3,enum=sports.SportType" json:"type,omitempty"`
	Visible *bool       `protobuf:"varint,2,opt,name=visible,proto3,oneof" json:"visible,omitempty"`
}

func (x *ListEventsRequestFilter) Reset() {
	*x = ListEventsRequestFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsRequestFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRequestFilter) ProtoMessage() {}

func (x *ListEventsRequestFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsRequestFilter.ProtoReflect.Descriptor instead.
func (*ListEventsRequestFilter) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{4}
}

func (x *ListEventsRequestFilter) GetType() []SportType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ListEventsRequestFilter) GetVisible() bool {
	if x != nil && x.Visible != nil {
		return *x.Visible
	}
	return false
}

//
type ListEventsRequestSortBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyName string `protobuf:"bytes,1,opt,name=property_name,json=propertyName,proto3" json:"property_name,omitempty"`
	Descending   bool   `protobuf:"varint,2,opt,name=descending,proto3" json:"descending,omitempty"`
}

func (x *ListEventsRequestSortBy) Reset() {
	*x = ListEventsRequestSortBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsRequestSortBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRequestSortBy) ProtoMessage() {}

func (x *ListEventsRequestSortBy) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsRequestSortBy.ProtoReflect.Descriptor instead.
func (*ListEventsRequestSortBy) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{5}
}

func (x *ListEventsRequestSortBy) GetPropertyName() string {
	if x != nil {
		return x.PropertyName
	}
	return ""
}

func (x *ListEventsRequestSortBy) GetDescending() bool {
	if x != nil {
		return x.Descending
	}
	return false
}

// An event resource.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID represents a unique identifier for the event.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name is the official name given to the event.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Type of sport the event is about
	Type SportType `protobuf:"varint,3,opt,name=type,proto3,enum=sports.SportType" json:"type,omitempty"`
	// Location represents the location of the event.
	Location string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	// Visible represents whether or not the event is visible.
	Visible bool `protobuf:"varint,5,opt,name=visible,proto3" json:"visible,omitempty"`
	// AdvertisedStartTime is the time the event is advertised to run.
	AdvertisedStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=advertised_start_time,json=advertisedStartTime,proto3" json:"advertised_start_time,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{6}
}

func (x *Event) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetType() SportType {
	if x != nil {
		return x.Type
	}
	return SportType_RUGBY
}

func (x *Event) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Event) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Event) GetAdvertisedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AdvertisedStartTime
	}
	return nil
}

var File_sports_sports_proto protoreflect.FileDescriptor

var file_sports_sports_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79,
	0x22, 0x3b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x6b, 0x0a,
	0x17, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e,
	0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0x5e, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65,
	0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0xd8, 0x01, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e,
	0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x38, 0x0a, 0x09, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x55, 0x47, 0x42, 0x59, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x52, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x46,
	0x4c, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4f, 0x43, 0x43, 0x45, 0x52, 0x10, 0x03, 0x32,
	0xc7, 0x01, 0x0a, 0x06, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x5f, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x2d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sports_sports_proto_rawDescOnce sync.Once
	file_sports_sports_proto_rawDescData = file_sports_sports_proto_rawDesc
)

func file_sports_sports_proto_rawDescGZIP() []byte {
	file_sports_sports_proto_rawDescOnce.Do(func() {
		file_sports_sports_proto_rawDescData = protoimpl.X.CompressGZIP(file_sports_sports_proto_rawDescData)
	})
	return file_sports_sports_proto_rawDescData
}

var file_sports_sports_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sports_sports_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sports_sports_proto_goTypes = []interface{}{
	(SportType)(0),                  // 0: sports.SportType
	(*GetEventRequest)(nil),         // 1: sports.GetEventRequest
	(*GetEventResponse)(nil),        // 2: sports.GetEventResponse
	(*ListEventsRequest)(nil),       // 3: sports.ListEventsRequest
	(*ListEventsResponse)(nil),      // 4: sports.ListEventsResponse
	(*ListEventsRequestFilter)(nil), // 5: sports.ListEventsRequestFilter
	(*ListEventsRequestSortBy)(nil), // 6: sports.ListEventsRequestSortBy
	(*Event)(nil),                   // 7: sports.Event
	(*timestamppb.Timestamp)(nil),   // 8: google.protobuf.Timestamp
}
var file_sports_sports_proto_depIdxs = []int32{
	7, // 0: sports.GetEventResponse.event:type_name -> sports.Event
	5, // 1: sports.ListEventsRequest.filter:type_name -> sports.ListEventsRequestFilter
	6, // 2: sports.ListEventsRequest.sortBy:type_name -> sports.ListEventsRequestSortBy
	7, // 3: sports.ListEventsResponse.events:type_name -> sports.Event
	0, // 4: sports.ListEventsRequestFilter.type:type_name -> sports.SportType
	0, // 5: sports.Event.type:type_name -> sports.SportType
	8, // 6: sports.Event.advertised_start_time:type_name -> google.protobuf.Timestamp
	3, // 7: sports.Sports.ListEvents:input_type -> sports.ListEventsRequest
	1, // 8: sports.Sports.GetEvent:input_type -> sports.GetEventRequest
	4, // 9: sports.Sports.ListEvents:output_type -> sports.ListEventsResponse
	2, // 10: sports.Sports.GetEvent:output_type -> sports.GetEventResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_sports_sports_proto_init() }
func file_sports_sports_proto_init() {
	if File_sports_sports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sports_sports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventRequest); i {
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
		file_sports_sports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventResponse); i {
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
		file_sports_sports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsRequest); i {
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
		file_sports_sports_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsResponse); i {
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
		file_sports_sports_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsRequestFilter); i {
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
		file_sports_sports_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsRequestSortBy); i {
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
		file_sports_sports_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
	file_sports_sports_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sports_sports_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sports_sports_proto_goTypes,
		DependencyIndexes: file_sports_sports_proto_depIdxs,
		EnumInfos:         file_sports_sports_proto_enumTypes,
		MessageInfos:      file_sports_sports_proto_msgTypes,
	}.Build()
	File_sports_sports_proto = out.File
	file_sports_sports_proto_rawDesc = nil
	file_sports_sports_proto_goTypes = nil
	file_sports_sports_proto_depIdxs = nil
}