package completion_test

import (
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/takumin/path-filter-action/internal/command/completion"
	"github.com/takumin/path-filter-action/internal/config"
)

func TestNewCommands(t *testing.T) {
	cfg := config.NewConfig()
	flags := []cli.Flag{}
	cmd := completion.NewCommands(cfg, flags)

	if cmd.Name != "completion" {
		t.Errorf("expected command name to be 'completion', but got '%s'", cmd.Name)
	}

	if cmd.Usage != "command completion" {
		t.Errorf("expected command usage to be 'command completion', but got '%s'", cmd.Usage)
	}

	for _, subcmd := range cmd.Subcommands {
		if subcmd == nil {
			t.Errorf("expected subcommand to not be nil")
		}
	}
}
