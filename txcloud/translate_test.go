package txcloud

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextTranslate(t *testing.T) {
	cli := newCli()

	targetText, err := cli.TextTranslate(context.Background(), "欢乐番薯")
	assert.Nil(t, err)
	t.Log(targetText)
}
