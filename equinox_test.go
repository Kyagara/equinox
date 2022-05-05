package equinox_test

import (
	"fmt"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEquinoxClient(t *testing.T) {
	tests := []struct {
		name    string
		want    *equinox.Equinox
		wantErr error
		key     string
	}{
		{
			name: "success",
			want: &equinox.Equinox{},
			key:  "RIOT_API_KEY",
		},
		{
			name:    "nil key",
			wantErr: fmt.Errorf("API Key not provided"),
			key:     "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotData, gotErr := equinox.NewClient(test.key)

			if test.name != "success" {
				require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

				if test.wantErr == nil {
					assert.Equal(t, test.want, gotData)
				}
			} else {
				require.NotEmpty(t, gotData, "expecting not empty client")
			}
		})
	}
}

func TestNewEquinoxClientWithConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *equinox.Equinox
		wantErr error
		config  *api.EquinoxConfig
	}{
		{
			name:   "success",
			want:   &equinox.Equinox{},
			config: internal.NewTestEquinoxConfig(),
		},
		{
			name:    "nil config",
			wantErr: fmt.Errorf("equinox configuration not provided"),
			config:  nil,
		},
		{
			name:    "api key nil",
			wantErr: fmt.Errorf("API Key not provided"),
			config: &api.EquinoxConfig{
				LogLevel: api.DebugLevel,
				Timeout:  10,
				Retry:    true,
			},
		},
		{
			name:    "cluster nil",
			wantErr: fmt.Errorf("cluster not provided"),
			config: &api.EquinoxConfig{
				Key:      "RIOT_API_KEY",
				LogLevel: api.DebugLevel,
				Timeout:  10,
				Retry:    true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotData, gotErr := equinox.NewClientWithConfig(test.config)

			if test.name != "success" {
				require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

				if test.wantErr == nil {
					assert.Equal(t, test.want, gotData)
				}
			} else {
				require.NotEmpty(t, gotData, "expecting not empty client")
			}
		})
	}
}
