// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: bilibili.proto

package bilibili

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

type DataType int32

const (
	DataType_Video      DataType = 0
	DataType_Season     DataType = 1
	DataType_Media      DataType = 2
	DataType_Audio      DataType = 3
	DataType_Article    DataType = 4
	DataType_Collection DataType = 5
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "Video",
		1: "Season",
		2: "Media",
		3: "Audio",
		4: "Article",
		5: "Collection",
	}
	DataType_value = map[string]int32{
		"Video":      0,
		"Season":     1,
		"Media":      2,
		"Audio":      3,
		"Article":    4,
		"Collection": 5,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_bilibili_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_proto_rawDescGZIP(), []int{0}
}

type Region int32

const (
	Region_CN   Region = 0
	Region_INTL Region = 1
	Region_HK   Region = 2
	Region_TW   Region = 3
	Region_TH   Region = 4
)

// Enum value maps for Region.
var (
	Region_name = map[int32]string{
		0: "CN",
		1: "INTL",
		2: "HK",
		3: "TW",
		4: "TH",
	}
	Region_value = map[string]int32{
		"CN":   0,
		"INTL": 1,
		"HK":   2,
		"TW":   3,
		"TH":   4,
	}
)

func (x Region) Enum() *Region {
	p := new(Region)
	*p = x
	return p
}

func (x Region) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Region) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_proto_enumTypes[1].Descriptor()
}

func (Region) Type() protoreflect.EnumType {
	return &file_bilibili_proto_enumTypes[1]
}

