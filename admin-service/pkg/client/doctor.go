package client

import (
	"context"
	"fmt"
	"healy-admin/pkg/config"
	papb "healy-admin/pkg/pb/doctor"
	"healy-admin/pkg/utils/models"
	"time"

	"google.golang.org/grpc"
)

type doctorClient struct {
	Client papb.DoctorClient
}

func NewdoctorClient(cfg *config.Config) *doctorClient {
	fmt.Println("Connecting to doctor:", cfg.DOCTOR_SVC)

	cc, err := grpc.Dial("doctor-srv:8888", grpc.WithInsecure(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	pbClient := papb.NewDoctorClient(cc)

	return &doctorClient{
		Client: pbClient,
	}

}
func (c *doctorClient) CheckDoctor(doctorid int) (bool, error) {
	ok, err := c.Client.Checkdoctor(context.Background(), &papb.Doctorreq{
		DoctorId: int32(doctorid),
	})
	if err != nil {
		return false, err
	}
	return ok.Bool, nil

}

func (c *doctorClient) DoctorDetailforBooking(doctorid int) (models.BookingDoctorDetails, error) {
	fmt.Println("Docter details client bookingDetails")
	res, err := c.Client.DoctorDetailforBooking(context.Background(), &papb.Doctorreq{DoctorId: int32(doctorid)})
	fmt.Println("Docter details client bookingDetails received", res)

	if err != nil {
		return models.BookingDoctorDetails{}, err
	}
	return models.BookingDoctorDetails{
		Doctorid:    uint(res.DoctorId),
		DoctorName:  res.DoctorName,
		DoctorEmail: res.DoctorEmail,
		Fees:        uint64(res.Fees),
	}, nil
}
