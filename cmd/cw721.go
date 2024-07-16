package cmd

import (
	"errors"
	"fmt"

	"github.com/SlayerAnsh/vfs-resolver-api/gql"
)

func Cw721Command(contract string, command []string) (any, error) {
	if len(command) == 1 {
		value, err := gql.QuerySmart(contract, fmt.Sprintf(gql.CW721_TOKEN_QUERY, command[0][1:len(command[0])-1]))
		if err != nil {
			return nil, err
		}
		return value, nil
	}
	return nil, errors.New("Not Implemented")
}
