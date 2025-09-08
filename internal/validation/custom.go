package validation

import (
	"fmt"
	"log"
	"regexp"

	"github.com/nyaruka/phonenumbers"
)

func ValidateE164Phone(phoneNumber string) (*phonenumbers.PhoneNumber, error) {
	// Parse the E.164 number
	var phoneRegex = regexp.MustCompile(`^\+?\d+( \d+){0,2}$`)
	if !phoneRegex.MatchString(phoneNumber) {
		log.Printf("Invalid phone number format: %s", phoneNumber)
		return nil, fmt.Errorf("invalid phone number format")
	}
	num, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		log.Printf("Error parsing phone number: %v", err)
		return nil, err
	}
	return num, nil
}
