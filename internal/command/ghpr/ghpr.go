package ghpr

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/path-filter-action/internal/config"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "github-token",
			Aliases:     []string{"t"},
			Usage:       "github token",
			EnvVars:     []string{"GITHUB_TOKEN"},
			Value:       cfg.GitHubToken,
			Destination: &cfg.GitHubToken,
		},
	}...)
	return &cli.Command{
		Name:    "github-pull-request",
		Aliases: []string{"ghpr"},
		Usage:   "GitHub Pull Request changed files",
		Flags:   flags,
		Action:  action(cfg),
	}
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return nil
	}
}
