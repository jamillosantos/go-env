package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv_WithPrefix(t *testing.T) {
	wantPrefix := "prefix"
	wantOverrides := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	env := NewEnv().WithOverrides(wantOverrides).WithPrefix(wantPrefix)
	assert.Equal(t, wantPrefix, env.prefix)
	assert.Equal(t, wantOverrides, env.overrides)
}

func TestEnv_WithOverrides(t *testing.T) {
	wantPrefix := "prefix"
	wantOverrides := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	env := NewEnv().WithPrefix(wantPrefix).WithOverrides(wantOverrides)
	assert.Equal(t, wantPrefix, env.prefix)
	assert.Equal(t, wantOverrides, env.overrides)
}
