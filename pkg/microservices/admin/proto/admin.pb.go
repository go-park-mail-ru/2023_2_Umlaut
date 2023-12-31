// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: admin.proto

package proto

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

type Feedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId     int32  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Rating     int32  `protobuf:"varint,3,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Liked      string `protobuf:"bytes,4,opt,name=Liked,proto3" json:"Liked,omitempty"`
	NeedFix    string `protobuf:"bytes,5,opt,name=NeedFix,proto3" json:"NeedFix,omitempty"`
	CommentFix string `protobuf:"bytes,6,opt,name=CommentFix,proto3" json:"CommentFix,omitempty"`
	Comment    string `protobuf:"bytes,7,opt,name=Comment,proto3" json:"Comment,omitempty"`
	Show       bool   `protobuf:"varint,8,opt,name=Show,proto3" json:"Show,omitempty"`
}

func (x *Feedback) Reset() {
	*x = Feedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feedback) ProtoMessage() {}

func (x *Feedback) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feedback.ProtoReflect.Descriptor instead.
func (*Feedback) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0}
}

func (x *Feedback) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Feedback) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Feedback) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Feedback) GetLiked() string {
	if x != nil {
		return x.Liked
	}
	return ""
}

func (x *Feedback) GetNeedFix() string {
	if x != nil {
		return x.NeedFix
	}
	return ""
}

func (x *Feedback) GetCommentFix() string {
	if x != nil {
		return x.CommentFix
	}
	return ""
}

func (x *Feedback) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Feedback) GetShow() bool {
	if x != nil {
		return x.Show
	}
	return false
}

type Recommendation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId int32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Rating int32 `protobuf:"varint,3,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Show   bool  `protobuf:"varint,4,opt,name=Show,proto3" json:"Show,omitempty"`
}

func (x *Recommendation) Reset() {
	*x = Recommendation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recommendation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recommendation) ProtoMessage() {}

func (x *Recommendation) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recommendation.ProtoReflect.Descriptor instead.
func (*Recommendation) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{1}
}

func (x *Recommendation) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Recommendation) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Recommendation) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Recommendation) GetShow() bool {
	if x != nil {
		return x.Show
	}
	return false
}

type FeedbackStatistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvgRating   float32       `protobuf:"fixed32,1,opt,name=AvgRating,proto3" json:"AvgRating,omitempty"`
	RatingCount []int32       `protobuf:"varint,2,rep,packed,name=RatingCount,proto3" json:"RatingCount,omitempty"`
	LikedMap    []*LikedMap   `protobuf:"bytes,3,rep,name=LikedMap,proto3" json:"LikedMap,omitempty"`
	NeedFixMap  []*NeedFixMap `protobuf:"bytes,4,rep,name=NeedFixMap,proto3" json:"NeedFixMap,omitempty"`
	Comments    []string      `protobuf:"bytes,5,rep,name=Comments,proto3" json:"Comments,omitempty"`
}

func (x *FeedbackStatistic) Reset() {
	*x = FeedbackStatistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedbackStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedbackStatistic) ProtoMessage() {}

func (x *FeedbackStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedbackStatistic.ProtoReflect.Descriptor instead.
func (*FeedbackStatistic) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{2}
}

func (x *FeedbackStatistic) GetAvgRating() float32 {
	if x != nil {
		return x.AvgRating
	}
	return 0
}

func (x *FeedbackStatistic) GetRatingCount() []int32 {
	if x != nil {
		return x.RatingCount
	}
	return nil
}

func (x *FeedbackStatistic) GetLikedMap() []*LikedMap {
	if x != nil {
		return x.LikedMap
	}
	return nil
}

func (x *FeedbackStatistic) GetNeedFixMap() []*NeedFixMap {
	if x != nil {
		return x.NeedFixMap
	}
	return nil
}

func (x *FeedbackStatistic) GetComments() []string {
	if x != nil {
		return x.Comments
	}
	return nil
}

type AdminEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdminEmpty) Reset() {
	*x = AdminEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminEmpty) ProtoMessage() {}

func (x *AdminEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminEmpty.ProtoReflect.Descriptor instead.
func (*AdminEmpty) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{3}
}

type LikedMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Liked string `protobuf:"bytes,1,opt,name=Liked,proto3" json:"Liked,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *LikedMap) Reset() {
	*x = LikedMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikedMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikedMap) ProtoMessage() {}

func (x *LikedMap) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikedMap.ProtoReflect.Descriptor instead.
func (*LikedMap) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{4}
}

func (x *LikedMap) GetLiked() string {
	if x != nil {
		return x.Liked
	}
	return ""
}

func (x *LikedMap) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type NeedFixMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NeedFix       string         `protobuf:"bytes,1,opt,name=NeedFix,proto3" json:"NeedFix,omitempty"`
	NeedFixObject *NeedFixObject `protobuf:"bytes,2,opt,name=NeedFixObject,proto3" json:"NeedFixObject,omitempty"`
}

func (x *NeedFixMap) Reset() {
	*x = NeedFixMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NeedFixMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NeedFixMap) ProtoMessage() {}

func (x *NeedFixMap) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NeedFixMap.ProtoReflect.Descriptor instead.
func (*NeedFixMap) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{5}
}

func (x *NeedFixMap) GetNeedFix() string {
	if x != nil {
		return x.NeedFix
	}
	return ""
}

func (x *NeedFixMap) GetNeedFixObject() *NeedFixObject {
	if x != nil {
		return x.NeedFixObject
	}
	return nil
}

type NeedFixObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count      int32    `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	CommentFix []string `protobuf:"bytes,2,rep,name=CommentFix,proto3" json:"CommentFix,omitempty"`
}

func (x *NeedFixObject) Reset() {
	*x = NeedFixObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NeedFixObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NeedFixObject) ProtoMessage() {}

func (x *NeedFixObject) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NeedFixObject.ProtoReflect.Descriptor instead.
func (*NeedFixObject) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{6}
}

func (x *NeedFixObject) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *NeedFixObject) GetCommentFix() []string {
	if x != nil {
		return x.CommentFix
	}
	return nil
}

type RecommendationStatistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvgRecommend   float32 `protobuf:"fixed32,1,opt,name=AvgRecommend,proto3" json:"AvgRecommend,omitempty"`
	NPS            float32 `protobuf:"fixed32,2,opt,name=NPS,proto3" json:"NPS,omitempty"`
	RecommendCount []int32 `protobuf:"varint,3,rep,packed,name=RecommendCount,proto3" json:"RecommendCount,omitempty"`
}

func (x *RecommendationStatistic) Reset() {
	*x = RecommendationStatistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendationStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationStatistic) ProtoMessage() {}

func (x *RecommendationStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationStatistic.ProtoReflect.Descriptor instead.
func (*RecommendationStatistic) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{7}
}

func (x *RecommendationStatistic) GetAvgRecommend() float32 {
	if x != nil {
		return x.AvgRecommend
	}
	return 0
}

func (x *RecommendationStatistic) GetNPS() float32 {
	if x != nil {
		return x.NPS
	}
	return 0
}

func (x *RecommendationStatistic) GetRecommendCount() []int32 {
	if x != nil {
		return x.RecommendCount
	}
	return nil
}

