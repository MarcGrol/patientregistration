package uuiding

import (
	"github.com/MarcGrol/patientregistration/lib/api/uuider"
	"github.com/google/uuid"
)

type basicuuider struct{}

func New() uuider.UuidGenerator {
	return &basicuuider{}
}

func (u basicuuider) GenerateUuid() string {
	return uuid.NewString()
}
