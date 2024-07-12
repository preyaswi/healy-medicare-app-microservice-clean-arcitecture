package service

import (
	"context"
	"fmt"
	pb "healy-admin/pkg/pb/admin"
	interfaces "healy-admin/pkg/usecase/interface"
	"healy-admin/pkg/utils/models"
	"time"
)

type AdminServer struct {
	adminUseCase interfaces.AdminUseCase
	pb.UnimplementedAdminServer
}

func NewAdminServer(useCase interfaces.AdminUseCase) pb.AdminServer {

	return &AdminServer{
		adminUseCase: useCase,
	}

}
func (ad *AdminServer) AdminSignup(ctx context.Context, req *pb.AdminSignupRequest) (*pb.AdminSignupResponse, error) {
	adminSignup := models.AdminSignUp{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
	}

	res, err := ad.adminUseCase.AdminSignUp(adminSignup)
	if err != nil {
		return &pb.AdminSignupResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(res.Admin.ID),
		Firstname: res.Admin.Firstname,
		Lastname:  res.Admin.Lastname,
		Email:     res.Admin.Email,
	}
	return &pb.AdminSignupResponse{
		Status:       201,
		AdminDetails: adminDetails,
		Token:        res.Token,
	}, nil
}

func (ad *AdminServer) AdminLogin(ctx context.Context, Req *pb.AdminLoginInRequest) (*pb.AdminLoginResponse, error) {
	adminLogin := models.AdminLogin{
		Email:    Req.Email,
		Password: Req.Password,
	}
	admin, err := ad.adminUseCase.LoginHandler(adminLogin)
	if err != nil {
		return &pb.AdminLoginResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(admin.Admin.ID),
		Firstname: admin.Admin.Firstname,
		Lastname:  admin.Admin.Lastname,
		Email:     admin.Admin.Email,
	}
	return &pb.AdminLoginResponse{
		Status:       200,
		AdminDetails: adminDetails,
		Token:        admin.Token,
	}, nil
}
func (ad *AdminServer) AddTobookings(ctx context.Context, req *pb.Bookingreq) (*pb.Bookingres, error) {
	err := ad.adminUseCase.AddToBooking(req.PatientId, int(req.DoctorId))
	if err != nil {
		return &pb.Bookingres{}, err
	}
	return &pb.Bookingres{}, nil
}
func (ad *AdminServer) CancelBookings(ctx context.Context, req *pb.Canbookingreq) (*pb.Bookingres, error) {
	err := ad.adminUseCase.CancelBooking(req.PatientId, int(req.BookingId))
	if err != nil {
		return &pb.Bookingres{}, err
	}
	return &pb.Bookingres{}, nil
}
func (ad *AdminServer) MakePaymentRazorpay(ctx context.Context, req *pb.PaymentReq) (*pb.PaymentRes, error) {
	bookingid := models.Paymentreq{
		Bookingid: uint(req.BookingId),
	}
	paymentDetails, razorId, err := ad.adminUseCase.MakePaymentRazorpay(int(bookingid.Bookingid))
	if err != nil {
		return &pb.PaymentRes{}, err
	}
	paymentDetail := &pb.PaymentDetails{
		BookingId:     uint32(paymentDetails.BookingId),
		PatientName:     paymentDetails.PatientId,
		DoctorId:      uint32(paymentDetails.DoctorId),
		DoctorName:    paymentDetails.DoctorName,
		DoctorEmail:   paymentDetails.DoctorEmail,
		Fees:          paymentDetails.Fees,
		PaymentStatus: paymentDetails.PaymentStatus,
	}
	return &pb.PaymentRes{
		PaymentDetails: paymentDetail,
		Razorid:        razorId,
	}, nil

}
func (ad *AdminServer) VerifyPayment(ctx context.Context, req *pb.PaymentReq) (*pb.Verifyres, error) {
	err := ad.adminUseCase.VerifyPayment(int(req.BookingId))
	if err != nil {
		return &pb.Verifyres{}, err
	}
	return &pb.Verifyres{}, nil
}
func (ad *AdminServer) GetPaidPatients(ctx context.Context, req *pb.GetPaidPatientsRequest) (*pb.GetPaidPatientsResponse, error) {
	bookedPatients, err := ad.adminUseCase.GetPaidPatients(int(req.DoctorId))
	if err != nil {
		return &pb.GetPaidPatientsResponse{}, err
	}
	pbBookedPatients := make([]*pb.BookedPatient, len(bookedPatients))
	for i, bp := range bookedPatients {
		pbBookedPatients[i] = &pb.BookedPatient{
			BookingId:     uint32(bp.BookingId),
			SlotId: uint32(bp.SlotId),
			PaymentStatus: bp.PaymentStatus,
			PatientDetail: &pb.Patient{
				Id:            uint32(bp.Patientdetail.Id),
				Fullname:      bp.Patientdetail.Fullname,
				Email:         bp.Patientdetail.Email,
				Gender:        bp.Patientdetail.Gender,
				Contactnumber: bp.Patientdetail.Contactnumber,
			},
		}
	}

	return &pb.GetPaidPatientsResponse{
		BookedPatients: pbBookedPatients,
	}, nil
}
func (ad *AdminServer) CreatePrescription(ctx context.Context, req *pb.CreatePrescriptionRequest) (*pb.CreatePrescriptionResponse, error) {
	prescriptionreq := models.PrescriptionRequest{
		DoctorID:  int(req.DoctorId),
		BookingID: int(req.BookingId),
		Medicine:  req.Medicine,
		Dosage:    req.Dosage,
		Notes:     req.Notes,
	}
	prescription, err := ad.adminUseCase.CreatePrescription(prescriptionreq)
	if err != nil {
		return &pb.CreatePrescriptionResponse{}, err
	}
	return &pb.CreatePrescriptionResponse{
		Id:        uint32(prescription.ID),
		BookingId: uint32(prescription.BookingID),
		DoctorId:  uint32(prescription.DoctorID),
		PatientId: prescription.PatientID,
		Medicine:  prescription.Medicine,
		Dosage:    prescription.Dosage,
		Notes:     prescription.Notes,
	}, nil
}
func (ad *AdminServer) SetDoctorAvailability(ctx context.Context, req *pb.SetDoctorAvailabilityRequest) (*pb.SetDoctorAvailabilityResponse, error) {
	doctorId := req.DoctorId
	Date, _ := time.Parse("2006-01-02", req.Date)
	startTime, _ := time.Parse("15:04", req.StartTime)
	endtime, _ := time.Parse("15:04", req.EndTime)

	status, err := ad.adminUseCase.SetDoctorAvailability(models.SetAvailability{
		DoctorId:  int(doctorId),
		Date:      Date,
		StartTime: startTime,
		EndTime:   endtime,
	})
	if err != nil {
		return &pb.SetDoctorAvailabilityResponse{}, err
	}
	return &pb.SetDoctorAvailabilityResponse{
		Status: status,
	}, nil

}
func (ad *AdminServer) GetDoctorAvailability(ctx context.Context, req *pb.GetDoctorAvailabilityRequest) (*pb.GetDoctorAvailabilityResponse, error) {
	doctorId := req.DoctorId
	date, _ := time.Parse("2006-01-02", req.Date)
	slots, err := ad.adminUseCase.GetDoctorAvailability(int(doctorId), date)
	if err != nil {
		return &pb.GetDoctorAvailabilityResponse{}, nil
	}
	var pbSlots []*pb.Slot
	for _, slot := range slots {
		pbSlots = append(pbSlots, &pb.Slot{
			SlotId:   uint32(slot.Slot_id),
			Time:     slot.Time,
			IsBooked: slot.IsBooked,
		})
	}
	return &pb.GetDoctorAvailabilityResponse{
		Slots: pbSlots,
	}, nil

}
func (ad *AdminServer) BookSlot(ctx context.Context, req *pb.BookSlotreq) (*pb.BookSlotres, error) {
	err := ad.adminUseCase.BookSlot(req.PatientId, int(req.BookingId), int(req.SlotId))
	if err != nil {
		return &pb.BookSlotres{}, err
	}
	return &pb.BookSlotres{}, nil
}
func (ad *AdminServer) BookDoctor(ctx context.Context, req *pb.BookDoctorreq) (*pb.PaymentRes, error) {
	bookingdetails, razorid, err := ad.adminUseCase.BookDoctor(req.PatientId, int(req.SlotId))
	if err != nil {
		return &pb.PaymentRes{}, err
	}
	fmt.Println(bookingdetails.PatientId,"patient name in adminservice")
	paymentDetail := &pb.PaymentDetails{
		BookingId:     uint32(bookingdetails.BookingId),
		PatientName:     bookingdetails.PatientId,
		DoctorId:      uint32(bookingdetails.DoctorId),
		DoctorName:    bookingdetails.DoctorName,
		DoctorEmail:   bookingdetails.DoctorEmail,
		Fees:          bookingdetails.Fees,
		PaymentStatus: bookingdetails.PaymentStatus,
	}
	return &pb.PaymentRes{
		PaymentDetails: paymentDetail,
		Razorid:        razorid,
	}, nil
}
func (ad *AdminServer) VerifyandCalenderCreation(ctx context.Context, req *pb.VerifyPaymentandcalenderreq) (*pb.VerifyPaymentandcalenderres, error) {
	err := ad.adminUseCase.VerifyandCalenderCreation(int(req.BookingId), req.PaymentId, req.RazorId)
	if err != nil {
		return &pb.VerifyPaymentandcalenderres{}, err
	}
	return &pb.VerifyPaymentandcalenderres{}, nil
}
