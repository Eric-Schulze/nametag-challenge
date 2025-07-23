package init

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"updater-server/api"
	"updater-server/internal/common/logger"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	DataFilePath    string `yaml:"data_file_path"`
	DataFileName    string `yaml:"data_file_name"`
	MinimumLogLevel string `yaml:"minimum_log_level"`
}

var SETTING_FILEPATH = filepath.Join("internal", "init", "config.yaml")

func Start(ctx context.Context, w io.Writer, args []string) error {
	settingsFile, err := os.Open(SETTING_FILEPATH)
	if err != nil {
		return err
	}
	defer settingsFile.Close()

	settingBytes, err := io.ReadAll(settingsFile)
	if err != nil {
		return err
	}

	var settings Settings
	if err := yaml.Unmarshal([]byte(settingBytes), &settings); err != nil {
		return fmt.Errorf("cannot unmarshal settings data: %v", err)
	}

	logger := logger.Logger{
		Output:      w,
		MinLogLevel: logger.LogInfo,
	}

	updaterHandler := api.NewUpdaterHandler(&logger, settings.DataFilePath, settings.DataFileName)

	// Create a new router
	router := initRouter(logger, *updaterHandler)

	logger.Info("Starting updater server on port 4040")
	http.ListenAndServe(":4040", router)

	return nil
}
