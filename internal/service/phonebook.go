package service

import (
	"context"
	"fmt"
	"phone-number-manager/internal/db"
	"phone-number-manager/internal/models"

	"github.com/nyaruka/phonenumbers"
)

type PhoneBookService struct {
	db db.PhoneBookRepository
}

func NewPhoneBookService(db db.PhoneBookRepository) *PhoneBookService {
	return &PhoneBookService{
		db: db,
	}
}

func (s *PhoneBookService) CreatePhoneBookEntry(ctx context.Context, entry *models.PhoneBook, num *phonenumbers.PhoneNumber) error {

	entry.CountryCode, entry.AreaCode, entry.LocalPhoneNumber = getAdditionalDetails(num)
	return s.db.CreateEntry(ctx, entry)
}

func removeSpaces(s string) string {
	result := ""
	for _, char := range s {
		if char != ' ' {
			result += string(char)
		}
	}
	return result
}
func getAdditionalDetails(num *phonenumbers.PhoneNumber) (string, string, string) {

	// Get national number as string
	nationalNumber := fmt.Sprintf("%d", num.GetNationalNumber())

	// Get countryCode (e.g., "US")
	countryCode := phonenumbers.GetRegionCodeForNumber(num)

	// Get length of geographical area code
	areaCodeLength := phonenumbers.GetLengthOfGeographicalAreaCode(num)

	// Extract area code and local number
	areaCode := nationalNumber[:areaCodeLength]
	localNumber := nationalNumber[areaCodeLength:]

	return countryCode, areaCode, localNumber
}
