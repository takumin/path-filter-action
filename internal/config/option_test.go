package config_test

import (
	"reflect"
	"testing"

	"github.com/takumin/path-filter-action/internal/config"
)

func TestLogLevel(t *testing.T) {
	want := &config.Config{LogLevel: "TEST"}
	got := &config.Config{}
	config.LogLevel("TEST").Apply(got)
	if !reflect.DeepEqual(want, got) {
		t.Error("expected config struct to be equal, but got not equal")
	}
}

func TestLogFormat(t *testing.T) {
	want := &config.Config{LogFormat: "TEST"}
	got := &config.Config{}
	config.LogFormat("TEST").Apply(got)
	if !reflect.DeepEqual(want, got) {
		t.Error("expected config struct to be equal, but got not equal")
	}
}

func TestGitHubToken(t *testing.T) {
	want := &config.Config{GitHubToken: "TEST"}
	got := &config.Config{}
	config.Variable("TEST").Apply(got)
	if !reflect.DeepEqual(want, got) {
		t.Error("expected config struct to be equal, but got not equal")
	}
}
