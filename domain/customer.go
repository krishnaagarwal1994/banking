package domain

type Customer struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"full_name" xml:"full_name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
	DateOfBirth string `json:"date_of_birth" xml:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}
