package model

// UserCreator structure for user creator in collection
type UserCreator struct {
	ID         string `json:"id,omitempty" jsonapi:"primary,id"`
	Name       string `json:"name,omitempty" jsonapi:"attr,name"`
	Email      string `json:"email,omitempty" jsonapi:"attr,email"`
	IP         string `json:"ip,omitempty" jsonapi:"attr,ip"`
	JobTitle   string `json:"jobTitle,omitempty" jsonapi:"attr,jobTitle"`
	Department string `json:"department,omitempty" jsonapi:"attr,department"`
	Phone      string `json:"phone,omitempty" jsonapi:"attr,phone"`
}

// UserModified structure for user modify in collection
type UserModified struct {
	ID         string `json:"id,omitempty" jsonapi:"primary,id"`
	Name       string `json:"name,omitempty" jsonapi:"attr,name"`
	Email      string `json:"email,omitempty" jsonapi:"attr,email"`
	IP         string `json:"ip,omitempty" jsonapi:"attr,ip"`
	JobTitle   string `json:"jobTitle,omitempty" jsonapi:"attr,jobTitle"`
	Department string `json:"department,omitempty" jsonapi:"attr,department"`
	Phone      string `json:"phone,omitempty" jsonapi:"attr,phone"`
}