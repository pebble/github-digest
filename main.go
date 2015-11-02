package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"time"
	"html/template"

	"github.com/pebble/github-digest/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/pebble/github-digest/githubdigest"
)

func dateArg(cutoff int) time.Time {
	invertedCutoff := math.Abs(float64(cutoff)) * -1
	return time.Now().AddDate(0, 0, int(invertedCutoff))
}

func main() {
	app := cli.NewApp()
	app.Name = "github-digest"
	app.Version = "0.0.1"
	app.Usage = "Report on github pull request activity"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "cutoff",
			Value: 21,
			Usage: "Days of pulls to consider",
			EnvVar: "GITHUB_CUTOFF",
		},
		cli.IntFlag{
			Name:  "closed-cutoff",
			Value: 1,
			Usage: "Days of merged pulls to consider",
			EnvVar: "GITHUB_CLOSED_CUTUFF",
		},
		cli.StringFlag{
			Name:   "oauth",
			Usage:  "Github OAuth token",
			EnvVar: "GITHUB_OAUTH_TOKEN",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Dump JSON instead of HTML",
		},
		cli.StringSliceFlag{
			Name: "repos",
			EnvVar: "GITHUB_REPOS",
		},
	}

	app.Action = func(c *cli.Context) {
		oauthToken := c.String("oauth")
		if oauthToken == "" {
			fmt.Println("Ouath token is required. Create one at https://github.com/settings/tokens")
			cli.ShowAppHelp(c)
			return
		}

		repos := c.StringSlice("repos")
		if len(repos) == 0 {
			repos = c.Args()
		}

		if len(repos) == 0 {
			fmt.Println("Specify at least 1 repo")
			cli.ShowAppHelp(c)
			return
		}

		statCutoff := dateArg(c.Int("cutoff"))
		closedCutoff := dateArg(c.Int("closed-cutoff"))

		digester := githubdigest.NewDigester(oauthToken)
		stats, err := digester.GetDigest(repos, statCutoff, closedCutoff)
		if err != nil {
			fmt.Println(err)
			return
		}

		if c.Bool("json") {
			statsJson, _ := json.Marshal(stats)
			fmt.Println(string(statsJson))
		} else {
			t, err := template.ParseFiles("report.html")
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}
			t.Execute(os.Stdout, stats)
		}
	}

	app.Run(os.Args)
}
