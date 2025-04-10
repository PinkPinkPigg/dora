// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: idls/services/executor.proto

package gen

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

type TaskType int32

const (
	TaskType_PYTHON TaskType = 0
	TaskType_SHELL  TaskType = 1
)

// Enum value maps for TaskType.
var (
	TaskType_name = map[int32]string{
		0: "PYTHON",
		1: "SHELL",
	}
	TaskType_value = map[string]int32{
		"PYTHON": 0,
		"SHELL":  1,
	}
)

func (x TaskType) Enum() *TaskType {
	p := new(TaskType)
	*p = x
	return p
}

func (x TaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_idls_services_executor_proto_enumTypes[0].Descriptor()
}

func (TaskType) Type() protoreflect.EnumType {
	return &file_idls_services_executor_proto_enumTypes[0]
}

func (x TaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskType.Descriptor instead.
func (TaskType) EnumDescriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{0}
}

type TaskFrequency int32

const (
	TaskFrequency_HOUR TaskFrequency = 0
	TaskFrequency_DAY  TaskFrequency = 1
	TaskFrequency_WEEK TaskFrequency = 2
)

// Enum value maps for TaskFrequency.
var (
	TaskFrequency_name = map[int32]string{
		0: "HOUR",
		1: "DAY",
		2: "WEEK",
	}
	TaskFrequency_value = map[string]int32{
		"HOUR": 0,
		"DAY":  1,
		"WEEK": 2,
	}
)

func (x TaskFrequency) Enum() *TaskFrequency {
	p := new(TaskFrequency)
	*p = x
	return p
}

func (x TaskFrequency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskFrequency) Descriptor() protoreflect.EnumDescriptor {
	return file_idls_services_executor_proto_enumTypes[1].Descriptor()
}

func (TaskFrequency) Type() protoreflect.EnumType {
	return &file_idls_services_executor_proto_enumTypes[1]
}

func (x TaskFrequency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskFrequency.Descriptor instead.
func (TaskFrequency) EnumDescriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{1}
}

type TaskStatus int32

const (
	TaskStatus_ONLINE  TaskStatus = 0 //任务在线
	TaskStatus_OFFLINE TaskStatus = 1 //任务离线
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "ONLINE",
		1: "OFFLINE",
	}
	TaskStatus_value = map[string]int32{
		"ONLINE":  0,
		"OFFLINE": 1,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_idls_services_executor_proto_enumTypes[2].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_idls_services_executor_proto_enumTypes[2]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{2}
}

type InstanceStatus int32

const (
	InstanceStatus_WAITING_TO_SUBMIT InstanceStatus = 0 //创建待提交
	InstanceStatus_EXECUTING         InstanceStatus = 1 //执行中
	InstanceStatus_SUCCESS           InstanceStatus = 2 //成功
	InstanceStatus_FAIL_WAIT_RETRY   InstanceStatus = 3 //失败待重试
	InstanceStatus_FAIL              InstanceStatus = 4 //失败状态
	InstanceStatus_CANCEL            InstanceStatus = 5 //被取消
)

// Enum value maps for InstanceStatus.
var (
	InstanceStatus_name = map[int32]string{
		0: "WAITING_TO_SUBMIT",
		1: "EXECUTING",
		2: "SUCCESS",
		3: "FAIL_WAIT_RETRY",
		4: "FAIL",
		5: "CANCEL",
	}
	InstanceStatus_value = map[string]int32{
		"WAITING_TO_SUBMIT": 0,
		"EXECUTING":         1,
		"SUCCESS":           2,
		"FAIL_WAIT_RETRY":   3,
		"FAIL":              4,
		"CANCEL":            5,
	}
)

func (x InstanceStatus) Enum() *InstanceStatus {
	p := new(InstanceStatus)
	*p = x
	return p
}

func (x InstanceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstanceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_idls_services_executor_proto_enumTypes[3].Descriptor()
}

func (InstanceStatus) Type() protoreflect.EnumType {
	return &file_idls_services_executor_proto_enumTypes[3]
}

