package cmd

import (
	"context"
	"fmt"
	"world-pop/internal/data"
	"world-pop/internal/updater"

	altsrc "github.com/urfave/cli-altsrc/v3"
	yaml "github.com/urfave/cli-altsrc/v3/yaml"
	"github.com/urfave/cli/v3"
)

func BuildCommand(ctx context.Context, args []string, settingFilePath string, dataManager *data.CountryDataManager, updaterClient *updater.UpdaterClient) *cli.Command {
	var countryNameOrCode string
	var dataFilePath string

	cmd := &cli.Command{
		Name:    "world-pop",
		Version: updaterClient.CurrentVersion,
		Usage:   "Query world population data",
		Commands: []*cli.Command{
			{
				Name:        "country",
				Aliases:     []string{"c"},
				Usage:       "country <name or code> - get population data for a specific country",
				UsageText:   "country <name or code>",
				Description: "Get population data for a specific country",
				ArgsUsage:   "<name or code>",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name:        "countryNameOrCode",
						Destination: &countryNameOrCode,
					},
				},
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cmd.FullName()

					countryData, err := dataManager.GetCountryData(countryNameOrCode)
					if err != nil {
						dataManager.Logger.Error("Failed to get country data", "error", err)
						return err
					}

					countryData.Print(cmd.Root().Writer)

					return nil
				},
			},
			{
				Name:            "latest",
				Aliases:         []string{"l"},
				Usage:           "latest - get the latest version available of the world-pop service",
				UsageText:       "latest",
				Description:     "Get the latest version available of the world-pop service",
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
				Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
					fmt.Fprintf(cmd.Root().Writer, "Fetching latest version number from the updater server...\n")
					return nil, nil
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cmd.FullName()

					if success, err := updaterClient.Ping(); err != nil || !success {
						updaterClient.Logger.Error("failed to ping updater service", "error", err)
					}

					version, err := updaterClient.GetLatestVersionNumber()
					if err != nil {
						updaterClient.Logger.Error("failed to get latest version", "error", err)
						return err
					}

					fmt.Fprintf(cmd.Root().Writer, "Latest version available from host: %s\n", version)

					return nil
				},
			},
			{
				Name:            "check-update",
				Aliases:         []string{"c"},
				Usage:           "check-update - check if the latest version is currently installed",
				UsageText:       "check-update",
				Description:     "Check if the latest version of the world-pop app is currently installed",
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if success, err := updaterClient.Ping(); err != nil || !success {
						updaterClient.Logger.Error("failed to ping updater service", "error", err)
					}

					isLatest, latestVersion, err := updaterClient.IsLatestVersionCurrentlyInstalled()
					if err != nil {
						updaterClient.Logger.Error("failed to get latest version", "error", err)
						return err
					}

					if isLatest {
						fmt.Fprintf(cmd.Root().Writer, "%s service is running the latest version: %s\n", cmd.Root().Name, updaterClient.CurrentVersion)
					} else {
						fmt.Fprintf(cmd.Root().Writer, "%s service is behind the latest version.\nCurrent Version: %s\nLatest Version: %s\n", cmd.Root().Name, updaterClient.CurrentVersion, latestVersion)
					}

					return nil
				},
			},
			{
				Name:            "update",
				Aliases:         []string{"u"},
				Usage:           "update - install the latest version available of the world-pop service",
				UsageText:       "update",
				Description:     "Install the latest version available of the world-pop service",
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
				Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
					fmt.Fprintf(cmd.Root().Writer, "Fetching latest version number from the updater server...\n")
					return nil, nil
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if success, err := updaterClient.Ping(); err != nil || !success {
						updaterClient.Logger.Error("failed to ping updater service", "error", err)
					}

					isLatest, err := updaterClient.UpdateService()
					if err != nil {
						updaterClient.Logger.Error("failed to update service", "error", err)
						return err
					}

					if isLatest {
						fmt.Fprintf(cmd.Root().Writer, "No update required. %s service is running the latest version: %s\n", cmd.Root().Name, updaterClient.CurrentVersion)
					} else {
						fmt.Fprintf(cmd.Root().Writer, "%s service has been updated to the latest version: %s\n", cmd.Root().Name, updaterClient.CurrentVersion)
					}

					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "data-file",
				Aliases:     []string{"d"},
				Usage:       "file path to the world population data",
				Destination: &dataFilePath,
				Sources:     cli.NewValueSourceChain(yaml.YAML("data_file_path", altsrc.StringSourcer(settingFilePath))),
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			cli.DefaultAppComplete(ctx, cmd)
			cli.ShowAppHelp(cmd)
			cli.ShowVersion(cmd)

			return nil
		},
	}

	return cmd
}
