package utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrl2ObjectKey(t *testing.T) {
	ctx := context.Background()
	fName0, _ := Url2ObjectKey(ctx, "https://example.com/path/to/resource.png?a=b")
	assert.Equal(t, fName0, "resource.png")

	fName, _ := Url2ObjectKey(ctx, "https://example.com/path/to/resource.png")
	assert.Equal(t, fName, "resource.png")

	fName2, _ := Url2ObjectKey(ctx, "example.com/path/to/resource.png")
	assert.Equal(t, fName2, "resource.png")

	fName3, _ := Url2ObjectKey(ctx, "resource.png")
	assert.Equal(t, fName3, "resource.png")

	fName4, _ := Url2ObjectKey(ctx, "")
	assert.Equal(t, fName4, "")
}
