package commands

import (
	"log"
	"os"

	"github.com/Masmerc/fsp/pkg/actions"
	"github.com/urfave/cli/v2"
)

func Commands() {
	app := &cli.App{
		Name:  "fsp",
		Usage: "sprint planning helper (v.0.1.0)",
	}

	app.Commands = []*cli.Command{
		{
			Name:  "bulk-create",
			Usage: "create issues from a csv file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "file",
					Aliases:  []string{"f"},
					Usage:    "path to input csv file",
					Required: true,
				},
				&cli.StringSliceFlag{
					Name:    "labels",
					Aliases: []string{"l"},
					Usage:   "labels to add",
				},
				&cli.StringFlag{
					Name:    "project-id",
					Aliases: []string{"p"},
					Usage:   "project id",
				},
			},
			Action: actions.BulkCreateIssues,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
