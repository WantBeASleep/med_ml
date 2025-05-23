// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/grpc/clients/uzi.proto

package uzi

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UziSrv_CreateDevice_FullMethodName                  = "/UziSrv/createDevice"
	UziSrv_GetDeviceList_FullMethodName                 = "/UziSrv/getDeviceList"
	UziSrv_CreateUzi_FullMethodName                     = "/UziSrv/createUzi"
	UziSrv_GetUziById_FullMethodName                    = "/UziSrv/getUziById"
	UziSrv_GetUzisByExternalId_FullMethodName           = "/UziSrv/getUzisByExternalId"
	UziSrv_GetUzisByAuthor_FullMethodName               = "/UziSrv/getUzisByAuthor"
	UziSrv_GetEchographicByUziId_FullMethodName         = "/UziSrv/getEchographicByUziId"
	UziSrv_UpdateUzi_FullMethodName                     = "/UziSrv/updateUzi"
	UziSrv_UpdateEchographic_FullMethodName             = "/UziSrv/updateEchographic"
	UziSrv_DeleteUzi_FullMethodName                     = "/UziSrv/deleteUzi"
	UziSrv_GetImagesByUziId_FullMethodName              = "/UziSrv/getImagesByUziId"
	UziSrv_GetNodesByUziId_FullMethodName               = "/UziSrv/getNodesByUziId"
	UziSrv_UpdateNode_FullMethodName                    = "/UziSrv/updateNode"
	UziSrv_CreateSegment_FullMethodName                 = "/UziSrv/createSegment"
	UziSrv_GetSegmentsByNodeId_FullMethodName           = "/UziSrv/getSegmentsByNodeId"
	UziSrv_UpdateSegment_FullMethodName                 = "/UziSrv/updateSegment"
	UziSrv_CreateNodeWithSegments_FullMethodName        = "/UziSrv/createNodeWithSegments"
	UziSrv_GetNodesWithSegmentsByImageId_FullMethodName = "/UziSrv/getNodesWithSegmentsByImageId"
	UziSrv_DeleteNode_FullMethodName                    = "/UziSrv/deleteNode"
	UziSrv_DeleteSegment_FullMethodName                 = "/UziSrv/deleteSegment"
)

