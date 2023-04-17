package util

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestingNewLOLClient() (*lol.LOLClient, error) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	if err != nil {
		return nil, err
	}

	return lol.NewLOLClient(internalClient), nil
}

type TestCase[Model any, Parameters any] struct {
	Name        string
	Code        int
	Want        *Model
	WantError   error
	AccessToken string
	Parameters  *Parameters
	Options     map[string]interface{}
}

func (testCase TestCase[Model, Parameters]) MockResponse(url string, region string, accessToken string) {
	mock := gock.New(fmt.Sprintf(api.BaseURLFormat, region)).
		Get(url).
		Reply(testCase.Code).
		JSON(testCase.Want)

	if accessToken != "" {
		mock.SetHeader("Authorization", "accessToken")
	}
}

func (testCase TestCase[Model, Parameters]) CheckResponse(t *testing.T, gotData *Model, gotErr error) {
	require.Equal(t, testCase.WantError, gotErr, fmt.Sprintf("want err %v, got %v", testCase.WantError, gotErr))

	if testCase.WantError == nil {
		assert.Equal(t, testCase.Want, gotData)
	}
}

func GetEndpointTestCases[Model any, Parameters any](model Model, parameters *Parameters) []TestCase[Model, Parameters] {
	return []TestCase[Model, Parameters]{
		{
			Name:        "found",
			Code:        http.StatusOK,
			Want:        &model,
			WantError:   nil,
			AccessToken: "",
			Parameters:  parameters,
		},
		{
			Name:        "not found",
			Code:        http.StatusNotFound,
			Want:        nil,
			WantError:   api.ErrNotFound,
			AccessToken: "",
			Parameters:  parameters,
		},
	}
}