func (x Region) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Region.Descriptor instead.
func (Region) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_proto_rawDescGZIP(), []int{1}
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg         string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Type        DataType `protobuf:"varint,3,opt,name=Type,proto3,enum=bilibili.DataType" json:"Type,omitempty"`
	Picture     *string  `protobuf:"bytes,4,opt,name=Picture,proto3,oneof" json:"Picture,omitempty"`
	BV          *string  `protobuf:"bytes,5,opt,name=BV,proto3,oneof" json:"BV,omitempty"`
	AV          *int64   `protobuf:"varint,6,opt,name=AV,proto3,oneof" json:"AV,omitempty"`
	Title       *string  `protobuf:"bytes,7,opt,name=Title,proto3,oneof" json:"Title,omitempty"`
	Author      *string  `protobuf:"bytes,8,opt,name=Author,proto3,oneof" json:"Author,omitempty"`
	CreateTime  *int64   `protobuf:"varint,9,opt,name=CreateTime,proto3,oneof" json:"CreateTime,omitempty"`
	PublicTime  *int64   `protobuf:"varint,10,opt,name=PublicTime,proto3,oneof" json:"PublicTime,omitempty"`
	Duration    *int64   `protobuf:"varint,11,opt,name=Duration,proto3,oneof" json:"Duration,omitempty"`
	Description *string  `protobuf:"bytes,12,opt,name=Description,proto3,oneof" json:"Description,omitempty"`
	Dynamic     *string  `protobuf:"bytes,13,opt,name=Dynamic,proto3,oneof" json:"Dynamic,omitempty"`
	//Season and Episode
	Evaluate                *string `protobuf:"bytes,14,opt,name=Evaluate,proto3,oneof" json:"Evaluate,omitempty"`
	Area                    *string `protobuf:"bytes,15,opt,name=Area,proto3,oneof" json:"Area,omitempty"`
	SerialStatusDescription *string `protobuf:"bytes,16,opt,name=SerialStatusDescription,proto3,oneof" json:"SerialStatusDescription,omitempty"`
	ShareURL                *string `protobuf:"bytes,17,opt,name=ShareURL,proto3,oneof" json:"ShareURL,omitempty"`
	MediaType               *string `protobuf:"bytes,18,opt,name=MediaType,proto3,oneof" json:"MediaType,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_bilibili_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Info) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Info) GetType() DataType {
	if x != nil {
		return x.Type
	}
	return DataType_Video
}

func (x *Info) GetPicture() string {
	if x != nil && x.Picture != nil {
		return *x.Picture
	}
	return ""
}

func (x *Info) GetBV() string {
	if x != nil && x.BV != nil {
		return *x.BV
	}
	return ""
}

func (x *Info) GetAV() int64 {
	if x != nil && x.AV != nil {
		return *x.AV
	}
	return 0
}

func (x *Info) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Info) GetAuthor() string {
	if x != nil && x.Author != nil {
		return *x.Author
	}
	return ""
}

func (x *Info) GetCreateTime() int64 {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return 0
}

func (x *Info) GetPublicTime() int64 {
	if x != nil && x.PublicTime != nil {
		return *x.PublicTime
	}
	return 0
}

func (x *Info) GetDuration() int64 {
	if x != nil && x.Duration != nil {
		return *x.Duration
	}
	return 0
}

func (x *Info) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Info) GetDynamic() string {
	if x != nil && x.Dynamic != nil {
		return *x.Dynamic
	}
	return ""
}

func (x *Info) GetEvaluate() string {
	if x != nil && x.Evaluate != nil {
		return *x.Evaluate
	}
	return ""
}

func (x *Info) GetArea() string {
	if x != nil && x.Area != nil {
		return *x.Area
	}
	return ""
}

func (x *Info) GetSerialStatusDescription() string {
	if x != nil && x.SerialStatusDescription != nil {
		return *x.SerialStatusDescription
	}
	return ""
}

func (x *Info) GetShareURL() string {
	if x != nil && x.ShareURL != nil {
		return *x.ShareURL
	}
	return ""
}

func (x *Info) GetMediaType() string {
	if x != nil && x.MediaType != nil {
		return *x.MediaType
	}
	return ""
}

type Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL             string    `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	ID              *string   `protobuf:"bytes,2,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	CID             *string   `protobuf:"bytes,6,opt,name=CID,proto3,oneof" json:"CID,omitempty"`
	CheckCollection bool      `protobuf:"varint,3,opt,name=CheckCollection,proto3" json:"CheckCollection,omitempty"`
	Type            *DataType `protobuf:"varint,4,opt,name=Type,proto3,enum=bilibili.DataType,oneof" json:"Type,omitempty"`
	Region          *Region   `protobuf:"varint,5,opt,name=Region,proto3,enum=bilibili.Region,oneof" json:"Region,omitempty"`
}

func (x *Param) Reset() {
	*x = Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Param.ProtoReflect.Descriptor instead.
func (*Param) Descriptor() ([]byte, []int) {
	return file_bilibili_proto_rawDescGZIP(), []int{1}
}

func (x *Param) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *Param) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

func (x *Param) GetCID() string {
	if x != nil && x.CID != nil {
		return *x.CID
	}
	return ""
}

func (x *Param) GetCheckCollection() bool {
	if x != nil {
		return x.CheckCollection
	}
	return false
}

func (x *Param) GetType() DataType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return DataType_Video
}

func (x *Param) GetRegion() Region {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return Region_CN
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code            int32        `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg             string       `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Type            DataType     `protobuf:"varint,3,opt,name=Type,proto3,enum=bilibili.DataType" json:"Type,omitempty"`
	ID              string       `protobuf:"bytes,4,opt,name=ID,proto3" json:"ID,omitempty"`
	Author          *string      `protobuf:"bytes,6,opt,name=Author,proto3,oneof" json:"Author,omitempty"`
	IsEnd           *bool        `protobuf:"varint,7,opt,name=IsEnd,proto3,oneof" json:"IsEnd,omitempty"`
	Detail          []*QueryInfo `protobuf:"bytes,8,rep,name=Detail,proto3" json:"Detail,omitempty"`
	CollectionTitle *string      `protobuf:"bytes,5,opt,name=CollectionTitle,proto3,oneof" json:"CollectionTitle,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_bilibili_proto_rawDescGZIP(), []int{2}
}

