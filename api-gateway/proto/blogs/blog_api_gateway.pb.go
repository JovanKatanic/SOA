// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: blog_api_gateway.proto

package blogs

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

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title        string                 `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description  string                 `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	CreationDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=CreationDate,proto3" json:"CreationDate,omitempty"`
	Status       int32                  `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	UserId       int32                  `protobuf:"varint,6,opt,name=UserId,proto3" json:"UserId,omitempty"`
	RatingSum    int32                  `protobuf:"varint,8,opt,name=RatingSum,proto3" json:"RatingSum,omitempty"`
	Ratings      []*Rating              `protobuf:"bytes,9,rep,name=Ratings,proto3" json:"Ratings,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *Blog) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Blog) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

func (x *Blog) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Blog) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Blog) GetRatingSum() int32 {
	if x != nil {
		return x.RatingSum
	}
	return 0
}

func (x *Blog) GetRatings() []*Rating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int32                  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	CreationDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=CreationDate,proto3" json:"CreationDate,omitempty"`
	RatingValue  int32                  `protobuf:"varint,3,opt,name=RatingValue,proto3" json:"RatingValue,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *Rating) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Rating) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

func (x *Rating) GetRatingValue() int32 {
	if x != nil {
		return x.RatingValue
	}
	return 0
}

type GetBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBlogRequest) Reset() {
	*x = GetBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogRequest) ProtoMessage() {}

func (x *GetBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogRequest.ProtoReflect.Descriptor instead.
func (*GetBlogRequest) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlogRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Emptyyy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Emptyyy) Reset() {
	*x = Emptyyy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emptyyy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emptyyy) ProtoMessage() {}

func (x *Emptyyy) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emptyyy.ProtoReflect.Descriptor instead.
func (*Emptyyy) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{3}
}

type ListBlog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blogs []*Blog `protobuf:"bytes,1,rep,name=Blogs,proto3" json:"Blogs,omitempty"`
}

func (x *ListBlog) Reset() {
	*x = ListBlog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlog) ProtoMessage() {}

func (x *ListBlog) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlog.ProtoReflect.Descriptor instead.
func (*ListBlog) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *ListBlog) GetBlogs() []*Blog {
	if x != nil {
		return x.Blogs
	}
	return nil
}

type GetBlogStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State int32 `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *GetBlogStatus) Reset() {
	*x = GetBlogStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogStatus) ProtoMessage() {}

func (x *GetBlogStatus) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogStatus.ProtoReflect.Descriptor instead.
func (*GetBlogStatus) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *GetBlogStatus) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

type UpdateRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId int32 `protobuf:"varint,1,opt,name=blogId,proto3" json:"blogId,omitempty"`
	UserId int32 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Value  int32 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateRatingRequest) Reset() {
	*x = UpdateRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRatingRequest) ProtoMessage() {}

func (x *UpdateRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRatingRequest.ProtoReflect.Descriptor instead.
func (*UpdateRatingRequest) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRatingRequest) GetBlogId() int32 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

func (x *UpdateRatingRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateRatingRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type DeleteRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId int32 `protobuf:"varint,1,opt,name=blogId,proto3" json:"blogId,omitempty"`
	UserId int32 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DeleteRatingRequest) Reset() {
	*x = DeleteRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRatingRequest) ProtoMessage() {}

func (x *DeleteRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRatingRequest.ProtoReflect.Descriptor instead.
func (*DeleteRatingRequest) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRatingRequest) GetBlogId() int32 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