// UziSrvClient is the client API for UziSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UziSrvClient interface {
	// DEVICE
	CreateDevice(ctx context.Context, in *CreateDeviceIn, opts ...grpc.CallOption) (*CreateDeviceOut, error)
	GetDeviceList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDeviceListOut, error)
	// UZI
	CreateUzi(ctx context.Context, in *CreateUziIn, opts ...grpc.CallOption) (*CreateUziOut, error)
	GetUziById(ctx context.Context, in *GetUziByIdIn, opts ...grpc.CallOption) (*GetUziByIdOut, error)
	GetUzisByExternalId(ctx context.Context, in *GetUzisByExternalIdIn, opts ...grpc.CallOption) (*GetUzisByExternalIdOut, error)
	GetUzisByAuthor(ctx context.Context, in *GetUzisByAuthorIn, opts ...grpc.CallOption) (*GetUzisByAuthorOut, error)
	GetEchographicByUziId(ctx context.Context, in *GetEchographicByUziIdIn, opts ...grpc.CallOption) (*GetEchographicByUziIdOut, error)
	UpdateUzi(ctx context.Context, in *UpdateUziIn, opts ...grpc.CallOption) (*UpdateUziOut, error)
	UpdateEchographic(ctx context.Context, in *UpdateEchographicIn, opts ...grpc.CallOption) (*UpdateEchographicOut, error)
	DeleteUzi(ctx context.Context, in *DeleteUziIn, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// IMAGE
	GetImagesByUziId(ctx context.Context, in *GetImagesByUziIdIn, opts ...grpc.CallOption) (*GetImagesByUziIdOut, error)
	// NODE
	GetNodesByUziId(ctx context.Context, in *GetNodesByUziIdIn, opts ...grpc.CallOption) (*GetNodesByUziIdOut, error)
	UpdateNode(ctx context.Context, in *UpdateNodeIn, opts ...grpc.CallOption) (*UpdateNodeOut, error)
	// SEGMENT
	CreateSegment(ctx context.Context, in *CreateSegmentIn, opts ...grpc.CallOption) (*CreateSegmentOut, error)
	GetSegmentsByNodeId(ctx context.Context, in *GetSegmentsByNodeIdIn, opts ...grpc.CallOption) (*GetSegmentsByNodeIdOut, error)
	UpdateSegment(ctx context.Context, in *UpdateSegmentIn, opts ...grpc.CallOption) (*UpdateSegmentOut, error)
	// доменные области слишком сильно пересекаются, вынесено в одну надобласть
	// NODE-SEGMENT
	CreateNodeWithSegments(ctx context.Context, in *CreateNodeWithSegmentsIn, opts ...grpc.CallOption) (*CreateNodeWithSegmentsOut, error)
	GetNodesWithSegmentsByImageId(ctx context.Context, in *GetNodesWithSegmentsByImageIdIn, opts ...grpc.CallOption) (*GetNodesWithSegmentsByImageIdOut, error)
	DeleteNode(ctx context.Context, in *DeleteNodeIn, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteSegment(ctx context.Context, in *DeleteSegmentIn, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type uziSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewUziSrvClient(cc grpc.ClientConnInterface) UziSrvClient {
	return &uziSrvClient{cc}
}

func (c *uziSrvClient) CreateDevice(ctx context.Context, in *CreateDeviceIn, opts ...grpc.CallOption) (*CreateDeviceOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDeviceOut)
	err := c.cc.Invoke(ctx, UziSrv_CreateDevice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetDeviceList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDeviceListOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDeviceListOut)
	err := c.cc.Invoke(ctx, UziSrv_GetDeviceList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) CreateUzi(ctx context.Context, in *CreateUziIn, opts ...grpc.CallOption) (*CreateUziOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUziOut)
	err := c.cc.Invoke(ctx, UziSrv_CreateUzi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetUziById(ctx context.Context, in *GetUziByIdIn, opts ...grpc.CallOption) (*GetUziByIdOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUziByIdOut)
	err := c.cc.Invoke(ctx, UziSrv_GetUziById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetUzisByExternalId(ctx context.Context, in *GetUzisByExternalIdIn, opts ...grpc.CallOption) (*GetUzisByExternalIdOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUzisByExternalIdOut)
	err := c.cc.Invoke(ctx, UziSrv_GetUzisByExternalId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetUzisByAuthor(ctx context.Context, in *GetUzisByAuthorIn, opts ...grpc.CallOption) (*GetUzisByAuthorOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUzisByAuthorOut)
	err := c.cc.Invoke(ctx, UziSrv_GetUzisByAuthor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetEchographicByUziId(ctx context.Context, in *GetEchographicByUziIdIn, opts ...grpc.CallOption) (*GetEchographicByUziIdOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEchographicByUziIdOut)
	err := c.cc.Invoke(ctx, UziSrv_GetEchographicByUziId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) UpdateUzi(ctx context.Context, in *UpdateUziIn, opts ...grpc.CallOption) (*UpdateUziOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUziOut)
	err := c.cc.Invoke(ctx, UziSrv_UpdateUzi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) UpdateEchographic(ctx context.Context, in *UpdateEchographicIn, opts ...grpc.CallOption) (*UpdateEchographicOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEchographicOut)
	err := c.cc.Invoke(ctx, UziSrv_UpdateEchographic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) DeleteUzi(ctx context.Context, in *DeleteUziIn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UziSrv_DeleteUzi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetImagesByUziId(ctx context.Context, in *GetImagesByUziIdIn, opts ...grpc.CallOption) (*GetImagesByUziIdOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetImagesByUziIdOut)
	err := c.cc.Invoke(ctx, UziSrv_GetImagesByUziId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetNodesByUziId(ctx context.Context, in *GetNodesByUziIdIn, opts ...grpc.CallOption) (*GetNodesByUziIdOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNodesByUziIdOut)
	err := c.cc.Invoke(ctx, UziSrv_GetNodesByUziId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) UpdateNode(ctx context.Context, in *UpdateNodeIn, opts ...grpc.CallOption) (*UpdateNodeOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNodeOut)
	err := c.cc.Invoke(ctx, UziSrv_UpdateNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) CreateSegment(ctx context.Context, in *CreateSegmentIn, opts ...grpc.CallOption) (*CreateSegmentOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSegmentOut)
	err := c.cc.Invoke(ctx, UziSrv_CreateSegment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetSegmentsByNodeId(ctx context.Context, in *GetSegmentsByNodeIdIn, opts ...grpc.CallOption) (*GetSegmentsByNodeIdOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSegmentsByNodeIdOut)
	err := c.cc.Invoke(ctx, UziSrv_GetSegmentsByNodeId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) UpdateSegment(ctx context.Context, in *UpdateSegmentIn, opts ...grpc.CallOption) (*UpdateSegmentOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSegmentOut)
	err := c.cc.Invoke(ctx, UziSrv_UpdateSegment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) CreateNodeWithSegments(ctx context.Context, in *CreateNodeWithSegmentsIn, opts ...grpc.CallOption) (*CreateNodeWithSegmentsOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNodeWithSegmentsOut)
	err := c.cc.Invoke(ctx, UziSrv_CreateNodeWithSegments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) GetNodesWithSegmentsByImageId(ctx context.Context, in *GetNodesWithSegmentsByImageIdIn, opts ...grpc.CallOption) (*GetNodesWithSegmentsByImageIdOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNodesWithSegmentsByImageIdOut)
	err := c.cc.Invoke(ctx, UziSrv_GetNodesWithSegmentsByImageId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) DeleteNode(ctx context.Context, in *DeleteNodeIn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UziSrv_DeleteNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uziSrvClient) DeleteSegment(ctx context.Context, in *DeleteSegmentIn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UziSrv_DeleteSegment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UziSrvServer is the server API for UziSrv service.
// All implementations must embed UnimplementedUziSrvServer
// for forward compatibility.
type UziSrvServer interface {
	// DEVICE
	CreateDevice(context.Context, *CreateDeviceIn) (*CreateDeviceOut, error)
	GetDeviceList(context.Context, *emptypb.Empty) (*GetDeviceListOut, error)
	// UZI
	CreateUzi(context.Context, *CreateUziIn) (*CreateUziOut, error)
	GetUziById(context.Context, *GetUziByIdIn) (*GetUziByIdOut, error)
	GetUzisByExternalId(context.Context, *GetUzisByExternalIdIn) (*GetUzisByExternalIdOut, error)
	GetUzisByAuthor(context.Context, *GetUzisByAuthorIn) (*GetUzisByAuthorOut, error)
	GetEchographicByUziId(context.Context, *GetEchographicByUziIdIn) (*GetEchographicByUziIdOut, error)
	UpdateUzi(context.Context, *UpdateUziIn) (*UpdateUziOut, error)
	UpdateEchographic(context.Context, *UpdateEchographicIn) (*UpdateEchographicOut, error)
	DeleteUzi(context.Context, *DeleteUziIn) (*emptypb.Empty, error)
	// IMAGE
	GetImagesByUziId(context.Context, *GetImagesByUziIdIn) (*GetImagesByUziIdOut, error)
	// NODE
	GetNodesByUziId(context.Context, *GetNodesByUziIdIn) (*GetNodesByUziIdOut, error)
	UpdateNode(context.Context, *UpdateNodeIn) (*UpdateNodeOut, error)
	// SEGMENT
	CreateSegment(context.Context, *CreateSegmentIn) (*CreateSegmentOut, error)
	GetSegmentsByNodeId(context.Context, *GetSegmentsByNodeIdIn) (*GetSegmentsByNodeIdOut, error)
	UpdateSegment(context.Context, *UpdateSegmentIn) (*UpdateSegmentOut, error)
	// доменные области слишком сильно пересекаются, вынесено в одну надобласть
	// NODE-SEGMENT
	CreateNodeWithSegments(context.Context, *CreateNodeWithSegmentsIn) (*CreateNodeWithSegmentsOut, error)
	GetNodesWithSegmentsByImageId(context.Context, *GetNodesWithSegmentsByImageIdIn) (*GetNodesWithSegmentsByImageIdOut, error)
	DeleteNode(context.Context, *DeleteNodeIn) (*emptypb.Empty, error)
	DeleteSegment(context.Context, *DeleteSegmentIn) (*emptypb.Empty, error)
	mustEmbedUnimplementedUziSrvServer()
}

// UnimplementedUziSrvServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUziSrvServer struct{}

func (UnimplementedUziSrvServer) CreateDevice(context.Context, *CreateDeviceIn) (*CreateDeviceOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedUziSrvServer) GetDeviceList(context.Context, *emptypb.Empty) (*GetDeviceListOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceList not implemented")
}
func (UnimplementedUziSrvServer) CreateUzi(context.Context, *CreateUziIn) (*CreateUziOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUzi not implemented")
}
func (UnimplementedUziSrvServer) GetUziById(context.Context, *GetUziByIdIn) (*GetUziByIdOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUziById not implemented")
}
func (UnimplementedUziSrvServer) GetUzisByExternalId(context.Context, *GetUzisByExternalIdIn) (*GetUzisByExternalIdOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUzisByExternalId not implemented")
}
func (UnimplementedUziSrvServer) GetUzisByAuthor(context.Context, *GetUzisByAuthorIn) (*GetUzisByAuthorOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUzisByAuthor not implemented")
}
func (UnimplementedUziSrvServer) GetEchographicByUziId(context.Context, *GetEchographicByUziIdIn) (*GetEchographicByUziIdOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEchographicByUziId not implemented")
}
func (UnimplementedUziSrvServer) UpdateUzi(context.Context, *UpdateUziIn) (*UpdateUziOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUzi not implemented")
}
func (UnimplementedUziSrvServer) UpdateEchographic(context.Context, *UpdateEchographicIn) (*UpdateEchographicOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEchographic not implemented")
}
func (UnimplementedUziSrvServer) DeleteUzi(context.Context, *DeleteUziIn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUzi not implemented")
}
func (UnimplementedUziSrvServer) GetImagesByUziId(context.Context, *GetImagesByUziIdIn) (*GetImagesByUziIdOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImagesByUziId not implemented")
}
func (UnimplementedUziSrvServer) GetNodesByUziId(context.Context, *GetNodesByUziIdIn) (*GetNodesByUziIdOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodesByUziId not implemented")
}
func (UnimplementedUziSrvServer) UpdateNode(context.Context, *UpdateNodeIn) (*UpdateNodeOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNode not implemented")
}
func (UnimplementedUziSrvServer) CreateSegment(context.Context, *CreateSegmentIn) (*CreateSegmentOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSegment not implemented")
}
func (UnimplementedUziSrvServer) GetSegmentsByNodeId(context.Context, *GetSegmentsByNodeIdIn) (*GetSegmentsByNodeIdOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSegmentsByNodeId not implemented")
}
func (UnimplementedUziSrvServer) UpdateSegment(context.Context, *UpdateSegmentIn) (*UpdateSegmentOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSegment not implemented")
}
func (UnimplementedUziSrvServer) CreateNodeWithSegments(context.Context, *CreateNodeWithSegmentsIn) (*CreateNodeWithSegmentsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeWithSegments not implemented")
}
func (UnimplementedUziSrvServer) GetNodesWithSegmentsByImageId(context.Context, *GetNodesWithSegmentsByImageIdIn) (*GetNodesWithSegmentsByImageIdOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodesWithSegmentsByImageId not implemented")
}
func (UnimplementedUziSrvServer) DeleteNode(context.Context, *DeleteNodeIn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}
func (UnimplementedUziSrvServer) DeleteSegment(context.Context, *DeleteSegmentIn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSegment not implemented")
}
func (UnimplementedUziSrvServer) mustEmbedUnimplementedUziSrvServer() {}
func (UnimplementedUziSrvServer) testEmbeddedByValue()                {}

// UnsafeUziSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UziSrvServer will
// result in compilation errors.
type UnsafeUziSrvServer interface {
	mustEmbedUnimplementedUziSrvServer()
}

func RegisterUziSrvServer(s grpc.ServiceRegistrar, srv UziSrvServer) {
	// If the following call pancis, it indicates UnimplementedUziSrvServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UziSrv_ServiceDesc, srv)
}

func _UziSrv_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_CreateDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).CreateDevice(ctx, req.(*CreateDeviceIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetDeviceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetDeviceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetDeviceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetDeviceList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_CreateUzi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUziIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).CreateUzi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_CreateUzi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).CreateUzi(ctx, req.(*CreateUziIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetUziById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUziByIdIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetUziById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetUziById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetUziById(ctx, req.(*GetUziByIdIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetUzisByExternalId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUzisByExternalIdIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetUzisByExternalId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetUzisByExternalId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetUzisByExternalId(ctx, req.(*GetUzisByExternalIdIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetUzisByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUzisByAuthorIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetUzisByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetUzisByAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetUzisByAuthor(ctx, req.(*GetUzisByAuthorIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetEchographicByUziId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEchographicByUziIdIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetEchographicByUziId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetEchographicByUziId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetEchographicByUziId(ctx, req.(*GetEchographicByUziIdIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_UpdateUzi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUziIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).UpdateUzi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_UpdateUzi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).UpdateUzi(ctx, req.(*UpdateUziIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_UpdateEchographic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEchographicIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).UpdateEchographic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_UpdateEchographic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).UpdateEchographic(ctx, req.(*UpdateEchographicIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_DeleteUzi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUziIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).DeleteUzi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_DeleteUzi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).DeleteUzi(ctx, req.(*DeleteUziIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetImagesByUziId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImagesByUziIdIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetImagesByUziId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetImagesByUziId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetImagesByUziId(ctx, req.(*GetImagesByUziIdIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetNodesByUziId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodesByUziIdIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetNodesByUziId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetNodesByUziId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetNodesByUziId(ctx, req.(*GetNodesByUziIdIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_UpdateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).UpdateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_UpdateNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).UpdateNode(ctx, req.(*UpdateNodeIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_CreateSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSegmentIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).CreateSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_CreateSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).CreateSegment(ctx, req.(*CreateSegmentIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetSegmentsByNodeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSegmentsByNodeIdIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetSegmentsByNodeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetSegmentsByNodeId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetSegmentsByNodeId(ctx, req.(*GetSegmentsByNodeIdIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_UpdateSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSegmentIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).UpdateSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_UpdateSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).UpdateSegment(ctx, req.(*UpdateSegmentIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_CreateNodeWithSegments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeWithSegmentsIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).CreateNodeWithSegments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_CreateNodeWithSegments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).CreateNodeWithSegments(ctx, req.(*CreateNodeWithSegmentsIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_GetNodesWithSegmentsByImageId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodesWithSegmentsByImageIdIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).GetNodesWithSegmentsByImageId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_GetNodesWithSegmentsByImageId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).GetNodesWithSegmentsByImageId(ctx, req.(*GetNodesWithSegmentsByImageIdIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_DeleteNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).DeleteNode(ctx, req.(*DeleteNodeIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UziSrv_DeleteSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSegmentIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UziSrvServer).DeleteSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UziSrv_DeleteSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UziSrvServer).DeleteSegment(ctx, req.(*DeleteSegmentIn))
	}
	return interceptor(ctx, in, info, handler)
}

// UziSrv_ServiceDesc is the grpc.ServiceDesc for UziSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UziSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UziSrv",
	HandlerType: (*UziSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createDevice",
			Handler:    _UziSrv_CreateDevice_Handler,
		},
		{
			MethodName: "getDeviceList",
			Handler:    _UziSrv_GetDeviceList_Handler,
		},
		{
			MethodName: "createUzi",
			Handler:    _UziSrv_CreateUzi_Handler,
		},
		{
			MethodName: "getUziById",
			Handler:    _UziSrv_GetUziById_Handler,
		},
		{
			MethodName: "getUzisByExternalId",
			Handler:    _UziSrv_GetUzisByExternalId_Handler,
		},
		{
			MethodName: "getUzisByAuthor",
			Handler:    _UziSrv_GetUzisByAuthor_Handler,
		},
		{
			MethodName: "getEchographicByUziId",
			Handler:    _UziSrv_GetEchographicByUziId_Handler,
		},
		{
			MethodName: "updateUzi",
			Handler:    _UziSrv_UpdateUzi_Handler,
		},
		{
			MethodName: "updateEchographic",
			Handler:    _UziSrv_UpdateEchographic_Handler,
		},
		{
			MethodName: "deleteUzi",
			Handler:    _UziSrv_DeleteUzi_Handler,
		},
		{
			MethodName: "getImagesByUziId",
			Handler:    _UziSrv_GetImagesByUziId_Handler,
		},
		{
			MethodName: "getNodesByUziId",
			Handler:    _UziSrv_GetNodesByUziId_Handler,
		},
		{
			MethodName: "updateNode",
			Handler:    _UziSrv_UpdateNode_Handler,
		},
		{
			MethodName: "createSegment",
			Handler:    _UziSrv_CreateSegment_Handler,
		},
		{
			MethodName: "getSegmentsByNodeId",
			Handler:    _UziSrv_GetSegmentsByNodeId_Handler,
		},
		{
			MethodName: "updateSegment",
			Handler:    _UziSrv_UpdateSegment_Handler,
		},
		{
			MethodName: "createNodeWithSegments",
			Handler:    _UziSrv_CreateNodeWithSegments_Handler,
		},
		{
			MethodName: "getNodesWithSegmentsByImageId",
			Handler:    _UziSrv_GetNodesWithSegmentsByImageId_Handler,
		},
		{
			MethodName: "deleteNode",
			Handler:    _UziSrv_DeleteNode_Handler,
		},
		{
			MethodName: "deleteSegment",
			Handler:    _UziSrv_DeleteSegment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc/clients/uzi.proto",
}
