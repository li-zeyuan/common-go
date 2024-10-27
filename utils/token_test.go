package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(
		2393222912,
		time.Hour*24*300,
		"",
	)
	assert.Equal(t, err, nil)
	t.Log(token)
}
