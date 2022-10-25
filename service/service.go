package main

import (
	"context"
	"fmt"
	"github.com/MarcGrol/patientregistration/lib/api/datastorer"
	"github.com/MarcGrol/patientregistration/lib/api/emailsender"
	"github.com/MarcGrol/patientregistration/regprotobuf"
)

type RegistrationService struct {
	patientStore datastorer.PatientStorer
	emailSender  emailsender.EmailSender
	regprotobuf.UnimplementedRegistrationServiceServer
}

func NewRegistrationService(patientStore datastorer.PatientStorer, emailSender emailsender.EmailSender) *RegistrationService {
	return &RegistrationService{
		patientStore: patientStore,
		emailSender:  emailSender,
	}
}

func (rs *RegistrationService) RegisterPatient(ctx context.Context, req *regprotobuf.RegisterPatientRequest) (*regprotobuf.RegisterPatientResponse, error) {
	// TODO
	return nil, fmt.Errorf("Not implemented")
}

// TODO add CompletePatientRegistration
