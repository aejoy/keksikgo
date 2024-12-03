package api

import (
	"context"
	"github.com/aejoy/keksikgo/responses"
)

func (api *API) Balance(context context.Context) (balance responses.Balance, err error) {
	return balance, api.Call(context, "balance", api.NewMapRequiredParams(nil), &balance)
}
