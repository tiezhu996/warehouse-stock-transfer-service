package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	c := Load()
	if c.AppName == "" {
		t.Fatal("AppName should not be empty")
	}
}