func (x InstanceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstanceStatus.Descriptor instead.
func (InstanceStatus) EnumDescriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{3}
}

type InstanceBase struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                                   //实例唯一id
	CreateTimestamp uint64                 `protobuf:"varint,2,opt,name=create_timestamp,json=createTimestamp,proto3" json:"create_timestamp,omitempty"` //创建时间
	ModifyTimestamp uint64                 `protobuf:"varint,3,opt,name=modify_timestamp,json=modifyTimestamp,proto3" json:"modify_timestamp,omitempty"` //修改时间
	ExecuteTime     uint32                 `protobuf:"varint,4,opt,name=execute_time,json=executeTime,proto3" json:"execute_time,omitempty"`             //执行次数
	Status          InstanceStatus         `protobuf:"varint,5,opt,name=status,proto3,enum=executor.InstanceStatus" json:"status,omitempty"`             //实例当前状态
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *InstanceBase) Reset() {
	*x = InstanceBase{}
	mi := &file_idls_services_executor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstanceBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceBase) ProtoMessage() {}

func (x *InstanceBase) ProtoReflect() protoreflect.Message {
	mi := &file_idls_services_executor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceBase.ProtoReflect.Descriptor instead.
func (*InstanceBase) Descriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{0}
}

func (x *InstanceBase) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InstanceBase) GetCreateTimestamp() uint64 {
	if x != nil {
		return x.CreateTimestamp
	}
	return 0
}

func (x *InstanceBase) GetModifyTimestamp() uint64 {
	if x != nil {
		return x.ModifyTimestamp
	}
	return 0
}

func (x *InstanceBase) GetExecuteTime() uint32 {
	if x != nil {
		return x.ExecuteTime
	}
	return 0
}

func (x *InstanceBase) GetStatus() InstanceStatus {
	if x != nil {
		return x.Status
	}
	return InstanceStatus_WAITING_TO_SUBMIT
}

type TaskBase struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                           //任务唯一id
	Type            TaskType               `protobuf:"varint,2,opt,name=type,proto3,enum=executor.TaskType" json:"type,omitempty"`                //任务类型
	Frequency       TaskFrequency          `protobuf:"varint,3,opt,name=frequency,proto3,enum=executor.TaskFrequency" json:"frequency,omitempty"` //执行周期
	Priority        uint32                 `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`                               //任务优先级
	Status          TaskStatus             `protobuf:"varint,5,opt,name=status,proto3,enum=executor.TaskStatus" json:"status,omitempty"`          //任务状态
	Description     string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`                          //任务描述
	CreateTimestamp uint64                 `protobuf:"varint,7,opt,name=create_timestamp,json=createTimestamp,proto3" json:"create_timestamp,omitempty"`
	ModifyTimestamp uint64                 `protobuf:"varint,8,opt,name=modify_timestamp,json=modifyTimestamp,proto3" json:"modify_timestamp,omitempty"`
	ScriptPath      string                 `protobuf:"bytes,9,opt,name=script_path,json=scriptPath,proto3" json:"script_path,omitempty"`                                                                               //脚本存储地址
	RetryConfig     map[string]string      `protobuf:"bytes,10,rep,name=retry_config,json=retryConfig,proto3" json:"retry_config,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` //重试配置
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *TaskBase) Reset() {
	*x = TaskBase{}
	mi := &file_idls_services_executor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskBase) ProtoMessage() {}

func (x *TaskBase) ProtoReflect() protoreflect.Message {
	mi := &file_idls_services_executor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskBase.ProtoReflect.Descriptor instead.
func (*TaskBase) Descriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{1}
}

func (x *TaskBase) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskBase) GetType() TaskType {
	if x != nil {
		return x.Type
	}
	return TaskType_PYTHON
}

func (x *TaskBase) GetFrequency() TaskFrequency {
	if x != nil {
		return x.Frequency
	}
	return TaskFrequency_HOUR
}

func (x *TaskBase) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *TaskBase) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_ONLINE
}

