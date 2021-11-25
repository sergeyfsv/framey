package speedtest

import (
	"reflect"
	"testing"
	"time"

	"go.framey.io/speedtest/internal/core"
	"go.framey.io/speedtest/internal/types"
)

var (
	configuration     types.Configuration
	zeroConfiguration types.Configuration
)

func init() {
	configuration = types.DefaultConfiguration()

	zeroConfiguration = types.DefaultConfiguration()
	zeroConfiguration.ConfigTimeout = 1 * time.Millisecond
}

func TestSpeedMeasurement_WhenConvertToString_ExpectValidValue(t *testing.T) {
	tests := []struct {
		name     string
		instance core.SpeedMeasurement
		want     string
	}{
		{
			name:     "SpeedMeasurement(FastdotcomEndpoint)",
			instance: NewMeasurement(types.FastdotcomEndpoint, &configuration),
			want:     "fast.com",
		},
		{
			name:     "SpeedMeasurement(SpeedtestdotnetEndpoint)",
			instance: NewMeasurement(types.SpeedtestdotnetEndpoint, &configuration),
			want:     "speedtest.net",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.instance.String(); got != tt.want {
				t.Errorf("SpeedMeasurement.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitMeasurement_WhenCreates_ExpectValidInstance(t *testing.T) {
	type args struct {
		endpointType types.SpeedMeasurementEndpointType
		cfg          *types.Configuration
	}
	tests := []struct {
		name string
		args args
		want core.SpeedMeasurement
	}{
		{
			name: "InitMeasurement(FastdotcomEndpoint)",
			args: args{
				endpointType: types.FastdotcomEndpoint,
				cfg:          &configuration,
			},
		},
		{
			name: "InitMeasurement(SpeedtestdotnetEndpoint)",
			args: args{
				endpointType: types.SpeedtestdotnetEndpoint,
				cfg:          &configuration,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMeasurement(tt.args.endpointType, tt.args.cfg)
			if !reflect.DeepEqual(*got.Configuration, configuration) {
				t.Errorf("InitMeasurement() = %v, Modified configuration.", got.Configuration)
			}
			if nil == got.Endpoint {
				t.Errorf("InitMeasurement() = %v, Invalid endpoint interface.", got.Endpoint)
			}
		})
	}
}

func TestSpeedMeasurement_WhenConfiguraitonConfigTimeoutIsZero_ExpectUploadFail(t *testing.T) {
	tests := []struct {
		name          string
		instance      core.SpeedMeasurement
		expectedSpeed float64
		expectError   bool
	}{
		{
			name:          "UploadZeroConfiguration(FastdotcomEndpoint)",
			instance:      NewMeasurement(types.FastdotcomEndpoint, &zeroConfiguration),
			expectedSpeed: 0,
			expectError:   true,
		},
		{
			name:          "UploadZeroConfiguration(SpeedtestdotnetEndpoint)",
			instance:      NewMeasurement(types.SpeedtestdotnetEndpoint, &zeroConfiguration),
			expectedSpeed: 0,
			expectError:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSpeed, err := Upload(&tt.instance)
			if (err != nil) != tt.expectError {
				t.Errorf("SpeedMeasurement.Upload(%v) error = %v, expectError %v", tt.instance, err, tt.expectError)
				return
			}
			if gotSpeed != tt.expectedSpeed {
				t.Errorf("SpeedMeasurement.Upload(%v) = %v, want %v", tt.instance, gotSpeed, tt.expectedSpeed)
			}
		})
	}
}

func TestSpeedMeasurement_WhenConfiguraitonConfigTimeoutIsZero_ExpectDownloadFail(t *testing.T) {
	tests := []struct {
		name          string
		instance      core.SpeedMeasurement
		expectedSpeed float64
		expectError   bool
	}{
		{
			name:          "DownloadZeroConfiguration(FastdotcomEndpoint)",
			instance:      NewMeasurement(types.FastdotcomEndpoint, &zeroConfiguration),
			expectedSpeed: 0,
			expectError:   true,
		},
		{
			name:          "DownloadZeroConfiguration(SpeedtestdotnetEndpoint)",
			instance:      NewMeasurement(types.SpeedtestdotnetEndpoint, &zeroConfiguration),
			expectedSpeed: 0,
			expectError:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSpeed, err := Download(&tt.instance)
			if (err != nil) != tt.expectError {
				t.Errorf("SpeedMeasurement.Download(%v) error = %v, expectError %v", tt.instance, err, tt.expectError)
				return
			}
			if gotSpeed != tt.expectedSpeed {
				t.Errorf("SpeedMeasurement.Download(%v) = %v, want %v", tt.instance, gotSpeed, tt.expectedSpeed)
			}
		})
	}
}

func TestSpeedMeasurement_WhenConfiguraitonIsValid_ExpectUploadSuccess(t *testing.T) {
	tests := []struct {
		name          string
		instance      core.SpeedMeasurement
		expectedSpeed float64
		expectError   bool
	}{
		{
			name:          "UploadValidConfig(FastdotcomEndpoint)",
			instance:      NewMeasurement(types.FastdotcomEndpoint, &configuration),
			expectedSpeed: 0,
			expectError:   false,
		},
		{
			name:          "UploadValidConfig(SpeedtestdotnetEndpoint)",
			instance:      NewMeasurement(types.SpeedtestdotnetEndpoint, &configuration),
			expectedSpeed: 0,
			expectError:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSpeed, err := Upload(&tt.instance)
			if (err != nil) != tt.expectError {
				t.Errorf("SpeedMeasurement.Upload(%v) error = %v, expectError %v", tt.instance, err, tt.expectError)
				return
			}
			if gotSpeed == 0 {
				t.Errorf("SpeedMeasurement.Upload(%v) = %v, Speed is zero", tt.instance, gotSpeed)
			}
		})
	}
}

func TestSpeedMeasurement_WhenConfiguraitonIsValid_ExpectDownloadSuccess(t *testing.T) {
	tests := []struct {
		name          string
		instance      core.SpeedMeasurement
		expectedSpeed float64
		expectError   bool
	}{
		{
			name:          "DownloadValidConfig(FastdotcomEndpoint)",
			instance:      NewMeasurement(types.FastdotcomEndpoint, &configuration),
			expectedSpeed: 0,
			expectError:   false,
		},
		{
			name:          "DownloadValidConfig(SpeedtestdotnetEndpoint)",
			instance:      NewMeasurement(types.SpeedtestdotnetEndpoint, &configuration),
			expectedSpeed: 0,
			expectError:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSpeed, err := Download(&tt.instance)
			if (err != nil) != tt.expectError {
				t.Errorf("SpeedMeasurement.Download(%v) error = %v, expectError %v", tt.instance, err, tt.expectError)
				return
			}
			if gotSpeed == 0 {
				t.Errorf("SpeedMeasurement.Download(%v) = %v, Speed is zero", tt.instance, gotSpeed)
			}
		})
	}
}

/*
func TestSpeedMeasurement_Download(t *testing.T) {
	tests := []struct {
		name     string
		instance SpeedMeasurement
		want     float64
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.instance.Download()
			if (err != nil) != tt.wantErr {
				t.Errorf("SpeedMeasurement.Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SpeedMeasurement.Download() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
*/
