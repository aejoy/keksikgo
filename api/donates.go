package api

import (
	"context"
	"github.com/aejoy/keksikgo/responses"
	"github.com/aejoy/keksikgo/types"
)

func (api *API) GetDonates(context context.Context, getDonates types.GetDonates) (donates responses.Donates, err error) {
	getDonates.SetRequiredFields(api.Version, api.Group, api.Token)
	return donates, api.Call(context, "donates/get", getDonates, &donates)
}

func (api *API) GetLastDonates(context context.Context, getLastDonates types.GetLastDonates) (donates responses.Donates, err error) {
	getLastDonates.SetRequiredFields(api.Version, api.Group, api.Token)
	return donates, api.Call(context, "donates/get-last", getLastDonates, &donates)
}

func (api *API) ChangeStatus(context context.Context, changeStatus types.ChangeStatus) (status responses.Status, err error) {
	changeStatus.SetRequiredFields(api.Version, api.Group, api.Token)
	return status, api.Call(context, "donates/change-status", changeStatus, &status)
}

func (api *API) Answer(context context.Context, answer types.Answer) (info responses.Answer, err error) {
	answer.SetRequiredFields(api.Version, api.Group, api.Token)
	return info, api.Call(context, "donates/answer", answer, &info)
}

func (api *API) ChangeReward(context context.Context, changeReward types.ChangeReward) (info responses.ChangeReward, err error) {
	changeReward.SetRequiredFields(api.Version, api.Group, api.Token)
	return info, api.Call(context, "donates/change-reward-status", changeReward, &info)
}
