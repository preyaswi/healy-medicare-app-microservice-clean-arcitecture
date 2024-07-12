package client

import (
	"context"
	"fmt"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/config"
	pb "healy-apigateway/pkg/pb/patient"
	models "healy-apigateway/pkg/utils"

	"google.golang.org/grpc"
)

type patientClient struct {
	Client pb.PatientClient
}

func NewPatientClient(cfg config.Config) interfaces.PatientClient {
	grpcConnection, err := grpc.Dial(cfg.PatientSvc, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}
	grpcClient := pb.NewPatientClient(grpcConnection)
	return &patientClient{
		Client: grpcClient,
	}
}
func (p *patientClient) GoogleSignIn(googleID, email, name, accesstoken, refreshtoken, tokenexpiry string) (models.TokenPatient, error) {
	res, err := p.Client.GoogleSignIn(context.Background(), &pb.GoogleSignInRequest{
		GoogleId:     googleID,
		Email:        email,
		Name:         name,
		AccessToken:  accesstoken,
		RefreshToken: refreshtoken,
		Tokenexpiry:  tokenexpiry,
	})
	if err != nil {
		return models.TokenPatient{}, err
	}

	patientDetails := models.GoogleSignupdetailResponse{
		Id:       res.PatientDetails.Id,
		GoogleId: res.PatientDetails.GoogleId,
		FullName: res.PatientDetails.Fullname,
		Email:    res.PatientDetails.Email,
	}
	return models.TokenPatient{
		Patient:      patientDetails,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}

func (p *patientClient) PatientDetails(user_id string) (models.SignupdetailResponse, error) {
	res, err := p.Client.IndPatientDetails(context.Background(), &pb.Idreq{UserId: user_id})
	if err != nil {
		return models.SignupdetailResponse{}, err
	}
	return models.SignupdetailResponse{
		Id:            res.Id,
		Fullname:      res.Fullname,
		Email:         res.Email,
		Gender:        res.Gender,
		Contactnumber: res.Contactnumber,
	}, nil
}
func (p *patientClient) UpdatePatientDetails(pa models.PatientDetails, patient_id string) (models.PatientDetails, error) {
	res, err := p.Client.UpdatePatientDetails(context.Background(), &pb.PatientDetails{
		Id:            patient_id,
		Fullname:      pa.Fullname,
		Email:         pa.Email,
		Gender:        pa.Gender,
		Contactnumber: pa.Contactnumber,
	})
	if err != nil {
		return models.PatientDetails{}, err
	}
	return models.PatientDetails{
		Fullname:      res.Fullname,
		Email:         res.Email,
		Gender:        res.Gender,
		Contactnumber: res.Contactnumber,
	}, nil
}

func (p *patientClient) ListPatients() ([]models.SignupdetailResponse, error) {
	res, err := p.Client.ListPatients(context.Background(), &pb.Req{})
	if err != nil {
		return []models.SignupdetailResponse{}, err
	}
	patientslist := make([]models.SignupdetailResponse, len(res.Pa))
	for i, patient := range res.Pa {
		patientdetail := models.SignupdetailResponse{
			Id:            patient.Id,
			Fullname:      patient.Fullname,
			Email:         patient.Email,
			Gender:        patient.Gender,
			Contactnumber: patient.Contactnumber,
		}
		patientslist[i] = patientdetail
	}
	return patientslist, nil
}
