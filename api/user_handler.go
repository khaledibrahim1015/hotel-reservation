package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/khaledibrahim1015/hotel-reservation/db"
	"github.com/khaledibrahim1015/hotel-reservation/types"
)

// GeneralResult General Response
type GeneralResult map[string]interface{}

type UserHandler struct {
	userStore db.UserStore //  interface
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (usrH *UserHandler) HandleGetUser(ctx *fiber.Ctx) error {
	var (
		id      = ctx.Params("id")
		context = context.Background()
	)

	user, err := usrH.userStore.GetUserByID(context, id)
	if err != nil {
		return err
	}

	return ctx.JSON(GeneralResult{"data": user})
}

func (usrH *UserHandler) HandleGetUsers(ctx *fiber.Ctx) error {
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
