package emailVerification

import (
	"github.com/xlzd/gotp"
)

func CodeGenerator() string{
	secret:=gotp.RandomSecret(16)
	//fmt.Println(secret)
	totp := gotp.NewDefaultTOTP(secret)
	code:= totp.Now()
	return code
}
