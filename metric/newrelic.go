package metric

import (
	"fmt"
	newrelic "github.com/newrelic/go-agent"
)

func InitializeNewRelic(serviceName string, licenseKey string) (newrelic.Application, error) {
	if licenseKey == "" {
		return nil, fmt.Errorf("LicenseKey cannot be empty")
	}
	newRelicApp, err := newrelic.NewApplication(newrelic.NewConfig(serviceName, licenseKey))
	if err != nil {
		return nil, fmt.Errorf("error creating newrelic application: %v", err)
	}
	return newRelicApp, nil
}

