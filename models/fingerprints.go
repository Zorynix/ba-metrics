package models

type Fingerprints struct {
	Ip           string `json:"ip"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country"`
	Timezone     string `json:"timezone"`
	Referer      string `json:"referer,omitempty"`
	Browser      string `json:"browser"`
	Localization string `json:"localization,omitempty"`
	Model        string `json:"model,omitempty"`
	Platform     string `json:"Platform,omitempty"`
	OS           string `json:"os"`
}
