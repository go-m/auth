package otp

import (
	"github.com/go-m/auth/base"
	"regexp"
	"time"
)

type Config struct {
	LoginHandler base.ILoginHandler

	CodeExpiration                    time.Duration
	MaxRequestRetries                 int
	MaxVerifyRetries                  int
	ValidationExpiration              time.Duration
	ResetMaxVerifyRetriesOnNewRequest bool
	MobileValidationRegexPattern      *string
	ValidateMobile                    func(mobile string) error
	ValidateOtp                       func(otp *OTP, code string) error

	MobileValidationRegex *regexp.Regexp
}
