package tests

import (
	"github.com/aejoy/keksikgo/api"
	"go.uber.org/zap"
	"os"
	"strconv"
)

func newAPI() (*api.API, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	groupID, err := strconv.Atoi(os.Getenv("GROUP_ID"))
	if err != nil {
		return nil, err
	}

	return api.New(groupID, os.Getenv("TOKEN"), logger), nil
}
