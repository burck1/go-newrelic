package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func readJSONObjectFile(path string) (map[string]interface{}, error) {
	fileBytes, err := readFileBytes(path)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	if err := json.Unmarshal(fileBytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func readFileBytes(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func writeToFile(body io.ReadCloser, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, body); err != nil {
		return err
	}

	return nil
}

func writeToFileIndent(body io.ReadCloser, path string) error {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	var data interface{}

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return err
	}

	json, err := json.MarshalIndent(&data, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, json, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func getApis(data map[string]interface{}) ([]interface{}, error) {
	apis, ok := data["apis"]
	if !ok {
		return nil, &swaggerJSONError{"$.apis", "not found"}
	}

	apisArray, ok := apis.([]interface{})
	if !ok {
		return nil, &swaggerJSONError{"$.apis", "is an unexpected type"}
	}

	return apisArray, nil
}

type swaggerJSONError struct {
	jsonPath string
	message  string
}

func (e *swaggerJSONError) Error() string {
	return fmt.Sprintf("%s %s in the swagger json", e.jsonPath, e.message)
}

func getPath(i int, api interface{}) (string, error) {
	apiObj, ok := api.(map[string]interface{})
	if !ok {
		return "", &swaggerJSONError{fmt.Sprintf("$.apis[%d]", i), "is an unexpected type"}
	}

	path, ok := apiObj["path"]
	if !ok {
		return "", &swaggerJSONError{fmt.Sprintf("$.apis[%d].path", i), "not found"}
	}

	pathStr, ok := path.(string)
	if !ok {
		return "", &swaggerJSONError{fmt.Sprintf("$.apis[%d].path", i), "is an unexpected type"}
	}

	return pathStr, nil
}

type pathError struct {
	path string
}

func (e *pathError) Error() string {
	return fmt.Sprintf("%s is invalid", e.path)
}

func getFilenameFromURL(url string) (string, error) {
	i := strings.LastIndex(url, "/")
	if i == -1 {
		return "", &pathError{url}
	}
	return url[i+1:], nil
}
