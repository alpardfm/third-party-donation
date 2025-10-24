package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// VerifySignature memverifikasi signature dari Saweria
func VerifySignature(payload interface{}, signature, secret string) (bool, error) {
	// Convert payload to JSON bytes
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}

	// Create HMAC SHA256 hash
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(payloadBytes)
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	// Compare signatures
	return hmac.Equal([]byte(signature), []byte(expectedSignature)), nil
}
