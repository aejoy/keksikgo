package api

import (
	"context"
	"github.com/aejoy/keksikgo/responses"
	"github.com/aejoy/keksikgo/types"
)

func (api *API) GetPayments(context context.Context, ids []int) (payments responses.Payments, err error) {
	return payments, api.Call(context, "payments/get", api.NewMapRequiredParams(map[string]any{"ids": ids}), &payments)
}

func (api *API) CreatePayments(context context.Context, createPayments types.CreatePayments) (payments responses.CreatePayments, err error) {
	createPayments.SetRequiredFields(api.Version, api.Group, api.Token)
	return payments, api.Call(context, "payments/create", createPayments, &payments)
}
