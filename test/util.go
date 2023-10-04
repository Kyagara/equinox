package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/lor"
	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/clients/tft"
	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

type TestCase[Model any, Parameters any] struct {
	Name        string
	Code        int
	Want        *Model
	WantError   error
	AccessToken string
	Parameters  *Parameters
	Options     map[string]interface{}
}

func (testCase TestCase[Model, Parameters]) MockGetResponse(url string, region string, accessToken string) {
	mock := gock.New(fmt.Sprintf(api.BaseURLFormat, region)).
		Get(url).
		Reply(testCase.Code).
		JSON(testCase.Want)

	if accessToken != "" {
		mock.SetHeader("Authorization", "accessToken")
	}
}

func (testCase TestCase[Model, Parameters]) MockPostResponse(url string, region string, accessToken string) {
	gock.New(fmt.Sprintf(api.BaseURLFormat, region)).
		Post(url).
		Reply(testCase.Code).
		JSON(testCase.Want)
}

func (testCase TestCase[Model, Parameters]) CheckResponse(t *testing.T, gotData *Model, gotErr error) {
	require.Equal(t, testCase.WantError, gotErr, fmt.Sprintf("want error '%v', got '%v'", testCase.WantError, gotErr))

	if testCase.WantError == nil {
		require.Equal(t, testCase.Want, gotData, "result not expected")
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

func TestingNewRiotClient() (*riot.RiotClient, error) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	if err != nil {
		return nil, err
	}

	return riot.NewRiotClient(internalClient), nil
}

func TestingNewLOLClient() (*lol.LOLClient, error) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	if err != nil {
		return nil, err
	}

	return lol.NewLOLClient(internalClient), nil
}

func TestingNewTFTClient() (*tft.TFTClient, error) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	if err != nil {
		return nil, err
	}

	return tft.NewTFTClient(internalClient), nil
}

func TestingNewVALClient() (*val.VALClient, error) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	if err != nil {
		return nil, err
	}

	return val.NewVALClient(internalClient), nil
}

func TestingNewLORClient() (*lor.LORClient, error) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	if err != nil {
		return nil, err
	}

	return lor.NewLORClient(internalClient), nil
}
