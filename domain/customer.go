package domain

import "banking/dto"

type Customer struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"full_name" xml:"full_name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
	DateOfBirth string `json:"date_of_birth" xml:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

func (c Customer) statusText() string {
	var status = "inactive"
	if c.Status == "1" {
		status = "active"
	}
	return status
}

func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusText(),
	}
}
