package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	a := []Config{
		{
			HashFile: []string{
				"Gemfile.lock",
			},
			Restore: []string{
				"vendor/bundle",
			},
		},
		{
			HashFile: []string{
				"composer.lock",
			},
			Restore: []string{
				"vendor",
			},
		},
		{
			HashFile: []string{
				"packages.json",
			},
			Restore: []string{
				"node_modules",
			},
		},
	}
	c, err := LoadConfig("tests/cache.yml")
	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, a, c, "Loaded the configuration.")
	}
}
