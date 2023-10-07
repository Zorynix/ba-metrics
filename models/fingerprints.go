package models

type Fingerprints struct {
	Ip        string `json:"ip"`
	UserAgent string `json:"userAgent"`
	Referer   string `json:"referer,omitempty"`
}
