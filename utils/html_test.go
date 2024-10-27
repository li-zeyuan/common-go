package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHtmlUnescapeAndTrim(t *testing.T) {
	assert.Equal(t, HtmlUnescapeAndTrim(``), "")
	assert.Equal(t, HtmlUnescapeAndTrim(`<p>&nbsp;hello world</p>`), " hello world")
}
