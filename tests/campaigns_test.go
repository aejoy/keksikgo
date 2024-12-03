package tests

import (
	"context"
	"github.com/aejoy/keksikgo/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCampaigns(t *testing.T) {
	api, err := newAPI()
	if err != nil {
		t.Error(err)
	}

	t.Run("GetCampaigns", func(t *testing.T) {
		campaigns, err := api.GetCampaigns(context.TODO(), nil)
		assert.Equal(t, 0, campaigns.Error, campaigns.Message)
		assert.Equal(t, nil, err, err)
	})

	t.Run("GetActiveCampaign", func(t *testing.T) {
		activeCampaign, err := api.GetActiveCampaign(context.TODO())
		assert.Equal(t, 0, activeCampaign.Error, activeCampaign.Message)
		assert.Equal(t, nil, err, err)
	})

	t.Run("GetCampaignRewards", func(t *testing.T) {
		campaignRewards, err := api.GetCampaignRewards(context.TODO(), 21725)
		assert.Equal(t, 0, campaignRewards.Error, campaignRewards.Message)
		assert.Equal(t, nil, err, err)
	})

	t.Run("ChangeCampaign", func(t *testing.T) {
		changeCampaign, err := api.ChangeCampaign(context.TODO(), types.ChangeCampaign{
			ID:    21725,
			Title: "Supernova",
		})
		assert.Equal(t, 0, changeCampaign.Error, changeCampaign.Message)
		assert.Equal(t, nil, err, err)
	})

	t.Run("ChangeCampaignReward", func(t *testing.T) {
		changeCampaignReward, err := api.ChangeCampaignReward(context.TODO(), types.ChangeCampaignReward{
			ID:     21725,
			Amount: 149,
		})
		assert.Equal(t, 0, changeCampaignReward.Error, changeCampaignReward.Message)
		assert.Equal(t, nil, err, err)
	})
}
