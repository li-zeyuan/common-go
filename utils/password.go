package utils

import (
	"regexp"

	"github.com/li-zeyuan/common-go/mylogger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

func GenPwd(ctx context.Context, pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		mylogger.Error(ctx, "generate pwd fail", zap.Error(err))
		return nil, err
	}

	return hash, nil
}

func ComparePwd(hashPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		return false
	}

	return true
}

func PwdFormat(pwd string) bool {
	b, err := regexp.MatchString(`^[0-9a-zA-Z]*$`, pwd)
	if err != nil {
		return false
	}

	return b
}
