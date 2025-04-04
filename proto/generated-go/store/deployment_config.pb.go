// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: store/deployment_config.proto

package store

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LabelSelectorRequirement_OperatorType int32

const (
	// The operator is not specified.
	LabelSelectorRequirement_OPERATOR_TYPE_UNSPECIFIED LabelSelectorRequirement_OperatorType = 0
	// The operator is "In".
	LabelSelectorRequirement_IN LabelSelectorRequirement_OperatorType = 1
	// The operator is "Exists".
	LabelSelectorRequirement_EXISTS LabelSelectorRequirement_OperatorType = 2
	// The operator is "Not In".
	LabelSelectorRequirement_NOT_IN LabelSelectorRequirement_OperatorType = 3
)

// Enum value maps for LabelSelectorRequirement_OperatorType.
var (
	LabelSelectorRequirement_OperatorType_name = map[int32]string{
		0: "OPERATOR_TYPE_UNSPECIFIED",
		1: "IN",
		2: "EXISTS",
		3: "NOT_IN",
	}
	LabelSelectorRequirement_OperatorType_value = map[string]int32{
		"OPERATOR_TYPE_UNSPECIFIED": 0,
		"IN":                        1,
		"EXISTS":                    2,
		"NOT_IN":                    3,
	}
)

func (x LabelSelectorRequirement_OperatorType) Enum() *LabelSelectorRequirement_OperatorType {
	p := new(LabelSelectorRequirement_OperatorType)
	*p = x
	return p
}

func (x LabelSelectorRequirement_OperatorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LabelSelectorRequirement_OperatorType) Descriptor() protoreflect.EnumDescriptor {
	return file_store_deployment_config_proto_enumTypes[0].Descriptor()
}

func (LabelSelectorRequirement_OperatorType) Type() protoreflect.EnumType {
	return &file_store_deployment_config_proto_enumTypes[0]
}

func (x LabelSelectorRequirement_OperatorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LabelSelectorRequirement_OperatorType.Descriptor instead.
func (LabelSelectorRequirement_OperatorType) EnumDescriptor() ([]byte, []int) {
	return file_store_deployment_config_proto_rawDescGZIP(), []int{5, 0}
}

type DeploymentConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Schedule      *Schedule              `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeploymentConfig) Reset() {
	*x = DeploymentConfig{}
	mi := &file_store_deployment_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploymentConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentConfig) ProtoMessage() {}

func (x *DeploymentConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_deployment_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentConfig.ProtoReflect.Descriptor instead.
func (*DeploymentConfig) Descriptor() ([]byte, []int) {
	return file_store_deployment_config_proto_rawDescGZIP(), []int{0}
}

func (x *DeploymentConfig) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

// Schedule is the message for deployment schedule.
type Schedule struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Deployments   []*ScheduleDeployment  `protobuf:"bytes,1,rep,name=deployments,proto3" json:"deployments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	mi := &file_store_deployment_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_store_deployment_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_store_deployment_config_proto_rawDescGZIP(), []int{1}
}

func (x *Schedule) GetDeployments() []*ScheduleDeployment {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type ScheduleDeployment struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The title of the deployment (stage) in a schedule.
	Title         string          `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Id            string          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Spec          *DeploymentSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScheduleDeployment) Reset() {
	*x = ScheduleDeployment{}
	mi := &file_store_deployment_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduleDeployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleDeployment) ProtoMessage() {}

func (x *ScheduleDeployment) ProtoReflect() protoreflect.Message {
	mi := &file_store_deployment_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleDeployment.ProtoReflect.Descriptor instead.
func (*ScheduleDeployment) Descriptor() ([]byte, []int) {
	return file_store_deployment_config_proto_rawDescGZIP(), []int{2}
}

func (x *ScheduleDeployment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ScheduleDeployment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ScheduleDeployment) GetSpec() *DeploymentSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type DeploymentSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Selector      *LabelSelector         `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeploymentSpec) Reset() {
	*x = DeploymentSpec{}
	mi := &file_store_deployment_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploymentSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentSpec) ProtoMessage() {}

