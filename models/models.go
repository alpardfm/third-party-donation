package models

import (
	"encoding/json"
	"time"
)

// WebhookResponse response untuk webhook
type WebhookResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type SaweriaDonationPayload struct {
	AmountRaw     int         `json:"amount_raw"`
	CreatedAt     string      `json:"created_at"`
	Cut           int         `json:"cut"`
	DonatorEmail  string      `json:"donator_email"`
	DonatorIsUser bool        `json:"donator_is_user"`
	DonatorName   string      `json:"donator_name"`
	Etc           DonationEtc `json:"etc"`
	ID            string      `json:"id"`
	Message       string      `json:"message"`
	Type          string      `json:"type"`
	Version       string      `json:"version"`
}

// DonationEtc merepresentasikan additional data dalam etc field
type DonationEtc struct {
	AmountToDisplay int `json:"amount_to_display"`
}

// SaweriaWebhookPayload merepresentasikan full webhook payload
type SaweriaWebhookPayload struct {
	Event string                 `json:"event"`
	Data  SaweriaDonationPayload `json:"data"`
}

type SaweriaDonation struct {
	AmountRaw     int         `json:"amount_raw"`
	CreatedAt     time.Time   `json:"created_at"`
	Cut           int         `json:"cut"`
	DonatorEmail  string      `json:"donator_email"`
	DonatorIsUser bool        `json:"donator_is_user"`
	DonatorName   string      `json:"donator_name"`
	Etc           DonationEtc `json:"etc"`
	ID            string      `json:"id"`
	Message       string      `json:"message"`
	Type          string      `json:"type"`
	Version       string      `json:"version"`
	ReceivedAt    time.Time   `json:"received_at"` // Tambahan untuk tracking
}

// Custom unmarshal untuk handle time parsing
func (d *SaweriaDonation) UnmarshalJSON(data []byte) error {
	type Alias SaweriaDonation
	aux := &struct {
		CreatedAt string `json:"created_at"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse waktu dari string ke time.Time
	parsedTime, err := time.Parse(time.RFC3339, aux.CreatedAt)
	if err != nil {
		// Fallback ke waktu sekarang jika parsing gagal
		d.CreatedAt = time.Now()
	} else {
		d.CreatedAt = parsedTime
	}

	d.ReceivedAt = time.Now()
	return nil
}
