package cmd

import (
	"errors"

	"github.com/SlayerAnsh/vfs-resolver-api/gql"
)

func ApplyCommand(contract string, command []string) (any, error) {
	adoType, err := gql.GetAdoType(contract)
	if err != nil {
		adoType = "account"
	}
	if adoType == "primitive" {
		return PrimtiveCommand(contract, command)
	}
	if adoType == "cw721" {
		return Cw721Command(contract, command)
	}
	return nil, errors.New("Not Implemented")
}
