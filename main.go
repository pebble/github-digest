package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/pebble/github-digest/githubdigest"
	"github.com/pebble/github-digest/Godeps/_workspace/src/github.com/codegangsta/cli"
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
			Name:   "cutoff",
			Value:  21,
			Usage:  "Days of pulls to consider",
			EnvVar: "GITHUB_CUTOFF",
		},
		cli.IntFlag{
			Name:   "closed-cutoff",
			Value:  1,
			Usage:  "Days of merged pulls to consider",
			EnvVar: "GITHUB_CLOSED_CUTOFF",
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
			Name:   "repos",
			EnvVar: "GITHUB_REPOS",
		},
		cli.StringFlag{
			Name:   "mail-to",
			Usage:  "Email recipient",
			EnvVar: "DIGEST_MAIL_TO",
		},
		cli.StringFlag{
			Name:   "mail-from",
			Usage:  "Email sender",
			Value:  "noreply@pebble.com",
			EnvVar: "DIGEST_MAIL_FROM",
		},
		cli.StringFlag{
			Name:   "mailgun",
			Usage:  "MailGun API key",
			EnvVar: "DIGEST_MAILGUN_API_KEY",
		},
		cli.StringFlag{
			Name:   "mailgun-domain",
			Usage:  "MailGun domain",
			Value:  "getpebble.com",
			EnvVar: "DIGEST_MAILGUN_DOMAIN",
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

		// Collect stats:
		digester := githubdigest.NewDigester(oauthToken)
		stats, err := digester.GetDigest(repos, statCutoff, closedCutoff)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Generate report:
		report, err := githubdigest.GenerateReport(c, stats)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Output:
		githubdigest.SendReport(c, *report)
	}

	app.Run(os.Args)
}
