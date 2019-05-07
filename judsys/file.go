package judsys

type File struct {
	Header  string   `json:"header"`
	Objects []string `json:"objects"`
}
