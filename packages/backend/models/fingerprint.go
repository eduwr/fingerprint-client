package models

type FingerprintData struct {
	Screen         string `json:"screen"`
	Timezone       string `json:"timezone"`
	Language       string `json:"language"`
	MaxTouchPoints int16  `json:"maxTouchPoints"`
	Gpu            string `json:"gpu"`
	Canvas         string `json:"canvas"`
}
