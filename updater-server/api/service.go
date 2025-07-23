package api

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"updater-server/internal/common/logger"

	"gopkg.in/yaml.v3"
)

type UpdaterService struct {
	// In production, the repo would be an interface to a database or external service.
	// For this challenge, I am using the os file system to simulate the repository.
	//repo          	UpdaterRepository
	DataFilePath string
	DataFileName string
	logger       *logger.Logger
}

type Metadata struct {
	Version		string
	Checksum	string
}

func NewUpdaterService(dataFilePath, dataFileName string, logger *logger.Logger) *UpdaterService {
	return &UpdaterService{
		DataFilePath: dataFilePath,
		DataFileName: dataFileName,
		logger:       logger,
	}
}

var SERVICE_NAME = "world-pop"

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

func (service *UpdaterService) GetCurrentChecksum(serviceName string) (string, error) {
	serviceBinaryPath := fmt.Sprintf("%s/%s", service.DataFilePath, serviceName)
	metadataPath := serviceBinaryPath + "_metadata.yaml"

	metadataFile, err := os.ReadFile(metadataPath)
    if err != nil {
		service.logger.Error("reading metadata file for service", "service", serviceName, "error", err)
		return "", err
    }

	metadata := Metadata{}

	if err = yaml.Unmarshal(metadataFile, &metadata); err != nil {
		service.logger.Error("parsing metadata file for service", "service", serviceName, "error", err)
		return "", err
    }

	return metadata.Checksum, nil
}

func (service *UpdaterService) UpdateData(reader io.ReadCloser) error {
	binaryPath := path.Join(service.DataFilePath, SERVICE_NAME)
	files, _ := os.ReadDir(service.DataFilePath)
    for _, file := range files {
        os.RemoveAll(path.Join(service.DataFilePath, file.Name()))
    }

	destination, err := os.Create(binaryPath)
	if err != nil {
		service.logger.Error("creating new binary file", "error", err)
		return err
	}
	defer destination.Close()

	if _, err := io.Copy(destination, reader); err != nil {
		service.logger.Error("copying data into new binary file", "error", err)
		return err
	}

	version, err := service.GetLatestVersionNumber(SERVICE_NAME)
	if err != nil {
		service.logger.Error("getting version from uploaded file", "error", err)
		return err
	}

	checksum, err := service.generateChecksum(binaryPath)
	if err != nil {
		service.logger.Error("generating checksum for uploaded binary", "error", err)
		return err
	}

	metadata := Metadata{
		Version: version,
		Checksum: checksum,
	}

	metadataYaml, err := yaml.Marshal(&metadata)
    if err != nil {
		service.logger.Error("compiling metadata for uploaded file", "error", err)
		return err
    }

	if err := os.WriteFile(binaryPath + "_metadata.yaml", metadataYaml, 0644); err != nil {
		service.logger.Error("writing metadata into new file", "error", err)
		return err
	}

	service.logger.Info("Binary data and metadata updated successfully", "New Version", version)

	return nil
}

func (service *UpdaterService) generateChecksum(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		service.logger.Error("opening file for checksum generation", "error", err)
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		service.logger.Error("copying file to hasher", "error", err)
		return "", err
	}

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil)), nil
}