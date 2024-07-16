package path

import (
	"strings"

	"github.com/SlayerAnsh/vfs-resolver-api/cmd"
	"github.com/SlayerAnsh/vfs-resolver-api/gql"
)

func ResolvePath(p string, vfs string) (any, error) {
	parts := strings.Split(p, "/")

	vfsPath := ""
	var result any

	for i := 0; i < len(parts); {
		part := parts[i]
		if strings.HasPrefix(part, "[") {

			adoAddress, err := gql.VfsResolvePath(vfs, vfsPath)

			if err != nil {
				return nil, err
			}

			subcommands := 0

			for j := 0; j < len(part); j++ {
				if part[j] != '[' {
					break
				}
				subcommands++
			}

			cmdResult, err := cmd.ApplyCommand(adoAddress, parts[i:i+subcommands])

			if err != nil {
				return nil, err
			}

			result = cmdResult
			vfsPath = "/" + result.(string)

			i += subcommands

		} else {
			result = nil
			if part == "" {
				i++
				continue
			}
			vfsPath = vfsPath + "/" + part
			i++
		}
	}

	if result == nil && vfsPath != "" {
		address, err := gql.VfsResolvePath(vfs, vfsPath)
		if err != nil {
			return nil, err
		}
		result = address
	}

	return result, nil
}
