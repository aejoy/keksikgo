package main

import (
	"fmt"
	"github.com/aejoy/keksikgo/api"
	"github.com/aejoy/keksikgo/callback"
	"github.com/aejoy/keksikgo/update"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"log"
	"os"
	"strconv"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	groupID, err := strconv.Atoi(os.Getenv("GROUP_ID"))
	if err != nil {
		panic(err)
	}

	cake := api.New(groupID, os.Getenv("TOKEN"), logger)

	app := fiber.New()
	session := callback.New(cake)

	session.Donate(func(_ *api.API, donate update.Donate) {
		fmt.Printf("New donate: %+v\n", donate)
	})

	app.Post("/callback/:confirmation", session.Fiber)

	log.Fatalln(app.Listen(":3000"))
}
