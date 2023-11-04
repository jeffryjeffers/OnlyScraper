package actions

import (
	"OnlyScraper/internal/api"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func SubsAction(c *cli.Context) error {
	in := c.Context.Value("instance").(*api.OnlyFans)
	logger := c.Context.Value("logger").(pterm.Logger)

	subscriptions, err := in.GetSubscriptions()
	if err != nil {
		return err
	}
	var models []string
	for _, sub := range subscriptions.List {
		models = append(models, sub.Username)
	}

	logger.Info("Active subscriptions", logger.ArgsFromMap(map[string]any{
		"count":  len(models),
		"models": models,
	}))
	return nil
}