func (x *Query) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Query) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Query) GetType() DataType {
	if x != nil {
		return x.Type
	}
	return DataType_Video
}

func (x *Query) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Query) GetAuthor() string {
	if x != nil && x.Author != nil {
		return *x.Author
	}
	return ""
}

func (x *Query) GetIsEnd() bool {
	if x != nil && x.IsEnd != nil {
		return *x.IsEnd
	}
	return false
}

func (x *Query) GetDetail() []*QueryInfo {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *Query) GetCollectionTitle() string {
	if x != nil && x.CollectionTitle != nil {
		return *x.CollectionTitle
	}
	return ""
}

type QueryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index  int32   `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	ID     string  `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	BVID   *string `protobuf:"bytes,3,opt,name=BVID,proto3,oneof" json:"BVID,omitempty"`
	Author *string `protobuf:"bytes,4,opt,name=Author,proto3,oneof" json:"Author,omitempty"`
	Region *Region `protobuf:"varint,5,opt,name=Region,proto3,enum=bilibili.Region,oneof" json:"Region,omitempty"`
}

func (x *QueryInfo) Reset() {
	*x = QueryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryInfo) ProtoMessage() {}

func (x *QueryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryInfo.ProtoReflect.Descriptor instead.
func (*QueryInfo) Descriptor() ([]byte, []int) {
	return file_bilibili_proto_rawDescGZIP(), []int{3}
}

func (x *QueryInfo) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *QueryInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *QueryInfo) GetBVID() string {
	if x != nil && x.BVID != nil {
		return *x.BVID
	}
	return ""
}

func (x *QueryInfo) GetAuthor() string {
	if x != nil && x.Author != nil {
		return *x.Author
	}
	return ""
}

func (x *QueryInfo) GetRegion() Region {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return Region_CN
}

type DownloadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32       `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg    string      `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Detail []*PartInfo `protobuf:"bytes,3,rep,name=Detail,proto3" json:"Detail,omitempty"`
	Type   DataType    `protobuf:"varint,4,opt,name=Type,proto3,enum=bilibili.DataType" json:"Type,omitempty"`
}

func (x *DownloadInfo) Reset() {
	*x = DownloadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadInfo) ProtoMessage() {}

func (x *DownloadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadInfo.ProtoReflect.Descriptor instead.
func (*DownloadInfo) Descriptor() ([]byte, []int) {
	return file_bilibili_proto_rawDescGZIP(), []int{4}
}

func (x *DownloadInfo) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DownloadInfo) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DownloadInfo) GetDetail() []*PartInfo {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *DownloadInfo) GetType() DataType {
	if x != nil {
		return x.Type
	}
	return DataType_Video
}

type PartInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index        int32       `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	ID           string      `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"` //CID or EPID
	CID          string      `protobuf:"bytes,3,opt,name=CID,proto3" json:"CID,omitempty"`
	Title        string      `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`       // Video Title or Episode Title
	SubTitle     string      `protobuf:"bytes,5,opt,name=SubTitle,proto3" json:"SubTitle,omitempty"` // Part Title or Episode Title
	VideoQuality string      `protobuf:"bytes,6,opt,name=VideoQuality,proto3" json:"VideoQuality,omitempty"`
	VideoURL     string      `protobuf:"bytes,7,opt,name=VideoURL,proto3" json:"VideoURL,omitempty"`
	AudioURL     string      `protobuf:"bytes,8,opt,name=AudioURL,proto3" json:"AudioURL,omitempty"`
	Subtitles    []*Subtitle `protobuf:"bytes,9,rep,name=Subtitles,proto3" json:"Subtitles,omitempty"`
}

func (x *PartInfo) Reset() {
	*x = PartInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartInfo) ProtoMessage() {}

