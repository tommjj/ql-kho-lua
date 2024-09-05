package config

import "testing"

func TestConfig(t *testing.T) {
	conf, err := New()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(conf.App)
	t.Log(conf.Auth)
	t.Log(conf.Http)
	t.Log(conf.Logger)
}
