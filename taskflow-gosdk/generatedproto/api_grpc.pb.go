// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: api.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CoordinatorService_UpdateTaskStatus_FullMethodName = "/grpc.CoordinatorService/UpdateTaskStatus"
)

// CoordinatorServiceClient is the client API for CoordinatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoordinatorServiceClient interface {
	UpdateTaskStatus(ctx context.Context, in *UpdateTaskStatusRequest, opts ...grpc.CallOption) (*UpdateTaskStatusResponse, error)
}

type coordinatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinatorServiceClient(cc grpc.ClientConnInterface) CoordinatorServiceClient {
	return &coordinatorServiceClient{cc}
}

func (c *coordinatorServiceClient) UpdateTaskStatus(ctx context.Context, in *UpdateTaskStatusRequest, opts ...grpc.CallOption) (*UpdateTaskStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTaskStatusResponse)
	err := c.cc.Invoke(ctx, CoordinatorService_UpdateTaskStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServiceServer is the server API for CoordinatorService service.
// All implementations must embed UnimplementedCoordinatorServiceServer
// for forward compatibility.
type CoordinatorServiceServer interface {
	UpdateTaskStatus(context.Context, *UpdateTaskStatusRequest) (*UpdateTaskStatusResponse, error)
	mustEmbedUnimplementedCoordinatorServiceServer()
}

// UnimplementedCoordinatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCoordinatorServiceServer struct{}

func (UnimplementedCoordinatorServiceServer) UpdateTaskStatus(context.Context, *UpdateTaskStatusRequest) (*UpdateTaskStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskStatus not implemented")
}
func (UnimplementedCoordinatorServiceServer) mustEmbedUnimplementedCoordinatorServiceServer() {}
func (UnimplementedCoordinatorServiceServer) testEmbeddedByValue()                            {}

// UnsafeCoordinatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoordinatorServiceServer will
// result in compilation errors.
type UnsafeCoordinatorServiceServer interface {
	mustEmbedUnimplementedCoordinatorServiceServer()
}

func RegisterCoordinatorServiceServer(s grpc.ServiceRegistrar, srv CoordinatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedCoordinatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CoordinatorService_ServiceDesc, srv)
}

