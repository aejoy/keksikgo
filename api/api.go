package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type API struct {
	Version int
	Group   int
	Token   string
	Client  *http.Client
	Logger  *zap.Logger
}

func New(group int, token string, logger *zap.Logger) *API {
	api := &API{
		Version: 1,
		Group:   group,
		Token:   token,
		Client:  http.DefaultClient,
		Logger:  logger,
	}

	return api
}

func (api *API) Call(context context.Context, method string, params any, data any) error {
	body, err := json.Marshal(params)
	if err != nil {
		api.Logger.Error(err.Error(), zap.String("type", "MarshalParams"))
		return err
	}

	api.Logger.Info("NewMethod",
		zap.String("method", method),
		zap.String("body", string(body)))

	request, err := http.NewRequestWithContext(
		context, http.MethodGet,
		"https://api.keksik.io/"+method,
		bytes.NewBuffer(body),
	)
	if err != nil {
		api.Logger.Error(err.Error(), zap.String("type", "NewRequest"))
		return err
	}

	response, err := api.Client.Do(request)
	if err != nil {
		api.Logger.Error(err.Error(), zap.String("type", "DoRequest"))
		return err
	}

	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		api.Logger.Error(err.Error(), zap.String("type", "DecodeResponse"))
		return err
	}

	return nil
}

func (api *API) NewMapRequiredParams(keyValues map[string]any) map[string]any {
	params := map[string]any{
		"v":     api.Version,
		"group": api.Group,
		"token": api.Token,
	}

	for key, value := range keyValues {
		params[key] = value
	}

	return params
}