func (x *DeploymentSpec) ProtoReflect() protoreflect.Message {
	mi := &file_store_deployment_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentSpec.ProtoReflect.Descriptor instead.
func (*DeploymentSpec) Descriptor() ([]byte, []int) {
	return file_store_deployment_config_proto_rawDescGZIP(), []int{3}
}

func (x *DeploymentSpec) GetSelector() *LabelSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

type LabelSelector struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// match_expressions is a list of label selector requirements. The requirements are ANDed.
	MatchExpressions []*LabelSelectorRequirement `protobuf:"bytes,1,rep,name=match_expressions,json=matchExpressions,proto3" json:"match_expressions,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *LabelSelector) Reset() {
	*x = LabelSelector{}
	mi := &file_store_deployment_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelSelector) ProtoMessage() {}

func (x *LabelSelector) ProtoReflect() protoreflect.Message {
	mi := &file_store_deployment_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelSelector.ProtoReflect.Descriptor instead.
func (*LabelSelector) Descriptor() ([]byte, []int) {
	return file_store_deployment_config_proto_rawDescGZIP(), []int{4}
}

func (x *LabelSelector) GetMatchExpressions() []*LabelSelectorRequirement {
	if x != nil {
		return x.MatchExpressions
	}
	return nil
}

type LabelSelectorRequirement struct {
	state    protoimpl.MessageState                `protogen:"open.v1"`
	Key      string                                `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Operator LabelSelectorRequirement_OperatorType `protobuf:"varint,2,opt,name=operator,proto3,enum=bytebase.store.LabelSelectorRequirement_OperatorType" json:"operator,omitempty"`
	// Values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Values        []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelSelectorRequirement) Reset() {
	*x = LabelSelectorRequirement{}
	mi := &file_store_deployment_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelSelectorRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelSelectorRequirement) ProtoMessage() {}

func (x *LabelSelectorRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_store_deployment_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelSelectorRequirement.ProtoReflect.Descriptor instead.
func (*LabelSelectorRequirement) Descriptor() ([]byte, []int) {
	return file_store_deployment_config_proto_rawDescGZIP(), []int{5}
}

func (x *LabelSelectorRequirement) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LabelSelectorRequirement) GetOperator() LabelSelectorRequirement_OperatorType {
	if x != nil {
		return x.Operator
	}
	return LabelSelectorRequirement_OPERATOR_TYPE_UNSPECIFIED
}

func (x *LabelSelectorRequirement) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_store_deployment_config_proto protoreflect.FileDescriptor

var file_store_deployment_config_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22,
	0x48, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x50, 0x0a, 0x08, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x6e, 0x0a, 0x12, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x4b, 0x0a, 0x0e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x39, 0x0a,
	0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x66, 0x0a, 0x0d, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x55, 0x0a, 0x11, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x10,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xe6, 0x01, 0x0a, 0x18, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x51, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x35, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x0c, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x4e, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x03, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_store_deployment_config_proto_rawDescOnce sync.Once
	file_store_deployment_config_proto_rawDescData []byte
)

func file_store_deployment_config_proto_rawDescGZIP() []byte {
	file_store_deployment_config_proto_rawDescOnce.Do(func() {
		file_store_deployment_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_store_deployment_config_proto_rawDesc), len(file_store_deployment_config_proto_rawDesc)))
	})
	return file_store_deployment_config_proto_rawDescData
}

var file_store_deployment_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_deployment_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_store_deployment_config_proto_goTypes = []any{
	(LabelSelectorRequirement_OperatorType)(0), // 0: bytebase.store.LabelSelectorRequirement.OperatorType
	(*DeploymentConfig)(nil),                   // 1: bytebase.store.DeploymentConfig
	(*Schedule)(nil),                           // 2: bytebase.store.Schedule
	(*ScheduleDeployment)(nil),                 // 3: bytebase.store.ScheduleDeployment
	(*DeploymentSpec)(nil),                     // 4: bytebase.store.DeploymentSpec
	(*LabelSelector)(nil),                      // 5: bytebase.store.LabelSelector
	(*LabelSelectorRequirement)(nil),           // 6: bytebase.store.LabelSelectorRequirement
}
var file_store_deployment_config_proto_depIdxs = []int32{
	2, // 0: bytebase.store.DeploymentConfig.schedule:type_name -> bytebase.store.Schedule
	3, // 1: bytebase.store.Schedule.deployments:type_name -> bytebase.store.ScheduleDeployment
	4, // 2: bytebase.store.ScheduleDeployment.spec:type_name -> bytebase.store.DeploymentSpec
	5, // 3: bytebase.store.DeploymentSpec.selector:type_name -> bytebase.store.LabelSelector
	6, // 4: bytebase.store.LabelSelector.match_expressions:type_name -> bytebase.store.LabelSelectorRequirement
	0, // 5: bytebase.store.LabelSelectorRequirement.operator:type_name -> bytebase.store.LabelSelectorRequirement.OperatorType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_store_deployment_config_proto_init() }
func file_store_deployment_config_proto_init() {
	if File_store_deployment_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_store_deployment_config_proto_rawDesc), len(file_store_deployment_config_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_deployment_config_proto_goTypes,
		DependencyIndexes: file_store_deployment_config_proto_depIdxs,
		EnumInfos:         file_store_deployment_config_proto_enumTypes,
		MessageInfos:      file_store_deployment_config_proto_msgTypes,
	}.Build()
	File_store_deployment_config_proto = out.File
	file_store_deployment_config_proto_goTypes = nil
	file_store_deployment_config_proto_depIdxs = nil
}
