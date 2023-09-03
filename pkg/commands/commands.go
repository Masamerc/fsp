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
		Usage: "frozen sprint planning helper",
	}

	app.Commands = []*cli.Command{
		{
			Name:  "bulk-create",
			Usage: "create issues from a csv file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "file",
					Aliases: []string{"f"},
					Usage:   "path to input csv file",
				},
				&cli.StringSliceFlag{
					Name:    "labels",
					Aliases: []string{"l"},
					Usage:   "labels to add",
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
