package handler

import (
	"time"

	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/services/pkg/gorm"
)

var (
	ErrMissingCode     = errors.BadRequest("MISSING_CODE", "Missing code")
	ErrMissingIdentity = errors.BadRequest("MISSING_IDENTITY", "Missing identity")
	ErrInvalidCode     = errors.BadRequest("INVALID_CODE", "Invalid code")
	ErrExpiredCode     = errors.BadRequest("EXPIRED_CODE", "Expired code")

	DefaultTTL = time.Minute * 5
)

type Codes struct {
	gorm.Helper
	Time func() time.Time
}

type Code struct {
	Code      string `gorm:"index:codeIdentity"`
	Identity  string `gorm:"index:codeIdentity"`
	ExpiresAt time.Time
}