func (x *PartInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartInfo.ProtoReflect.Descriptor instead.
func (*PartInfo) Descriptor() ([]byte, []int) {
	return file_bilibili_proto_rawDescGZIP(), []int{5}
}

func (x *PartInfo) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *PartInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *PartInfo) GetCID() string {
	if x != nil {
		return x.CID
	}
	return ""
}

func (x *PartInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PartInfo) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *PartInfo) GetVideoQuality() string {
	if x != nil {
		return x.VideoQuality
	}
	return ""
}

func (x *PartInfo) GetVideoURL() string {
	if x != nil {
		return x.VideoURL
	}
	return ""
}

func (x *PartInfo) GetAudioURL() string {
	if x != nil {
		return x.AudioURL
	}
	return ""
}

func (x *PartInfo) GetSubtitles() []*Subtitle {
	if x != nil {
		return x.Subtitles
	}
	return nil
}

type Subtitle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index       int32  `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Locale      string `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	LocaleText  string `protobuf:"bytes,3,opt,name=localeText,proto3" json:"localeText,omitempty"`
	SubtitleUrl string `protobuf:"bytes,4,opt,name=subtitleUrl,proto3" json:"subtitleUrl,omitempty"`
}

func (x *Subtitle) Reset() {
	*x = Subtitle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subtitle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subtitle) ProtoMessage() {}

func (x *Subtitle) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subtitle.ProtoReflect.Descriptor instead.
func (*Subtitle) Descriptor() ([]byte, []int) {
	return file_bilibili_proto_rawDescGZIP(), []int{6}
}

func (x *Subtitle) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Subtitle) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Subtitle) GetLocaleText() string {
	if x != nil {
		return x.LocaleText
	}
	return ""
}

func (x *Subtitle) GetSubtitleUrl() string {
	if x != nil {
		return x.SubtitleUrl
	}
	return ""
}

var File_bilibili_proto protoreflect.FileDescriptor

var file_bilibili_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x22, 0x86, 0x06, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x07, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x13, 0x0a, 0x02, 0x42, 0x56, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x02,
	0x42, 0x56, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x41, 0x56, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x02, 0x52, 0x02, 0x41, 0x56, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x06, 0x52, 0x0a, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x48, 0x07,
	0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x08, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x07, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a, 0x52, 0x08, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0b, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a,
	0x17, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0c,
	0x52, 0x17, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0d,
	0x52, 0x08, 0x53, 0x68, 0x61, 0x72, 0x65, 0x55, 0x52, 0x4c, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a,
	0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x0e, 0x52, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x42, 0x56, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x41, 0x56, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x44,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x41, 0x72, 0x65, 0x61, 0x42, 0x1a, 0x0a, 0x18,
	0x5f, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x55, 0x52, 0x4c, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x22, 0xee, 0x01, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12,
	0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x43, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x03, 0x43, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0f, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x48, 0x02, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x2d, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x48, 0x03, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x43, 0x49, 0x44,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x22, 0xa2, 0x02, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x06,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x49, 0x73, 0x45,
	0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x05, 0x49, 0x73, 0x45, 0x6e,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x2d, 0x0a, 0x0f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0f, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x49, 0x73, 0x45, 0x6e, 0x64, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x09, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x17, 0x0a,
	0x04, 0x42, 0x56, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x42,
	0x56, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x48, 0x02, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x42, 0x56, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x22, 0x88, 0x01, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x82, 0x02, 0x0a,
	0x08, 0x50, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x10, 0x0a, 0x03, 0x43, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x75, 0x62, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x51, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x55, 0x52, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x55, 0x52, 0x4c, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x55, 0x52, 0x4c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x55, 0x52, 0x4c, 0x12,
	0x30, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x53, 0x75,
	0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x09, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x73, 0x22, 0x7a, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x65, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x2a, 0x54, 0x0a,
	0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x10, 0x05, 0x2a, 0x32, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x0a,
	0x02, 0x43, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x54, 0x4c, 0x10, 0x01, 0x12,
	0x06, 0x0a, 0x02, 0x48, 0x4b, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x54, 0x57, 0x10, 0x03, 0x12,
	0x06, 0x0a, 0x02, 0x54, 0x48, 0x10, 0x04, 0x32, 0xa7, 0x01, 0x0a, 0x08, 0x42, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x12, 0x33, 0x0a, 0x0f, 0x44, 0x6f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0f, 0x2e, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x16, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x0e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_proto_rawDescOnce sync.Once
	file_bilibili_proto_rawDescData = file_bilibili_proto_rawDesc
)

