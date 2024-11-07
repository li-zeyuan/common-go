package utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenPwd(t *testing.T) {
	pwdB, err := GenPwd(context.Background(), "admin")
	assert.Equal(t, err, nil)
	t.Log(string(pwdB))
}

func TestComparePwd(t *testing.T) {
	b := ComparePwd("$2a$10$pT.o43FrURN0eg/L7pX8fePkLM8zVBSPMV1TRf4nRqAFIHrUxT/j.", "1111")
	assert.True(t, b)
}
