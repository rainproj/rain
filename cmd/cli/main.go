package cli

import (
	c "github.com/urfave/cli/v2"

	"github.com/rainproj/rain/cmd/cli/commands/build"
	ini "github.com/rainproj/rain/cmd/cli/commands/init"
	"github.com/rainproj/rain/cmd/cli/commands/push"
)

// Execute cli
func Execute(version string, args []string) error {
	app := &c.App{
		Name:     "rain",
		HelpName: "rain",
		Usage:    "Build and deliver Docker images easily",
		Version:  version,
		Authors: []*c.Author{{
			Name:  "Weslei Juan Moser Pereira",
			Email: "wesleimsr@gmail.com",
		}},
		Commands: []*c.Command{
			ini.Command,
			build.Command,
			push.Command,
		},
	}

	return app.Run(args)
}