func (x *TaskBase) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskBase) GetCreateTimestamp() uint64 {
	if x != nil {
		return x.CreateTimestamp
	}
	return 0
}

func (x *TaskBase) GetModifyTimestamp() uint64 {
	if x != nil {
		return x.ModifyTimestamp
	}
	return 0
}

func (x *TaskBase) GetScriptPath() string {
	if x != nil {
		return x.ScriptPath
	}
	return ""
}

func (x *TaskBase) GetRetryConfig() map[string]string {
	if x != nil {
		return x.RetryConfig
	}
	return nil
}

type ExecuteInstanceRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Task              *TaskBase              `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	ScheduleTimestamp uint64                 `protobuf:"varint,2,opt,name=schedule_timestamp,json=scheduleTimestamp,proto3" json:"schedule_timestamp,omitempty"` //业务分区日期
	Force             bool                   `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`                                                  //是否强制执行
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ExecuteInstanceRequest) Reset() {
	*x = ExecuteInstanceRequest{}
	mi := &file_idls_services_executor_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteInstanceRequest) ProtoMessage() {}

func (x *ExecuteInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idls_services_executor_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteInstanceRequest.ProtoReflect.Descriptor instead.
func (*ExecuteInstanceRequest) Descriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{2}
}

func (x *ExecuteInstanceRequest) GetTask() *TaskBase {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *ExecuteInstanceRequest) GetScheduleTimestamp() uint64 {
	if x != nil {
		return x.ScheduleTimestamp
	}
	return 0
}

func (x *ExecuteInstanceRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type ExecuteInstanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int32                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`                                                                        //是否成功进入队列
	Extra         map[string]string      `protobuf:"bytes,2,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` //额外信息map
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecuteInstanceResponse) Reset() {
	*x = ExecuteInstanceResponse{}
	mi := &file_idls_services_executor_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteInstanceResponse) ProtoMessage() {}

func (x *ExecuteInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idls_services_executor_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteInstanceResponse.ProtoReflect.Descriptor instead.
func (*ExecuteInstanceResponse) Descriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{3}
}

func (x *ExecuteInstanceResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ExecuteInstanceResponse) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type CancelInstanceRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Task              *TaskBase              `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	ScheduleTimestamp uint64                 `protobuf:"varint,2,opt,name=schedule_timestamp,json=scheduleTimestamp,proto3" json:"schedule_timestamp,omitempty"` //业务分区日期
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CancelInstanceRequest) Reset() {
	*x = CancelInstanceRequest{}
	mi := &file_idls_services_executor_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelInstanceRequest) ProtoMessage() {}

func (x *CancelInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idls_services_executor_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelInstanceRequest.ProtoReflect.Descriptor instead.
func (*CancelInstanceRequest) Descriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{4}
}

func (x *CancelInstanceRequest) GetTask() *TaskBase {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *CancelInstanceRequest) GetScheduleTimestamp() uint64 {
	if x != nil {
		return x.ScheduleTimestamp
	}
	return 0
}

type CancelInstanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int32                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`                                                                        //是否成功进入队列
	Extra         map[string]string      `protobuf:"bytes,2,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` //额外信息map
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelInstanceResponse) Reset() {
	*x = CancelInstanceResponse{}
	mi := &file_idls_services_executor_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelInstanceResponse) ProtoMessage() {}

func (x *CancelInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idls_services_executor_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelInstanceResponse.ProtoReflect.Descriptor instead.
func (*CancelInstanceResponse) Descriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{5}
}

