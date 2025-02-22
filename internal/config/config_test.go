package config_test

import (
	"testing"

	"github.com/tommjj/ql-kho-lua/internal/config"
)

func TestConfig(t *testing.T) {
	conf, err := config.New()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", conf)
}
