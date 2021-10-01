package emailVerification

import (
	"github.com/xlzd/gotp"
)

// CodeGenerator generates code that user can use to verify account
func CodeGenerator() string{
	secret:=gotp.RandomSecret(16)
	totp := gotp.NewDefaultTOTP(secret)
	code:= totp.Now()
	return code
}
