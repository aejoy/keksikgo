package callback

import (
	"encoding/json"
	"fmt"
	"github.com/aejoy/keksikgo/update"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func (callback *Callback) Fiber(context fiber.Ctx) error {
	update := update.Update{}
	if err := json.Unmarshal(context.Body(), &update); err != nil {
		callback.API.Logger.Error(err.Error(), zap.String("type", "UnmarshallUpdate"))
		return err
	}

	if callback.API.Logger != nil {
		callback.API.Logger.Info(string(context.Body()), zap.String("type", "NewUpdate"))
	}

	if update.Type == "confirmation" {
		confirmation := context.Params("confirmation", "no_confirmation_param")
		return context.SendString(fmt.Sprintf(`{"status": "ok", "code": "%s"}`, confirmation))
	} else {
		switch update.Type {
		case "new_donate":
			callback.Scenes.Donate(callback.API, update.Donate)
		case "payment_status":
			callback.Scenes.Payment(callback.API, update.Payment)
		}
	}

	return context.SendString(`{"status": "ok"}`)
}
