package client

import (
	"context"
	"fmt"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/config"
	models "healy-apigateway/pkg/utils"

	pb "healy-apigateway/pkg/pb/admin"

	"google.golang.org/grpc"
)

type adminClient struct {
	Client pb.AdminClient
}

func NewAdminClient(cfg config.Config) interfaces.AdminClient {

	grpcConnection, err := grpc.Dial(cfg.AdminSvc, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewAdminClient(grpcConnection)

	return &adminClient{
		Client: grpcClient,
	}

}
func (ad *adminClient) AdminSignUp(admindeatils models.AdminSignUp) (models.TokenAdmin, error) {
	admin, err := ad.Client.AdminSignup(context.Background(), &pb.AdminSignupRequest{
		Firstname: admindeatils.Firstname,
		Lastname:  admindeatils.Lastname,
		Email:     admindeatils.Email,
		Password:  admindeatils.Password,
	})
	if err != nil {
		return models.TokenAdmin{}, err
	}
	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID:        uint(admin.AdminDetails.Id),
			Firstname: admin.AdminDetails.Firstname,
			Lastname:  admin.AdminDetails.Lastname,
			Email:     admin.AdminDetails.Email,
		},
		Token: admin.Token,
	}, nil
}

func (ad *adminClient) AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error) {
	admin, err := ad.Client.AdminLogin(context.Background(), &pb.AdminLoginInRequest{
		Email:    adminDetails.Email,
		Password: adminDetails.Password,
	})

	if err != nil {
		return models.TokenAdmin{}, err
	}
	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID:        uint(admin.AdminDetails.Id),
			Firstname: admin.AdminDetails.Firstname,
			Lastname:  admin.AdminDetails.Lastname,
			Email:     admin.AdminDetails.Email,
		},
		Token: admin.Token,
	}, nil
}
func (ad *adminClient) GetPaidPatients(doctor_id int) ([]models.Patient, error) {
	res, err := ad.Client.GetPaidPatients(context.Background(), &pb.GetPaidPatientsRequest{DoctorId: int32(doctor_id)})
	if err != nil {
		return []models.Patient{}, err
	}
	patientsDetails := make([]models.Patient, len(res.BookedPatients))
	for i, bookedPatient := range res.BookedPatients {
		patientsDetails[i] = models.Patient{
			BookingId:     uint(bookedPatient.BookingId),
			SlotId: uint(bookedPatient.SlotId),
			PaymentStatus: bookedPatient.PaymentStatus,
			PatientId:     uint(bookedPatient.PatientDetail.Id),
			Fullname:      bookedPatient.PatientDetail.Fullname,
			Email:         bookedPatient.PatientDetail.Email,
			Gender:        bookedPatient.PatientDetail.Gender,
			Contactnumber: bookedPatient.PatientDetail.Contactnumber,
		}
	}
	return patientsDetails, nil
}
func (ad *adminClient) CreatePrescription(prescription models.PrescriptionRequest) (models.CreatedPrescription, error) {
	res, err := ad.Client.CreatePrescription(context.Background(), &pb.CreatePrescriptionRequest{
		BookingId: uint32(prescription.BookingID),
		DoctorId:  uint32(prescription.DoctorID),
		Medicine:  prescription.Medicine,
		Dosage:    prescription.Dosage,
		Notes:     prescription.Notes,
	})
	if err != nil {
		return models.CreatedPrescription{}, err
	}
	return models.CreatedPrescription{
		Id:        int(res.Id),
		DoctorID:  int(res.DoctorId),
		PatientID: res.PatientId,
		BookingID: int(res.BookingId),
		Medicine:  res.Medicine,
		Dosage:    res.Dosage,
		Notes:     res.Notes,
	}, nil

}
func (ad *adminClient) SetDoctorAvailability(setreq models.SetAvailability, doctorId int) (string, error) {
	res, err := ad.Client.SetDoctorAvailability(context.Background(), &pb.SetDoctorAvailabilityRequest{
		DoctorId:  uint32(doctorId),
		Date:      setreq.Date,
		StartTime: setreq.StartTime,
		EndTime:   setreq.EndTime,
	})
	if err != nil {
		return "", err
	}
	return res.Status, nil
}
func (ad *adminClient) GetDoctorAvailability(doctorid int, date string) ([]models.GetAvailability, error) {
	res, err := ad.Client.GetDoctorAvailability(context.Background(), &pb.GetDoctorAvailabilityRequest{
		DoctorId: uint32(doctorid),
		Date:     date,
	})
	if err != nil {
		return []models.GetAvailability{}, err
	}
	var availabilities []models.GetAvailability
	for _, slot := range res.Slots {
		availabilities = append(availabilities, models.GetAvailability{
			Slot_id:   slot.SlotId,
			Time:      slot.Time,
			Is_booked: slot.IsBooked,
		})
	}

	return availabilities, nil
}

func (ad *adminClient) BookDoctor(patientid string, slotid int) (models.CombinedBookingDetails, string, error) {
	res, err := ad.Client.BookDoctor(context.Background(), &pb.BookDoctorreq{
		PatientId: patientid,
		SlotId:    uint32(slotid),
	})
	if err != nil {
		return models.CombinedBookingDetails{}, "", err
	}
	paymentdetail := models.CombinedBookingDetails{
		BookingId:     uint(res.PaymentDetails.BookingId),
		PatientName:   res.PaymentDetails.PatientId,
		DoctorId:      uint(res.PaymentDetails.DoctorId),
		DoctorName:    res.PaymentDetails.DoctorName,
		DoctorEmail:   res.PaymentDetails.DoctorEmail,
		Fees:          res.PaymentDetails.Fees,
		PaymentStatus: res.PaymentDetails.PaymentStatus,
	}
	razor_id := res.Razorid
	return paymentdetail, razor_id, nil
}
func (ad *adminClient) VerifyandCalenderCreation(bookingId int, paymentid, razorid string) error {
	_, err := ad.Client.VerifyandCalenderCreation(context.Background(), &pb.VerifyPaymentandcalenderreq{
		BookingId: uint32(bookingId),
		PaymentId: paymentid,
		RazorId:   razorid,
	})
	if err != nil {
		return err
	}
	return nil

}
