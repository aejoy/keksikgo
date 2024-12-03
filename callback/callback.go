package callback

import "github.com/aejoy/keksikgo/api"

type Callback struct {
	API    *api.API
	Scenes Scenes
}

func New(api *api.API) *Callback {
	return &Callback{API: api}
}
