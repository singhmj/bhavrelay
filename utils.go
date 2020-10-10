package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// parseJSON : Reads the json from the file, and unmarshal it into the dataToParseIn(it should be a pointer to your struct)
func parseJSON(filepath string, dataToParseIn interface{}) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, dataToParseIn)
	return err
}

func readCSV(filepath string, dataToParseIn interface{}) error {
	// TODO:
	return nil
}

// sendHTTPRequest :
func sendHTTPRequest(request *http.Request) (*http.Response, error) {
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// TODO: The API can return other types of status codes as well, so handle them accordingly
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api returned status code: %v", resp.StatusCode)
	}
	return resp, nil
}

// readHTTPResponseBody ...
func readHTTPResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// ReadZip : Reads an entire zip from []bytes, and returns filewise data filled in a map
// where key is the filename, and value is the filedata
func ReadZip(data []byte) (map[string][]byte, error) {
	// filename vs data
	readFiles := make(map[string][]byte)
	size := int64(len(data))
	reader, err := zip.NewReader(bytes.NewReader(data), size)
	if err != nil {
		return nil, fmt.Errorf("failed to create zip reader, err: %v", err)
	}

	for _, file := range reader.File {
		fileData, err := readZipFile(file)
		if err != nil {
			return nil, err
		}

		readFiles[file.FileInfo().Name()] = fileData
	}

	return readFiles, nil
}

func readZipFile(file *zip.File) ([]byte, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
