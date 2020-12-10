package plugin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupTest(ctx *Context) error {
	return nil
}

func TestRegister(t *testing.T) {
	err := Register("test", setupTest)
	assert.Nil(t, err, "Register a plugin func should be nil")
}

func TestGet(t *testing.T) {
	Register("test", setupTest)
	plugin := Get("test")
	assert.NotNil(t, plugin, "Get a plugin should not be nil")
}
