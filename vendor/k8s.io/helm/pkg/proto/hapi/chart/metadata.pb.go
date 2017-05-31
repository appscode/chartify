// Code generated by protoc-gen-go.
// source: hapi/chart/metadata.proto
// DO NOT EDIT!

package chart

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Metadata_Engine int32

const (
	Metadata_UNKNOWN Metadata_Engine = 0
	Metadata_GOTPL   Metadata_Engine = 1
)

var Metadata_Engine_name = map[int32]string{
	0: "UNKNOWN",
	1: "GOTPL",
}
var Metadata_Engine_value = map[string]int32{
	"UNKNOWN": 0,
	"GOTPL":   1,
}

func (x Metadata_Engine) String() string {
	return proto.EnumName(Metadata_Engine_name, int32(x))
}
func (Metadata_Engine) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

// Maintainer describes a Chart maintainer.
type Maintainer struct {
	// Name is a user name or organization name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Email is an optional email address to contact the named maintainer
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *Maintainer) Reset()                    { *m = Maintainer{} }
func (m *Maintainer) String() string            { return proto.CompactTextString(m) }
func (*Maintainer) ProtoMessage()               {}
func (*Maintainer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Maintainer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Maintainer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// 	Metadata for a Chart file. This models the structure of a Chart.yaml file.
//
// 	Spec: https://k8s.io/helm/blob/master/docs/design/chart_format.md#the-chart-file
type Metadata struct {
	// The name of the chart
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The URL to a relevant project page, git repo, or contact person
	Home string `protobuf:"bytes,2,opt,name=home" json:"home,omitempty"`
	// Source is the URL to the source code of this chart
	Sources []string `protobuf:"bytes,3,rep,name=sources" json:"sources,omitempty"`
	// A SemVer 2 conformant version string of the chart
	Version string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	// A one-sentence description of the chart
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	// A list of string keywords
	Keywords []string `protobuf:"bytes,6,rep,name=keywords" json:"keywords,omitempty"`
	// A list of name and URL/email address combinations for the maintainer(s)
	Maintainers []*Maintainer `protobuf:"bytes,7,rep,name=maintainers" json:"maintainers,omitempty"`
	// The name of the template engine to use. Defaults to 'gotpl'.
	Engine string `protobuf:"bytes,8,opt,name=engine" json:"engine,omitempty"`
	// The URL to an icon file.
	Icon string `protobuf:"bytes,9,opt,name=icon" json:"icon,omitempty"`
	// The API Version of this chart.
	ApiVersion string `protobuf:"bytes,10,opt,name=apiVersion" json:"apiVersion,omitempty"`
	// The condition to check to enable chart
	Condition string `protobuf:"bytes,11,opt,name=condition" json:"condition,omitempty"`
	// The tags to check to enable chart
	Tags string `protobuf:"bytes,12,opt,name=tags" json:"tags,omitempty"`
	// The version of the application enclosed inside of this chart.
	AppVersion string `protobuf:"bytes,13,opt,name=appVersion" json:"appVersion,omitempty"`
	// Whether or not this chart is deprecated
	Deprecated bool `protobuf:"varint,14,opt,name=deprecated" json:"deprecated,omitempty"`
	// TillerVersion is a SemVer constraints on what version of Tiller is required.
	// See SemVer ranges here: https://github.com/Masterminds/semver#basic-comparisons
	TillerVersion string `protobuf:"bytes,15,opt,name=tillerVersion" json:"tillerVersion,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metadata) GetHome() string {
	if m != nil {
		return m.Home
	}
	return ""
}

func (m *Metadata) GetSources() []string {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Metadata) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Metadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Metadata) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

func (m *Metadata) GetMaintainers() []*Maintainer {
	if m != nil {
		return m.Maintainers
	}
	return nil
}

func (m *Metadata) GetEngine() string {
	if m != nil {
		return m.Engine
	}
	return ""
}

func (m *Metadata) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Metadata) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *Metadata) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

func (m *Metadata) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *Metadata) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *Metadata) GetDeprecated() bool {
	if m != nil {
		return m.Deprecated
	}
	return false
}

func (m *Metadata) GetTillerVersion() string {
	if m != nil {
		return m.TillerVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*Maintainer)(nil), "hapi.chart.Maintainer")
	proto.RegisterType((*Metadata)(nil), "hapi.chart.Metadata")
	proto.RegisterEnum("hapi.chart.Metadata_Engine", Metadata_Engine_name, Metadata_Engine_value)
}

func init() { proto.RegisterFile("hapi/chart/metadata.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0x5f, 0x4b, 0xe3, 0x40,
	0x14, 0xc5, 0x37, 0x9b, 0xe6, 0xdf, 0xcd, 0x76, 0xb7, 0x0c, 0x4b, 0x99, 0x5d, 0x44, 0x42, 0xf1,
	0x21, 0x4f, 0x29, 0x28, 0x88, 0xcf, 0x82, 0xf8, 0xa0, 0x6d, 0x25, 0xf8, 0x07, 0x7c, 0x1b, 0x93,
	0x4b, 0x3b, 0xd8, 0xcc, 0x84, 0xc9, 0xa8, 0xf8, 0x7d, 0xfd, 0x20, 0x32, 0x93, 0xa4, 0x8d, 0xe0,
	0xdb, 0x3d, 0xe7, 0xe4, 0xfe, 0xc2, 0xb9, 0x0c, 0xfc, 0xdb, 0xb0, 0x9a, 0xcf, 0x8b, 0x0d, 0x53,
	0x7a, 0x5e, 0xa1, 0x66, 0x25, 0xd3, 0x2c, 0xab, 0x95, 0xd4, 0x92, 0x80, 0x89, 0x32, 0x1b, 0xcd,
	0x4e, 0x01, 0x16, 0x8c, 0x0b, 0xcd, 0xb8, 0x40, 0x45, 0x08, 0x8c, 0x04, 0xab, 0x90, 0x3a, 0x89,
	0x93, 0x46, 0xb9, 0x9d, 0xc9, 0x5f, 0xf0, 0xb0, 0x62, 0x7c, 0x4b, 0x7f, 0x5a, 0xb3, 0x15, 0xb3,
	0x0f, 0x17, 0xc2, 0x45, 0x87, 0xfd, 0x76, 0x8d, 0xc0, 0x68, 0x23, 0x2b, 0xec, 0xb6, 0xec, 0x4c,
	0x28, 0x04, 0x8d, 0x7c, 0x51, 0x05, 0x36, 0xd4, 0x4d, 0xdc, 0x34, 0xca, 0x7b, 0x69, 0x92, 0x57,
	0x54, 0x0d, 0x97, 0x82, 0x8e, 0xec, 0x42, 0x2f, 0x49, 0x02, 0x71, 0x89, 0x4d, 0xa1, 0x78, 0xad,
	0x4d, 0xea, 0xd9, 0x74, 0x68, 0x91, 0xff, 0x10, 0x3e, 0xe3, 0xfb, 0x9b, 0x54, 0x65, 0x43, 0x7d,
	0x8b, 0xdd, 0x69, 0x72, 0x06, 0x71, 0xb5, 0xab, 0xd7, 0xd0, 0x20, 0x71, 0xd3, 0xf8, 0x78, 0x9a,
	0xed, 0x0f, 0x90, 0xed, 0xdb, 0xe7, 0xc3, 0x4f, 0xc9, 0x14, 0x7c, 0x14, 0x6b, 0x2e, 0x90, 0x86,
	0xf6, 0x97, 0x9d, 0x32, 0xbd, 0x78, 0x21, 0x05, 0x8d, 0xda, 0x5e, 0x66, 0x26, 0x87, 0x00, 0xac,
	0xe6, 0xf7, 0x5d, 0x01, 0xb0, 0xc9, 0xc0, 0x21, 0x07, 0x10, 0x15, 0x52, 0x94, 0xdc, 0x36, 0x88,
	0x6d, 0xbc, 0x37, 0x0c, 0x51, 0xb3, 0x75, 0x43, 0x7f, 0xb5, 0x44, 0x33, 0xb7, 0xc4, 0xba, 0x27,
	0x8e, 0x7b, 0x62, 0xef, 0x98, 0xbc, 0xc4, 0x5a, 0x61, 0xc1, 0x34, 0x96, 0xf4, 0x77, 0xe2, 0xa4,
	0x61, 0x3e, 0x70, 0xc8, 0x11, 0x8c, 0x35, 0xdf, 0x6e, 0x51, 0xf5, 0x88, 0x3f, 0x16, 0xf1, 0xd5,
	0x9c, 0x25, 0xe0, 0x5f, 0xb4, 0xad, 0x62, 0x08, 0xee, 0x96, 0x57, 0xcb, 0xd5, 0xc3, 0x72, 0xf2,
	0x83, 0x44, 0xe0, 0x5d, 0xae, 0x6e, 0x6f, 0xae, 0x27, 0xce, 0x79, 0xf0, 0xe8, 0xd9, 0x33, 0x3d,
	0xf9, 0xf6, 0xe9, 0x9c, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0xea, 0xb5, 0x4c, 0xbe, 0x57, 0x02,
	0x00, 0x00,
}
