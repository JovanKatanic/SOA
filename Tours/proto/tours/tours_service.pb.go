// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: tours/tours_service.proto

package tours

import (
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

type Tour struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Difficulty   int32                  `protobuf:"varint,4,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Tags         []string               `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Status       int32                  `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Price        float64                `protobuf:"fixed64,7,opt,name=price,proto3" json:"price,omitempty"`
	AuthorId     int32                  `protobuf:"varint,8,opt,name=authorId,proto3" json:"authorId,omitempty"`
	Equipment    []int32                `protobuf:"varint,9,rep,packed,name=equipment,proto3" json:"equipment,omitempty"`
	DistanceInKm float64                `protobuf:"fixed64,10,opt,name=distanceInKm,proto3" json:"distanceInKm,omitempty"`
	ArchivedDate *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=archivedDate,proto3" json:"archivedDate,omitempty"`
	PublishDate  *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=publishDate,proto3" json:"publishDate,omitempty"`
	Durations    []*TourDuration        `protobuf:"bytes,13,rep,name=durations,proto3" json:"durations,omitempty"`
	Keypoints    []*Keypoint            `protobuf:"bytes,14,rep,name=keypoints,proto3" json:"keypoints,omitempty"`
	Image        string                 `protobuf:"bytes,15,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Tour) Reset() {
	*x = Tour{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tour) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tour) ProtoMessage() {}

func (x *Tour) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tour.ProtoReflect.Descriptor instead.
func (*Tour) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{0}
}

func (x *Tour) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tour) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tour) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Tour) GetDifficulty() int32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *Tour) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Tour) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Tour) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Tour) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *Tour) GetEquipment() []int32 {
	if x != nil {
		return x.Equipment
	}
	return nil
}

func (x *Tour) GetDistanceInKm() float64 {
	if x != nil {
		return x.DistanceInKm
	}
	return 0
}

func (x *Tour) GetArchivedDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ArchivedDate
	}
	return nil
}

func (x *Tour) GetPublishDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishDate
	}
	return nil
}

func (x *Tour) GetDurations() []*TourDuration {
	if x != nil {
		return x.Durations
	}
	return nil
}

func (x *Tour) GetKeypoints() []*Keypoint {
	if x != nil {
		return x.Keypoints
	}
	return nil
}

func (x *Tour) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type TourDuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeInSeconds  uint32 `protobuf:"varint,1,opt,name=timeInSeconds,proto3" json:"timeInSeconds,omitempty"`
	Transportation int32  `protobuf:"varint,2,opt,name=transportation,proto3" json:"transportation,omitempty"`
}

func (x *TourDuration) Reset() {
	*x = TourDuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TourDuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TourDuration) ProtoMessage() {}

func (x *TourDuration) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TourDuration.ProtoReflect.Descriptor instead.
func (*TourDuration) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{1}
}

func (x *TourDuration) GetTimeInSeconds() uint32 {
	if x != nil {
		return x.TimeInSeconds
	}
	return 0
}

func (x *TourDuration) GetTransportation() int32 {
	if x != nil {
		return x.Transportation
	}
	return 0
}

type Keypoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description    string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Image          string  `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Latitude       float64 `protobuf:"fixed64,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude      float64 `protobuf:"fixed64,6,opt,name=longitude,proto3" json:"longitude,omitempty"`
	TourId         int32   `protobuf:"varint,7,opt,name=tourId,proto3" json:"tourId,omitempty"`
	PositionInTour int32   `protobuf:"varint,8,opt,name=positionInTour,proto3" json:"positionInTour,omitempty"`
	Secret         string  `protobuf:"bytes,9,opt,name=secret,proto3" json:"secret,omitempty"`
	Discriminator  string  `protobuf:"bytes,10,opt,name=discriminator,proto3" json:"discriminator,omitempty"`
}

func (x *Keypoint) Reset() {
	*x = Keypoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keypoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keypoint) ProtoMessage() {}

func (x *Keypoint) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keypoint.ProtoReflect.Descriptor instead.
func (*Keypoint) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{2}
}

func (x *Keypoint) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Keypoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Keypoint) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Keypoint) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Keypoint) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Keypoint) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Keypoint) GetTourId() int32 {
	if x != nil {
		return x.TourId
	}
	return 0
}

func (x *Keypoint) GetPositionInTour() int32 {
	if x != nil {
		return x.PositionInTour
	}
	return 0
}

func (x *Keypoint) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Keypoint) GetDiscriminator() string {
	if x != nil {
		return x.Discriminator
	}
	return ""
}

type CreateTourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tour *Tour `protobuf:"bytes,1,opt,name=tour,proto3" json:"tour,omitempty"`
}

func (x *CreateTourRequest) Reset() {
	*x = CreateTourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTourRequest) ProtoMessage() {}

func (x *CreateTourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTourRequest.ProtoReflect.Descriptor instead.
func (*CreateTourRequest) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTourRequest) GetTour() *Tour {
	if x != nil {
		return x.Tour
	}
	return nil
}

type CreateTourResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tour *Tour `protobuf:"bytes,1,opt,name=tour,proto3" json:"tour,omitempty"`
}

func (x *CreateTourResponse) Reset() {
	*x = CreateTourResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTourResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTourResponse) ProtoMessage() {}

func (x *CreateTourResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTourResponse.ProtoReflect.Descriptor instead.
func (*CreateTourResponse) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTourResponse) GetTour() *Tour {
	if x != nil {
		return x.Tour
	}
	return nil
}

type GetTourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTourRequest) Reset() {
	*x = GetTourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTourRequest) ProtoMessage() {}

func (x *GetTourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTourRequest.ProtoReflect.Descriptor instead.
func (*GetTourRequest) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetTourRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTourResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tour *Tour `protobuf:"bytes,1,opt,name=tour,proto3" json:"tour,omitempty"`
}

func (x *GetTourResponse) Reset() {
	*x = GetTourResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTourResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTourResponse) ProtoMessage() {}

func (x *GetTourResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTourResponse.ProtoReflect.Descriptor instead.
func (*GetTourResponse) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetTourResponse) GetTour() *Tour {
	if x != nil {
		return x.Tour
	}
	return nil
}

type GetToursByAuthorIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetToursByAuthorIdRequest) Reset() {
	*x = GetToursByAuthorIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetToursByAuthorIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetToursByAuthorIdRequest) ProtoMessage() {}

func (x *GetToursByAuthorIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetToursByAuthorIdRequest.ProtoReflect.Descriptor instead.
func (*GetToursByAuthorIdRequest) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetToursByAuthorIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetToursByAuthorIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tour []*Tour `protobuf:"bytes,1,rep,name=tour,proto3" json:"tour,omitempty"`
}

func (x *GetToursByAuthorIdResponse) Reset() {
	*x = GetToursByAuthorIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetToursByAuthorIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetToursByAuthorIdResponse) ProtoMessage() {}

func (x *GetToursByAuthorIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetToursByAuthorIdResponse.ProtoReflect.Descriptor instead.
func (*GetToursByAuthorIdResponse) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetToursByAuthorIdResponse) GetTour() []*Tour {
	if x != nil {
		return x.Tour
	}
	return nil
}

type UpdateTourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tour *Tour `protobuf:"bytes,1,opt,name=tour,proto3" json:"tour,omitempty"`
}

func (x *UpdateTourRequest) Reset() {
	*x = UpdateTourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTourRequest) ProtoMessage() {}

func (x *UpdateTourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTourRequest.ProtoReflect.Descriptor instead.
func (*UpdateTourRequest) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateTourRequest) GetTour() *Tour {
	if x != nil {
		return x.Tour
	}
	return nil
}

type UpdateTourResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tour *Tour `protobuf:"bytes,1,opt,name=tour,proto3" json:"tour,omitempty"`
}

func (x *UpdateTourResponse) Reset() {
	*x = UpdateTourResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTourResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTourResponse) ProtoMessage() {}

func (x *UpdateTourResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTourResponse.ProtoReflect.Descriptor instead.
func (*UpdateTourResponse) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateTourResponse) GetTour() *Tour {
	if x != nil {
		return x.Tour
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{11}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tours []*Tour `protobuf:"bytes,1,rep,name=tours,proto3" json:"tours,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tours_tours_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tours_tours_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_tours_tours_service_proto_rawDescGZIP(), []int{12}
}

