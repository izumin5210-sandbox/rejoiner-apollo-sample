// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package qiita_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetUserRequest struct {
	// Required. User ID.
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type User struct {
	// Optional. self-description.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. Facebook ID.
	FacebookId string `protobuf:"bytes,2,opt,name=facebook_id,json=facebookId,proto3" json:"facebook_id,omitempty"`
	// Required. Followees count.
	FolloweesCount uint32 `protobuf:"varint,3,opt,name=followees_count,json=followeesCount,proto3" json:"followees_count,omitempty"`
	// Required. Followers count.
	FollowersCount uint32 `protobuf:"varint,4,opt,name=followers_count,json=followersCount,proto3" json:"followers_count,omitempty"`
	// Optional. GitHub ID.
	GithubLoginName string `protobuf:"bytes,5,opt,name=github_login_name,json=githubLoginName,proto3" json:"github_login_name,omitempty"`
	// Required. User ID.
	Id string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	// Required. How many items a user posted on qiita.com (Items on Qiita Team are not included).
	ItemsCount uint32 `protobuf:"varint,7,opt,name=items_count,json=itemsCount,proto3" json:"items_count,omitempty"`
	// Optional. LinkedIn ID.
	LinkedinId string `protobuf:"bytes,8,opt,name=linkedin_id,json=linkedinId,proto3" json:"linkedin_id,omitempty"`
	// Optional. Location
	Location string `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	// Optional. Customized user name.
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Organization which a user belongs to.
	Organization string `protobuf:"bytes,11,opt,name=organization,proto3" json:"organization,omitempty"`
	// Required. Unique integer ID.
	PermanentId uint32 `protobuf:"varint,12,opt,name=permanent_id,json=permanentId,proto3" json:"permanent_id,omitempty"`
	// Required. Profile image URL.
	ProfileImageUrl string `protobuf:"bytes,13,opt,name=profile_image_url,json=profileImageUrl,proto3" json:"profile_image_url,omitempty"`
	// Required. A flag whether this user is configured as team-only.
	TeamOnly bool `protobuf:"varint,14,opt,name=team_only,json=teamOnly,proto3" json:"team_only,omitempty"`
	// Optional. Twitter screen name.
	TwitterScreenName string `protobuf:"bytes,15,opt,name=twitter_screen_name,json=twitterScreenName,proto3" json:"twitter_screen_name,omitempty"`
	// Optional.  Website URL.
	WebsiteUrl           string   `protobuf:"bytes,16,opt,name=website_url,json=websiteUrl,proto3" json:"website_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *User) GetFacebookId() string {
	if m != nil {
		return m.FacebookId
	}
	return ""
}

func (m *User) GetFolloweesCount() uint32 {
	if m != nil {
		return m.FolloweesCount
	}
	return 0
}

func (m *User) GetFollowersCount() uint32 {
	if m != nil {
		return m.FollowersCount
	}
	return 0
}

func (m *User) GetGithubLoginName() string {
	if m != nil {
		return m.GithubLoginName
	}
	return ""
}

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetItemsCount() uint32 {
	if m != nil {
		return m.ItemsCount
	}
	return 0
}

func (m *User) GetLinkedinId() string {
	if m != nil {
		return m.LinkedinId
	}
	return ""
}

func (m *User) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func (m *User) GetPermanentId() uint32 {
	if m != nil {
		return m.PermanentId
	}
	return 0
}

func (m *User) GetProfileImageUrl() string {
	if m != nil {
		return m.ProfileImageUrl
	}
	return ""
}

func (m *User) GetTeamOnly() bool {
	if m != nil {
		return m.TeamOnly
	}
	return false
}

func (m *User) GetTwitterScreenName() string {
	if m != nil {
		return m.TwitterScreenName
	}
	return ""
}

func (m *User) GetWebsiteUrl() string {
	if m != nil {
		return m.WebsiteUrl
	}
	return ""
}

