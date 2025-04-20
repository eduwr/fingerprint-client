package models

type FingerprintData struct {
	Screen         string `json:"screen" validate:"required"`
	Timezone       string `json:"timezone" validate:"required"`
	Language       string `json:"language" validate:"required,min=2"`
	MaxTouchPoints int8   `json:"maxTouchPoints" validate:"min=0,max=10"`
	Gpu            string `json:"gpu" validate:"required"`
	Canvas         string `json:"canvas" validate:"required"`
}
