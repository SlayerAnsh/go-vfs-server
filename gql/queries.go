package gql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetChain(chain string) (map[string]interface{}, error) {
	jsonStr := fmt.Sprintf(`{"query":"%s","variables":{"chain":"%s"}}`, CHAIN_QUERY, chain)
	postBody := []byte(jsonStr)

	println(jsonStr)

	req, err := http.NewRequest(http.MethodPost, GQL_URL, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var jsonData any
	err = json.NewDecoder(res.Body).Decode(&jsonData)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	gqlerr := jsonData.(map[string]any)["errors"]
	if gqlerr != nil {
		return nil, fmt.Errorf("%s", gqlerr)
	}
	return jsonData.(map[string]any)["data"].(map[string]any)["chainConfigs"].(map[string]any)["config"].(map[string]interface{}), nil
}

func QuerySmart(contract string, query string) (any, error) {
	jsonStr := fmt.Sprintf(`{"query":"%s","variables":{"contract":"%s","query":"%s"}}`, SMART_QUERY, contract, query)
	postBody := []byte(jsonStr)

	println(jsonStr)

	req, err := http.NewRequest(http.MethodPost, GQL_URL, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var jsonData map[string]any
	err = json.NewDecoder(res.Body).Decode(&jsonData)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	gqlerr := jsonData["errors"]
	if gqlerr != nil {
		return nil, fmt.Errorf("%s", gqlerr)
	}
	gqlerr = jsonData["error"]
	if gqlerr != nil {
		return nil, fmt.Errorf("%s", gqlerr)
	}
	return jsonData["data"].(map[string]any)["ADO"].(map[string]any)["adoSmart"].(map[string]any)["queryResult"], nil
}

func GetVfs(contract string) (string, error) {
	res, err := QuerySmart(contract, fmt.Sprintf(KERNEL_KEY_ADDRESS_QUERY, "vfs"))
	return res.(string), err
}

func VfsResolvePath(contract string, path string) (string, error) {
	if !strings.HasPrefix(path, "/home") && !strings.HasPrefix(path, "~") {
		path = "/home" + path
	}
	res, err := QuerySmart(contract, fmt.Sprintf(VFS_RESOLVE_PATH_QUERY, path))
	if err != nil {
		return "", err
	}
	return res.(string), nil
}

func GetAdoType(contract string) (string, error) {
	res, err := QuerySmart(contract, fmt.Sprintf(ADO_TYPE_QUERY))
	return res.(map[string]any)["ado_type"].(string), err
}
