package models

// PhoneNumber example model
type PhoneBook struct {
	ID               string `json:"id,omitempty" bson:"_id,omitempty"`
	PhoneNumber      string `json:"number" validate:"required" bson:"number"`
	CountryCode      string `json:"countryCode,omitempty" bson:"countryCode"`
	AreaCode         string `json:"areaCode,omitempty" bson:"areaCode"`
	LocalPhoneNumber string `json:"localPhoneNumber,omitempty" bson:"localPhoneNumber"`
}