func file_bilibili_proto_rawDescGZIP() []byte {
	file_bilibili_proto_rawDescOnce.Do(func() {
		file_bilibili_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_proto_rawDescData)
	})
	return file_bilibili_proto_rawDescData
}

var file_bilibili_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bilibili_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bilibili_proto_goTypes = []interface{}{
	(DataType)(0),        // 0: bilibili.DataType
	(Region)(0),          // 1: bilibili.Region
	(*Info)(nil),         // 2: bilibili.Info
	(*Param)(nil),        // 3: bilibili.Param
	(*Query)(nil),        // 4: bilibili.Query
	(*QueryInfo)(nil),    // 5: bilibili.QueryInfo
	(*DownloadInfo)(nil), // 6: bilibili.DownloadInfo
	(*PartInfo)(nil),     // 7: bilibili.PartInfo
	(*Subtitle)(nil),     // 8: bilibili.Subtitle
}
var file_bilibili_proto_depIdxs = []int32{
	0,  // 0: bilibili.Info.Type:type_name -> bilibili.DataType
	0,  // 1: bilibili.Param.Type:type_name -> bilibili.DataType
	1,  // 2: bilibili.Param.Region:type_name -> bilibili.Region
	0,  // 3: bilibili.Query.Type:type_name -> bilibili.DataType
	5,  // 4: bilibili.Query.Detail:type_name -> bilibili.QueryInfo
	1,  // 5: bilibili.QueryInfo.Region:type_name -> bilibili.Region
	7,  // 6: bilibili.DownloadInfo.Detail:type_name -> bilibili.PartInfo
	0,  // 7: bilibili.DownloadInfo.Type:type_name -> bilibili.DataType
	8,  // 8: bilibili.PartInfo.Subtitles:type_name -> bilibili.Subtitle
	3,  // 9: bilibili.Bilibili.DoDownloadQuery:input_type -> bilibili.Param
	3,  // 10: bilibili.Bilibili.GetDownloadInfo:input_type -> bilibili.Param
	3,  // 11: bilibili.Bilibili.GetInfo:input_type -> bilibili.Param
	4,  // 12: bilibili.Bilibili.DoDownloadQuery:output_type -> bilibili.Query
	6,  // 13: bilibili.Bilibili.GetDownloadInfo:output_type -> bilibili.DownloadInfo
	2,  // 14: bilibili.Bilibili.GetInfo:output_type -> bilibili.Info
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_bilibili_proto_init() }
func file_bilibili_proto_init() {
	if File_bilibili_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_bilibili_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Param); i {
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
		file_bilibili_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_bilibili_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryInfo); i {
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
		file_bilibili_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadInfo); i {
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
		file_bilibili_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartInfo); i {
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
		file_bilibili_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subtitle); i {
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
	file_bilibili_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_bilibili_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_bilibili_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_bilibili_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bilibili_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_proto_goTypes,
		DependencyIndexes: file_bilibili_proto_depIdxs,
		EnumInfos:         file_bilibili_proto_enumTypes,
		MessageInfos:      file_bilibili_proto_msgTypes,
	}.Build()
	File_bilibili_proto = out.File
	file_bilibili_proto_rawDesc = nil
	file_bilibili_proto_goTypes = nil
	file_bilibili_proto_depIdxs = nil
}
