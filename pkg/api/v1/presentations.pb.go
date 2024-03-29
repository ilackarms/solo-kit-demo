// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ilackarms/solo-kit-demo/api/v1/presentations.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//Presentation defines the Presentation Template that should be created.
//Create, delete, and update these resources to create update and delete new presentations.
type Presentation struct {
	// required solo-kit field
	// the presence of this field turns this resource into a solo-kit resource
	Metadata core.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata"`
	// optional status, allows us to report back to the user
	// any validation / status
	Status core.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status"`
	// title of the presentation
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// title of the content slide
	ContentSlideTitle string `protobuf:"bytes,4,opt,name=content_slide_title,json=contentSlideTitle,proto3" json:"content_slide_title,omitempty"`
	// content of the content slide
	ContentSlideContent  string   `protobuf:"bytes,5,opt,name=content_slide_content,json=contentSlideContent,proto3" json:"content_slide_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Presentation) Reset()         { *m = Presentation{} }
func (m *Presentation) String() string { return proto.CompactTextString(m) }
func (*Presentation) ProtoMessage()    {}
func (*Presentation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3044ac96d3b325c5, []int{0}
}
func (m *Presentation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Presentation.Unmarshal(m, b)
}
func (m *Presentation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Presentation.Marshal(b, m, deterministic)
}
func (m *Presentation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Presentation.Merge(m, src)
}
func (m *Presentation) XXX_Size() int {
	return xxx_messageInfo_Presentation.Size(m)
}
func (m *Presentation) XXX_DiscardUnknown() {
	xxx_messageInfo_Presentation.DiscardUnknown(m)
}

var xxx_messageInfo_Presentation proto.InternalMessageInfo

func (m *Presentation) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Presentation) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Presentation) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Presentation) GetContentSlideTitle() string {
	if m != nil {
		return m.ContentSlideTitle
	}
	return ""
}

func (m *Presentation) GetContentSlideContent() string {
	if m != nil {
		return m.ContentSlideContent
	}
	return ""
}

//
//GoogleApiKey is a secret containing an API Key for talking to Google Slides API
type GoogleApiKey struct {
	// required solo-kit field
	// the presence of this field turns this resource into a solo-kit resource
	Metadata core.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata"`
	// credentials.json encoded as a base-64 string
	CredentialsBase64    string   `protobuf:"bytes,2,opt,name=credentials_base64,json=credentialsBase64,proto3" json:"credentials_base64,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleApiKey) Reset()         { *m = GoogleApiKey{} }
func (m *GoogleApiKey) String() string { return proto.CompactTextString(m) }
func (*GoogleApiKey) ProtoMessage()    {}
func (*GoogleApiKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_3044ac96d3b325c5, []int{1}
}
func (m *GoogleApiKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleApiKey.Unmarshal(m, b)
}
func (m *GoogleApiKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleApiKey.Marshal(b, m, deterministic)
}
func (m *GoogleApiKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleApiKey.Merge(m, src)
}
func (m *GoogleApiKey) XXX_Size() int {
	return xxx_messageInfo_GoogleApiKey.Size(m)
}
func (m *GoogleApiKey) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleApiKey.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleApiKey proto.InternalMessageInfo

func (m *GoogleApiKey) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *GoogleApiKey) GetCredentialsBase64() string {
	if m != nil {
		return m.CredentialsBase64
	}
	return ""
}

func init() {
	proto.RegisterType((*Presentation)(nil), "demos.solo.io.Presentation")
	proto.RegisterType((*GoogleApiKey)(nil), "demos.solo.io.GoogleApiKey")
}

func init() {
	proto.RegisterFile("github.com/ilackarms/solo-kit-demo/api/v1/presentations.proto", fileDescriptor_3044ac96d3b325c5)
}

var fileDescriptor_3044ac96d3b325c5 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0x6f, 0xb9, 0x85, 0xdc, 0x3b, 0x17, 0x72, 0x63, 0x41, 0xd3, 0x60, 0xa2, 0x84, 0x15,
	0x9b, 0xb6, 0x01, 0x8c, 0x31, 0x26, 0x2e, 0xc4, 0x85, 0x0b, 0x63, 0x62, 0xc0, 0x95, 0x1b, 0x32,
	0xb4, 0x27, 0x75, 0xd2, 0x3f, 0xa7, 0xe9, 0x1c, 0x4c, 0xd8, 0xf2, 0x0a, 0xfa, 0x10, 0x3e, 0x8a,
	0x4f, 0xe1, 0xc2, 0x37, 0xe0, 0x0d, 0x4c, 0xa7, 0x2d, 0xc2, 0x0e, 0xe3, 0x6e, 0x4e, 0xbe, 0xdf,
	0x37, 0xe7, 0x3b, 0x73, 0x86, 0x5d, 0xf8, 0x82, 0x1e, 0xe7, 0x33, 0xdb, 0xc5, 0xc8, 0x11, 0x21,
	0x77, 0x03, 0x9e, 0x46, 0xd2, 0x91, 0x18, 0xa2, 0x15, 0x08, 0xb2, 0x3c, 0x88, 0xd0, 0xe1, 0x89,
	0x70, 0x9e, 0xfa, 0x4e, 0x92, 0x82, 0x84, 0x98, 0x38, 0x09, 0x8c, 0xa5, 0x9d, 0xa4, 0x48, 0x68,
	0x34, 0x32, 0x40, 0xda, 0x19, 0x6f, 0x0b, 0x6c, 0xb7, 0x7c, 0xf4, 0x51, 0x29, 0x4e, 0x76, 0xca,
	0xa1, 0x76, 0x7f, 0xa3, 0x87, 0xba, 0x59, 0xe0, 0xba, 0x43, 0x79, 0x79, 0x04, 0xc4, 0x3d, 0x4e,
	0xbc, 0xb0, 0x38, 0x3b, 0x58, 0x24, 0x71, 0x9a, 0xcb, 0x6f, 0xf4, 0x28, 0xeb, 0xdc, 0xd2, 0x7d,
	0xae, 0xb0, 0xfa, 0xdd, 0xc6, 0x4c, 0xc6, 0x19, 0xfb, 0x53, 0xc6, 0x30, 0xb5, 0x8e, 0xd6, 0xfb,
	0x37, 0x38, 0xb0, 0x5d, 0x4c, 0xa1, 0x1c, 0xcf, 0xbe, 0x2d, 0xd4, 0x91, 0xfe, 0xf6, 0x7e, 0xfc,
	0x6b, 0xbc, 0xa6, 0x8d, 0x01, 0xab, 0xe5, 0x69, 0xcc, 0x8a, 0xf2, 0xb5, 0xb6, 0x7d, 0x13, 0xa5,
	0x15, 0xae, 0x82, 0x34, 0x5a, 0xac, 0x4a, 0x82, 0x42, 0x30, 0x7f, 0x77, 0xb4, 0xde, 0xdf, 0x71,
	0x5e, 0x18, 0x36, 0x6b, 0xba, 0x18, 0x13, 0xc4, 0x34, 0x95, 0xa1, 0xf0, 0x60, 0x9a, 0x33, 0xba,
	0x62, 0xf6, 0x0a, 0x69, 0x92, 0x29, 0xf7, 0x8a, 0x1f, 0xb0, 0xfd, 0x6d, 0xbe, 0xa8, 0xcc, 0xaa,
	0x72, 0x34, 0x37, 0x1d, 0x57, 0xf9, 0xf9, 0xfc, 0x70, 0xb9, 0xd2, 0x75, 0x56, 0x49, 0xd2, 0xe5,
	0x4a, 0xff, 0x6f, 0x34, 0xb6, 0xf6, 0xda, 0x7d, 0xd1, 0x58, 0xfd, 0x1a, 0xd1, 0x0f, 0xe1, 0x32,
	0x11, 0x37, 0xb0, 0xf8, 0xc1, 0xab, 0x58, 0xcc, 0x70, 0x53, 0xf0, 0x20, 0x26, 0xc1, 0x43, 0x39,
	0x9d, 0x71, 0x09, 0xa7, 0x27, 0xea, 0x85, 0xb2, 0x51, 0xbe, 0x94, 0x91, 0x12, 0xca, 0x58, 0x7e,
	0x90, 0xc7, 0xf2, 0x55, 0x08, 0x9e, 0x88, 0x00, 0x16, 0x72, 0x34, 0x7c, 0xfd, 0x38, 0xd2, 0x1e,
	0xac, 0x1d, 0x7e, 0x6b, 0x12, 0xf8, 0xc5, 0xc2, 0x67, 0x35, 0xb5, 0xe8, 0xe1, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xec, 0xc0, 0x62, 0x55, 0xe5, 0x02, 0x00, 0x00,
}

func (this *Presentation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Presentation)
	if !ok {
		that2, ok := that.(Presentation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.ContentSlideTitle != that1.ContentSlideTitle {
		return false
	}
	if this.ContentSlideContent != that1.ContentSlideContent {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GoogleApiKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GoogleApiKey)
	if !ok {
		that2, ok := that.(GoogleApiKey)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if this.CredentialsBase64 != that1.CredentialsBase64 {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
