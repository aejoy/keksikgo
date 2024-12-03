package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalance(t *testing.T) {
	api, err := newAPI()
	if err != nil {
		t.Error(err)
	}

	t.Run("Balance", func(t *testing.T) {
		donates, err := api.Balance(context.TODO())
		assert.Equal(t, 0, donates.Error, donates.Message)
		assert.Equal(t, nil, err, err)
	})
}
