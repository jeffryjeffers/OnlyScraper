package main

import (
	"OnlyScraper/internal/actions"
	"OnlyScraper/internal/api"
	"OnlyScraper/internal/config"
	"context"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	err := (&cli.App{
		Name:  "OnlyScraper",
		Usage: "A OnlyFans scraper",
		Before: func(c *cli.Context) error {
			logger := pterm.Logger{
				Formatter:  pterm.LogFormatterColorful,
				Writer:     os.Stdout,
				Level:      pterm.LogLevelInfo,
				ShowTime:   true,
				TimeFormat: "15:04:05",
				MaxWidth:   100,
				KeyStyles: map[string]pterm.Style{
					"error":  *pterm.NewStyle(pterm.FgRed, pterm.Bold),
					"err":    *pterm.NewStyle(pterm.FgRed, pterm.Bold),
					"caller": *pterm.NewStyle(pterm.FgGray, pterm.Bold),
				},
			}

			auth, err := config.LoadAuth("auth.json")
			if err != nil {
				return err
			}
			instance, err := api.Init(auth)
			if err != nil {
				return err
			}
			cfg, err := config.LoadConfig("config.json")
			if err != nil {
				return err
			}

			c.Context = context.WithValue(c.Context, "config", cfg)
			c.Context = context.WithValue(c.Context, "logger", logger)
			c.Context = context.WithValue(c.Context, "auth", auth)
			c.Context = context.WithValue(c.Context, "instance", instance)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:   "subs",
				Action: actions.SubsAction,
			},
			{
				Name:   "scrape",
				Action: actions.ScrapeAction,
				Flags: []cli.Flag{&cli.StringFlag{
					Name:  "model",
					Usage: "Specify which model(s) to scrape",
				}},
			},
		},
	}).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
