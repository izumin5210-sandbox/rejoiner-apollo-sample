// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package github_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
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

func (m *GetUserRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

type User struct {
	Login                   string               `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Id                      int64                `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	NodeId                  string               `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	AvatarUrl               string               `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	HtmlUrl                 string               `protobuf:"bytes,5,opt,name=html_url,json=htmlUrl,proto3" json:"html_url,omitempty"`
	GravatarId              string               `protobuf:"bytes,6,opt,name=gravatar_id,json=gravatarId,proto3" json:"gravatar_id,omitempty"`
	Name                    string               `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Company                 string               `protobuf:"bytes,8,opt,name=company,proto3" json:"company,omitempty"`
	Blog                    string               `protobuf:"bytes,9,opt,name=blog,proto3" json:"blog,omitempty"`
	Location                string               `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
	Email                   string               `protobuf:"bytes,11,opt,name=email,proto3" json:"email,omitempty"`
	Hireable                bool                 `protobuf:"varint,12,opt,name=hireable,proto3" json:"hireable,omitempty"`
	Bio                     string               `protobuf:"bytes,13,opt,name=bio,proto3" json:"bio,omitempty"`
	PublicRepos             uint32               `protobuf:"varint,14,opt,name=public_repos,json=publicRepos,proto3" json:"public_repos,omitempty"`
	PublicGists             uint32               `protobuf:"varint,15,opt,name=public_gists,json=publicGists,proto3" json:"public_gists,omitempty"`
	Followers               uint32               `protobuf:"varint,16,opt,name=followers,proto3" json:"followers,omitempty"`
	Following               uint32               `protobuf:"varint,17,opt,name=following,proto3" json:"following,omitempty"`
	CreatedAt               *timestamp.Timestamp `protobuf:"bytes,18,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt               *timestamp.Timestamp `protobuf:"bytes,19,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	SuspendedAt             *timestamp.Timestamp `protobuf:"bytes,20,opt,name=suspended_at,json=suspendedAt,proto3" json:"suspended_at,omitempty"`
	Type                    string               `protobuf:"bytes,21,opt,name=type,proto3" json:"type,omitempty"`
	SiteAdmin               bool                 `protobuf:"varint,22,opt,name=site_admin,json=siteAdmin,proto3" json:"site_admin,omitempty"`
	TotalPrivateRepos       uint32               `protobuf:"varint,23,opt,name=total_private_repos,json=totalPrivateRepos,proto3" json:"total_private_repos,omitempty"`
	OwnedPrivateRepos       uint32               `protobuf:"varint,24,opt,name=owned_private_repos,json=ownedPrivateRepos,proto3" json:"owned_private_repos,omitempty"`
	PrivateGists            uint32               `protobuf:"varint,25,opt,name=private_gists,json=privateGists,proto3" json:"private_gists,omitempty"`
	DiskUsage               uint32               `protobuf:"varint,26,opt,name=disk_usage,json=diskUsage,proto3" json:"disk_usage,omitempty"`
	Collaborators           uint32               `protobuf:"varint,27,opt,name=collaborators,proto3" json:"collaborators,omitempty"`
	TwoFactorAuthentication bool                 `protobuf:"varint,28,opt,name=two_factor_authentication,json=twoFactorAuthentication,proto3" json:"two_factor_authentication,omitempty"`
	Plan                    *Plan                `protobuf:"bytes,29,opt,name=plan,proto3" json:"plan,omitempty"`
	LdapDn                  string               `protobuf:"bytes,30,opt,name=ldap_dn,json=ldapDn,proto3" json:"ldap_dn,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
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

func (m *User) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *User) GetHtmlUrl() string {
	if m != nil {
		return m.HtmlUrl
	}
	return ""
}

func (m *User) GetGravatarId() string {
	if m != nil {
		return m.GravatarId
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetBlog() string {
	if m != nil {
		return m.Blog
	}
	return ""
}

func (m *User) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetHireable() bool {
	if m != nil {
		return m.Hireable
	}
	return false
}

func (m *User) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *User) GetPublicRepos() uint32 {
	if m != nil {
		return m.PublicRepos
	}
	return 0
}

func (m *User) GetPublicGists() uint32 {
	if m != nil {
		return m.PublicGists
	}
	return 0
}

func (m *User) GetFollowers() uint32 {
	if m != nil {
		return m.Followers
	}
	return 0
}

func (m *User) GetFollowing() uint32 {
	if m != nil {
		return m.Following
	}
	return 0
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetSuspendedAt() *timestamp.Timestamp {
	if m != nil {
		return m.SuspendedAt
	}
	return nil
}

func (m *User) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *User) GetSiteAdmin() bool {
	if m != nil {
		return m.SiteAdmin
	}
	return false
}

func (m *User) GetTotalPrivateRepos() uint32 {
	if m != nil {
		return m.TotalPrivateRepos
	}
	return 0
}

func (m *User) GetOwnedPrivateRepos() uint32 {
	if m != nil {
		return m.OwnedPrivateRepos
	}
	return 0
}

func (m *User) GetPrivateGists() uint32 {
	if m != nil {
		return m.PrivateGists
	}
	return 0
}

func (m *User) GetDiskUsage() uint32 {
	if m != nil {
		return m.DiskUsage
	}
	return 0
}

func (m *User) GetCollaborators() uint32 {
	if m != nil {
		return m.Collaborators
	}
	return 0
}

func (m *User) GetTwoFactorAuthentication() bool {
	if m != nil {
		return m.TwoFactorAuthentication
	}
	return false
}

func (m *User) GetPlan() *Plan {
	if m != nil {
		return m.Plan
	}
	return nil
}

func (m *User) GetLdapDn() string {
	if m != nil {
		return m.LdapDn
	}
	return ""
}

type ListUsersRequest struct {
	// Required. User IDs.
	Logins               []string `protobuf:"bytes,1,rep,name=logins,proto3" json:"logins,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetLogins() []string {
	if m != nil {
		return m.Logins
	}
	return nil
}

type ListUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "github.GetUserRequest")
	proto.RegisterType((*User)(nil), "github.User")
	proto.RegisterType((*ListUsersRequest)(nil), "github.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "github.ListUsersResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 770 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcf, 0x6f, 0xe4, 0x34,
	0x14, 0xc7, 0x37, 0xfd, 0x31, 0xd3, 0x79, 0x33, 0x2d, 0xad, 0xbb, 0xb4, 0xee, 0xb0, 0x65, 0x87,
	0x01, 0xa1, 0x11, 0x52, 0x13, 0xb6, 0x08, 0x21, 0x40, 0x1c, 0x66, 0x85, 0x58, 0xad, 0xb4, 0x87,
	0x2a, 0xd0, 0x0b, 0x97, 0xc8, 0x89, 0xdd, 0xd4, 0xe0, 0xd8, 0xc6, 0x76, 0x3a, 0x2c, 0x67, 0xfe,
	0x5d, 0xfe, 0x07, 0x64, 0x3b, 0x99, 0xce, 0xac, 0x56, 0xda, 0x53, 0xfc, 0xbe, 0xdf, 0xcf, 0x4b,
	0x62, 0xfb, 0xbd, 0x07, 0xd0, 0x5a, 0x66, 0x52, 0x6d, 0x94, 0x53, 0x68, 0x50, 0x73, 0x77, 0xdf,
	0x96, 0xd3, 0xe7, 0xb5, 0x52, 0xb5, 0x60, 0x59, 0x50, 0xcb, 0xf6, 0x2e, 0x73, 0xbc, 0x61, 0xd6,
	0x91, 0x46, 0x47, 0x70, 0x7a, 0x12, 0xc1, 0x4c, 0x0b, 0x22, 0xa3, 0x34, 0xff, 0x12, 0x8e, 0x5e,
	0x31, 0x77, 0x6b, 0x99, 0xc9, 0xd9, 0x5f, 0x2d, 0xb3, 0x0e, 0x3d, 0x85, 0x7d, 0xa1, 0x6a, 0x2e,
	0x71, 0x32, 0x4b, 0x16, 0xa3, 0x3c, 0x06, 0xf3, 0xff, 0x86, 0xb0, 0xe7, 0xa9, 0xf7, 0xdb, 0xe8,
	0x08, 0x76, 0x38, 0xc5, 0x3b, 0xb3, 0x64, 0xb1, 0x9b, 0xef, 0x70, 0x8a, 0xce, 0x61, 0x28, 0x15,
	0x65, 0x05, 0xa7, 0x78, 0x37, 0x70, 0x03, 0x1f, 0xbe, 0xa6, 0xe8, 0x12, 0x80, 0x3c, 0x10, 0x47,
	0x4c, 0xd1, 0x1a, 0x81, 0xf7, 0x82, 0x37, 0x8a, 0xca, 0xad, 0x11, 0xe8, 0x02, 0x0e, 0xee, 0x5d,
	0x23, 0x82, 0xb9, 0x1f, 0xcc, 0xa1, 0x8f, 0xbd, 0xf5, 0x1c, 0xc6, 0xb5, 0xe9, 0x72, 0x39, 0xc5,
	0x83, 0xe0, 0x42, 0x2f, 0xbd, 0xa6, 0x08, 0xc1, 0x9e, 0x24, 0x0d, 0xc3, 0xc3, 0xe0, 0x84, 0x35,
	0xc2, 0x30, 0xac, 0x54, 0xa3, 0x89, 0x7c, 0x8b, 0x0f, 0xe2, 0xeb, 0xba, 0xd0, 0xd3, 0xa5, 0x50,
	0x35, 0x1e, 0x45, 0xda, 0xaf, 0xd1, 0x14, 0x0e, 0x84, 0xaa, 0x88, 0xe3, 0x4a, 0x62, 0x08, 0xfa,
	0x3a, 0xf6, 0xfb, 0x66, 0x0d, 0xe1, 0x02, 0x8f, 0xe3, 0xbe, 0x43, 0xe0, 0x33, 0xee, 0xb9, 0x61,
	0xa4, 0x14, 0x0c, 0x4f, 0x66, 0xc9, 0xe2, 0x20, 0x5f, 0xc7, 0xe8, 0x18, 0x76, 0x4b, 0xae, 0xf0,
	0x61, 0xe0, 0xfd, 0x12, 0x7d, 0x06, 0x13, 0xdd, 0x96, 0x82, 0x57, 0x85, 0x61, 0x5a, 0x59, 0x7c,
	0x34, 0x4b, 0x16, 0x87, 0xf9, 0x38, 0x6a, 0xb9, 0x97, 0x36, 0x90, 0x9a, 0x5b, 0x67, 0xf1, 0x47,
	0x9b, 0xc8, 0x2b, 0x2f, 0xa1, 0x67, 0x30, 0xba, 0x53, 0x42, 0xa8, 0x15, 0x33, 0x16, 0x1f, 0x07,
	0xff, 0x51, 0x78, 0x74, 0xb9, 0xac, 0xf1, 0xc9, 0xa6, 0xcb, 0x65, 0x8d, 0xbe, 0x07, 0xa8, 0x0c,
	0x23, 0x8e, 0xd1, 0x82, 0x38, 0x8c, 0x66, 0xc9, 0x62, 0x7c, 0x3d, 0x4d, 0x63, 0xdd, 0xa4, 0x7d,
	0xdd, 0xa4, 0xbf, 0xf5, 0x75, 0x93, 0x8f, 0x3a, 0x7a, 0xe9, 0x7c, 0x6a, 0xab, 0x69, 0x9f, 0x7a,
	0xfa, 0xe1, 0xd4, 0x8e, 0x5e, 0x3a, 0xf4, 0x13, 0x4c, 0x6c, 0x6b, 0x35, 0x93, 0x34, 0x26, 0x3f,
	0xfd, 0x60, 0xf2, 0x78, 0xcd, 0x2f, 0x9d, 0xbf, 0x2a, 0xf7, 0x56, 0x33, 0xfc, 0x71, 0xbc, 0x2a,
	0xbf, 0xf6, 0x75, 0x64, 0xb9, 0x63, 0x05, 0xa1, 0x0d, 0x97, 0xf8, 0x2c, 0x1c, 0xfd, 0xc8, 0x2b,
	0x4b, 0x2f, 0xa0, 0x14, 0x4e, 0x9d, 0x72, 0x44, 0x14, 0xda, 0xf0, 0x07, 0xe2, 0x58, 0x77, 0xe0,
	0xe7, 0xe1, 0x3c, 0x4e, 0x82, 0x75, 0x13, 0x9d, 0x78, 0xec, 0x29, 0x9c, 0xaa, 0x95, 0x64, 0xf4,
	0x1d, 0x1e, 0x47, 0x3e, 0x58, 0x5b, 0xfc, 0xe7, 0x70, 0xd8, 0x93, 0xf1, 0x9e, 0x2e, 0x02, 0x39,
	0xe9, 0xc4, 0x78, 0x51, 0x97, 0x00, 0x94, 0xdb, 0x3f, 0x8b, 0xd6, 0x92, 0x9a, 0xe1, 0x69, 0xbc,
	0x0b, 0xaf, 0xdc, 0x7a, 0x01, 0x7d, 0x01, 0x87, 0x95, 0x12, 0x82, 0x94, 0xca, 0x10, 0xa7, 0x8c,
	0xc5, 0x9f, 0x04, 0x62, 0x5b, 0x44, 0x3f, 0xc0, 0x85, 0x5b, 0xa9, 0xe2, 0x8e, 0x54, 0x4e, 0x99,
	0x82, 0xb4, 0xee, 0x9e, 0x49, 0xc7, 0xbb, 0x22, 0x7d, 0x16, 0xf6, 0x7d, 0xee, 0x56, 0xea, 0x97,
	0xe0, 0x2f, 0xb7, 0x6c, 0x34, 0x83, 0x3d, 0xdf, 0xea, 0xf8, 0x32, 0x9c, 0xf7, 0x24, 0x8d, 0xed,
	0x9f, 0xde, 0x08, 0x22, 0xf3, 0xe0, 0xf8, 0x3e, 0x15, 0x94, 0xe8, 0x82, 0x4a, 0xfc, 0x69, 0xec,
	0x53, 0x1f, 0xfe, 0x2c, 0xe7, 0x5f, 0xc1, 0xf1, 0x1b, 0x6e, 0xc3, 0x60, 0xb0, 0xfd, 0x64, 0x38,
	0x83, 0x41, 0xe8, 0x76, 0x8b, 0x93, 0xd9, 0x6e, 0x60, 0x43, 0x34, 0xff, 0x0e, 0x4e, 0x36, 0x58,
	0xab, 0x95, 0xb4, 0x0c, 0xcd, 0x61, 0xdf, 0x8f, 0xa8, 0xc8, 0x6e, 0x7c, 0x3c, 0x8c, 0x9a, 0x68,
	0x5d, 0xff, 0x9b, 0xc0, 0xd8, 0xc7, 0xbf, 0x32, 0xf3, 0xc0, 0x2b, 0x86, 0x5e, 0xc0, 0xb0, 0x1b,
	0x46, 0xe8, 0xac, 0xe7, 0xb7, 0xa7, 0xd3, 0x74, 0xeb, 0x3d, 0xf3, 0x27, 0xe8, 0x25, 0x8c, 0xd6,
	0xdf, 0x46, 0xb8, 0x37, 0xdf, 0xfd, 0xf5, 0xe9, 0xc5, 0x7b, 0x9c, 0xf8, 0xa3, 0xf3, 0x27, 0x2f,
	0x15, 0x5c, 0x52, 0xf6, 0x90, 0xf2, 0x7f, 0xda, 0x86, 0xcb, 0x6f, 0xaf, 0x5f, 0x7c, 0x9d, 0x5a,
	0x22, 0x69, 0xa9, 0xfe, 0xee, 0x92, 0x6e, 0x92, 0xdf, 0xdf, 0x74, 0xe9, 0x95, 0x6a, 0xb2, 0x47,
	0xee, 0xaa, 0xe3, 0x32, 0xc3, 0xfe, 0x50, 0x5c, 0x32, 0x73, 0x45, 0xb4, 0x6f, 0xb5, 0x2b, 0x4b,
	0x1a, 0x2d, 0x58, 0x46, 0x34, 0xcf, 0x6a, 0x95, 0xc5, 0xdc, 0x1f, 0xe3, 0xa3, 0xd0, 0x65, 0x39,
	0x08, 0x15, 0xff, 0xcd, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xbe, 0xd9, 0xeb, 0xc5, 0x05,
	0x00, 0x00,
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
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/github.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/github.UserService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*User, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
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
		FullMethod: "/github.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.UserService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UserService_ListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
