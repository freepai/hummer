package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShortUrl(t *testing.T) {
	su := NewShortUrl("default", "X1", "http://www.baidu.com")
	assert.NotNil(t, su, "new object should not be nil")
}
