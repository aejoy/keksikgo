package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPayments(t *testing.T) {
	api, err := newAPI()
	if err != nil {
		t.Error(err)
	}

	t.Run("GetPayments", func(t *testing.T) {
		payments, err := api.GetPayments(context.TODO(), []int{151312, 161722})
		assert.Equal(t, 0, payments.Error, payments.Message)
		assert.Equal(t, nil, err, err)
	})
}
