package main

import (
	"log"

	"github.com/MarcGrol/patientregistration/lib/impl/datastoring"
	"github.com/MarcGrol/patientregistration/lib/impl/emailsending"
	"github.com/MarcGrol/patientregistration/regprotobuf"
)

func main() {
	patientStore := datastoring.New()
	emailSender := emailsending.New()

	service := NewRegistrationService(patientStore, emailSender)
	err := regprotobuf.StartGrpcServer(regprotobuf.DefaultPort, service)
	if err != nil {
		log.Fatalf("Error starting registration server: %s", err)
	}
}
