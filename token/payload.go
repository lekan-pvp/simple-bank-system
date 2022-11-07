package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrExpireToken  = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// PayLoad contains the payload data of the token
type PayLoad struct {
	ID        uuid.UUID `json:"id`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with specific username and duration
func NewPayload(username string, duration time.Duration) (*PayLoad, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &PayLoad{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *PayLoad) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpireToken
	}
	return nil
}
