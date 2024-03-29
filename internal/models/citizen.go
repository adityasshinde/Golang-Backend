package models

type Citizen struct {
	Id          string `json:"id,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	DateOfBirth string `json:"dob,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Pincode     string `json:"pincode,omitempty"`
}
type NewCitizen struct {
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	DateOfBirth string `json:"dob,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Pincode     string `json:"pincode,omitempty"`
}
