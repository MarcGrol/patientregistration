package pincoding

import (
	"math/rand"
	"time"

	"github.com/MarcGrol/patientregistration/lib/api/pincoder"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type randomPincodeGenerator struct{}

func New() pincoder.PincodeGenerator {
	return &randomPincodeGenerator{}
}

func (pc randomPincodeGenerator) GeneratePincode() int {
	return rand.Intn(10000)
}
