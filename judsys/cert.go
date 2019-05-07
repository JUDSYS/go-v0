package judsys

import "time"

type Cert struct {
	Subject   Subject     `json:"subject"`
	Issuer    string      `json:"issuer"`
	NotBefore time.Time   `json:"notBefore"`
	NotAfter  time.Time   `json:"notAfter"`
	Keys      []PublicKey `json:"keys"`
}
