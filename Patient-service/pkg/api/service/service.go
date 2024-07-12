package service

import (
	"context"
	"patient-service/pkg/models"
	"patient-service/pkg/pb"
	usecaseint "patient-service/pkg/usecase/interface"
)

type PatientServer struct {
	patientUseCase usecaseint.PatientUseCase
	pb.UnimplementedPatientServer
}

func NewPatientServer(useCase usecaseint.PatientUseCase) pb.PatientServer {
	return &PatientServer{
		patientUseCase: useCase,
	}

}
func (p *PatientServer) GoogleSignIn(ctx context.Context, req *pb.GoogleSignInRequest) (*pb.PatientSignUpResponse, error) {

	res, err := p.patientUseCase.GoogleSignIn(
		req.GoogleId, req.Name, req.Email, req.AccessToken, req.RefreshToken, req.Tokenexpiry,
	)
	if err != nil {
		return &pb.PatientSignUpResponse{}, err
	}
	return &pb.PatientSignUpResponse{
		PatientDetails: &pb.GoogleSignInResponse{
			Id:       res.Patient.Id,
			GoogleId: res.Patient.GoogleId,
			Fullname: res.Patient.FullName,
			Email:    res.Patient.Email,
		},
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}
func (p *PatientServer) IndPatientDetails(ctx context.Context, req *pb.Idreq) (*pb.PatientDetails, error) {
	doctor, err := p.patientUseCase.IndPatientDetails(req.UserId)
	if err != nil {
		return &pb.PatientDetails{}, err
	}
	return &pb.PatientDetails{
		Id:            doctor.Id,
		Fullname:      doctor.Fullname,
		Email:         doctor.Email,
		Gender:        doctor.Gender,
		Contactnumber: doctor.Contactnumber,
	}, nil
}
func (p *PatientServer) UpdatePatientDetails(ctx context.Context, req *pb.PatientDetails) (*pb.PatientDetails, error) {
	patient := models.SignupdetailResponse{
		Id:            req.Id,
		Fullname:      req.Fullname,
		Email:         req.Email,
		Gender:        req.Gender,
		Contactnumber: req.Contactnumber,
	}
	res, err := p.patientUseCase.UpdatePatientDetails(patient)
	if err != nil {
		return &pb.PatientDetails{}, err
	}
	return &pb.PatientDetails{
		Fullname:      res.Fullname,
		Email:         res.Email,
		Gender:        res.Gender,
		Contactnumber: res.Contactnumber,
	}, nil
}

func (p *PatientServer) ListPatients(ctx context.Context, req *pb.Req) (*pb.Listpares, error) {
	listed, err := p.patientUseCase.ListPatients()
	if err != nil {
		return &pb.Listpares{}, err
	}
	patientlist := make([]*pb.PatientDetails, len(listed))
	for i, patient := range listed {
		patientlist[i] = &pb.PatientDetails{
			Id:            patient.Id,
			Fullname:      patient.Fullname,
			Email:         patient.Email,
			Gender:        patient.Gender,
			Contactnumber: patient.Contactnumber,
		}
	}
	return &pb.Listpares{
		Pali: patientlist,
	}, nil

}
func (p *PatientServer) GetPatientGoogleDetailsByID(ctx context.Context, req *pb.Idreq) (*pb.GooglePatientDetails, error) {
	res, err := p.patientUseCase.GetPatientGoogleDetailsByID(req.UserId)
	if err != nil {
		return &pb.GooglePatientDetails{}, err
	}
	return &pb.GooglePatientDetails{
		Googleid:     res.GoogleID,
		Email:        res.GoogleEmail,
		Accesstoken:  res.AccessToken,
		Refreshtoken: res.RefreshToken,
		Tokenexpiry:  res.TokenExpiry,
	}, nil
}
func (p *PatientServer) UpdatePatientGoogleToken(ctx context.Context, req *pb.UpdateGoogleTokenReq) (*pb.UpdateGoogleTokenRes, error) {
	err := p.patientUseCase.UpdatePatientGoogleToken(req.GoogleID, req.AccessToken, req.RefreshToken, req.TokenExpiry)
	if err != nil {
		return &pb.UpdateGoogleTokenRes{}, err
	}
	return &pb.UpdateGoogleTokenRes{}, nil
}
