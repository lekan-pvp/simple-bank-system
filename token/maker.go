package token

import "time"

// Maker is the interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	// VerifiyToken checks if the token is valid or not
	VerifyToken(token string) (*PayLoad, error)
}
