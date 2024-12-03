package tests

import (
	"context"
	"github.com/aejoy/keksikgo/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDonates(t *testing.T) {
	api, err := newAPI()
	if err != nil {
		t.Error(err)
	}

	t.Run("GetDonates", func(t *testing.T) {
		donates, err := api.GetDonates(context.TODO(), types.GetDonates{Count: 2})
		assert.Equal(t, 0, donates.Error, donates.Message)
		assert.Equal(t, nil, err, err)
	})

	t.Run("GetLastDonates", func(t *testing.T) {
		donates, err := api.GetLastDonates(context.TODO(), types.GetLastDonates{})
		assert.Equal(t, 0, donates.Error, donates.Message)
		assert.Equal(t, nil, err, err)
	})

	t.Run("ChangeStatus", func(t *testing.T) {
		changeStatus, err := api.ChangeStatus(context.TODO(), types.ChangeStatus{
			ID:     1457245,
			Status: "public",
		})
		assert.Equal(t, 0, changeStatus.Error, changeStatus.Message)
		assert.Equal(t, nil, err, err)
	})

	t.Run("Answer", func(t *testing.T) {
		answer, err := api.Answer(context.TODO(), types.Answer{
			ID:     1457245,
			Answer: "Thanks for!",
		})
		assert.Equal(t, 0, answer.Error, answer.Message)
		assert.Equal(t, nil, err, err)
	})

	t.Run("ChangeReward", func(t *testing.T) {
		changeReward, err := api.ChangeReward(context.TODO(), types.ChangeReward{
			ID:     1457245,
			Status: "sended",
		})
		assert.Equal(t, 0, changeReward.Error, changeReward.Message)
		assert.Equal(t, nil, err, err)
	})
}
