package services

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"mine/multipart"
	"net/http"
)

func ParseResume(filePath string) ([]byte, error) {
	url := "https://api.apilayer.com/resume_parser/upload"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	reqBody := new(bytes.Buffer)
	writer := multipart.NewWriter(reqBody)
	part, _ := writer.CreateFormFile("resume", filePath)
	part.Write(file)
	writer.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, reqBody)
	req.Header.Add("content-Type", writer.FormDataContentType())
	req.Header.Add("apiKey", "Your API KEY")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}