type ListFollowersRequest struct {
	// Required. User ID.
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFollowersRequest) Reset()         { *m = ListFollowersRequest{} }
func (m *ListFollowersRequest) String() string { return proto.CompactTextString(m) }
func (*ListFollowersRequest) ProtoMessage()    {}
func (*ListFollowersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *ListFollowersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFollowersRequest.Unmarshal(m, b)
}
func (m *ListFollowersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFollowersRequest.Marshal(b, m, deterministic)
}
func (m *ListFollowersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFollowersRequest.Merge(m, src)
}
func (m *ListFollowersRequest) XXX_Size() int {
	return xxx_messageInfo_ListFollowersRequest.Size(m)
}
func (m *ListFollowersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFollowersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFollowersRequest proto.InternalMessageInfo

func (m *ListFollowersRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ListFollowersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFollowersResponse) Reset()         { *m = ListFollowersResponse{} }
func (m *ListFollowersResponse) String() string { return proto.CompactTextString(m) }
func (*ListFollowersResponse) ProtoMessage()    {}
func (*ListFollowersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *ListFollowersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFollowersResponse.Unmarshal(m, b)
}
func (m *ListFollowersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFollowersResponse.Marshal(b, m, deterministic)
}
func (m *ListFollowersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFollowersResponse.Merge(m, src)
}
func (m *ListFollowersResponse) XXX_Size() int {
	return xxx_messageInfo_ListFollowersResponse.Size(m)
}
func (m *ListFollowersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFollowersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFollowersResponse proto.InternalMessageInfo

func (m *ListFollowersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type ListFolloweesRequest struct {
	// Required. User ID.
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFolloweesRequest) Reset()         { *m = ListFolloweesRequest{} }
func (m *ListFolloweesRequest) String() string { return proto.CompactTextString(m) }
func (*ListFolloweesRequest) ProtoMessage()    {}
func (*ListFolloweesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *ListFolloweesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFolloweesRequest.Unmarshal(m, b)
}
func (m *ListFolloweesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFolloweesRequest.Marshal(b, m, deterministic)
}
func (m *ListFolloweesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFolloweesRequest.Merge(m, src)
}
func (m *ListFolloweesRequest) XXX_Size() int {
	return xxx_messageInfo_ListFolloweesRequest.Size(m)
}
func (m *ListFolloweesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFolloweesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFolloweesRequest proto.InternalMessageInfo

func (m *ListFolloweesRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ListFolloweesResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFolloweesResponse) Reset()         { *m = ListFolloweesResponse{} }
func (m *ListFolloweesResponse) String() string { return proto.CompactTextString(m) }
func (*ListFolloweesResponse) ProtoMessage()    {}
func (*ListFolloweesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *ListFolloweesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFolloweesResponse.Unmarshal(m, b)
}
func (m *ListFolloweesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFolloweesResponse.Marshal(b, m, deterministic)
}
func (m *ListFolloweesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFolloweesResponse.Merge(m, src)
}
func (m *ListFolloweesResponse) XXX_Size() int {
	return xxx_messageInfo_ListFolloweesResponse.Size(m)
}
func (m *ListFolloweesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFolloweesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFolloweesResponse proto.InternalMessageInfo

func (m *ListFolloweesResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type ListStockersRequest struct {
	// Required. Item ID.
	ItemId               string   `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListStockersRequest) Reset()         { *m = ListStockersRequest{} }
func (m *ListStockersRequest) String() string { return proto.CompactTextString(m) }
func (*ListStockersRequest) ProtoMessage()    {}
func (*ListStockersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *ListStockersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStockersRequest.Unmarshal(m, b)
}
func (m *ListStockersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStockersRequest.Marshal(b, m, deterministic)
}
func (m *ListStockersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStockersRequest.Merge(m, src)
}
func (m *ListStockersRequest) XXX_Size() int {
	return xxx_messageInfo_ListStockersRequest.Size(m)
}
func (m *ListStockersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStockersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListStockersRequest proto.InternalMessageInfo

func (m *ListStockersRequest) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

type ListStockersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListStockersResponse) Reset()         { *m = ListStockersResponse{} }
func (m *ListStockersResponse) String() string { return proto.CompactTextString(m) }
func (*ListStockersResponse) ProtoMessage()    {}
func (*ListStockersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *ListStockersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStockersResponse.Unmarshal(m, b)
}
func (m *ListStockersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStockersResponse.Marshal(b, m, deterministic)
}
func (m *ListStockersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStockersResponse.Merge(m, src)
}
func (m *ListStockersResponse) XXX_Size() int {
	return xxx_messageInfo_ListStockersResponse.Size(m)
}
func (m *ListStockersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStockersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListStockersResponse proto.InternalMessageInfo

func (m *ListStockersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "qiita.GetUserRequest")
	proto.RegisterType((*User)(nil), "qiita.User")
	proto.RegisterType((*ListFollowersRequest)(nil), "qiita.ListFollowersRequest")
	proto.RegisterType((*ListFollowersResponse)(nil), "qiita.ListFollowersResponse")
	proto.RegisterType((*ListFolloweesRequest)(nil), "qiita.ListFolloweesRequest")
	proto.RegisterType((*ListFolloweesResponse)(nil), "qiita.ListFolloweesResponse")
	proto.RegisterType((*ListStockersRequest)(nil), "qiita.ListStockersRequest")
	proto.RegisterType((*ListStockersResponse)(nil), "qiita.ListStockersResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xa6, 0x5d, 0xb7, 0x75, 0x2f, 0x5d, 0xc7, 0x3c, 0x26, 0xa2, 0x6e, 0xd2, 0xba, 0x5e, 0x28,
	0x48, 0x4b, 0x60, 0x88, 0x03, 0x70, 0x03, 0x09, 0x14, 0x69, 0x02, 0xd4, 0x69, 0x17, 0x2e, 0x91,
	0x9b, 0xbc, 0x15, 0xd3, 0xc4, 0xce, 0x6c, 0x77, 0x63, 0xfb, 0x3f, 0x5c, 0xf8, 0x95, 0xc8, 0x76,
	0xb2, 0xa5, 0xa5, 0x02, 0x76, 0xa9, 0xea, 0xef, 0xfb, 0xde, 0xf7, 0xbd, 0xf7, 0xe2, 0x04, 0x60,
	0xa6, 0x50, 0x06, 0x85, 0x14, 0x5a, 0x90, 0xd5, 0x0b, 0xc6, 0x34, 0x1d, 0x3c, 0x85, 0xee, 0x47,
	0xd4, 0x67, 0x0a, 0xe5, 0x08, 0x2f, 0x66, 0xa8, 0x34, 0x79, 0x0c, 0xeb, 0x46, 0x16, 0xb3, 0xd4,
	0x6f, 0xf4, 0x1b, 0xc3, 0x8d, 0xd1, 0x9a, 0x39, 0x46, 0xe9, 0xe0, 0x57, 0x0b, 0x5a, 0x46, 0x48,
	0xfa, 0xe0, 0xa5, 0xa8, 0x12, 0xc9, 0x0a, 0xcd, 0x04, 0x2f, 0x55, 0x75, 0x88, 0x1c, 0x80, 0x77,
	0x4e, 0x13, 0x1c, 0x0b, 0x31, 0x35, 0x3e, 0x4d, 0xab, 0x80, 0x0a, 0x8a, 0x52, 0xf2, 0x04, 0xb6,
	0xce, 0x45, 0x96, 0x89, 0x2b, 0x44, 0x15, 0x27, 0x62, 0xc6, 0xb5, 0xbf, 0xd2, 0x6f, 0x0c, 0x37,
	0x47, 0xdd, 0x5b, 0xf8, 0xbd, 0x41, 0x6b, 0x42, 0x59, 0x09, 0x5b, 0x73, 0x42, 0x59, 0x0a, 0x9f,
	0xc1, 0xf6, 0x84, 0xe9, 0x6f, 0xb3, 0x71, 0x9c, 0x89, 0x09, 0xe3, 0x31, 0xa7, 0x39, 0xfa, 0xab,
	0x36, 0x78, 0xcb, 0x11, 0x27, 0x06, 0xff, 0x44, 0x73, 0x24, 0x5d, 0x68, 0xb2, 0xd4, 0x5f, 0xb3,
	0x64, 0x93, 0xa5, 0xa6, 0x5d, 0xa6, 0x31, 0xaf, 0x02, 0xd6, 0x6d, 0x00, 0x58, 0xc8, 0x99, 0x1f,
	0x80, 0x97, 0x31, 0x3e, 0xc5, 0x94, 0x71, 0x33, 0x4f, 0xdb, 0xcd, 0x53, 0x41, 0x51, 0x4a, 0x7a,
	0xd0, 0xce, 0x44, 0x42, 0xed, 0x3e, 0x36, 0x2c, 0x7b, 0x7b, 0x26, 0x04, 0x5a, 0xb6, 0x19, 0xb0,
	0xb8, 0xfd, 0x4f, 0x06, 0xd0, 0x11, 0x72, 0x42, 0x39, 0xbb, 0x71, 0x35, 0x9e, 0xe5, 0xe6, 0x30,
	0x72, 0x08, 0x9d, 0x02, 0x65, 0x4e, 0x39, 0x72, 0x6d, 0x52, 0x3b, 0xb6, 0x2d, 0xef, 0x16, 0x8b,
	0x52, 0x33, 0x74, 0x21, 0xc5, 0x39, 0xcb, 0x30, 0x66, 0x39, 0x9d, 0x60, 0x3c, 0x93, 0x99, 0xbf,
	0xe9, 0x86, 0x2e, 0x89, 0xc8, 0xe0, 0x67, 0x32, 0x23, 0x7b, 0xb0, 0xa1, 0x91, 0xe6, 0xb1, 0xe0,
	0xd9, 0xb5, 0xdf, 0xed, 0x37, 0x86, 0xed, 0x51, 0xdb, 0x00, 0x9f, 0x79, 0x76, 0x4d, 0x02, 0xd8,
	0xd1, 0x57, 0x4c, 0x6b, 0x94, 0xb1, 0x4a, 0x24, 0x62, 0xb9, 0xbf, 0x2d, 0x6b, 0xb5, 0x5d, 0x52,
	0xa7, 0x96, 0xb1, 0x1b, 0x3c, 0x00, 0xef, 0x0a, 0xc7, 0x8a, 0x69, 0x17, 0xf9, 0xd0, 0x2d, 0xa4,
	0x84, 0xce, 0x64, 0x36, 0x08, 0xe1, 0xd1, 0x09, 0x53, 0xfa, 0x43, 0xf5, 0x90, 0xfe, 0x79, 0xbb,
	0xde, 0xc0, 0xee, 0x42, 0x81, 0x2a, 0x04, 0x57, 0x48, 0x0e, 0x61, 0xd5, 0x48, 0x94, 0xdf, 0xe8,
	0xaf, 0x0c, 0xbd, 0x63, 0x2f, 0xb0, 0x17, 0x37, 0xb0, 0x57, 0xd6, 0x31, 0x0b, 0x61, 0x78, 0xdf,
	0x30, 0xbc, 0x57, 0x58, 0x00, 0x3b, 0xa6, 0xf6, 0x54, 0x8b, 0x64, 0x3a, 0x3f, 0x98, 0xb9, 0x30,
	0xb5, 0x2c, 0x73, 0x8c, 0xd2, 0xc1, 0x6b, 0xd7, 0xdc, 0x9d, 0xfe, 0xbf, 0xa3, 0x8e, 0x7f, 0x36,
	0xc1, 0x33, 0xe7, 0x53, 0x94, 0x97, 0x2c, 0x41, 0x12, 0xc2, 0x7a, 0xf9, 0xb2, 0x92, 0xdd, 0x52,
	0x3e, 0xff, 0xf2, 0xf6, 0xea, 0x2e, 0x83, 0x07, 0xe4, 0x04, 0x36, 0xe7, 0x96, 0x4a, 0xf6, 0x4a,
	0x7e, 0xd9, 0xb3, 0xe9, 0xed, 0x2f, 0x27, 0x5d, 0xbf, 0x7f, 0xb8, 0xe1, 0x52, 0x37, 0xfc, 0x9b,
	0x1b, 0xd6, 0xdd, 0x22, 0xe8, 0xd4, 0xf7, 0x42, 0x7a, 0x35, 0xfd, 0xc2, 0x72, 0x7b, 0x7b, 0x4b,
	0xb9, 0xca, 0xea, 0xdd, 0x14, 0xf6, 0x53, 0xbc, 0x0c, 0xd8, 0xcd, 0x2c, 0x67, 0xfc, 0xd5, 0xf1,
	0x8b, 0xe7, 0x81, 0xa2, 0x3c, 0x1d, 0x8b, 0x1f, 0xae, 0xec, 0x4b, 0xe3, 0x6b, 0xe4, 0xea, 0x13,
	0x91, 0x87, 0x77, 0xaa, 0xa3, 0x52, 0x15, 0x4a, 0xfc, 0x2e, 0x18, 0x47, 0x79, 0x44, 0x0b, 0xd3,
	0xe8, 0x91, 0xa2, 0x79, 0x91, 0x61, 0x48, 0x0b, 0x16, 0x4e, 0x44, 0x68, 0x4b, 0xdf, 0xda, 0xdf,
	0xb8, 0x18, 0x8f, 0xd7, 0xec, 0xf7, 0xf3, 0xe5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xbb,
	0x9c, 0xf3, 0x4d, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	ListFollowers(ctx context.Context, in *ListFollowersRequest, opts ...grpc.CallOption) (*ListFollowersResponse, error)
	ListFollowees(ctx context.Context, in *ListFolloweesRequest, opts ...grpc.CallOption) (*ListFolloweesResponse, error)
	ListStockers(ctx context.Context, in *ListStockersRequest, opts ...grpc.CallOption) (*ListStockersResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/qiita.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListFollowers(ctx context.Context, in *ListFollowersRequest, opts ...grpc.CallOption) (*ListFollowersResponse, error) {
	out := new(ListFollowersResponse)
	err := c.cc.Invoke(ctx, "/qiita.UserService/ListFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListFollowees(ctx context.Context, in *ListFolloweesRequest, opts ...grpc.CallOption) (*ListFolloweesResponse, error) {
	out := new(ListFolloweesResponse)
	err := c.cc.Invoke(ctx, "/qiita.UserService/ListFollowees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListStockers(ctx context.Context, in *ListStockersRequest, opts ...grpc.CallOption) (*ListStockersResponse, error) {
	out := new(ListStockersResponse)
	err := c.cc.Invoke(ctx, "/qiita.UserService/ListStockers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*User, error)
	ListFollowers(context.Context, *ListFollowersRequest) (*ListFollowersResponse, error)
	ListFollowees(context.Context, *ListFolloweesRequest) (*ListFolloweesResponse, error)
	ListStockers(context.Context, *ListStockersRequest) (*ListStockersResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) ListFollowers(ctx context.Context, req *ListFollowersRequest) (*ListFollowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFollowers not implemented")
}
func (*UnimplementedUserServiceServer) ListFollowees(ctx context.Context, req *ListFolloweesRequest) (*ListFolloweesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFollowees not implemented")
}
func (*UnimplementedUserServiceServer) ListStockers(ctx context.Context, req *ListStockersRequest) (*ListStockersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStockers not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qiita.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFollowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qiita.UserService/ListFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListFollowers(ctx, req.(*ListFollowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListFollowees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFolloweesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListFollowees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qiita.UserService/ListFollowees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListFollowees(ctx, req.(*ListFolloweesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListStockers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStockersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListStockers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qiita.UserService/ListStockers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListStockers(ctx, req.(*ListStockersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qiita.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "ListFollowers",
			Handler:    _UserService_ListFollowers_Handler,
		},
		{
			MethodName: "ListFollowees",
			Handler:    _UserService_ListFollowees_Handler,
		},
		{
			MethodName: "ListStockers",
			Handler:    _UserService_ListStockers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
