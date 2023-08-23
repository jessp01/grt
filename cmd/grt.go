package main

import (
	"log"
	"os"
	"time"

	"github.com/jessp01/grt"
	"github.com/urfave/cli"
)

const (
    repoHost = "@REPO_HOST@"
    repoOrg =  "@REPO_ORG@"
    repoName = "@REPO_NAME@"
)
var inputFilePath string
var outputFilePath string
var newRepo string

func populateAppMetadata(app *cli.App) {
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[input-file]{{end}}
   {{if len .Authors}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}{{ "\n" }}
   {{end}}{{end}}{{if .Copyright }}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COPYRIGHT:
   {{.Copyright}}
   {{end}}
`
	app.Usage = "GO Repo Template"
	app.Version = "0.21.0"
	app.EnableBashCompletion = true
	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}
	app.Compiled = time.Now()
	app.Description = "Bootstrapping code for GO repos.\n"
	app.Authors = []cli.Author{
		{
			Name:  "Jesse Portnoy",
			Email: "jesse@packman.io",
		},
	}
	app.Copyright = "(c) packman.io"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "input-file, i",
			Usage:       "Path to input file.\n",
			Destination: &inputFilePath,
		},
		cli.StringFlag{
			Name:        "output-file, o",
			Usage:       "Path to output file\n",
			Destination: &outputFilePath,
		},
		cli.StringFlag{
			Name:        "new-repo-name, r",
			Usage:       "Name of the new repo. Will replace references in input file",
			Destination: &newRepo,
		},
	}
}

func main() {
	app := cli.NewApp()
	populateAppMetadata(app)

	app.Action = func(c *cli.Context) error {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		if inputFilePath == "" {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}

		grt.ReplaceTokens(inputFilePath, outputFilePath, repoHost, repoOrg, repoName, newRepo)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
