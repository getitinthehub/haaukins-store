// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreClient interface {
	//Insert
	AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	AddTeam(ctx context.Context, in *AddTeamRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	//Select
	GetEvents(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error)
	GetEventByUser(ctx context.Context, in *GetEventByUserReq, opts ...grpc.CallOption) (*GetEventResponse, error)
	GetEventTeams(ctx context.Context, in *GetEventTeamsRequest, opts ...grpc.CallOption) (*GetEventTeamsResponse, error)
	GetEventStatus(ctx context.Context, in *GetEventStatusRequest, opts ...grpc.CallOption) (*EventStatusStore, error)
	IsEventExists(ctx context.Context, in *GetEventByTagReq, opts ...grpc.CallOption) (*GetEventByTagResp, error)
	GetTimeSeries(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetTimeSeriesResponse, error)
	DropEvent(ctx context.Context, in *DropEventReq, opts ...grpc.CallOption) (*DropEventResp, error)
	GetEventID(ctx context.Context, in *GetEventIDReq, opts ...grpc.CallOption) (*GetEventIDResp, error)
	SetEventStatus(ctx context.Context, in *SetEventStatusRequest, opts ...grpc.CallOption) (*EventStatusStore, error)
	//Update
	UpdateCloseEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	UpdateTeamSolvedChallenge(ctx context.Context, in *UpdateTeamSolvedChallengeRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	UpdateTeamLastAccess(ctx context.Context, in *UpdateTeamLastAccessRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	UpdateTeamPassword(ctx context.Context, in *UpdateTeamPassRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete
	DeleteTeam(ctx context.Context, in *DelTeamRequest, opts ...grpc.CallOption) (*DelTeamResp, error)
}

type storeClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreClient(cc grpc.ClientConnInterface) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, "/store.Store/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) AddTeam(ctx context.Context, in *AddTeamRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, "/store.Store/AddTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetEvents(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error) {
	out := new(GetEventResponse)
	err := c.cc.Invoke(ctx, "/store.Store/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetEventByUser(ctx context.Context, in *GetEventByUserReq, opts ...grpc.CallOption) (*GetEventResponse, error) {
	out := new(GetEventResponse)
	err := c.cc.Invoke(ctx, "/store.Store/GetEventByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetEventTeams(ctx context.Context, in *GetEventTeamsRequest, opts ...grpc.CallOption) (*GetEventTeamsResponse, error) {
	out := new(GetEventTeamsResponse)
	err := c.cc.Invoke(ctx, "/store.Store/GetEventTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetEventStatus(ctx context.Context, in *GetEventStatusRequest, opts ...grpc.CallOption) (*EventStatusStore, error) {
	out := new(EventStatusStore)
	err := c.cc.Invoke(ctx, "/store.Store/GetEventStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) IsEventExists(ctx context.Context, in *GetEventByTagReq, opts ...grpc.CallOption) (*GetEventByTagResp, error) {
	out := new(GetEventByTagResp)
	err := c.cc.Invoke(ctx, "/store.Store/IsEventExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetTimeSeries(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetTimeSeriesResponse, error) {
	out := new(GetTimeSeriesResponse)
	err := c.cc.Invoke(ctx, "/store.Store/GetTimeSeries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) DropEvent(ctx context.Context, in *DropEventReq, opts ...grpc.CallOption) (*DropEventResp, error) {
	out := new(DropEventResp)
	err := c.cc.Invoke(ctx, "/store.Store/DropEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetEventID(ctx context.Context, in *GetEventIDReq, opts ...grpc.CallOption) (*GetEventIDResp, error) {
	out := new(GetEventIDResp)
	err := c.cc.Invoke(ctx, "/store.Store/GetEventID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) SetEventStatus(ctx context.Context, in *SetEventStatusRequest, opts ...grpc.CallOption) (*EventStatusStore, error) {
	out := new(EventStatusStore)
	err := c.cc.Invoke(ctx, "/store.Store/SetEventStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) UpdateCloseEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/store.Store/UpdateCloseEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) UpdateTeamSolvedChallenge(ctx context.Context, in *UpdateTeamSolvedChallengeRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/store.Store/UpdateTeamSolvedChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) UpdateTeamLastAccess(ctx context.Context, in *UpdateTeamLastAccessRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/store.Store/UpdateTeamLastAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) UpdateTeamPassword(ctx context.Context, in *UpdateTeamPassRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/store.Store/UpdateTeamPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) DeleteTeam(ctx context.Context, in *DelTeamRequest, opts ...grpc.CallOption) (*DelTeamResp, error) {
	out := new(DelTeamResp)
	err := c.cc.Invoke(ctx, "/store.Store/DeleteTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
// All implementations must embed UnimplementedStoreServer
// for forward compatibility
type StoreServer interface {
	//Insert
	AddEvent(context.Context, *AddEventRequest) (*InsertResponse, error)
	AddTeam(context.Context, *AddTeamRequest) (*InsertResponse, error)
	//Select
	GetEvents(context.Context, *GetEventRequest) (*GetEventResponse, error)
	GetEventByUser(context.Context, *GetEventByUserReq) (*GetEventResponse, error)
	GetEventTeams(context.Context, *GetEventTeamsRequest) (*GetEventTeamsResponse, error)
	GetEventStatus(context.Context, *GetEventStatusRequest) (*EventStatusStore, error)
	IsEventExists(context.Context, *GetEventByTagReq) (*GetEventByTagResp, error)
	GetTimeSeries(context.Context, *EmptyRequest) (*GetTimeSeriesResponse, error)
	DropEvent(context.Context, *DropEventReq) (*DropEventResp, error)
	GetEventID(context.Context, *GetEventIDReq) (*GetEventIDResp, error)
	SetEventStatus(context.Context, *SetEventStatusRequest) (*EventStatusStore, error)
	//Update
	UpdateCloseEvent(context.Context, *UpdateEventRequest) (*UpdateResponse, error)
	UpdateTeamSolvedChallenge(context.Context, *UpdateTeamSolvedChallengeRequest) (*UpdateResponse, error)
	UpdateTeamLastAccess(context.Context, *UpdateTeamLastAccessRequest) (*UpdateResponse, error)
	UpdateTeamPassword(context.Context, *UpdateTeamPassRequest) (*UpdateResponse, error)
	// Delete
	DeleteTeam(context.Context, *DelTeamRequest) (*DelTeamResp, error)
	mustEmbedUnimplementedStoreServer()
}

// UnimplementedStoreServer must be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (UnimplementedStoreServer) AddEvent(context.Context, *AddEventRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (UnimplementedStoreServer) AddTeam(context.Context, *AddTeamRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTeam not implemented")
}
func (UnimplementedStoreServer) GetEvents(context.Context, *GetEventRequest) (*GetEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedStoreServer) GetEventByUser(context.Context, *GetEventByUserReq) (*GetEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventByUser not implemented")
}
func (UnimplementedStoreServer) GetEventTeams(context.Context, *GetEventTeamsRequest) (*GetEventTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventTeams not implemented")
}
func (UnimplementedStoreServer) GetEventStatus(context.Context, *GetEventStatusRequest) (*EventStatusStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventStatus not implemented")
}
func (UnimplementedStoreServer) IsEventExists(context.Context, *GetEventByTagReq) (*GetEventByTagResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsEventExists not implemented")
}
func (UnimplementedStoreServer) GetTimeSeries(context.Context, *EmptyRequest) (*GetTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeries not implemented")
}
func (UnimplementedStoreServer) DropEvent(context.Context, *DropEventReq) (*DropEventResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropEvent not implemented")
}
func (UnimplementedStoreServer) GetEventID(context.Context, *GetEventIDReq) (*GetEventIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventID not implemented")
}
func (UnimplementedStoreServer) SetEventStatus(context.Context, *SetEventStatusRequest) (*EventStatusStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEventStatus not implemented")
}
func (UnimplementedStoreServer) UpdateCloseEvent(context.Context, *UpdateEventRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCloseEvent not implemented")
}
func (UnimplementedStoreServer) UpdateTeamSolvedChallenge(context.Context, *UpdateTeamSolvedChallengeRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeamSolvedChallenge not implemented")
}
func (UnimplementedStoreServer) UpdateTeamLastAccess(context.Context, *UpdateTeamLastAccessRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeamLastAccess not implemented")
}
func (UnimplementedStoreServer) UpdateTeamPassword(context.Context, *UpdateTeamPassRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeamPassword not implemented")
}
func (UnimplementedStoreServer) DeleteTeam(context.Context, *DelTeamRequest) (*DelTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (UnimplementedStoreServer) mustEmbedUnimplementedStoreServer() {}

// UnsafeStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreServer will
// result in compilation errors.
type UnsafeStoreServer interface {
	mustEmbedUnimplementedStoreServer()
}

func RegisterStoreServer(s grpc.ServiceRegistrar, srv StoreServer) {
	s.RegisterService(&Store_ServiceDesc, srv)
}

func _Store_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).AddEvent(ctx, req.(*AddEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_AddTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).AddTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/AddTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).AddTeam(ctx, req.(*AddTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetEvents(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetEventByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventByUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetEventByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/GetEventByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetEventByUser(ctx, req.(*GetEventByUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetEventTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetEventTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/GetEventTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetEventTeams(ctx, req.(*GetEventTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetEventStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetEventStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/GetEventStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetEventStatus(ctx, req.(*GetEventStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_IsEventExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventByTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).IsEventExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/IsEventExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).IsEventExists(ctx, req.(*GetEventByTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/GetTimeSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetTimeSeries(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_DropEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).DropEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/DropEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).DropEvent(ctx, req.(*DropEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetEventID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetEventID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/GetEventID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetEventID(ctx, req.(*GetEventIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_SetEventStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEventStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).SetEventStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/SetEventStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).SetEventStatus(ctx, req.(*SetEventStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_UpdateCloseEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).UpdateCloseEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/UpdateCloseEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).UpdateCloseEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_UpdateTeamSolvedChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamSolvedChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).UpdateTeamSolvedChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/UpdateTeamSolvedChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).UpdateTeamSolvedChallenge(ctx, req.(*UpdateTeamSolvedChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_UpdateTeamLastAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamLastAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).UpdateTeamLastAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/UpdateTeamLastAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).UpdateTeamLastAccess(ctx, req.(*UpdateTeamLastAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_UpdateTeamPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).UpdateTeamPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/UpdateTeamPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).UpdateTeamPassword(ctx, req.(*UpdateTeamPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/DeleteTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).DeleteTeam(ctx, req.(*DelTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Store_ServiceDesc is the grpc.ServiceDesc for Store service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Store_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "store.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _Store_AddEvent_Handler,
		},
		{
			MethodName: "AddTeam",
			Handler:    _Store_AddTeam_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _Store_GetEvents_Handler,
		},
		{
			MethodName: "GetEventByUser",
			Handler:    _Store_GetEventByUser_Handler,
		},
		{
			MethodName: "GetEventTeams",
			Handler:    _Store_GetEventTeams_Handler,
		},
		{
			MethodName: "GetEventStatus",
			Handler:    _Store_GetEventStatus_Handler,
		},
		{
			MethodName: "IsEventExists",
			Handler:    _Store_IsEventExists_Handler,
		},
		{
			MethodName: "GetTimeSeries",
			Handler:    _Store_GetTimeSeries_Handler,
		},
		{
			MethodName: "DropEvent",
			Handler:    _Store_DropEvent_Handler,
		},
		{
			MethodName: "GetEventID",
			Handler:    _Store_GetEventID_Handler,
		},
		{
			MethodName: "SetEventStatus",
			Handler:    _Store_SetEventStatus_Handler,
		},
		{
			MethodName: "UpdateCloseEvent",
			Handler:    _Store_UpdateCloseEvent_Handler,
		},
		{
			MethodName: "UpdateTeamSolvedChallenge",
			Handler:    _Store_UpdateTeamSolvedChallenge_Handler,
		},
		{
			MethodName: "UpdateTeamLastAccess",
			Handler:    _Store_UpdateTeamLastAccess_Handler,
		},
		{
			MethodName: "UpdateTeamPassword",
			Handler:    _Store_UpdateTeamPassword_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _Store_DeleteTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store.proto",
}