func (x *GetAllResponse) GetTours() []*Tour {
	if x != nil {
		return x.Tours
	}
	return nil
}

var File_tours_tours_service_proto protoreflect.FileDescriptor

var file_tours_tours_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x03, 0x0a,
	0x04, 0x54, 0x6f, 0x75, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x65, 0x71,
	0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x6e, 0x4b, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x64,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x4b, 0x6d, 0x12, 0x3e, 0x0a, 0x0c, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54,
	0x6f, 0x75, 0x72, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4b, 0x65, 0x79, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x5c, 0x0a, 0x0c, 0x54, 0x6f, 0x75, 0x72, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x69,
	0x6d, 0x65, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x9e, 0x02, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x75, 0x72, 0x49, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x6f, 0x75, 0x72, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x54, 0x6f, 0x75, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x54, 0x6f, 0x75, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x22, 0x2e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x6f, 0x75,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x04,
	0x74, 0x6f, 0x75, 0x72, 0x22, 0x2f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x6f,
	0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x52,
	0x04, 0x74, 0x6f, 0x75, 0x72, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x6f,
	0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x52,
	0x04, 0x74, 0x6f, 0x75, 0x72, 0x22, 0x2b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72,
	0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x37, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x73, 0x42, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x04, 0x74, 0x6f, 0x75, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x04, 0x74, 0x6f, 0x75, 0x72, 0x22, 0x2e, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x04, 0x74, 0x6f, 0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x04, 0x74, 0x6f, 0x75, 0x72, 0x22, 0x2f, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x6f, 0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x04, 0x74, 0x6f, 0x75, 0x72, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x74, 0x6f, 0x75, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x05, 0x74,
	0x6f, 0x75, 0x72, 0x73, 0x32, 0xa9, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x75, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x75, 0x72, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x73, 0x42, 0x79, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75,
	0x72, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x73, 0x42, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72,
	0x12, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tours_tours_service_proto_rawDescOnce sync.Once
	file_tours_tours_service_proto_rawDescData = file_tours_tours_service_proto_rawDesc
)

