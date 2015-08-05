package config

import (
	"encoding/hex"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HashFile []string `yaml:"hash_file"`
	Restore  []string `yaml:"restore"`
}

// Load a file and build a list of Configs.
func LoadConfig(f string) ([]Config, error) {
	var c []Config

	filename, _ := filepath.Abs(f)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (c *Config) Hash() (string, error) {
	h, e := ComputeMd5(c.HashFile)
	if e != nil {
		return "", e
	}
	return hex.EncodeToString(h), nil
}

func (c *Config) HashFileFlat() string {
	return strings.Join(c.HashFile, ",")
}

func (c *Config) RestoreFlat() string {
	return strings.Join(c.Restore, ",")
}
