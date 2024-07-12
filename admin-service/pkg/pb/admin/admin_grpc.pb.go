// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/pb/admin/admin.proto

package admin

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

const (
	Admin_AdminSignup_FullMethodName               = "/admin.Admin/AdminSignup"
	Admin_AdminLogin_FullMethodName                = "/admin.Admin/AdminLogin"
	Admin_AddTobookings_FullMethodName             = "/admin.Admin/AddTobookings"
	Admin_CancelBookings_FullMethodName            = "/admin.Admin/CancelBookings"
	Admin_MakePaymentRazorpay_FullMethodName       = "/admin.Admin/MakePaymentRazorpay"
	Admin_VerifyPayment_FullMethodName             = "/admin.Admin/VerifyPayment"
	Admin_GetPaidPatients_FullMethodName           = "/admin.Admin/GetPaidPatients"
	Admin_CreatePrescription_FullMethodName        = "/admin.Admin/CreatePrescription"
	Admin_SetDoctorAvailability_FullMethodName     = "/admin.Admin/SetDoctorAvailability"
	Admin_GetDoctorAvailability_FullMethodName     = "/admin.Admin/GetDoctorAvailability"
	Admin_BookSlot_FullMethodName                  = "/admin.Admin/BookSlot"
	Admin_BookDoctor_FullMethodName                = "/admin.Admin/BookDoctor"
	Admin_VerifyandCalenderCreation_FullMethodName = "/admin.Admin/VerifyandCalenderCreation"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	AdminSignup(ctx context.Context, in *AdminSignupRequest, opts ...grpc.CallOption) (*AdminSignupResponse, error)
	AdminLogin(ctx context.Context, in *AdminLoginInRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
	AddTobookings(ctx context.Context, in *Bookingreq, opts ...grpc.CallOption) (*Bookingres, error)
	CancelBookings(ctx context.Context, in *Canbookingreq, opts ...grpc.CallOption) (*Bookingres, error)
	MakePaymentRazorpay(ctx context.Context, in *PaymentReq, opts ...grpc.CallOption) (*PaymentRes, error)
	VerifyPayment(ctx context.Context, in *PaymentReq, opts ...grpc.CallOption) (*Verifyres, error)
	GetPaidPatients(ctx context.Context, in *GetPaidPatientsRequest, opts ...grpc.CallOption) (*GetPaidPatientsResponse, error)
	CreatePrescription(ctx context.Context, in *CreatePrescriptionRequest, opts ...grpc.CallOption) (*CreatePrescriptionResponse, error)
	SetDoctorAvailability(ctx context.Context, in *SetDoctorAvailabilityRequest, opts ...grpc.CallOption) (*SetDoctorAvailabilityResponse, error)
	GetDoctorAvailability(ctx context.Context, in *GetDoctorAvailabilityRequest, opts ...grpc.CallOption) (*GetDoctorAvailabilityResponse, error)
	BookSlot(ctx context.Context, in *BookSlotreq, opts ...grpc.CallOption) (*BookSlotres, error)
	BookDoctor(ctx context.Context, in *BookDoctorreq, opts ...grpc.CallOption) (*PaymentRes, error)
	VerifyandCalenderCreation(ctx context.Context, in *VerifyPaymentandcalenderreq, opts ...grpc.CallOption) (*VerifyPaymentandcalenderres, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) AdminSignup(ctx context.Context, in *AdminSignupRequest, opts ...grpc.CallOption) (*AdminSignupResponse, error) {
	out := new(AdminSignupResponse)
	err := c.cc.Invoke(ctx, Admin_AdminSignup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminLogin(ctx context.Context, in *AdminLoginInRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	out := new(AdminLoginResponse)
	err := c.cc.Invoke(ctx, Admin_AdminLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AddTobookings(ctx context.Context, in *Bookingreq, opts ...grpc.CallOption) (*Bookingres, error) {
	out := new(Bookingres)
	err := c.cc.Invoke(ctx, Admin_AddTobookings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CancelBookings(ctx context.Context, in *Canbookingreq, opts ...grpc.CallOption) (*Bookingres, error) {
	out := new(Bookingres)
	err := c.cc.Invoke(ctx, Admin_CancelBookings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) MakePaymentRazorpay(ctx context.Context, in *PaymentReq, opts ...grpc.CallOption) (*PaymentRes, error) {
	out := new(PaymentRes)
	err := c.cc.Invoke(ctx, Admin_MakePaymentRazorpay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) VerifyPayment(ctx context.Context, in *PaymentReq, opts ...grpc.CallOption) (*Verifyres, error) {
	out := new(Verifyres)
	err := c.cc.Invoke(ctx, Admin_VerifyPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetPaidPatients(ctx context.Context, in *GetPaidPatientsRequest, opts ...grpc.CallOption) (*GetPaidPatientsResponse, error) {
	out := new(GetPaidPatientsResponse)
	err := c.cc.Invoke(ctx, Admin_GetPaidPatients_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CreatePrescription(ctx context.Context, in *CreatePrescriptionRequest, opts ...grpc.CallOption) (*CreatePrescriptionResponse, error) {
	out := new(CreatePrescriptionResponse)
	err := c.cc.Invoke(ctx, Admin_CreatePrescription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SetDoctorAvailability(ctx context.Context, in *SetDoctorAvailabilityRequest, opts ...grpc.CallOption) (*SetDoctorAvailabilityResponse, error) {
	out := new(SetDoctorAvailabilityResponse)
	err := c.cc.Invoke(ctx, Admin_SetDoctorAvailability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetDoctorAvailability(ctx context.Context, in *GetDoctorAvailabilityRequest, opts ...grpc.CallOption) (*GetDoctorAvailabilityResponse, error) {
	out := new(GetDoctorAvailabilityResponse)
	err := c.cc.Invoke(ctx, Admin_GetDoctorAvailability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) BookSlot(ctx context.Context, in *BookSlotreq, opts ...grpc.CallOption) (*BookSlotres, error) {
	out := new(BookSlotres)
	err := c.cc.Invoke(ctx, Admin_BookSlot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) BookDoctor(ctx context.Context, in *BookDoctorreq, opts ...grpc.CallOption) (*PaymentRes, error) {
	out := new(PaymentRes)
	err := c.cc.Invoke(ctx, Admin_BookDoctor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) VerifyandCalenderCreation(ctx context.Context, in *VerifyPaymentandcalenderreq, opts ...grpc.CallOption) (*VerifyPaymentandcalenderres, error) {
	out := new(VerifyPaymentandcalenderres)
	err := c.cc.Invoke(ctx, Admin_VerifyandCalenderCreation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	AdminSignup(context.Context, *AdminSignupRequest) (*AdminSignupResponse, error)
	AdminLogin(context.Context, *AdminLoginInRequest) (*AdminLoginResponse, error)
	AddTobookings(context.Context, *Bookingreq) (*Bookingres, error)
	CancelBookings(context.Context, *Canbookingreq) (*Bookingres, error)
	MakePaymentRazorpay(context.Context, *PaymentReq) (*PaymentRes, error)
	VerifyPayment(context.Context, *PaymentReq) (*Verifyres, error)
	GetPaidPatients(context.Context, *GetPaidPatientsRequest) (*GetPaidPatientsResponse, error)
	CreatePrescription(context.Context, *CreatePrescriptionRequest) (*CreatePrescriptionResponse, error)
	SetDoctorAvailability(context.Context, *SetDoctorAvailabilityRequest) (*SetDoctorAvailabilityResponse, error)
	GetDoctorAvailability(context.Context, *GetDoctorAvailabilityRequest) (*GetDoctorAvailabilityResponse, error)
	BookSlot(context.Context, *BookSlotreq) (*BookSlotres, error)
	BookDoctor(context.Context, *BookDoctorreq) (*PaymentRes, error)
	VerifyandCalenderCreation(context.Context, *VerifyPaymentandcalenderreq) (*VerifyPaymentandcalenderres, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) AdminSignup(context.Context, *AdminSignupRequest) (*AdminSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminSignup not implemented")
}
func (UnimplementedAdminServer) AdminLogin(context.Context, *AdminLoginInRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAdminServer) AddTobookings(context.Context, *Bookingreq) (*Bookingres, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTobookings not implemented")
}
func (UnimplementedAdminServer) CancelBookings(context.Context, *Canbookingreq) (*Bookingres, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBookings not implemented")
}
func (UnimplementedAdminServer) MakePaymentRazorpay(context.Context, *PaymentReq) (*PaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePaymentRazorpay not implemented")
}
func (UnimplementedAdminServer) VerifyPayment(context.Context, *PaymentReq) (*Verifyres, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPayment not implemented")
}
func (UnimplementedAdminServer) GetPaidPatients(context.Context, *GetPaidPatientsRequest) (*GetPaidPatientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaidPatients not implemented")
}
func (UnimplementedAdminServer) CreatePrescription(context.Context, *CreatePrescriptionRequest) (*CreatePrescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrescription not implemented")
}
func (UnimplementedAdminServer) SetDoctorAvailability(context.Context, *SetDoctorAvailabilityRequest) (*SetDoctorAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDoctorAvailability not implemented")
}
func (UnimplementedAdminServer) GetDoctorAvailability(context.Context, *GetDoctorAvailabilityRequest) (*GetDoctorAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctorAvailability not implemented")
}
func (UnimplementedAdminServer) BookSlot(context.Context, *BookSlotreq) (*BookSlotres, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookSlot not implemented")
}
func (UnimplementedAdminServer) BookDoctor(context.Context, *BookDoctorreq) (*PaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookDoctor not implemented")
}
func (UnimplementedAdminServer) VerifyandCalenderCreation(context.Context, *VerifyPaymentandcalenderreq) (*VerifyPaymentandcalenderres, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyandCalenderCreation not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_AdminSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminSignup(ctx, req.(*AdminSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminLogin(ctx, req.(*AdminLoginInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AddTobookings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bookingreq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AddTobookings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AddTobookings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AddTobookings(ctx, req.(*Bookingreq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CancelBookings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Canbookingreq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CancelBookings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_CancelBookings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CancelBookings(ctx, req.(*Canbookingreq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_MakePaymentRazorpay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).MakePaymentRazorpay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_MakePaymentRazorpay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).MakePaymentRazorpay(ctx, req.(*PaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_VerifyPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).VerifyPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_VerifyPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).VerifyPayment(ctx, req.(*PaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetPaidPatients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaidPatientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetPaidPatients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_GetPaidPatients_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetPaidPatients(ctx, req.(*GetPaidPatientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CreatePrescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreatePrescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_CreatePrescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreatePrescription(ctx, req.(*CreatePrescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SetDoctorAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDoctorAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SetDoctorAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_SetDoctorAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SetDoctorAvailability(ctx, req.(*SetDoctorAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetDoctorAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDoctorAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetDoctorAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_GetDoctorAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetDoctorAvailability(ctx, req.(*GetDoctorAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_BookSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookSlotreq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).BookSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_BookSlot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).BookSlot(ctx, req.(*BookSlotreq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_BookDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookDoctorreq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).BookDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_BookDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).BookDoctor(ctx, req.(*BookDoctorreq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_VerifyandCalenderCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPaymentandcalenderreq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).VerifyandCalenderCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_VerifyandCalenderCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).VerifyandCalenderCreation(ctx, req.(*VerifyPaymentandcalenderreq))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminSignup",
			Handler:    _Admin_AdminSignup_Handler,
		},
		{
			MethodName: "AdminLogin",
			Handler:    _Admin_AdminLogin_Handler,
		},
		{
			MethodName: "AddTobookings",
			Handler:    _Admin_AddTobookings_Handler,
		},
		{
			MethodName: "CancelBookings",
			Handler:    _Admin_CancelBookings_Handler,
		},
		{
			MethodName: "MakePaymentRazorpay",
			Handler:    _Admin_MakePaymentRazorpay_Handler,
		},
		{
			MethodName: "VerifyPayment",
			Handler:    _Admin_VerifyPayment_Handler,
		},
		{
			MethodName: "GetPaidPatients",
			Handler:    _Admin_GetPaidPatients_Handler,
		},
		{
			MethodName: "CreatePrescription",
			Handler:    _Admin_CreatePrescription_Handler,
		},
		{
			MethodName: "SetDoctorAvailability",
			Handler:    _Admin_SetDoctorAvailability_Handler,
		},
		{
			MethodName: "GetDoctorAvailability",
			Handler:    _Admin_GetDoctorAvailability_Handler,
		},
		{
			MethodName: "BookSlot",
			Handler:    _Admin_BookSlot_Handler,
		},
		{
			MethodName: "BookDoctor",
			Handler:    _Admin_BookDoctor_Handler,
		},
		{
			MethodName: "VerifyandCalenderCreation",
			Handler:    _Admin_VerifyandCalenderCreation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/admin/admin.proto",
}