func (x *CancelInstanceResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CancelInstanceResponse) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type GetInstanceInfoRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	TaskId            uint32                 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`                                  //任务id
	ScheduleTimestamp uint64                 `protobuf:"varint,2,opt,name=schedule_timestamp,json=scheduleTimestamp,proto3" json:"schedule_timestamp,omitempty"` //业务分区日期
	Frequency         TaskFrequency          `protobuf:"varint,3,opt,name=frequency,proto3,enum=executor.TaskFrequency" json:"frequency,omitempty"`              //执行周期
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetInstanceInfoRequest) Reset() {
	*x = GetInstanceInfoRequest{}
	mi := &file_idls_services_executor_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInstanceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstanceInfoRequest) ProtoMessage() {}

func (x *GetInstanceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idls_services_executor_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstanceInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInstanceInfoRequest) Descriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{6}
}

func (x *GetInstanceInfoRequest) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *GetInstanceInfoRequest) GetScheduleTimestamp() uint64 {
	if x != nil {
		return x.ScheduleTimestamp
	}
	return 0
}

func (x *GetInstanceInfoRequest) GetFrequency() TaskFrequency {
	if x != nil {
		return x.Frequency
	}
	return TaskFrequency_HOUR
}

type GetInstanceInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Instance      *InstanceBase          `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`                                                                     //实例基础信息
	Extra         map[string]string      `protobuf:"bytes,3,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` //额外信息map
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInstanceInfoResponse) Reset() {
	*x = GetInstanceInfoResponse{}
	mi := &file_idls_services_executor_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInstanceInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstanceInfoResponse) ProtoMessage() {}

func (x *GetInstanceInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idls_services_executor_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstanceInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInstanceInfoResponse) Descriptor() ([]byte, []int) {
	return file_idls_services_executor_proto_rawDescGZIP(), []int{7}
}

func (x *GetInstanceInfoResponse) GetInstance() *InstanceBase {
	if x != nil {
		return x.Instance
	}
	return nil
}

func (x *GetInstanceInfoResponse) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_idls_services_executor_proto protoreflect.FileDescriptor

