package validator

import (
	"encoding/xml"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/hguerra/jacoco-check/internal/models"
	"github.com/hguerra/jacoco-check/internal/printer"
)

func Validate(
	xmlReportPath string,
	filesChanged []string,
	coverageOverallCode,
	coverageNewCode float32,
) (string, error) {
	if strings.TrimSpace(xmlReportPath) == "" {
		return "", errors.New("xmlReportPath is mandatory")
	}

	if len(filesChanged) == 0 {
		return "", errors.New("filesChanged should not be empty")
	}

	xmlFile, err := os.Open(xmlReportPath)
	if err != nil {
		return "", err
	}
	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return "", err
	}

	var report models.Report
	err = xml.Unmarshal(byteValue, &report)
	if err != nil {
		return "", err
	}

	printer.Print([]string{"Jest: \"global\" coverage threshold for lines (90%) not met: 50%"})
	// errors.New("code coverage threshold not met")

	return "", nil
}
