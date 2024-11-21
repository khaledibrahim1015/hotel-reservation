package types

import (
	"fmt"

	"github.com/khaledibrahim1015/hotel-reservation/utils"
	"github.com/khaledibrahim1015/hotel-reservation/utils/validator"
)

type User struct {
	// ID        string `bson:"_id" json:"id,omitempty"`
	ID               string `bson:"_id,omitempty" json:"id"`
	FirstName        string `bson:"firstName" json:"firstName"`
	LastName         string `bson:"lastName"lastName"`
	Email            string `bson:"email" json:"email"`
	EncryptedPaaword string `bson:"encryptedPaaword" json:"-"`
}

// Define the interface (union type simulation)
type UserParamInterface interface{}

// request body
type CreateUserParam struct {
	FirstName string `validate:"required,min=3,max=50" json:"firstName"`
	LastName  string `validate:"required,min=3,max=50" json:"lastName"`
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required,min=7,max=50" json:"password"`
}

type UpdateUserParam struct {
	FirstName string `json:"firstName" validate:"min=3,max=50"`
	LastName  string `json:"lastName" validate:"min=3,max=50"`
	Email     string `json:"email" validate:"email"`
}

func MapUserFromParams(param interface{}) (*User, error) {

	switch v := param.(type) { // Type assertion to determine actual type
	case CreateUserParam:
		encryptedPassword, err := utils.HashPassword(v.Password)
		if err != nil {
			return nil, err
		}

		return &User{
			FirstName:        v.FirstName,
			LastName:         v.LastName,
			Email:            v.Email,
			EncryptedPaaword: encryptedPassword,
		}, nil
	case UpdateUserParam:
		user := &User{}
		if v.FirstName != "" {
			user.FirstName = v.FirstName
		}
		if v.LastName != "" {
			user.LastName = v.LastName
		}
		if v.Email != "" {
			user.Email = v.Email
		}

		return user, nil
	default:
		return nil, fmt.Errorf("unsupported parameter type")
	}

}

func (ve *CreateUserParam) Validate() error {

	v := validator.New()
	if err := v.Validate(ve); err != nil {
		return err
	}
	return nil
}

func (ve *UpdateUserParam) Validate() error {

	v := validator.New()
	if err := v.Validate(ve); err != nil {
		return err
	}
	return nil
}