type Complaint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ReporterUserId  int32                  `protobuf:"varint,2,opt,name=ReporterUserId,proto3" json:"ReporterUserId,omitempty"`
	ReportedUserId  int32                  `protobuf:"varint,3,opt,name=ReportedUserId,proto3" json:"ReportedUserId,omitempty"`
	ComplaintTypeId int32                  `protobuf:"varint,4,opt,name=ComplaintTypeId,proto3" json:"ComplaintTypeId,omitempty"`
	ComplaintText   string                 `protobuf:"bytes,5,opt,name=ComplaintText,proto3" json:"ComplaintText,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Complaint) Reset() {
	*x = Complaint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Complaint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Complaint) ProtoMessage() {}

func (x *Complaint) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Complaint.ProtoReflect.Descriptor instead.
func (*Complaint) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{8}
}

func (x *Complaint) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Complaint) GetReporterUserId() int32 {
	if x != nil {
		return x.ReporterUserId
	}
	return 0
}

func (x *Complaint) GetReportedUserId() int32 {
	if x != nil {
		return x.ReportedUserId
	}
	return 0
}

func (x *Complaint) GetComplaintTypeId() int32 {
	if x != nil {
		return x.ComplaintTypeId
	}
	return 0
}

func (x *Complaint) GetComplaintText() string {
	if x != nil {
		return x.ComplaintText
	}
	return ""
}

func (x *Complaint) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_admin_proto protoreflect.FileDescriptor

var file_admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x65, 0x64,
	0x46, 0x69, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x65, 0x65, 0x64, 0x46,
	0x69, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x78,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x46,
	0x69, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x53, 0x68, 0x6f, 0x77, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x53, 0x68, 0x6f, 0x77,
	0x22, 0x64, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x53, 0x68, 0x6f, 0x77, 0x22, 0xcf, 0x01, 0x0a, 0x11, 0x46, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09,
	0x41, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x41, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0b, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x08,
	0x4c, 0x69, 0x6b, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x52,
	0x08, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x12, 0x31, 0x0a, 0x0a, 0x4e, 0x65, 0x65,
	0x64, 0x46, 0x69, 0x78, 0x4d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x65, 0x64, 0x46, 0x69, 0x78, 0x4d, 0x61, 0x70,
	0x52, 0x0a, 0x4e, 0x65, 0x65, 0x64, 0x46, 0x69, 0x78, 0x4d, 0x61, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x36, 0x0a, 0x08, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x4d,
	0x61, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x62,
	0x0a, 0x0a, 0x4e, 0x65, 0x65, 0x64, 0x46, 0x69, 0x78, 0x4d, 0x61, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x4e, 0x65, 0x65, 0x64, 0x46, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e,
	0x65, 0x65, 0x64, 0x46, 0x69, 0x78, 0x12, 0x3a, 0x0a, 0x0d, 0x4e, 0x65, 0x65, 0x64, 0x46, 0x69,
	0x78, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x65, 0x64, 0x46, 0x69, 0x78, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x0d, 0x4e, 0x65, 0x65, 0x64, 0x46, 0x69, 0x78, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x45, 0x0a, 0x0d, 0x4e, 0x65, 0x65, 0x64, 0x46, 0x69, 0x78, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x78, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x78, 0x22, 0x77, 0x0a, 0x17, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x76, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x41, 0x76, 0x67, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x50, 0x53, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x4e, 0x50, 0x53, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xf5, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x8e, 0x04, 0x0a, 0x05, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x45, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74,
	0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x61, 0x69, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x1a, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_proto_rawDescOnce sync.Once
	file_admin_proto_rawDescData = file_admin_proto_rawDesc
)

func file_admin_proto_rawDescGZIP() []byte {
	file_admin_proto_rawDescOnce.Do(func() {
		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_proto_rawDescData)
	})
	return file_admin_proto_rawDescData
}

var file_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_admin_proto_goTypes = []interface{}{
	(*Feedback)(nil),                // 0: proto.Feedback
	(*Recommendation)(nil),          // 1: proto.Recommendation
	(*FeedbackStatistic)(nil),       // 2: proto.FeedbackStatistic
	(*AdminEmpty)(nil),              // 3: proto.AdminEmpty
	(*LikedMap)(nil),                // 4: proto.LikedMap
	(*NeedFixMap)(nil),              // 5: proto.NeedFixMap
	(*NeedFixObject)(nil),           // 6: proto.NeedFixObject
	(*RecommendationStatistic)(nil), // 7: proto.RecommendationStatistic
	(*Complaint)(nil),               // 8: proto.Complaint
	(*timestamppb.Timestamp)(nil),   // 9: google.protobuf.Timestamp
}
var file_admin_proto_depIdxs = []int32{
	4,  // 0: proto.FeedbackStatistic.LikedMap:type_name -> proto.LikedMap
	5,  // 1: proto.FeedbackStatistic.NeedFixMap:type_name -> proto.NeedFixMap
	6,  // 2: proto.NeedFixMap.NeedFixObject:type_name -> proto.NeedFixObject
	9,  // 3: proto.Complaint.CreatedAt:type_name -> google.protobuf.Timestamp
	3,  // 4: proto.Admin.GetFeedbackStatistic:input_type -> proto.AdminEmpty
	3,  // 5: proto.Admin.GetRecommendationStatistic:input_type -> proto.AdminEmpty
	0,  // 6: proto.Admin.CreateFeedback:input_type -> proto.Feedback
	1,  // 7: proto.Admin.CreateRecommendation:input_type -> proto.Recommendation
	1,  // 8: proto.Admin.CreateFeedFeedback:input_type -> proto.Recommendation
	3,  // 9: proto.Admin.GetNextComplaint:input_type -> proto.AdminEmpty
	8,  // 10: proto.Admin.DeleteComplaint:input_type -> proto.Complaint
	8,  // 11: proto.Admin.AcceptComplaint:input_type -> proto.Complaint
	2,  // 12: proto.Admin.GetFeedbackStatistic:output_type -> proto.FeedbackStatistic
	7,  // 13: proto.Admin.GetRecommendationStatistic:output_type -> proto.RecommendationStatistic
	3,  // 14: proto.Admin.CreateFeedback:output_type -> proto.AdminEmpty
	3,  // 15: proto.Admin.CreateRecommendation:output_type -> proto.AdminEmpty
	3,  // 16: proto.Admin.CreateFeedFeedback:output_type -> proto.AdminEmpty
	8,  // 17: proto.Admin.GetNextComplaint:output_type -> proto.Complaint
	3,  // 18: proto.Admin.DeleteComplaint:output_type -> proto.AdminEmpty
	3,  // 19: proto.Admin.AcceptComplaint:output_type -> proto.AdminEmpty
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feedback); i {
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
		file_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recommendation); i {
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
		file_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedbackStatistic); i {
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
		file_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminEmpty); i {
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
		file_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikedMap); i {
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
		file_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NeedFixMap); i {
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
		file_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NeedFixObject); i {
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
		file_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendationStatistic); i {
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
		file_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Complaint); i {
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
			RawDescriptor: file_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
		MessageInfos:      file_admin_proto_msgTypes,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_rawDesc = nil
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}
