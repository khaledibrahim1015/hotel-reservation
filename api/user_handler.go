package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khaledibrahim1015/hotel-reservation/types"
)

// GeneralResult General Response
type GeneralResult map[string]interface{}

func HandleGetUsers(ctx *fiber.Ctx) error {
	var user types.User = types.User{
		FirstName: "khaled",
		LastName:  "ibrahim",
	}
	users := []types.User{
		user,
		types.User{
			FirstName: "khaled2",
			LastName:  "ibrahim2",
		},
	}

	return ctx.JSON(GeneralResult{"data": users})

}
func HandleGetUser(ctx *fiber.Ctx) error {
	user := types.User{
		FirstName: "anothony",
		LastName:  "gg",
	}
	return ctx.JSON(GeneralResult{"data": user})
}
