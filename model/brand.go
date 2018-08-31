package model

import (
	"github.com/arangodb/go-driver"
)

// LogoBrand data structure to collect logo in brand
type LogoBrand struct {
	Type string `json:"type"`
	Size string `json:"size"`
	URL  string `json:"url"`
}

// Brand structure for database
type Brand struct {
	ID         driver.DocumentID `json:"_id,omitempty"`
	Key        string            `json:"_key,omitempty"`
	Rev        string            `json:"_rev,omitempty"`
	OldID      string            `json:"oldId"`
	Code       string            `json:"code"`
	Name       string            `json:"name"`
	NameLower  string            `json:"name_lower"`
	Logo       []LogoBrand       `json:"logo"`
	URL        string            `json:"url"`
	CreatedAt  string            `json:"createdAt"`
	CreatedBy  UserCreator      `json:"creator"`
	ModifiedAt string            `json:"modifiedAt"`
	ModifiedBy UserModified     `json:"editor"`
}