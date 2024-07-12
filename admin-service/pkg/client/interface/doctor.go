package interfaces

import "healy-admin/pkg/utils/models"

type NewDoctorClient interface {
	CheckDoctor(doctorid int) (bool, error)
	DoctorDetailforBooking(doctorid int) (models.BookingDoctorDetails, error)
}
