package metric

import (
	newrelic "github.com/newrelic/go-agent"
	"reflect"
	"testing"
)

func TestInitializeNewRelic(t *testing.T) {
	type args struct {
		serviceName string
		licenseKey  string
	}
	tests := []struct {
		name    string
		args    args
		want    newrelic.Application
		wantErr bool
	}{
		{
			"Error_Empty_LicenseKey",
			args{
				serviceName: "test",
				licenseKey:  "",
			},
			nil,
			true,
		},
		{
			"Error_LicenseKey_should_be_length_40",
			args{
				serviceName: "test",
				licenseKey:  "random",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitializeNewRelic(tt.args.serviceName, tt.args.licenseKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitializeNewRelic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeNewRelic() got = %v, want %v", got, tt.want)
			}
		})
	}
}