func (x *DeleteRatingRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId       int32                  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Username     string                 `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	ProfilePic   string                 `protobuf:"bytes,4,opt,name=ProfilePic,proto3" json:"ProfilePic,omitempty"`
	CreationDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=CreationDate,proto3" json:"CreationDate,omitempty"`
	Description  string                 `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	LastEditDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=LastEditDate,proto3" json:"LastEditDate,omitempty"`
	BlogId       int32                  `protobuf:"varint,8,opt,name=BlogId,proto3" json:"BlogId,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{8}
}

func (x *Comment) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Comment) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Comment) GetProfilePic() string {
	if x != nil {
		return x.ProfilePic
	}
	return ""
}

func (x *Comment) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

func (x *Comment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Comment) GetLastEditDate() *timestamppb.Timestamp {
	if x != nil {
		return x.LastEditDate
	}
	return nil
}

func (x *Comment) GetBlogId() int32 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

type GetCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId int32 `protobuf:"varint,1,opt,name=blogId,proto3" json:"blogId,omitempty"`
}

func (x *GetCommentRequest) Reset() {
	*x = GetCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentRequest) ProtoMessage() {}

func (x *GetCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentRequest.ProtoReflect.Descriptor instead.
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{9}
}

func (x *GetCommentRequest) GetBlogId() int32 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

type ListComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment `protobuf:"bytes,1,rep,name=Comments,proto3" json:"Comments,omitempty"`
}

func (x *ListComment) Reset() {
	*x = ListComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComment) ProtoMessage() {}

func (x *ListComment) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComment.ProtoReflect.Descriptor instead.
func (*ListComment) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{10}
}

func (x *ListComment) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type DeleteCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCommentRequest) Reset() {
	*x = DeleteCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_api_gateway_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRequest) ProtoMessage() {}

func (x *DeleteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_api_gateway_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return file_blog_api_gateway_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteCommentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_blog_api_gateway_proto protoreflect.FileDescriptor

var file_blog_api_gateway_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x01, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x53, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x06, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x09, 0x0a, 0x07, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x79, 0x79, 0x22, 0x27, 0x0a, 0x08, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x1b, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x05, 0x42,
	0x6c, 0x6f, 0x67, 0x73, 0x22, 0x25, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x5b, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x45, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xa7, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x12,
	0x3e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x32, 0x8a, 0x06, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x0f,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b,
	0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x2d, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67,
	0x1a, 0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x3a,
	0x01, 0x2a, 0x22, 0x06, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x08, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x79, 0x79, 0x1a, 0x09, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x0e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x3e, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x05,
	0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x1a, 0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x1a, 0x14, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x54, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x1a, 0x09, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f,
	0x67, 0x65, 0x74, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x7d, 0x12, 0x5f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67,
	0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x1a, 0x27, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x7d, 0x2f, 0x7b, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x7d, 0x2f, 0x7b, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x7d, 0x12, 0x54, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x42, 0x6c, 0x6f,
	0x67, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x2a, 0x1f, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x73, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x7d, 0x2f, 0x7b, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x38, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x08, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x57, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x79, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x41, 0x73, 0x79, 0x6e, 0x63,
	0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x38, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x08,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x1a, 0x08, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x2a, 0x0d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_api_gateway_proto_rawDescOnce sync.Once
	file_blog_api_gateway_proto_rawDescData = file_blog_api_gateway_proto_rawDesc
)

func file_blog_api_gateway_proto_rawDescGZIP() []byte {
	file_blog_api_gateway_proto_rawDescOnce.Do(func() {
		file_blog_api_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_api_gateway_proto_rawDescData)
	})
	return file_blog_api_gateway_proto_rawDescData
}

var file_blog_api_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_blog_api_gateway_proto_goTypes = []interface{}{
	(*Blog)(nil),                  // 0: Blog
	(*Rating)(nil),                // 1: Rating
	(*GetBlogRequest)(nil),        // 2: GetBlogRequest
	(*Emptyyy)(nil),               // 3: Emptyyy
	(*ListBlog)(nil),              // 4: ListBlog
	(*GetBlogStatus)(nil),         // 5: GetBlogStatus
	(*UpdateRatingRequest)(nil),   // 6: UpdateRatingRequest
	(*DeleteRatingRequest)(nil),   // 7: DeleteRatingRequest
	(*Comment)(nil),               // 8: Comment
	(*GetCommentRequest)(nil),     // 9: GetCommentRequest
	(*ListComment)(nil),           // 10: ListComment
	(*DeleteCommentRequest)(nil),  // 11: DeleteCommentRequest
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_blog_api_gateway_proto_depIdxs = []int32{
	12, // 0: Blog.CreationDate:type_name -> google.protobuf.Timestamp
	1,  // 1: Blog.Ratings:type_name -> Rating
	12, // 2: Rating.CreationDate:type_name -> google.protobuf.Timestamp
	0,  // 3: ListBlog.Blogs:type_name -> Blog
	12, // 4: Comment.CreationDate:type_name -> google.protobuf.Timestamp
	12, // 5: Comment.LastEditDate:type_name -> google.protobuf.Timestamp
	8,  // 6: ListComment.Comments:type_name -> Comment
	2,  // 7: BlogService.GetBlog:input_type -> GetBlogRequest
	0,  // 8: BlogService.CreateBlog:input_type -> Blog
	3,  // 9: BlogService.GetAllBlog:input_type -> Emptyyy
	0,  // 10: BlogService.UpdateOneBlog:input_type -> Blog
	5,  // 11: BlogService.GetAllBlogsByStatus:input_type -> GetBlogStatus
	6,  // 12: BlogService.UpdateRating:input_type -> UpdateRatingRequest
	7,  // 13: BlogService.DeleteRating:input_type -> DeleteRatingRequest
	8,  // 14: BlogService.CreateComment:input_type -> Comment
	9,  // 15: BlogService.GetCommentsByBlogIdAsync:input_type -> GetCommentRequest
	8,  // 16: BlogService.UpdateComment:input_type -> Comment
	11, // 17: BlogService.DeleteComment:input_type -> DeleteCommentRequest
	0,  // 18: BlogService.GetBlog:output_type -> Blog
	0,  // 19: BlogService.CreateBlog:output_type -> Blog
	4,  // 20: BlogService.GetAllBlog:output_type -> ListBlog
	0,  // 21: BlogService.UpdateOneBlog:output_type -> Blog
	4,  // 22: BlogService.GetAllBlogsByStatus:output_type -> ListBlog
	0,  // 23: BlogService.UpdateRating:output_type -> Blog
	0,  // 24: BlogService.DeleteRating:output_type -> Blog
	8,  // 25: BlogService.CreateComment:output_type -> Comment
	10, // 26: BlogService.GetCommentsByBlogIdAsync:output_type -> ListComment
	8,  // 27: BlogService.UpdateComment:output_type -> Comment
	8,  // 28: BlogService.DeleteComment:output_type -> Comment
	18, // [18:29] is the sub-list for method output_type
	7,  // [7:18] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_blog_api_gateway_proto_init() }
func file_blog_api_gateway_proto_init() {
	if File_blog_api_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blog_api_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
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
		file_blog_api_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
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
		file_blog_api_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogRequest); i {
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
		file_blog_api_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emptyyy); i {
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
		file_blog_api_gateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlog); i {
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
		file_blog_api_gateway_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogStatus); i {
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
		file_blog_api_gateway_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRatingRequest); i {
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
		file_blog_api_gateway_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRatingRequest); i {
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
		file_blog_api_gateway_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_blog_api_gateway_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentRequest); i {
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
		file_blog_api_gateway_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListComment); i {
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
		file_blog_api_gateway_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentRequest); i {
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
			RawDescriptor: file_blog_api_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_api_gateway_proto_goTypes,
		DependencyIndexes: file_blog_api_gateway_proto_depIdxs,
		MessageInfos:      file_blog_api_gateway_proto_msgTypes,
	}.Build()
	File_blog_api_gateway_proto = out.File
	file_blog_api_gateway_proto_rawDesc = nil
	file_blog_api_gateway_proto_goTypes = nil
	file_blog_api_gateway_proto_depIdxs = nil
}
