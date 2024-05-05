// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/learning/learning.proto

package learning

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

// LearningServiceClient is the client API for LearningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LearningServiceClient interface {
	CreateClassRooms(ctx context.Context, in *CreateClassRoomsRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	CreateClassRoom(ctx context.Context, in *CreateClassRoomRequest, opts ...grpc.CallOption) (*ClassRoom, error)
	GetClassRooms(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ClassRooms, error)
	GetMyClassRooms(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ClassRooms, error)
	GetClassRoom(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ClassRoom, error)
	DeleteClassRoom(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	DeleteClassRoomsByTeacher(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	// CHAPTER
	CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*Chapter, error)
	GetChaptersByClassRoom(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Chapters, error)
	UpdateChapter(ctx context.Context, in *UpdateChapterRequest, opts ...grpc.CallOption) (*Chapter, error)
	DeleteChapter(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	// LESSONS
	GetLessonsByChapter(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Lessons, error)
	CreateLesson(ctx context.Context, in *CreateLessonRequest, opts ...grpc.CallOption) (*Lesson, error)
	UpdateLesson(ctx context.Context, in *UpdateLessonRequest, opts ...grpc.CallOption) (*Lesson, error)
	DeleteLesson(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	GetLessonsByStudent(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StudentLessons, error)
	// DOCUMENT
	CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*Document, error)
	GetDocumentsByLesson(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Documents, error)
}

type learningServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLearningServiceClient(cc grpc.ClientConnInterface) LearningServiceClient {
	return &learningServiceClient{cc}
}

func (c *learningServiceClient) CreateClassRooms(ctx context.Context, in *CreateClassRoomsRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/learning.LearningService/CreateClassRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) CreateClassRoom(ctx context.Context, in *CreateClassRoomRequest, opts ...grpc.CallOption) (*ClassRoom, error) {
	out := new(ClassRoom)
	err := c.cc.Invoke(ctx, "/learning.LearningService/CreateClassRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) GetClassRooms(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ClassRooms, error) {
	out := new(ClassRooms)
	err := c.cc.Invoke(ctx, "/learning.LearningService/GetClassRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) GetMyClassRooms(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ClassRooms, error) {
	out := new(ClassRooms)
	err := c.cc.Invoke(ctx, "/learning.LearningService/GetMyClassRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) GetClassRoom(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ClassRoom, error) {
	out := new(ClassRoom)
	err := c.cc.Invoke(ctx, "/learning.LearningService/GetClassRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) DeleteClassRoom(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/learning.LearningService/DeleteClassRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) DeleteClassRoomsByTeacher(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/learning.LearningService/DeleteClassRoomsByTeacher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*Chapter, error) {
	out := new(Chapter)
	err := c.cc.Invoke(ctx, "/learning.LearningService/CreateChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) GetChaptersByClassRoom(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Chapters, error) {
	out := new(Chapters)
	err := c.cc.Invoke(ctx, "/learning.LearningService/GetChaptersByClassRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) UpdateChapter(ctx context.Context, in *UpdateChapterRequest, opts ...grpc.CallOption) (*Chapter, error) {
	out := new(Chapter)
	err := c.cc.Invoke(ctx, "/learning.LearningService/UpdateChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) DeleteChapter(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/learning.LearningService/DeleteChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) GetLessonsByChapter(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Lessons, error) {
	out := new(Lessons)
	err := c.cc.Invoke(ctx, "/learning.LearningService/GetLessonsByChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) CreateLesson(ctx context.Context, in *CreateLessonRequest, opts ...grpc.CallOption) (*Lesson, error) {
	out := new(Lesson)
	err := c.cc.Invoke(ctx, "/learning.LearningService/CreateLesson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) UpdateLesson(ctx context.Context, in *UpdateLessonRequest, opts ...grpc.CallOption) (*Lesson, error) {
	out := new(Lesson)
	err := c.cc.Invoke(ctx, "/learning.LearningService/UpdateLesson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) DeleteLesson(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/learning.LearningService/DeleteLesson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) GetLessonsByStudent(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StudentLessons, error) {
	out := new(StudentLessons)
	err := c.cc.Invoke(ctx, "/learning.LearningService/GetLessonsByStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, "/learning.LearningService/CreateVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) GetDocumentsByLesson(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Documents, error) {
	out := new(Documents)
	err := c.cc.Invoke(ctx, "/learning.LearningService/GetDocumentsByLesson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LearningServiceServer is the server API for LearningService service.
// All implementations must embed UnimplementedLearningServiceServer
// for forward compatibility
type LearningServiceServer interface {
	CreateClassRooms(context.Context, *CreateClassRoomsRequest) (*OperationStatus, error)
	CreateClassRoom(context.Context, *CreateClassRoomRequest) (*ClassRoom, error)
	GetClassRooms(context.Context, *IDRequest) (*ClassRooms, error)
	GetMyClassRooms(context.Context, *IDRequest) (*ClassRooms, error)
	GetClassRoom(context.Context, *IDRequest) (*ClassRoom, error)
	DeleteClassRoom(context.Context, *IDRequest) (*OperationStatus, error)
	DeleteClassRoomsByTeacher(context.Context, *IDRequest) (*OperationStatus, error)
	// CHAPTER
	CreateChapter(context.Context, *CreateChapterRequest) (*Chapter, error)
	GetChaptersByClassRoom(context.Context, *IDRequest) (*Chapters, error)
	UpdateChapter(context.Context, *UpdateChapterRequest) (*Chapter, error)
	DeleteChapter(context.Context, *IDRequest) (*OperationStatus, error)
	// LESSONS
	GetLessonsByChapter(context.Context, *IDRequest) (*Lessons, error)
	CreateLesson(context.Context, *CreateLessonRequest) (*Lesson, error)
	UpdateLesson(context.Context, *UpdateLessonRequest) (*Lesson, error)
	DeleteLesson(context.Context, *IDRequest) (*OperationStatus, error)
	GetLessonsByStudent(context.Context, *IDRequest) (*StudentLessons, error)
	// DOCUMENT
	CreateVideo(context.Context, *CreateVideoRequest) (*Document, error)
	GetDocumentsByLesson(context.Context, *IDRequest) (*Documents, error)
	mustEmbedUnimplementedLearningServiceServer()
}

// UnimplementedLearningServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLearningServiceServer struct {
}

func (UnimplementedLearningServiceServer) CreateClassRooms(context.Context, *CreateClassRoomsRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClassRooms not implemented")
}
func (UnimplementedLearningServiceServer) CreateClassRoom(context.Context, *CreateClassRoomRequest) (*ClassRoom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClassRoom not implemented")
}
func (UnimplementedLearningServiceServer) GetClassRooms(context.Context, *IDRequest) (*ClassRooms, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassRooms not implemented")
}
func (UnimplementedLearningServiceServer) GetMyClassRooms(context.Context, *IDRequest) (*ClassRooms, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyClassRooms not implemented")
}
func (UnimplementedLearningServiceServer) GetClassRoom(context.Context, *IDRequest) (*ClassRoom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassRoom not implemented")
}
func (UnimplementedLearningServiceServer) DeleteClassRoom(context.Context, *IDRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClassRoom not implemented")
}
func (UnimplementedLearningServiceServer) DeleteClassRoomsByTeacher(context.Context, *IDRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClassRoomsByTeacher not implemented")
}
func (UnimplementedLearningServiceServer) CreateChapter(context.Context, *CreateChapterRequest) (*Chapter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChapter not implemented")
}
func (UnimplementedLearningServiceServer) GetChaptersByClassRoom(context.Context, *IDRequest) (*Chapters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChaptersByClassRoom not implemented")
}
func (UnimplementedLearningServiceServer) UpdateChapter(context.Context, *UpdateChapterRequest) (*Chapter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChapter not implemented")
}
func (UnimplementedLearningServiceServer) DeleteChapter(context.Context, *IDRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChapter not implemented")
}
func (UnimplementedLearningServiceServer) GetLessonsByChapter(context.Context, *IDRequest) (*Lessons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLessonsByChapter not implemented")
}
func (UnimplementedLearningServiceServer) CreateLesson(context.Context, *CreateLessonRequest) (*Lesson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLesson not implemented")
}
func (UnimplementedLearningServiceServer) UpdateLesson(context.Context, *UpdateLessonRequest) (*Lesson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLesson not implemented")
}
func (UnimplementedLearningServiceServer) DeleteLesson(context.Context, *IDRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLesson not implemented")
}
func (UnimplementedLearningServiceServer) GetLessonsByStudent(context.Context, *IDRequest) (*StudentLessons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLessonsByStudent not implemented")
}
func (UnimplementedLearningServiceServer) CreateVideo(context.Context, *CreateVideoRequest) (*Document, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVideo not implemented")
}
func (UnimplementedLearningServiceServer) GetDocumentsByLesson(context.Context, *IDRequest) (*Documents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDocumentsByLesson not implemented")
}
func (UnimplementedLearningServiceServer) mustEmbedUnimplementedLearningServiceServer() {}

// UnsafeLearningServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LearningServiceServer will
// result in compilation errors.
type UnsafeLearningServiceServer interface {
	mustEmbedUnimplementedLearningServiceServer()
}

func RegisterLearningServiceServer(s grpc.ServiceRegistrar, srv LearningServiceServer) {
	s.RegisterService(&LearningService_ServiceDesc, srv)
}

func _LearningService_CreateClassRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClassRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).CreateClassRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/CreateClassRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).CreateClassRooms(ctx, req.(*CreateClassRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_CreateClassRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClassRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).CreateClassRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/CreateClassRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).CreateClassRoom(ctx, req.(*CreateClassRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_GetClassRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).GetClassRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/GetClassRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).GetClassRooms(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_GetMyClassRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).GetMyClassRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/GetMyClassRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).GetMyClassRooms(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_GetClassRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).GetClassRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/GetClassRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).GetClassRoom(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_DeleteClassRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).DeleteClassRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/DeleteClassRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).DeleteClassRoom(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_DeleteClassRoomsByTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).DeleteClassRoomsByTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/DeleteClassRoomsByTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).DeleteClassRoomsByTeacher(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_CreateChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).CreateChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/CreateChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).CreateChapter(ctx, req.(*CreateChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_GetChaptersByClassRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).GetChaptersByClassRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/GetChaptersByClassRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).GetChaptersByClassRoom(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_UpdateChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).UpdateChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/UpdateChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).UpdateChapter(ctx, req.(*UpdateChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_DeleteChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).DeleteChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/DeleteChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).DeleteChapter(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_GetLessonsByChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).GetLessonsByChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/GetLessonsByChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).GetLessonsByChapter(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_CreateLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLessonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).CreateLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/CreateLesson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).CreateLesson(ctx, req.(*CreateLessonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_UpdateLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLessonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).UpdateLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/UpdateLesson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).UpdateLesson(ctx, req.(*UpdateLessonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_DeleteLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).DeleteLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/DeleteLesson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).DeleteLesson(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_GetLessonsByStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).GetLessonsByStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/GetLessonsByStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).GetLessonsByStudent(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_CreateVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).CreateVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/CreateVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).CreateVideo(ctx, req.(*CreateVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_GetDocumentsByLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).GetDocumentsByLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learning.LearningService/GetDocumentsByLesson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).GetDocumentsByLesson(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LearningService_ServiceDesc is the grpc.ServiceDesc for LearningService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LearningService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "learning.LearningService",
	HandlerType: (*LearningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClassRooms",
			Handler:    _LearningService_CreateClassRooms_Handler,
		},
		{
			MethodName: "CreateClassRoom",
			Handler:    _LearningService_CreateClassRoom_Handler,
		},
		{
			MethodName: "GetClassRooms",
			Handler:    _LearningService_GetClassRooms_Handler,
		},
		{
			MethodName: "GetMyClassRooms",
			Handler:    _LearningService_GetMyClassRooms_Handler,
		},
		{
			MethodName: "GetClassRoom",
			Handler:    _LearningService_GetClassRoom_Handler,
		},
		{
			MethodName: "DeleteClassRoom",
			Handler:    _LearningService_DeleteClassRoom_Handler,
		},
		{
			MethodName: "DeleteClassRoomsByTeacher",
			Handler:    _LearningService_DeleteClassRoomsByTeacher_Handler,
		},
		{
			MethodName: "CreateChapter",
			Handler:    _LearningService_CreateChapter_Handler,
		},
		{
			MethodName: "GetChaptersByClassRoom",
			Handler:    _LearningService_GetChaptersByClassRoom_Handler,
		},
		{
			MethodName: "UpdateChapter",
			Handler:    _LearningService_UpdateChapter_Handler,
		},
		{
			MethodName: "DeleteChapter",
			Handler:    _LearningService_DeleteChapter_Handler,
		},
		{
			MethodName: "GetLessonsByChapter",
			Handler:    _LearningService_GetLessonsByChapter_Handler,
		},
		{
			MethodName: "CreateLesson",
			Handler:    _LearningService_CreateLesson_Handler,
		},
		{
			MethodName: "UpdateLesson",
			Handler:    _LearningService_UpdateLesson_Handler,
		},
		{
			MethodName: "DeleteLesson",
			Handler:    _LearningService_DeleteLesson_Handler,
		},
		{
			MethodName: "GetLessonsByStudent",
			Handler:    _LearningService_GetLessonsByStudent_Handler,
		},
		{
			MethodName: "CreateVideo",
			Handler:    _LearningService_CreateVideo_Handler,
		},
		{
			MethodName: "GetDocumentsByLesson",
			Handler:    _LearningService_GetDocumentsByLesson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/learning/learning.proto",
}
