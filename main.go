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
	inv_cutoff := math.Abs(float64(cutoff)) * -1
	return time.Now().AddDate(0, 0, int(inv_cutoff))
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
		},
		cli.IntFlag{
			Name:  "closed-cutoff",
			Value: 1,
			Usage: "Days of merged pulls to consider",
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
	}

	app.Action = func(c *cli.Context) {
		oauth_token := c.String("oauth")
		if oauth_token == "" {
			fmt.Println("Ouath token is required. Create one at https://github.com/settings/tokens")
			cli.ShowAppHelp(c)
			return
		}

		repos := c.Args()
		if len(repos) == 0 {
			fmt.Println("Specify at least 1 repo")
			cli.ShowAppHelp(c)
			return
		}

		stat_cutoff := dateArg(c.Int("cutoff"))
		closed_cutoff := dateArg(c.Int("closed-cutoff"))

		digester := githubdigest.NewDigester(oauth_token)
		stats := digester.GetDigest(repos, stat_cutoff, closed_cutoff)

		if c.Bool("json") {
			stats_json, _ := json.Marshal(stats)
			fmt.Println(string(stats_json))
		} else {
			t, err := template.ParseFiles("report.html")
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			t.Execute(os.Stdout, stats)
		}
	}

	app.Run(os.Args)
}