func _CoordinatorService_UpdateTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServiceServer).UpdateTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoordinatorService_UpdateTaskStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServiceServer).UpdateTaskStatus(ctx, req.(*UpdateTaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoordinatorService_ServiceDesc is the grpc.ServiceDesc for CoordinatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoordinatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.CoordinatorService",
	HandlerType: (*CoordinatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateTaskStatus",
			Handler:    _CoordinatorService_UpdateTaskStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	QueueService_EnqueueTask_FullMethodName   = "/grpc.QueueService/EnqueueTask"
	QueueService_DequeueTask_FullMethodName   = "/grpc.QueueService/DequeueTask"
	QueueService_AllQueuesInfo_FullMethodName = "/grpc.QueueService/AllQueuesInfo"
)

// QueueServiceClient is the client API for QueueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueueServiceClient interface {
	EnqueueTask(ctx context.Context, in *EnqueueTaskRequest, opts ...grpc.CallOption) (*EnqueueTaskResponse, error)
	DequeueTask(ctx context.Context, in *DequeueTaskRequest, opts ...grpc.CallOption) (*DequeueTaskResponse, error)
	AllQueuesInfo(ctx context.Context, in *AllQueuesInfoRequest, opts ...grpc.CallOption) (*AllQueuesInfoResponse, error)
}

type queueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueServiceClient(cc grpc.ClientConnInterface) QueueServiceClient {
	return &queueServiceClient{cc}
}

func (c *queueServiceClient) EnqueueTask(ctx context.Context, in *EnqueueTaskRequest, opts ...grpc.CallOption) (*EnqueueTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnqueueTaskResponse)
	err := c.cc.Invoke(ctx, QueueService_EnqueueTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) DequeueTask(ctx context.Context, in *DequeueTaskRequest, opts ...grpc.CallOption) (*DequeueTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DequeueTaskResponse)
	err := c.cc.Invoke(ctx, QueueService_DequeueTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) AllQueuesInfo(ctx context.Context, in *AllQueuesInfoRequest, opts ...grpc.CallOption) (*AllQueuesInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllQueuesInfoResponse)
	err := c.cc.Invoke(ctx, QueueService_AllQueuesInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueServiceServer is the server API for QueueService service.
// All implementations must embed UnimplementedQueueServiceServer
// for forward compatibility.
type QueueServiceServer interface {
	EnqueueTask(context.Context, *EnqueueTaskRequest) (*EnqueueTaskResponse, error)
	DequeueTask(context.Context, *DequeueTaskRequest) (*DequeueTaskResponse, error)
	AllQueuesInfo(context.Context, *AllQueuesInfoRequest) (*AllQueuesInfoResponse, error)
	mustEmbedUnimplementedQueueServiceServer()
}

// UnimplementedQueueServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueueServiceServer struct{}

func (UnimplementedQueueServiceServer) EnqueueTask(context.Context, *EnqueueTaskRequest) (*EnqueueTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnqueueTask not implemented")
}
func (UnimplementedQueueServiceServer) DequeueTask(context.Context, *DequeueTaskRequest) (*DequeueTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DequeueTask not implemented")
}
func (UnimplementedQueueServiceServer) AllQueuesInfo(context.Context, *AllQueuesInfoRequest) (*AllQueuesInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllQueuesInfo not implemented")
}
func (UnimplementedQueueServiceServer) mustEmbedUnimplementedQueueServiceServer() {}
func (UnimplementedQueueServiceServer) testEmbeddedByValue()                      {}

// UnsafeQueueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueueServiceServer will
// result in compilation errors.
type UnsafeQueueServiceServer interface {
	mustEmbedUnimplementedQueueServiceServer()
}

func RegisterQueueServiceServer(s grpc.ServiceRegistrar, srv QueueServiceServer) {
	// If the following call pancis, it indicates UnimplementedQueueServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&QueueService_ServiceDesc, srv)
}

func _QueueService_EnqueueTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).EnqueueTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueService_EnqueueTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).EnqueueTask(ctx, req.(*EnqueueTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_DequeueTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DequeueTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).DequeueTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueService_DequeueTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).DequeueTask(ctx, req.(*DequeueTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_AllQueuesInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllQueuesInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).AllQueuesInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueService_AllQueuesInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).AllQueuesInfo(ctx, req.(*AllQueuesInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueueService_ServiceDesc is the grpc.ServiceDesc for QueueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.QueueService",
	HandlerType: (*QueueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnqueueTask",
			Handler:    _QueueService_EnqueueTask_Handler,
		},
		{
			MethodName: "DequeueTask",
			Handler:    _QueueService_DequeueTask_Handler,
		},
		{
			MethodName: "AllQueuesInfo",
			Handler:    _QueueService_AllQueuesInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	JobService_CreateJob_FullMethodName                       = "/grpc.JobService/CreateJob"
	JobService_UpdateJobStatus_FullMethodName                 = "/grpc.JobService/UpdateJobStatus"
	JobService_TriggerJobReRun_FullMethodName                 = "/grpc.JobService/TriggerJobReRun"
	JobService_GetJobStatusWithId_FullMethodName              = "/grpc.JobService/GetJobStatusWithId"
	JobService_GetAllJobsOfParticularTaskQueue_FullMethodName = "/grpc.JobService/GetAllJobsOfParticularTaskQueue"
	JobService_GetAllRunningJobsOfTaskQueue_FullMethodName    = "/grpc.JobService/GetAllRunningJobsOfTaskQueue"
)

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobServiceClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error)
	UpdateJobStatus(ctx context.Context, in *UpdateJobStatusRequest, opts ...grpc.CallOption) (*UpdateJobStatusResponse, error)
	TriggerJobReRun(ctx context.Context, in *TriggerReRunRequest, opts ...grpc.CallOption) (*TriggerReRunResponse, error)
	GetJobStatusWithId(ctx context.Context, in *GetJobStatusWithIdRequest, opts ...grpc.CallOption) (*GetJobStatusWithIdResponse, error)
	GetAllJobsOfParticularTaskQueue(ctx context.Context, in *GetAllJobsOfParticularTaskQueueRequest, opts ...grpc.CallOption) (*GetAllJobsOfParticularTaskQueueResponse, error)
	GetAllRunningJobsOfTaskQueue(ctx context.Context, in *GetAllRunningJobsOfTaskQueueRequest, opts ...grpc.CallOption) (*GetAllRunningJobsOfTaskQueueResponse, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*CreateJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateJobResponse)
	err := c.cc.Invoke(ctx, JobService_CreateJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) UpdateJobStatus(ctx context.Context, in *UpdateJobStatusRequest, opts ...grpc.CallOption) (*UpdateJobStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateJobStatusResponse)
	err := c.cc.Invoke(ctx, JobService_UpdateJobStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) TriggerJobReRun(ctx context.Context, in *TriggerReRunRequest, opts ...grpc.CallOption) (*TriggerReRunResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TriggerReRunResponse)
	err := c.cc.Invoke(ctx, JobService_TriggerJobReRun_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJobStatusWithId(ctx context.Context, in *GetJobStatusWithIdRequest, opts ...grpc.CallOption) (*GetJobStatusWithIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetJobStatusWithIdResponse)
	err := c.cc.Invoke(ctx, JobService_GetJobStatusWithId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetAllJobsOfParticularTaskQueue(ctx context.Context, in *GetAllJobsOfParticularTaskQueueRequest, opts ...grpc.CallOption) (*GetAllJobsOfParticularTaskQueueResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllJobsOfParticularTaskQueueResponse)
	err := c.cc.Invoke(ctx, JobService_GetAllJobsOfParticularTaskQueue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetAllRunningJobsOfTaskQueue(ctx context.Context, in *GetAllRunningJobsOfTaskQueueRequest, opts ...grpc.CallOption) (*GetAllRunningJobsOfTaskQueueResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllRunningJobsOfTaskQueueResponse)
	err := c.cc.Invoke(ctx, JobService_GetAllRunningJobsOfTaskQueue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
// All implementations must embed UnimplementedJobServiceServer
// for forward compatibility.
type JobServiceServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*CreateJobResponse, error)
	UpdateJobStatus(context.Context, *UpdateJobStatusRequest) (*UpdateJobStatusResponse, error)
	TriggerJobReRun(context.Context, *TriggerReRunRequest) (*TriggerReRunResponse, error)
	GetJobStatusWithId(context.Context, *GetJobStatusWithIdRequest) (*GetJobStatusWithIdResponse, error)
	GetAllJobsOfParticularTaskQueue(context.Context, *GetAllJobsOfParticularTaskQueueRequest) (*GetAllJobsOfParticularTaskQueueResponse, error)
	GetAllRunningJobsOfTaskQueue(context.Context, *GetAllRunningJobsOfTaskQueueRequest) (*GetAllRunningJobsOfTaskQueueResponse, error)
	mustEmbedUnimplementedJobServiceServer()
}

// UnimplementedJobServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJobServiceServer struct{}

func (UnimplementedJobServiceServer) CreateJob(context.Context, *CreateJobRequest) (*CreateJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (UnimplementedJobServiceServer) UpdateJobStatus(context.Context, *UpdateJobStatusRequest) (*UpdateJobStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJobStatus not implemented")
}
func (UnimplementedJobServiceServer) TriggerJobReRun(context.Context, *TriggerReRunRequest) (*TriggerReRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerJobReRun not implemented")
}
func (UnimplementedJobServiceServer) GetJobStatusWithId(context.Context, *GetJobStatusWithIdRequest) (*GetJobStatusWithIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobStatusWithId not implemented")
}
func (UnimplementedJobServiceServer) GetAllJobsOfParticularTaskQueue(context.Context, *GetAllJobsOfParticularTaskQueueRequest) (*GetAllJobsOfParticularTaskQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllJobsOfParticularTaskQueue not implemented")
}
func (UnimplementedJobServiceServer) GetAllRunningJobsOfTaskQueue(context.Context, *GetAllRunningJobsOfTaskQueueRequest) (*GetAllRunningJobsOfTaskQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRunningJobsOfTaskQueue not implemented")
}
func (UnimplementedJobServiceServer) mustEmbedUnimplementedJobServiceServer() {}
func (UnimplementedJobServiceServer) testEmbeddedByValue()                    {}

// UnsafeJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServiceServer will
// result in compilation errors.
type UnsafeJobServiceServer interface {
	mustEmbedUnimplementedJobServiceServer()
}

func RegisterJobServiceServer(s grpc.ServiceRegistrar, srv JobServiceServer) {
	// If the following call pancis, it indicates UnimplementedJobServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&JobService_ServiceDesc, srv)
}

func _JobService_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_CreateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).CreateJob(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_UpdateJobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).UpdateJobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_UpdateJobStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).UpdateJobStatus(ctx, req.(*UpdateJobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_TriggerJobReRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerReRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).TriggerJobReRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_TriggerJobReRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).TriggerJobReRun(ctx, req.(*TriggerReRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJobStatusWithId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobStatusWithIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJobStatusWithId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_GetJobStatusWithId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJobStatusWithId(ctx, req.(*GetJobStatusWithIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetAllJobsOfParticularTaskQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllJobsOfParticularTaskQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetAllJobsOfParticularTaskQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_GetAllJobsOfParticularTaskQueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetAllJobsOfParticularTaskQueue(ctx, req.(*GetAllJobsOfParticularTaskQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetAllRunningJobsOfTaskQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRunningJobsOfTaskQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetAllRunningJobsOfTaskQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_GetAllRunningJobsOfTaskQueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetAllRunningJobsOfTaskQueue(ctx, req.(*GetAllRunningJobsOfTaskQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobService_ServiceDesc is the grpc.ServiceDesc for JobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _JobService_CreateJob_Handler,
		},
		{
			MethodName: "UpdateJobStatus",
			Handler:    _JobService_UpdateJobStatus_Handler,
		},
		{
			MethodName: "TriggerJobReRun",
			Handler:    _JobService_TriggerJobReRun_Handler,
		},
		{
			MethodName: "GetJobStatusWithId",
			Handler:    _JobService_GetJobStatusWithId_Handler,
		},
		{
			MethodName: "GetAllJobsOfParticularTaskQueue",
			Handler:    _JobService_GetAllJobsOfParticularTaskQueue_Handler,
		},
		{
			MethodName: "GetAllRunningJobsOfTaskQueue",
			Handler:    _JobService_GetAllRunningJobsOfTaskQueue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
