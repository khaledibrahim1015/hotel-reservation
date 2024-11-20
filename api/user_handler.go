package api

import (
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
		id = ctx.Params("id")
		// context = context.Background()
	)

	user, err := usrH.userStore.GetUserByID(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(GeneralResult{"data": user})
}

func (usrH *UserHandler) HandleGetUsers(ctx *fiber.Ctx) error {

	users, err := usrH.userStore.GetUsers(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(GeneralResult{"data": users})

}

func (usrH *UserHandler) HandlePostUser(ctx *fiber.Ctx) error {

	var params types.CreateUserParam

	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	//  validation
	if err := params.Validate(); err != nil {
		return err
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := usrH.userStore.InsertUser(ctx.Context(), user)
	if err != nil {
		return err
	}
	return ctx.JSON(GeneralResult{"data": insertedUser})

}
