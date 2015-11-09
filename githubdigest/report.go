package githubdigest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/pebble/github-digest/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/pebble/github-digest/Godeps/_workspace/src/github.com/mailgun/mailgun-go"
)

func GenerateReport(c *cli.Context, stats *GithubDigest) (*string, error) {
	if c.Bool("json") {
		statsJson, err := json.Marshal(stats)
		if err != nil {
			return nil, err
		}
		report := string(statsJson)
		return &report, nil
	}

	t, err := template.ParseFiles("report.html")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	t.Execute(&buf, stats)
	report := buf.String()
	return &report, nil
}

func SendReport(c *cli.Context, report string) error {
	mailgunApiKey := c.String("mailgun")
	emailTo := c.String("mail-to")
	if mailgunApiKey != "" && emailTo != "" {
		emailFrom := c.String("mail-from")
		mailgunDomain := c.String("mailgun-domain")
		mg := mailgun.NewMailgun(mailgunDomain, mailgunApiKey, "")
		m := mg.NewMessage(emailFrom,
			"GitHub PR digest",
			report,
			emailTo,
		)
		if !c.Bool("json") {
			m.SetHtml(report)
		}

		_, _, err := mg.Send(m)
		if err != nil {
			return err
		}
	} else {
		fmt.Println(report)
	}
	return nil
}
