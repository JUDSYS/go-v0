package judsys

import "image"

const SUBJECT_TYPE_NATURAL_PERSON = "natural person"
const SUBJECT_TYPE_LEGAL_PERSON = "legal person"
const SUBJECT_TYPE_OTHER = "other"

type Subject struct {
	Id          string        `json:"id"` // Ex: BRA/CPF:123.456.789-00
	Type        string        `json:"-"`
	DateOfBirth Date          `json:"-"`
	Names       []string      `json:"-"`
	Emails      []string      `json:"-"`
	Phones      []string      `json:"-"`
	Addresses   []string      `json:"-"`
	Pictures    []image.Image `json:"-"`
	Parents     []*Subject    `json:"-"`
	Guardians   []*Subject    `json:"-"`
	Passports   []Passport    `json:"-"`
}
