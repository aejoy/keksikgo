package callback

import (
	"github.com/aejoy/keksikgo/api"
	"github.com/aejoy/keksikgo/update"
)

type Scenes struct {
	Donate  Donate
	Payment Payment
}

type Donate func(*api.API, update.Donate)
type Payment func(*api.API, update.Payment)

func (callback *Callback) Donate(donate Donate) {
	callback.Scenes.Donate = donate
}

func (callback *Callback) Payment(payment Payment) {
	callback.Scenes.Payment = payment
}
