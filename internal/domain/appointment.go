package domain

type Appointment struct {
	Dentist     Dentist `json:"dentist"`
	Patient     Patient `json:"patient"`
	Date        string  `json:"date"`
	Time        string  `json:"time"`
	Description string  `json:"description"`
}