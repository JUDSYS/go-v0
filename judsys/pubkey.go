package judsys

const USAGE_SIGNING = "signing"
const USAGE_IDENTIFICATION = "identification"
const USAGE_ENCRYPTION = "encryption"
const USAGE_CERTIFICATION = "certification"
const USAGE_REVOCATION = "revocation"

type PublicKeyBase struct {
	ID        string `json:"id"`
	Algorithm string `json:"algorithm"`
	Usage     string `json:"usage"`
}

type PublicKey interface {
	ID() string
	Algoritm() string
	Usage() string
}
