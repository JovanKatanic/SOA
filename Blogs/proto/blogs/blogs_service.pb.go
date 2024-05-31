// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: blogs/blogs_service.proto

package blogs

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

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32     `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title        string    `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description  string    `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	CreationDate string    `protobuf:"bytes,4,opt,name=CreationDate,proto3" json:"CreationDate,omitempty"`
	Status       int32     `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	UserId       int32     `protobuf:"varint,6,opt,name=UserId,proto3" json:"UserId,omitempty"`
	RatingSum    int32     `protobuf:"varint,8,opt,name=RatingSum,proto3" json:"RatingSum,omitempty"`
	Ratings      []*Rating `protobuf:"bytes,9,rep,name=Ratings,proto3" json:"Ratings,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogs_blogs_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_blogs_blogs_service_proto_msgTypes[0]
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
	return file_blogs_blogs_service_proto_rawDescGZIP(), []int{0}
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

func (x *Blog) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
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

	UserId       int32  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	CreationDate string `protobuf:"bytes,2,opt,name=CreationDate,proto3" json:"CreationDate,omitempty"`
	RatingValue  int32  `protobuf:"varint,3,opt,name=RatingValue,proto3" json:"RatingValue,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogs_blogs_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_blogs_blogs_service_proto_msgTypes[1]
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
	return file_blogs_blogs_service_proto_rawDescGZIP(), []int{1}
}

func (x *Rating) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Rating) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
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
		mi := &file_blogs_blogs_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogRequest) ProtoMessage() {}

func (x *GetBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blogs_blogs_service_proto_msgTypes[2]
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
	return file_blogs_blogs_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlogRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogs_blogs_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_blogs_blogs_service_proto_msgTypes[3]
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
	return file_blogs_blogs_service_proto_rawDescGZIP(), []int{3}
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
		mi := &file_blogs_blogs_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlog) ProtoMessage() {}

func (x *ListBlog) ProtoReflect() protoreflect.Message {
	mi := &file_blogs_blogs_service_proto_msgTypes[4]
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
	return file_blogs_blogs_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListBlog) GetBlogs() []*Blog {
	if x != nil {
		return x.Blogs
	}
	return nil
}

var File_blogs_blogs_service_proto protoreflect.FileDescriptor

var file_blogs_blogs_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x04,
	0x42, 0x6c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x6d, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x6d, 0x12, 0x21,
	0x0a, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0x66, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x27, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67,
	0x12, 0x1b, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x05, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x32, 0x73, 0x0a,
	0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22,
	0x00, 0x12, 0x1c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12,
	0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x1a, 0x05, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x00, 0x12,
	0x21, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x06, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67,
	0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blogs_blogs_service_proto_rawDescOnce sync.Once
	file_blogs_blogs_service_proto_rawDescData = file_blogs_blogs_service_proto_rawDesc
)

func file_blogs_blogs_service_proto_rawDescGZIP() []byte {
	file_blogs_blogs_service_proto_rawDescOnce.Do(func() {
		file_blogs_blogs_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_blogs_blogs_service_proto_rawDescData)
	})
	return file_blogs_blogs_service_proto_rawDescData
}

var file_blogs_blogs_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_blogs_blogs_service_proto_goTypes = []interface{}{
	(*Blog)(nil),           // 0: Blog
	(*Rating)(nil),         // 1: Rating
	(*GetBlogRequest)(nil), // 2: GetBlogRequest
	(*Empty)(nil),          // 3: Empty
	(*ListBlog)(nil),       // 4: ListBlog
}
var file_blogs_blogs_service_proto_depIdxs = []int32{
	1, // 0: Blog.Ratings:type_name -> Rating
	0, // 1: ListBlog.Blogs:type_name -> Blog
	2, // 2: BlogService.GetBlog:input_type -> GetBlogRequest
	0, // 3: BlogService.CreateBlog:input_type -> Blog
	3, // 4: BlogService.GetAllBlog:input_type -> Empty
	0, // 5: BlogService.GetBlog:output_type -> Blog
	0, // 6: BlogService.CreateBlog:output_type -> Blog
	4, // 7: BlogService.GetAllBlog:output_type -> ListBlog
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_blogs_blogs_service_proto_init() }
func file_blogs_blogs_service_proto_init() {
	if File_blogs_blogs_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blogs_blogs_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_blogs_blogs_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_blogs_blogs_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_blogs_blogs_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_blogs_blogs_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blogs_blogs_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blogs_blogs_service_proto_goTypes,
		DependencyIndexes: file_blogs_blogs_service_proto_depIdxs,
		MessageInfos:      file_blogs_blogs_service_proto_msgTypes,
	}.Build()
	File_blogs_blogs_service_proto = out.File
	file_blogs_blogs_service_proto_rawDesc = nil
	file_blogs_blogs_service_proto_goTypes = nil
	file_blogs_blogs_service_proto_depIdxs = nil
}
