package api

import (
	"context"
	"github.com/aejoy/keksikgo/responses"
	"github.com/aejoy/keksikgo/types"
)

func (api *API) GetCampaigns(context context.Context, ids []int) (campaigns responses.Campaigns, err error) {
	return campaigns, api.Call(context, "campaigns/get", api.NewMapRequiredParams(map[string]any{"ids": ids}), &campaigns)
}

func (api *API) GetActiveCampaign(context context.Context) (activeCampaign responses.ActiveCampaign, err error) {
	return activeCampaign, api.Call(context, "campaigns/get-active", api.NewMapRequiredParams(nil), &activeCampaign)
}

func (api *API) GetCampaignRewards(context context.Context, campaign int) (campaignRewards responses.CampaignRewards, err error) {
	return campaignRewards, api.Call(context, "campaigns/get-rewards", api.NewMapRequiredParams(map[string]any{"campaign": campaign}), &campaignRewards)
}

func (api *API) ChangeCampaign(context context.Context, changeCampaigns types.ChangeCampaign) (campaigns responses.ChangeCampaign, err error) {
	changeCampaigns.SetRequiredFields(api.Version, api.Group, api.Token)
	return campaigns, api.Call(context, "campaigns/change", changeCampaigns, &campaigns)
}

func (api *API) ChangeCampaignReward(context context.Context, changeCampaignReward types.ChangeCampaignReward) (campaignReward responses.ChangeCampaign, err error) {
	changeCampaignReward.SetRequiredFields(api.Version, api.Group, api.Token)
	return campaignReward, api.Call(context, "campaigns/change-reward", changeCampaignReward, &campaignReward)
}
