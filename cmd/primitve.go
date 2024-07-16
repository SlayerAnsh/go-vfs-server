package cmd

import (
	"errors"
	"fmt"

	"github.com/SlayerAnsh/vfs-resolver-api/gql"
)

func PrimtiveCommand(contract string, command []string) (any, error) {
	if len(command) == 1 {
		value, err := gql.QuerySmart(contract, fmt.Sprintf(gql.PRIMITIVE_KEY_QUERY, command[0][1:len(command[0])-1]))
		if err != nil {
			return nil, err
		}
		value = value.(map[string]any)["value"]

		stringValue := value.(map[string]any)["string"]
		if stringValue != nil {
			return stringValue.(string), nil
		}

		addrValue := value.(map[string]any)["addr"]
		if addrValue != nil {
			return addrValue.(string), nil
		}
		return value, nil
	}
	return nil, errors.New("Not Implemented")
}
