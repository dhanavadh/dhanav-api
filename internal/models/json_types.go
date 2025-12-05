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