const file_idls_services_executor_proto_rawDesc = "" +
	"\n" +
	"\x1cidls/services/executor.proto\x12\bexecutor\"\xc9\x01\n" +
	"\fInstanceBase\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12)\n" +
	"\x10create_timestamp\x18\x02 \x01(\x04R\x0fcreateTimestamp\x12)\n" +
	"\x10modify_timestamp\x18\x03 \x01(\x04R\x0fmodifyTimestamp\x12!\n" +
	"\fexecute_time\x18\x04 \x01(\rR\vexecuteTime\x120\n" +
	"\x06status\x18\x05 \x01(\x0e2\x18.executor.InstanceStatusR\x06status\"\xe4\x03\n" +
	"\bTaskBase\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12&\n" +
	"\x04type\x18\x02 \x01(\x0e2\x12.executor.TaskTypeR\x04type\x125\n" +
	"\tfrequency\x18\x03 \x01(\x0e2\x17.executor.TaskFrequencyR\tfrequency\x12\x1a\n" +
	"\bpriority\x18\x04 \x01(\rR\bpriority\x12,\n" +
	"\x06status\x18\x05 \x01(\x0e2\x14.executor.TaskStatusR\x06status\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescription\x12)\n" +
	"\x10create_timestamp\x18\a \x01(\x04R\x0fcreateTimestamp\x12)\n" +
	"\x10modify_timestamp\x18\b \x01(\x04R\x0fmodifyTimestamp\x12\x1f\n" +
	"\vscript_path\x18\t \x01(\tR\n" +
	"scriptPath\x12F\n" +
	"\fretry_config\x18\n" +
	" \x03(\v2#.executor.TaskBase.RetryConfigEntryR\vretryConfig\x1a>\n" +
	"\x10RetryConfigEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\x85\x01\n" +
	"\x16ExecuteInstanceRequest\x12&\n" +
	"\x04task\x18\x01 \x01(\v2\x12.executor.TaskBaseR\x04task\x12-\n" +
	"\x12schedule_timestamp\x18\x02 \x01(\x04R\x11scheduleTimestamp\x12\x14\n" +
	"\x05force\x18\x03 \x01(\bR\x05force\"\xaf\x01\n" +
	"\x17ExecuteInstanceResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x05R\x06status\x12B\n" +
	"\x05extra\x18\x02 \x03(\v2,.executor.ExecuteInstanceResponse.ExtraEntryR\x05extra\x1a8\n" +
	"\n" +
	"ExtraEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"n\n" +
	"\x15CancelInstanceRequest\x12&\n" +
	"\x04task\x18\x01 \x01(\v2\x12.executor.TaskBaseR\x04task\x12-\n" +
	"\x12schedule_timestamp\x18\x02 \x01(\x04R\x11scheduleTimestamp\"\xad\x01\n" +
	"\x16CancelInstanceResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x05R\x06status\x12A\n" +
	"\x05extra\x18\x02 \x03(\v2+.executor.CancelInstanceResponse.ExtraEntryR\x05extra\x1a8\n" +
	"\n" +
	"ExtraEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\x97\x01\n" +
	"\x16GetInstanceInfoRequest\x12\x17\n" +
	"\atask_id\x18\x01 \x01(\rR\x06taskId\x12-\n" +
	"\x12schedule_timestamp\x18\x02 \x01(\x04R\x11scheduleTimestamp\x125\n" +
	"\tfrequency\x18\x03 \x01(\x0e2\x17.executor.TaskFrequencyR\tfrequency\"\xcb\x01\n" +
	"\x17GetInstanceInfoResponse\x122\n" +
	"\binstance\x18\x01 \x01(\v2\x16.executor.InstanceBaseR\binstance\x12B\n" +
	"\x05extra\x18\x03 \x03(\v2,.executor.GetInstanceInfoResponse.ExtraEntryR\x05extra\x1a8\n" +
	"\n" +
	"ExtraEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01*!\n" +
	"\bTaskType\x12\n" +
	"\n" +
	"\x06PYTHON\x10\x00\x12\t\n" +
	"\x05SHELL\x10\x01*,\n" +
	"\rTaskFrequency\x12\b\n" +
	"\x04HOUR\x10\x00\x12\a\n" +
	"\x03DAY\x10\x01\x12\b\n" +
	"\x04WEEK\x10\x02*%\n" +
	"\n" +
	"TaskStatus\x12\n" +
	"\n" +
	"\x06ONLINE\x10\x00\x12\v\n" +
	"\aOFFLINE\x10\x01*n\n" +
	"\x0eInstanceStatus\x12\x15\n" +
	"\x11WAITING_TO_SUBMIT\x10\x00\x12\r\n" +
	"\tEXECUTING\x10\x01\x12\v\n" +
	"\aSUCCESS\x10\x02\x12\x13\n" +
	"\x0fFAIL_WAIT_RETRY\x10\x03\x12\b\n" +
	"\x04FAIL\x10\x04\x12\n" +
	"\n" +
	"\x06CANCEL\x10\x052\x96\x02\n" +
	"\x0fExecutorService\x12V\n" +
	"\x0fExecuteInstance\x12 .executor.ExecuteInstanceRequest\x1a!.executor.ExecuteInstanceResponse\x12S\n" +
	"\x0eCancelInstance\x12\x1f.executor.CancelInstanceRequest\x1a .executor.CancelInstanceResponse\x12V\n" +
	"\x0fGetInstanceInfo\x12 .executor.GetInstanceInfoRequest\x1a!.executor.GetInstanceInfoResponseB\x19Z\x17/dora/services/executorb\x06proto3"

var (
	file_idls_services_executor_proto_rawDescOnce sync.Once
	file_idls_services_executor_proto_rawDescData []byte
)

func file_idls_services_executor_proto_rawDescGZIP() []byte {
	file_idls_services_executor_proto_rawDescOnce.Do(func() {
		file_idls_services_executor_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_idls_services_executor_proto_rawDesc), len(file_idls_services_executor_proto_rawDesc)))
	})
	return file_idls_services_executor_proto_rawDescData
}