func file_tours_tours_service_proto_rawDescGZIP() []byte {
	file_tours_tours_service_proto_rawDescOnce.Do(func() {
		file_tours_tours_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tours_tours_service_proto_rawDescData)
	})
	return file_tours_tours_service_proto_rawDescData
}

var file_tours_tours_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_tours_tours_service_proto_goTypes = []interface{}{
	(*Tour)(nil),                       // 0: Tour
	(*TourDuration)(nil),               // 1: TourDuration
	(*Keypoint)(nil),                   // 2: Keypoint
	(*CreateTourRequest)(nil),          // 3: CreateTourRequest
	(*CreateTourResponse)(nil),         // 4: CreateTourResponse
	(*GetTourRequest)(nil),             // 5: GetTourRequest
	(*GetTourResponse)(nil),            // 6: GetTourResponse
	(*GetToursByAuthorIdRequest)(nil),  // 7: GetToursByAuthorIdRequest
	(*GetToursByAuthorIdResponse)(nil), // 8: GetToursByAuthorIdResponse
	(*UpdateTourRequest)(nil),          // 9: UpdateTourRequest
	(*UpdateTourResponse)(nil),         // 10: UpdateTourResponse
	(*Empty)(nil),                      // 11: Empty
	(*GetAllResponse)(nil),             // 12: GetAllResponse
	(*timestamppb.Timestamp)(nil),      // 13: google.protobuf.Timestamp
}
var file_tours_tours_service_proto_depIdxs = []int32{
	13, // 0: Tour.archivedDate:type_name -> google.protobuf.Timestamp
	13, // 1: Tour.publishDate:type_name -> google.protobuf.Timestamp
	1,  // 2: Tour.durations:type_name -> TourDuration
	2,  // 3: Tour.keypoints:type_name -> Keypoint
	0,  // 4: CreateTourRequest.tour:type_name -> Tour
	0,  // 5: CreateTourResponse.tour:type_name -> Tour
	0,  // 6: GetTourResponse.tour:type_name -> Tour
	0,  // 7: GetToursByAuthorIdResponse.tour:type_name -> Tour
	0,  // 8: UpdateTourRequest.tour:type_name -> Tour
	0,  // 9: UpdateTourResponse.tour:type_name -> Tour
	0,  // 10: GetAllResponse.tours:type_name -> Tour
	3,  // 11: TourService.CreateTour:input_type -> CreateTourRequest
	5,  // 12: TourService.GetTourById:input_type -> GetTourRequest
	7,  // 13: TourService.GetToursByAuthorId:input_type -> GetToursByAuthorIdRequest
	9,  // 14: TourService.UpdateTour:input_type -> UpdateTourRequest
	11, // 15: TourService.GetAll:input_type -> Empty
	4,  // 16: TourService.CreateTour:output_type -> CreateTourResponse
	6,  // 17: TourService.GetTourById:output_type -> GetTourResponse
	8,  // 18: TourService.GetToursByAuthorId:output_type -> GetToursByAuthorIdResponse
	10, // 19: TourService.UpdateTour:output_type -> UpdateTourResponse
	12, // 20: TourService.GetAll:output_type -> GetAllResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tours_tours_service_proto_init() }
func file_tours_tours_service_proto_init() {
	if File_tours_tours_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tours_tours_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tour); i {
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
		file_tours_tours_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TourDuration); i {
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
		file_tours_tours_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keypoint); i {
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
		file_tours_tours_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTourRequest); i {
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
		file_tours_tours_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTourResponse); i {
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
		file_tours_tours_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTourRequest); i {
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
		file_tours_tours_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTourResponse); i {
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
		file_tours_tours_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetToursByAuthorIdRequest); i {
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
		file_tours_tours_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetToursByAuthorIdResponse); i {
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
		file_tours_tours_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTourRequest); i {
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
		file_tours_tours_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTourResponse); i {
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
		file_tours_tours_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_tours_tours_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
			RawDescriptor: file_tours_tours_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tours_tours_service_proto_goTypes,
		DependencyIndexes: file_tours_tours_service_proto_depIdxs,
		MessageInfos:      file_tours_tours_service_proto_msgTypes,
	}.Build()
	File_tours_tours_service_proto = out.File
	file_tours_tours_service_proto_rawDesc = nil
	file_tours_tours_service_proto_goTypes = nil
	file_tours_tours_service_proto_depIdxs = nil
}
