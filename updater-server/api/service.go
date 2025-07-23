package api

import (
	"fmt"
	"os/exec"
	"strings"
	"updater-server/internal/common/logger"
)

type UpdaterService struct {
	// In production, the repo would be an interface to a database or external service.
	// For this challenge, I am using the os file system to simulate the repository.
	//repo          	UpdaterRepository
	DataFilePath string
	DataFileName string
	logger       *logger.Logger
}

func NewUpdaterService(dataFilePath, dataFileName string, logger *logger.Logger) *UpdaterService {
	return &UpdaterService{
		DataFilePath: dataFilePath,
		DataFileName: dataFileName,
		logger:       logger,
	}
}

func (service *UpdaterService) GetLatestVersionNumber(serviceName string) (string, error) {
	serviceBinaryPath := fmt.Sprintf("%s/%s", service.DataFilePath, serviceName)
	cmd := exec.Command("go", "version", "-m", serviceBinaryPath)

	service.logger.Info("reading binary", "service", serviceName, "path", serviceBinaryPath)

	versionInfo, err := cmd.CombinedOutput()
	if err != nil {
		service.logger.Error("unable to read version information from binary file", "service", serviceName)
		return "", err
	}

	versionSubstring := "main.version="
	if !strings.Contains(string(versionInfo), versionSubstring) {
		service.logger.Error("binary file does not contain version information", "service", serviceName, "build variable", versionSubstring)
		return "", err
	}

	versionSplit := string(versionInfo)[strings.Index(string(versionInfo), versionSubstring) + len(versionSubstring):]

	version := versionSplit[:strings.Index(versionSplit, `"`)]

	return version, nil
}
