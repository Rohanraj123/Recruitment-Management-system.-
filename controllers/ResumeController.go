package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// ResumeParserResponse represents the structure of the response from the resume parser API
type ResumeParserResponse struct {
	Education []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"education"`
	Email      string `json:"email"`
	Experience []struct {
		Dates []string `json:"dates"`
		Name  string   `json:"name"`
		URL   string   `json:"url,omitempty"`
	} `json:"experience"`
	Name   string   `json:"name"`
	Phone  string   `json:"phone"`
	Skills []string `json:"skills"`
}

func UploadResume(c *gin.Context) {
	// Check if the user is an Applicant
	userType, exists := c.Get("user_type")
	if !exists || userType != "Applicant" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only applicants can upload resumes"})
		return
	}

	// Get the file from the request
	file, err := c.FormFile("resume")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Validate file type (PDF or DOCX)
	ext := filepath.Ext(file.Filename)
	if ext != ".pdf" && ext != ".docx" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only PDF or DOCX allowed"})
		return
	}

	// Define file storage path
	storagePath := filepath.Join("uploads", file.Filename)

	// Save the uploaded file
	if err := c.SaveUploadedFile(file, storagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
		return
	}

	// Call the API to parse the resume
	parsedResume, apiErr := parseResume(storagePath)
	if apiErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse resume: " + apiErr.Error()})
		return
	}

	// Return the parsed resume as JSON
	c.JSON(http.StatusOK, gin.H{
		"message":       "Resume uploaded and parsed successfully",
		"file_path":     storagePath,
		"parsed_resume": parsedResume,
	})
}

// parseResume sends the resume file to the third-party API for parsing
func parseResume(filePath string) (ResumeParserResponse, error) {
	var resumeResponse ResumeParserResponse

	// Read the file content
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return resumeResponse, fmt.Errorf("failed to read file: %v", err)
	}

	// Prepare the API request
	apiURL := "https://api.apilayer.com/resume_parser/upload"
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(fileContent))
	if err != nil {
		return resumeResponse, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", "YOUR_API_KEY_HERE") // Replace with your API key

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resumeResponse, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return resumeResponse, fmt.Errorf("API returned status: %s", resp.Status)
	}

	// Parse the response
	if err := json.NewDecoder(resp.Body).Decode(&resumeResponse); err != nil {
		return resumeResponse, fmt.Errorf("failed to decode response: %v", err)
	}

	return resumeResponse, nil
}
