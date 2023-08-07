package commands

import (
	"fmt"

	"github.com/karkitirtha10/simplebank/services"
)

type GenerateRSA struct{}

func (GenerateRSA) Name() string {
	return string("generate:rsa")
}

func (cMigtn GenerateRSA) Handle() {
	// _ := flag.NewFlagSet(cMigtn.Name(), flag.ExitOnError)
	rsa := services.RSAGeneartor{}
	rsa.Generate(2048)
	fmt.Printf("prvate/public key pair generated success fully at " + services.PRIVATE_KEY_FILE_NAME + " and " + services.PUBLIC_KEY_FILE_NAME)
}

func NewGenerateRSACommand() *GenerateRSA {
	return &GenerateRSA{}
}
