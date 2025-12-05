package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type StringArray []string

func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to scan StringArray")
	}
	return json.Unmarshal(bytes, s)
}

type ContentBlock struct {
	Type        string `json:"type"`
	Text        string `json:"text,omitempty"`
	Src         string `json:"src,omitempty"`
	Alt         string `json:"alt,omitempty"`
	Description string `json:"description,omitempty"`
	Title       string `json:"title,omitempty"`
	Target      string `json:"target,omitempty"`
}

type ContentBlocks []ContentBlock

type Education struct {
	School    string `json:"school"`
	Degree    string `json:"degree"`
	YearStart string `json:"year_start"`
	YearEnd   string `json:"year_end"`
}

type Experience struct {
	Company   string `json:"company"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	YearStart string `json:"year_start"`
	YearEnd   string `json:"year_end"`
}

type ExperienceList []Experience

type EducationList []Education
type SocialLinks map[string]string

// Content Block
func (c ContentBlock) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *ContentBlock) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to scan ContentBlock")
	}
	return json.Unmarshal(bytes, c)
}

// Education List
func (e EducationList) Value() (driver.Value, error) {
	return json.Marshal(e)
}

func (e *EducationList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to scan EducationList")
	}
	return json.Unmarshal(bytes, e)
}

// Experience List
func (e ExperienceList) Value() (driver.Value, error) {
	return json.Marshal(e)
}

func (e *ExperienceList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to scan ExperienceList")
	}
	return json.Unmarshal(bytes, e)
}

// Social Links
func (s SocialLinks) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *SocialLinks) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to scan SocialLinks")
	}
	return json.Unmarshal(bytes, s)
}
