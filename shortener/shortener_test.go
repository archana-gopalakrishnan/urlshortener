package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	longurl := "https://www.youtube.com/watch?v=ryp60Q397b4"
	shorturl := GenerateShortUrl(longurl)

	assert.Equal(t, shorturl, "MUwZtCT1")
}
