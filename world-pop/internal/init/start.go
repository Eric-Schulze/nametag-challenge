package init

import (
	"context"
	"embed"
	"fmt"
	"io"
	"path/filepath"

	"world-pop/cmd"
	"world-pop/internal/common/logger"
	"world-pop/internal/common/models"
	"world-pop/internal/data"
	"world-pop/internal/updater"

	"gopkg.in/yaml.v3"
)

var SETTING_FILEPATH = filepath.Join("internal", "init", "config.yaml")	

//go:embed config.yaml
var f embed.FS

func Start(ctx context.Context, w io.Writer, args []string, version string) error {
	settingBytes, err := f.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	var settings models.Settings
	if err := yaml.Unmarshal([]byte(settingBytes), &settings); err != nil {
		return fmt.Errorf("cannot unmarshal settings data: %v", err)
	}

	settings.SettingFilePath = SETTING_FILEPATH

	logger := logger.Logger{
		Output:      w,
		MinLogLevel: logger.LogInfo,
		LoggingEnabled: settings.EnableLogging,
	}

	dataManager, err := data.NewCountryDataManager(settings.DataFilePath, logger)
	if err != nil {
		logger.Error("failed to initialize data manager", "error", err)
		return err
	}

	updaterClient := &updater.UpdaterClient{
		BaseUrl:        settings.UpdaterServerUrl,
		Logger:         &logger,
		CurrentVersion: version,
	}

	command := cmd.BuildCommand(ctx, args, settings, dataManager, updaterClient)
	command.Writer = w

	if err := command.Run(ctx, args); err != nil {
		return err
	}

	return nil
}
