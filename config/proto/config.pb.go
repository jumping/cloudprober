// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/config/proto/config.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "github.com/google/cloudprober/probes/proto"
import proto3 "github.com/google/cloudprober/servers/proto"
import proto2 "github.com/google/cloudprober/surfacers/proto"
import proto6 "github.com/google/cloudprober/targets/proto"
import proto4 "github.com/google/cloudprober/targets/rds/server/proto"
import proto5 "github.com/google/cloudprober/targets/rtc/rtcreporter/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProberConfig struct {
	// Probes to run.
	Probe []*proto1.ProbeDef `protobuf:"bytes,1,rep,name=probe" json:"probe,omitempty"`
	// Surfacers are used to export probe results for further processing.
	// If no surfacer is configured, a prometheus and a file surfacer are
	// initialized:
	//  - Prometheus makes probe results available at http://<host>:9313/metrics.
	//  - File surfacer writes results to stdout.
	//
	// You can disable default surfacers (in case you want no surfacer at all), by
	// adding the following to your config:
	//   surfacer {}
	Surfacer []*proto2.SurfacerDef `protobuf:"bytes,2,rep,name=surfacer" json:"surfacer,omitempty"`
	// Servers to run inside cloudprober. These servers can serve as targets for
	// other probes.
	Server []*proto3.ServerDef `protobuf:"bytes,3,rep,name=server" json:"server,omitempty"`
	// Resource discovery server
	RdsServer *proto4.ServerConf `protobuf:"bytes,95,opt,name=rds_server,json=rdsServer" json:"rds_server,omitempty"`
	// Port for the default HTTP server. This port is also used for prometheus
	// exporter (URL /metrics). Default port is 9313. If not specified in the
	// config, default port can be overridden by the environment variable
	// CLOUDPROBER_PORT.
	Port *int32 `protobuf:"varint,96,opt,name=port" json:"port,omitempty"`
	// Host for the default HTTP server. Default listens on all addresses. If not
	// specified in the config, default port can be overridden by the environment
	// variable CLOUDPROBER_HOST.
	Host *string `protobuf:"bytes,101,opt,name=host" json:"host,omitempty"`
	// Probes are staggered across time to avoid executing all of them at the
	// same time. This behavior can be disabled by setting the following option
	// to true.
	DisableJitter *bool `protobuf:"varint,102,opt,name=disable_jitter,json=disableJitter,def=0" json:"disable_jitter,omitempty"`
	// How often to export system variables. To learn more about system variables:
	// http://godoc.org/github.com/google/cloudprober/sysvars.
	SysvarsIntervalMsec *int32 `protobuf:"varint,97,opt,name=sysvars_interval_msec,json=sysvarsIntervalMsec,def=10000" json:"sysvars_interval_msec,omitempty"`
	// Variables specified in this environment variable are exported as it is.
	// This is specifically useful to export information about system environment,
	// for example, docker image tag/digest-id, OS version etc. See
	// tools/cloudprober_startup.sh in the cloudprober directory for an example on
	// how to use these variables.
	SysvarsEnvVar *string `protobuf:"bytes,98,opt,name=sysvars_env_var,json=sysvarsEnvVar,def=SYSVARS" json:"sysvars_env_var,omitempty"`
	// Options for RTC reporter. RTC reporter reports information about the
	// current instance to a Runtime Config (RTC). This is useful if you want your
	// instance to be dynamically discoverable through RTC targets. This is
	// disabled by default.
	RtcReportOptions *proto5.RtcReportOptions `protobuf:"bytes,99,opt,name=rtc_report_options,json=rtcReportOptions" json:"rtc_report_options,omitempty"`
	// Global targets options. Per-probe options are specified within the probe
	// stanza.
	GlobalTargetsOptions *proto6.GlobalTargetsOptions `protobuf:"bytes,100,opt,name=global_targets_options,json=globalTargetsOptions" json:"global_targets_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ProberConfig) Reset()         { *m = ProberConfig{} }
func (m *ProberConfig) String() string { return proto.CompactTextString(m) }
func (*ProberConfig) ProtoMessage()    {}
func (*ProberConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_c2391771f7538cdb, []int{0}
}
func (m *ProberConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProberConfig.Unmarshal(m, b)
}
func (m *ProberConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProberConfig.Marshal(b, m, deterministic)
}
func (dst *ProberConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProberConfig.Merge(dst, src)
}
func (m *ProberConfig) XXX_Size() int {
	return xxx_messageInfo_ProberConfig.Size(m)
}
func (m *ProberConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProberConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProberConfig proto.InternalMessageInfo

const Default_ProberConfig_DisableJitter bool = false
const Default_ProberConfig_SysvarsIntervalMsec int32 = 10000
const Default_ProberConfig_SysvarsEnvVar string = "SYSVARS"

func (m *ProberConfig) GetProbe() []*proto1.ProbeDef {
	if m != nil {
		return m.Probe
	}
	return nil
}

func (m *ProberConfig) GetSurfacer() []*proto2.SurfacerDef {
	if m != nil {
		return m.Surfacer
	}
	return nil
}

func (m *ProberConfig) GetServer() []*proto3.ServerDef {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *ProberConfig) GetRdsServer() *proto4.ServerConf {
	if m != nil {
		return m.RdsServer
	}
	return nil
}

func (m *ProberConfig) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *ProberConfig) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *ProberConfig) GetDisableJitter() bool {
	if m != nil && m.DisableJitter != nil {
		return *m.DisableJitter
	}
	return Default_ProberConfig_DisableJitter
}

func (m *ProberConfig) GetSysvarsIntervalMsec() int32 {
	if m != nil && m.SysvarsIntervalMsec != nil {
		return *m.SysvarsIntervalMsec
	}
	return Default_ProberConfig_SysvarsIntervalMsec
}

func (m *ProberConfig) GetSysvarsEnvVar() string {
	if m != nil && m.SysvarsEnvVar != nil {
		return *m.SysvarsEnvVar
	}
	return Default_ProberConfig_SysvarsEnvVar
}

func (m *ProberConfig) GetRtcReportOptions() *proto5.RtcReportOptions {
	if m != nil {
		return m.RtcReportOptions
	}
	return nil
}

func (m *ProberConfig) GetGlobalTargetsOptions() *proto6.GlobalTargetsOptions {
	if m != nil {
		return m.GlobalTargetsOptions
	}
	return nil
}

func init() {
	proto.RegisterType((*ProberConfig)(nil), "cloudprober.ProberConfig")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/config/proto/config.proto", fileDescriptor_config_c2391771f7538cdb)
}

var fileDescriptor_config_c2391771f7538cdb = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x5b, 0x8b, 0xd4, 0x30,
	0x14, 0xc7, 0xa9, 0x6b, 0x75, 0x37, 0xe3, 0xaa, 0xc4, 0x0b, 0x61, 0x11, 0xa9, 0xfa, 0x52, 0x41,
	0xda, 0xd9, 0x79, 0x50, 0x77, 0xc0, 0x07, 0x5d, 0x45, 0x14, 0x16, 0x25, 0x95, 0x05, 0x9f, 0x62,
	0x9a, 0xa6, 0xdd, 0x4a, 0xb7, 0x19, 0x4e, 0x32, 0x05, 0x1f, 0xfd, 0xe6, 0xd2, 0x24, 0x1d, 0x3a,
	0xd2, 0xbd, 0x3c, 0x94, 0x9c, 0xdb, 0xef, 0x7f, 0x92, 0x73, 0x8a, 0xde, 0x54, 0xb5, 0x39, 0x5b,
	0xe7, 0x89, 0x50, 0xe7, 0x69, 0xa5, 0x54, 0xd5, 0xc8, 0x54, 0x34, 0x6a, 0x5d, 0xac, 0x40, 0xe5,
	0x12, 0x52, 0xa1, 0xda, 0xb2, 0xae, 0xd2, 0x15, 0x28, 0xa3, 0xbc, 0x93, 0x58, 0x07, 0xcf, 0x46,
	0x65, 0x07, 0x57, 0xa8, 0xd8, 0x43, 0x4f, 0xa8, 0x1c, 0xbc, 0xbd, 0x1c, 0xd4, 0x12, 0x3a, 0x09,
	0x93, 0xe4, 0xf2, 0x0a, 0x72, 0x0d, 0x25, 0x17, 0x17, 0xb0, 0x47, 0x97, 0xb3, 0x86, 0x43, 0x25,
	0xcd, 0x40, 0x7a, 0xcf, 0xa3, 0xc7, 0xd7, 0x43, 0xa1, 0xd0, 0xfe, 0xf2, 0x53, 0xfd, 0x4f, 0xae,
	0x29, 0x62, 0x44, 0xff, 0x81, 0x5c, 0x29, 0x30, 0x1b, 0xa5, 0x51, 0xc4, 0xc9, 0x3d, 0xff, 0x1b,
	0xa2, 0x3b, 0xdf, 0x2d, 0x7a, 0x6c, 0xbb, 0xe0, 0x05, 0x0a, 0xad, 0x14, 0x09, 0xa2, 0x9d, 0x78,
	0xb6, 0x78, 0x92, 0x8c, 0xd4, 0x13, 0xb7, 0x8c, 0xc4, 0x02, 0x1f, 0x65, 0x49, 0x5d, 0x29, 0x7e,
	0x87, 0x76, 0x87, 0x99, 0x91, 0x1b, 0x16, 0x7b, 0xb6, 0x85, 0x0d, 0xc9, 0x24, 0xf3, 0x46, 0xcf,
	0x6e, 0x10, 0xfc, 0x1a, 0xdd, 0x72, 0xef, 0x25, 0x3b, 0x16, 0x7e, 0xba, 0x0d, 0xbb, 0x3d, 0x26,
	0x99, 0x3d, 0x7b, 0xd2, 0x57, 0xe3, 0x0f, 0x08, 0x41, 0xa1, 0x99, 0x67, 0x59, 0x14, 0xc4, 0xb3,
	0xc5, 0x8b, 0x2d, 0x76, 0x98, 0x3f, 0x14, 0x03, 0xdf, 0xbf, 0x92, 0xee, 0x41, 0xa1, 0x9d, 0x8b,
	0x31, 0xba, 0xd9, 0xcf, 0x83, 0xfc, 0x8a, 0x82, 0x38, 0xa4, 0xd6, 0xee, 0x63, 0x67, 0x4a, 0x1b,
	0x22, 0xa3, 0x20, 0xde, 0xa3, 0xd6, 0xc6, 0xaf, 0xd0, 0xdd, 0xa2, 0xd6, 0x3c, 0x6f, 0x24, 0xfb,
	0x5d, 0x1b, 0x23, 0x81, 0x94, 0x51, 0x10, 0xef, 0x2e, 0xc3, 0x92, 0x37, 0x5a, 0xd2, 0x7d, 0x9f,
	0xfc, 0x6a, 0x73, 0xf8, 0x08, 0x3d, 0xd2, 0x7f, 0x74, 0xc7, 0x41, 0xb3, 0xba, 0x35, 0x12, 0x3a,
	0xde, 0xb0, 0x73, 0x2d, 0x05, 0xe1, 0x7d, 0x9b, 0x65, 0x78, 0x38, 0x9f, 0xcf, 0xe7, 0xf4, 0x81,
	0xaf, 0xf9, 0xe2, 0x4b, 0x4e, 0xb4, 0x14, 0x38, 0x45, 0xf7, 0x06, 0x54, 0xb6, 0x1d, 0xeb, 0x38,
	0x90, 0xbc, 0xbf, 0xc7, 0xf2, 0x76, 0xf6, 0x33, 0x3b, 0x7d, 0x4f, 0x33, 0xba, 0xef, 0xf3, 0x9f,
	0xda, 0xee, 0x94, 0x03, 0x66, 0x08, 0x83, 0x11, 0xcc, 0xed, 0x95, 0xa9, 0x95, 0xa9, 0x55, 0xab,
	0x89, 0xb0, 0xd3, 0x38, 0x9c, 0x9e, 0xc6, 0xe8, 0x2f, 0xa0, 0x46, 0x50, 0x6b, 0x7f, 0x73, 0x20,
	0xbd, 0x0f, 0xff, 0x45, 0x30, 0x43, 0x8f, 0xab, 0x46, 0xe5, 0xbc, 0x61, 0x5e, 0x60, 0xd3, 0xa4,
	0xb0, 0x4d, 0x5e, 0x4e, 0x36, 0xf9, 0x6c, 0x91, 0x1f, 0xce, 0x1b, 0xc4, 0x1f, 0x56, 0x13, 0xd1,
	0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xd9, 0xd8, 0x87, 0x48, 0x04, 0x00, 0x00,
}
