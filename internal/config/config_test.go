package config_test

import (
	"reflect"
	"testing"

	"github.com/takumin/path-filter-action/internal/config"
)

func TestNewConfig(t *testing.T) {
	if !reflect.DeepEqual(config.NewConfig(), &config.Config{}) {
		t.Error("expected config struct to be equal, but got not equal")
	}

	if !reflect.DeepEqual(config.NewConfig(config.LogLevel("TEST")), &config.Config{LogLevel: "TEST"}) {
		t.Error("expected config struct to be equal, but got not equal")
	}
}
