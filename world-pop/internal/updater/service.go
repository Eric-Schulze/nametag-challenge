package updater

import (
	"fmt"
	"io"
	"net/http"
	"world-pop/internal/common/logger"

	"github.com/sanbornm/go-selfupdate/selfupdate"
)

type UpdaterClient struct {
	BaseUrl        string
	CurrentVersion string
	Logger         *logger.Logger
}

var SERVICE_NAME = "world-pop"

func (client *UpdaterClient) Ping() (bool, error) {
	resp, err := http.Get(client.BaseUrl + "/ping")
	if err != nil {
		client.Logger.Error("error pinging updater server", "error", err)
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		client.Logger.Error("error pinging updater server", "status", resp.Status)
		return false, fmt.Errorf("error pinging updater server: %s", resp.Status)
	}

	return true, nil
}

func (client *UpdaterClient) GetLatestVersionNumber() (string, error) {
	versionUrl := fmt.Sprintf("%s/updater/latest?service-name=%s", client.BaseUrl, SERVICE_NAME)
	client.Logger.Info("Fetching latest version number from", "url", versionUrl)

	resp, err := http.Get(versionUrl)
	if err != nil {
		client.Logger.Error("error fetching latest version number", "error", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		client.Logger.Error("error fetching latest version number", "status", resp.Status)
		return "", fmt.Errorf("error fetching latest version number: %s", resp.Status)
	}

	version, err := io.ReadAll(resp.Body)
	if err != nil {
		client.Logger.Error("Error reading response body: %v", err)
		return "", err
	}

	return string(version), nil
}

func (client *UpdaterClient) IsLatestVersionCurrentlyInstalled() (bool, string, error) {
	latestVersion, err := client.GetLatestVersionNumber()
	if err != nil {
		client.Logger.Error("error getting latest version number", "error", err)
		return false, "", err
	}

	if latestVersion == client.CurrentVersion {
		client.Logger.Info("Current version is up to date", "version", client.CurrentVersion)
		return true, latestVersion, nil
	}

	client.Logger.Info("Current version is not up to date", "currentVersion", client.CurrentVersion, "latestVersion", latestVersion)
	return false, latestVersion, nil
}

func (client *UpdaterClient) UpdateService() (bool, error) {
	client.Logger.Info("Updating service to latest version", "currentVersion", client.CurrentVersion)

	isLatest, latestVersion, err := client.IsLatestVersionCurrentlyInstalled()
	if err != nil {
		client.Logger.Error("error checking if latest version is installed", "error", err)
		return false, err
	}

	if isLatest {
		client.Logger.Info("Service is already at the latest version", "version", client.CurrentVersion)
		return false, nil
	}

	client.Logger.Info("Service is not at the latest version", "Current Version", client.CurrentVersion, "Latest Version", latestVersion)

	var updater = &selfupdate.Updater{
		CurrentVersion: client.CurrentVersion,
		ApiURL:         fmt.Sprintf("%s/data/", client.BaseUrl),
		BinURL:         fmt.Sprintf("%s/data/", client.BaseUrl),
		DiffURL:        fmt.Sprintf("%s/data/", client.BaseUrl),
		Dir:            "./data/",
		CmdName:        "world-pop",
		ForceCheck:		true,
	}

	if err = updater.BackgroundRun(); err != nil {
		return false, err
	}

	client.CurrentVersion = updater.Info.Version

	client.Logger.Info("Service has been updated to the latest version", "Current Version", client.CurrentVersion)

	return true, nil
}
