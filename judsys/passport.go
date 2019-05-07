package judsys

const PASSPORT_TYPE = "P"
const SEX_MALE = "M"
const SEX_FEMALE = "F"
const SEX_INTERSEX = "X"

type Passport struct {
	Type           string `json:"type"`
	Country        string `json:"country"`
	Number         string `json:"number"`
	Surname        string `json:"surname"`
	GivenName      string `json:"givenName"`
	Nationality    string `json:"nationality"`
	DateOfBirth    Date   `json:"dateOfBirth"`
	PlaceOfBirth   string `json:"placeOfBirth"`
	PersonalNumber string `json:"personalNumber"`
	Sex            string `json:"sex"`
	DateOfIssue    Date   `json:"dateOfIssue"`
	DateOfExpiry   Date   `json:"dateOfExpiry"`
	MRZ            string `json:"MRZ"`
}