var file_idls_services_executor_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_idls_services_executor_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_idls_services_executor_proto_goTypes = []any{
	(TaskType)(0),                   // 0: executor.TaskType
	(TaskFrequency)(0),              // 1: executor.TaskFrequency
	(TaskStatus)(0),                 // 2: executor.TaskStatus
	(InstanceStatus)(0),             // 3: executor.InstanceStatus
	(*InstanceBase)(nil),            // 4: executor.InstanceBase
	(*TaskBase)(nil),                // 5: executor.TaskBase
	(*ExecuteInstanceRequest)(nil),  // 6: executor.ExecuteInstanceRequest
	(*ExecuteInstanceResponse)(nil), // 7: executor.ExecuteInstanceResponse
	(*CancelInstanceRequest)(nil),   // 8: executor.CancelInstanceRequest
	(*CancelInstanceResponse)(nil),  // 9: executor.CancelInstanceResponse
	(*GetInstanceInfoRequest)(nil),  // 10: executor.GetInstanceInfoRequest
	(*GetInstanceInfoResponse)(nil), // 11: executor.GetInstanceInfoResponse
	nil,                             // 12: executor.TaskBase.RetryConfigEntry
	nil,                             // 13: executor.ExecuteInstanceResponse.ExtraEntry
	nil,                             // 14: executor.CancelInstanceResponse.ExtraEntry
	nil,                             // 15: executor.GetInstanceInfoResponse.ExtraEntry
}
var file_idls_services_executor_proto_depIdxs = []int32{
	3,  // 0: executor.InstanceBase.status:type_name -> executor.InstanceStatus
	0,  // 1: executor.TaskBase.type:type_name -> executor.TaskType
	1,  // 2: executor.TaskBase.frequency:type_name -> executor.TaskFrequency
	2,  // 3: executor.TaskBase.status:type_name -> executor.TaskStatus
	12, // 4: executor.TaskBase.retry_config:type_name -> executor.TaskBase.RetryConfigEntry
	5,  // 5: executor.ExecuteInstanceRequest.task:type_name -> executor.TaskBase
	13, // 6: executor.ExecuteInstanceResponse.extra:type_name -> executor.ExecuteInstanceResponse.ExtraEntry
	5,  // 7: executor.CancelInstanceRequest.task:type_name -> executor.TaskBase
	14, // 8: executor.CancelInstanceResponse.extra:type_name -> executor.CancelInstanceResponse.ExtraEntry
	1,  // 9: executor.GetInstanceInfoRequest.frequency:type_name -> executor.TaskFrequency
	4,  // 10: executor.GetInstanceInfoResponse.instance:type_name -> executor.InstanceBase
	15, // 11: executor.GetInstanceInfoResponse.extra:type_name -> executor.GetInstanceInfoResponse.ExtraEntry
	6,  // 12: executor.ExecutorService.ExecuteInstance:input_type -> executor.ExecuteInstanceRequest
	8,  // 13: executor.ExecutorService.CancelInstance:input_type -> executor.CancelInstanceRequest
	10, // 14: executor.ExecutorService.GetInstanceInfo:input_type -> executor.GetInstanceInfoRequest
	7,  // 15: executor.ExecutorService.ExecuteInstance:output_type -> executor.ExecuteInstanceResponse
	9,  // 16: executor.ExecutorService.CancelInstance:output_type -> executor.CancelInstanceResponse
	11, // 17: executor.ExecutorService.GetInstanceInfo:output_type -> executor.GetInstanceInfoResponse
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_idls_services_executor_proto_init() }
func file_idls_services_executor_proto_init() {
	if File_idls_services_executor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_idls_services_executor_proto_rawDesc), len(file_idls_services_executor_proto_rawDesc)),
			NumEnums:      4,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idls_services_executor_proto_goTypes,
		DependencyIndexes: file_idls_services_executor_proto_depIdxs,
		EnumInfos:         file_idls_services_executor_proto_enumTypes,
		MessageInfos:      file_idls_services_executor_proto_msgTypes,
	}.Build()
	File_idls_services_executor_proto = out.File
	file_idls_services_executor_proto_goTypes = nil
	file_idls_services_executor_proto_depIdxs = nil
}
