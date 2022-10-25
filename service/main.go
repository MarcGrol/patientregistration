package main

import (
	"log"

	"github.com/MarcGrol/patientregistration/lib/impl/datastoring"
	"github.com/MarcGrol/patientregistration/lib/impl/emailsending"
	"github.com/MarcGrol/patientregistration/lib/impl/pincoding"
	"github.com/MarcGrol/patientregistration/lib/impl/uuiding"
	"github.com/MarcGrol/patientregistration/regprotobuf"
)

func main() {
	uuidGenerator := uuiding.New()
	patientStore := datastoring.New()
	pincoder := pincoding.New()
	emailSender := emailsending.New()

	service := NewRegistrationService(uuidGenerator, patientStore, pincoder, emailSender)
	err := regprotobuf.StartGrpcServer(regprotobuf.DefaultPort, service)
	if err != nil {
		log.Fatalf("Error starting registration server: %s", err)
	}
}
